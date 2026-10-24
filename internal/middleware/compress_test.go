package middleware

import (
	"compress/gzip"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/banua-coder/pico-api-go/internal/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var testCacheCfg = config.CacheConfig{MaxAge: 5 * time.Minute, StaleWhileRevalidate: time.Hour}

func bodyHandler(body string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(body))
	})
}

func serve(h http.Handler, method, path string, hdr map[string]string) *httptest.ResponseRecorder {
	req := httptest.NewRequest(method, path, nil)
	for k, v := range hdr {
		req.Header.Set(k, v)
	}
	w := httptest.NewRecorder()
	h.ServeHTTP(w, req)
	return w
}

func TestGzip_CompressesLargeBody(t *testing.T) {
	body := strings.Repeat(`{"a":1},`, 500)
	w := serve(Gzip(bodyHandler(body)), "GET", "/api/v1/x", map[string]string{"Accept-Encoding": "gzip"})

	assert.Equal(t, "gzip", w.Header().Get("Content-Encoding"))
	assert.Contains(t, w.Header().Get("Vary"), "Accept-Encoding")
	zr, err := gzip.NewReader(w.Body)
	require.NoError(t, err)
	got, _ := io.ReadAll(zr)
	assert.Equal(t, body, string(got))
}

func TestGzip_SkipsWithoutAcceptEncoding(t *testing.T) {
	body := strings.Repeat("x", 5000)
	w := serve(Gzip(bodyHandler(body)), "GET", "/api/v1/x", nil)

	assert.Empty(t, w.Header().Get("Content-Encoding"))
	assert.Contains(t, w.Header().Get("Vary"), "Accept-Encoding")
	assert.Equal(t, body, w.Body.String())
}

func TestGzip_SkipsTinyBody(t *testing.T) {
	w := serve(Gzip(bodyHandler(`{"ok":true}`)), "GET", "/api/v1/x", map[string]string{"Accept-Encoding": "gzip"})

	assert.Empty(t, w.Header().Get("Content-Encoding"))
	assert.Equal(t, `{"ok":true}`, w.Body.String())
}

func TestGzip_SkipsAlreadyEncoded(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Encoding", "br")
		_, _ = w.Write([]byte(strings.Repeat("x", 5000)))
	})
	w := serve(Gzip(h), "GET", "/api/v1/x", map[string]string{"Accept-Encoding": "gzip"})

	assert.Equal(t, "br", w.Header().Get("Content-Encoding"))
	assert.Len(t, w.Body.String(), 5000)
}

func TestGzip_RejectsGzipQZero(t *testing.T) {
	w := serve(Gzip(bodyHandler(strings.Repeat("x", 5000))), "GET", "/api/v1/x", map[string]string{"Accept-Encoding": "gzip;q=0"})
	assert.Empty(t, w.Header().Get("Content-Encoding"))
}

func TestGzip_NoBodyStatusUntouched(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(http.StatusNotModified) })
	w := serve(Gzip(h), "GET", "/api/v1/x", map[string]string{"Accept-Encoding": "gzip"})

	assert.Equal(t, http.StatusNotModified, w.Code)
	assert.Empty(t, w.Header().Get("Content-Encoding"))
}

func TestCache_SetsHeadersAndETag(t *testing.T) {
	w := serve(Cache(testCacheCfg)(bodyHandler(`{"a":1}`)), "GET", "/api/v1/national", nil)

	assert.Equal(t, "public, max-age=300, stale-while-revalidate=3600", w.Header().Get("Cache-Control"))
	assert.Regexp(t, `^W/"[0-9a-f]+"$`, w.Header().Get("ETag"))
	assert.Equal(t, `{"a":1}`, w.Body.String())
}

func TestCache_IfNoneMatchReturns304(t *testing.T) {
	h := Cache(testCacheCfg)(bodyHandler(`{"a":1}`))
	etag := serve(h, "GET", "/api/v1/national", nil).Header().Get("ETag")

	w := serve(h, "GET", "/api/v1/national", map[string]string{"If-None-Match": etag})
	assert.Equal(t, http.StatusNotModified, w.Code)
	assert.Empty(t, w.Body.String())
	assert.Equal(t, etag, w.Header().Get("ETag"))

	w = serve(h, "GET", "/api/v1/national", map[string]string{"If-None-Match": `W/"other"`})
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCache_HealthIsNoStore(t *testing.T) {
	w := serve(Cache(testCacheCfg)(bodyHandler(`{"ok":1}`)), "GET", "/api/v1/health", nil)

	assert.Equal(t, "no-store", w.Header().Get("Cache-Control"))
	assert.Empty(t, w.Header().Get("ETag"))
}

func TestCache_SkipsErrorsNonGETAndOtherPaths(t *testing.T) {
	notFound := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte("nope"))
	})
	w := serve(Cache(testCacheCfg)(notFound), "GET", "/api/v1/x", nil)
	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Empty(t, w.Header().Get("Cache-Control"))
	assert.Equal(t, "nope", w.Body.String())

	h := Cache(testCacheCfg)(bodyHandler("ok"))
	assert.Empty(t, serve(h, "POST", "/admin/cache/clear", nil).Header().Get("Cache-Control"))
	assert.Empty(t, serve(h, "GET", "/swagger/index.html", nil).Header().Get("Cache-Control"))
}

func TestGzipAndCache_Together(t *testing.T) {
	body := strings.Repeat(`{"a":1},`, 500)
	h := Gzip(Cache(testCacheCfg)(bodyHandler(body)))
	w := serve(h, "GET", "/api/v1/national", map[string]string{"Accept-Encoding": "gzip"})

	assert.Equal(t, "gzip", w.Header().Get("Content-Encoding"))
	assert.NotEmpty(t, w.Header().Get("ETag"))
	assert.NotEmpty(t, w.Header().Get("Cache-Control"))
}

// Over a real server, headers must be sent before the body, so Content-Encoding has to be set
// before WriteHeader reaches the wire (the recorder-based tests cannot catch this).
func TestGzipAndCache_OverRealServer(t *testing.T) {
	body := strings.Repeat(`{"a":1},`, 500)
	srv := httptest.NewServer(Gzip(Cache(testCacheCfg)(bodyHandler(body))))
	defer srv.Close()

	req, _ := http.NewRequest(http.MethodGet, srv.URL+"/api/v1/national", nil)
	req.Header.Set("Accept-Encoding", "gzip")
	resp, err := (&http.Transport{DisableCompression: true}).RoundTrip(req)
	require.NoError(t, err)
	defer func() { _ = resp.Body.Close() }()

	assert.Equal(t, "gzip", resp.Header.Get("Content-Encoding"))
	zr, err := gzip.NewReader(resp.Body)
	require.NoError(t, err)
	got, _ := io.ReadAll(zr)
	assert.Equal(t, body, string(got))

	req2, _ := http.NewRequest(http.MethodGet, srv.URL+"/api/v1/national", nil)
	req2.Header.Set("If-None-Match", resp.Header.Get("ETag"))
	resp2, err := (&http.Transport{DisableCompression: true}).RoundTrip(req2)
	require.NoError(t, err)
	defer func() { _ = resp2.Body.Close() }()
	assert.Equal(t, http.StatusNotModified, resp2.StatusCode)
}
