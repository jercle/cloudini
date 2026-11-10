package lib

import (
	"net/http"
	"os"
	"time"
)

// DefaultHTTPTimeout is the overall per-request timeout (connect, headers and
// body) applied to HTTPClient. Override with the CLD_HTTP_TIMEOUT environment
// variable, e.g. "30s" or "15m". A value of "0" disables the timeout.
const DefaultHTTPTimeout = 5 * time.Minute

// HTTPClient is the shared client used for all outbound API calls.
//
// http.DefaultClient has no timeout, so a stalled connection hangs the CLI
// forever, and its transport only keeps 2 idle connections per host, which
// forces new TLS handshakes when many goroutines call the same API.
var HTTPClient = newHTTPClient()

func newHTTPClient() *http.Client {
	timeout := DefaultHTTPTimeout
	if v := os.Getenv("CLD_HTTP_TIMEOUT"); v != "" {
		if d, err := time.ParseDuration(v); err == nil && d >= 0 {
			timeout = d
		}
	}

	transport := http.DefaultTransport.(*http.Transport).Clone()
	transport.MaxIdleConns = 256
	transport.MaxIdleConnsPerHost = 64

	return &http.Client{
		Timeout:   timeout,
		Transport: transport,
	}
}
