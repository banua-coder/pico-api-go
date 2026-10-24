package middleware

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/banua-coder/pico-api-go/internal/config"
)

const (
	apiPrefix   = "/api/v1"
	gzipMinSize = 1024
)

func isAPIRead(r *http.Request) bool {
	return (r.Method == http.MethodGet || r.Method == http.MethodHead) && strings.HasPrefix(r.URL.Path, apiPrefix)
}

// acceptsGzip follows RFC 9110 12.5.3: an explicit gzip entry wins over "*",
// and a q-value of 0 means "not acceptable". An unparsable q is treated as not acceptable.
func acceptsGzip(r *http.Request) bool {
	star, starSeen := false, false
	for _, part := range strings.Split(r.Header.Get("Accept-Encoding"), ",") {
		coding, params, _ := strings.Cut(strings.TrimSpace(part), ";")
		coding = strings.TrimSpace(coding)
		ok := qValueOK(params)
		switch {
		case strings.EqualFold(coding, "gzip"):
			return ok
		case coding == "*":
			star, starSeen = ok, true
		}
	}
	return starSeen && star
}

// qValueOK reports whether the q parameter (default 1) is greater than zero.
func qValueOK(params string) bool {
	for _, p := range strings.Split(params, ";") {
		k, v, found := strings.Cut(strings.TrimSpace(p), "=")
		if found && strings.EqualFold(strings.TrimSpace(k), "q") {
			q, err := strconv.ParseFloat(strings.TrimSpace(v), 64)
			return err == nil && q > 0
		}
	}
	return true
}

// gzipWriter decides on the first Write whether to compress.
// ponytail: decision uses only the first chunk's size; fine because the cache
// middleware writes the whole body at once. Revisit if streaming handlers appear.
type gzipWriter struct {
	http.ResponseWriter
	gz      *gzip.Writer
	decided bool
	status  int // held back until the first Write so Content-Encoding can still be set
}

func (g *gzipWriter) WriteHeader(code int) {
	if code == http.StatusNoContent || code == http.StatusNotModified {
		g.decided = true // no body, nothing to compress
		g.ResponseWriter.WriteHeader(code)
		return
	}
	g.status = code
}

func (g *gzipWriter) flushHeader() {
	if g.status != 0 {
		g.ResponseWriter.WriteHeader(g.status)
		g.status = 0
	}
}

func (g *gzipWriter) Write(p []byte) (int, error) {
	if !g.decided {
		g.decided = true
		h := g.Header()
		if len(p) >= gzipMinSize && h.Get("Content-Encoding") == "" {
			h.Set("Content-Encoding", "gzip")
			h.Del("Content-Length")
			g.gz = gzip.NewWriter(g.ResponseWriter)
		}
	}
	g.flushHeader()
	if g.gz != nil {
		return g.gz.Write(p)
	}
	return g.ResponseWriter.Write(p)
}

func (g *gzipWriter) Flush() {
	g.flushHeader()
	if g.gz != nil {
		_ = g.gz.Flush()
	}
	if f, ok := g.ResponseWriter.(http.Flusher); ok {
		f.Flush()
	}
}

// Gzip compresses /api/v1 GET responses for clients that accept gzip.
// Bodies under 1 KB and responses that already set Content-Encoding are left alone.
func Gzip(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !isAPIRead(r) {
			next.ServeHTTP(w, r)
			return
		}
		w.Header().Add("Vary", "Accept-Encoding")
		if !acceptsGzip(r) {
			next.ServeHTTP(w, r)
			return
		}
		gw := &gzipWriter{ResponseWriter: w}
		defer func() {
			gw.flushHeader()
			if gw.gz != nil {
				_ = gw.gz.Close()
			}
		}()
		next.ServeHTTP(gw, r)
	})
}

// bufferedWriter holds status and body so the cache middleware can hash it.
// Headers go straight to the real writer.
type bufferedWriter struct {
	http.ResponseWriter
	status int
	buf    bytes.Buffer
}

func (b *bufferedWriter) WriteHeader(code int) { b.status = code }

func (b *bufferedWriter) Write(p []byte) (int, error) { return b.buf.Write(p) }

// Cache adds Cache-Control and a weak ETag to successful /api/v1 GET responses
// and answers matching If-None-Match with 304. /health is never cached.
func Cache(cfg config.CacheConfig) func(http.Handler) http.Handler {
	policy := "no-cache"
	if cfg.MaxAge > 0 {
		policy = fmt.Sprintf("public, max-age=%d", int(cfg.MaxAge.Seconds()))
		if cfg.StaleWhileRevalidate > 0 {
			policy += fmt.Sprintf(", stale-while-revalidate=%d", int(cfg.StaleWhileRevalidate.Seconds()))
		}
	}
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if !isAPIRead(r) {
				next.ServeHTTP(w, r)
				return
			}
			if r.URL.Path == apiPrefix+"/health" {
				w.Header().Set("Cache-Control", "no-store")
				next.ServeHTTP(w, r)
				return
			}

			bw := &bufferedWriter{ResponseWriter: w, status: http.StatusOK}
			next.ServeHTTP(bw, r)

			if bw.status == http.StatusOK {
				sum := sha256.Sum256(bw.buf.Bytes())
				etag := `W/"` + hex.EncodeToString(sum[:8]) + `"`
				w.Header().Set("ETag", etag)
				w.Header().Set("Cache-Control", policy)
				if etagMatches(r.Header.Get("If-None-Match"), etag) {
					// Quota headers set upstream describe the 200 path; drop them on 304 so clients
					// do not cache a stale remaining count.
					for _, h := range []string{"X-RateLimit-Limit", "X-RateLimit-Remaining", "X-RateLimit-Reset"} {
						w.Header().Del(h)
					}
					w.WriteHeader(http.StatusNotModified)
					return
				}
			}
			w.WriteHeader(bw.status)
			if r.Method != http.MethodHead {
				_, _ = w.Write(bw.buf.Bytes())
			}
		})
	}
}

func etagMatches(header, etag string) bool {
	for _, candidate := range strings.Split(header, ",") {
		if c := strings.TrimSpace(candidate); c == "*" || c == etag {
			return true
		}
	}
	return false
}
