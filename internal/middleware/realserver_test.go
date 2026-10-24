package middleware

import (
	"compress/gzip"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func rawGet(t *testing.T, method, url string, hdr map[string]string) (*http.Response, []byte) {
	t.Helper()
	req, err := http.NewRequest(method, url, nil)
	require.NoError(t, err)
	for k, v := range hdr {
		req.Header.Set(k, v)
	}
	resp, err := (&http.Transport{DisableCompression: true}).RoundTrip(req)
	require.NoError(t, err)
	defer func() { _ = resp.Body.Close() }()
	b, _ := io.ReadAll(resp.Body)
	return resp, b
}

func newGzipCacheServer(body string) *httptest.Server {
	return httptest.NewServer(Gzip(Cache(testCacheCfg)(bodyHandler(body))))
}

func TestRealServer_HeadHasNoBodyButHeaders(t *testing.T) {
	srv := newGzipCacheServer(strings.Repeat(`{"a":1},`, 500))
	defer srv.Close()
	resp, b := rawGet(t, http.MethodHead, srv.URL+"/api/v1/national", map[string]string{"Accept-Encoding": "gzip"})
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Empty(t, b)
	assert.NotEmpty(t, resp.Header.Get("ETag"))
	assert.NotEmpty(t, resp.Header.Get("Cache-Control"))
}

func TestRealServer_GzipClientConditionalGets304(t *testing.T) {
	srv := newGzipCacheServer(strings.Repeat(`{"a":1},`, 500))
	defer srv.Close()
	first, _ := rawGet(t, http.MethodGet, srv.URL+"/api/v1/national", map[string]string{"Accept-Encoding": "gzip"})
	resp, b := rawGet(t, http.MethodGet, srv.URL+"/api/v1/national",
		map[string]string{"Accept-Encoding": "gzip", "If-None-Match": first.Header.Get("ETag")})
	assert.Equal(t, http.StatusNotModified, resp.StatusCode)
	assert.Empty(t, b)
	assert.Empty(t, resp.Header.Get("Content-Encoding"))
}

func TestRealServer_LargeBodyRoundTrips(t *testing.T) {
	body := strings.Repeat("0123456789abcdef", 1<<16) // 1 MiB
	srv := newGzipCacheServer(body)
	defer srv.Close()
	resp, b := rawGet(t, http.MethodGet, srv.URL+"/api/v1/national", map[string]string{"Accept-Encoding": "gzip"})
	assert.Equal(t, "gzip", resp.Header.Get("Content-Encoding"))
	zr, err := gzip.NewReader(strings.NewReader(string(b)))
	require.NoError(t, err)
	got, err := io.ReadAll(zr)
	require.NoError(t, err)
	assert.True(t, body == string(got))
}

func TestRealServer_AcceptEncodingQValues(t *testing.T) {
	srv := newGzipCacheServer(strings.Repeat(`{"a":1},`, 500))
	defer srv.Close()
	for in, want := range map[string]string{
		"GZIP": "gzip", "*": "gzip", "gzip;q=0.5": "gzip",
		"gzip;q=0": "", "gzip;q=0.000": "", "*;q=0": "",
	} {
		resp, _ := rawGet(t, http.MethodGet, srv.URL+"/api/v1/national", map[string]string{"Accept-Encoding": in})
		assert.Equal(t, want, resp.Header.Get("Content-Encoding"), "Accept-Encoding: %q", in)
	}
}

func TestCache_304DropsRateLimitHeaders(t *testing.T) {
	rl := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("X-RateLimit-Limit", "100")
			w.Header().Set("X-RateLimit-Remaining", "99")
			next.ServeHTTP(w, r)
		})
	}
	h := rl(Cache(testCacheCfg)(bodyHandler("ok")))
	first := serve(h, "GET", "/api/v1/x", nil)
	assert.Equal(t, "100", first.Header().Get("X-RateLimit-Limit"))
	w := serve(h, "GET", "/api/v1/x", map[string]string{"If-None-Match": first.Header().Get("ETag")})
	assert.Equal(t, http.StatusNotModified, w.Code)
	assert.Empty(t, w.Header().Get("X-RateLimit-Limit"))
	assert.Empty(t, w.Header().Get("X-RateLimit-Remaining"))
}
