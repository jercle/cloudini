package azure

import (
	json "encoding/json/v2"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/jercle/cloudini/lib"
)

type recordedRequest struct {
	Method        string
	Path          string
	Query         string
	Authorization string
	ContentType   string
	Body          string
}

// Starts a test API that records every request it receives.
//
// The response is picked from the request path:
//
//	/notfound  - 404 with a JSON error body
//	/forbidden - 403 with a JSON error body
//	anything else - 200, echoing the request method
func newRecordingApi(t *testing.T) (*httptest.Server, func() []recordedRequest) {
	t.Helper()
	var (
		mu       sync.Mutex
		requests []recordedRequest
	)
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		mu.Lock()
		requests = append(requests, recordedRequest{
			Method:        r.Method,
			Path:          r.URL.Path,
			Query:         r.URL.RawQuery,
			Authorization: r.Header.Get("Authorization"),
			ContentType:   r.Header.Get("Content-Type"),
			Body:          string(body),
		})
		mu.Unlock()

		w.Header().Set("X-Ms-Request-Id", "req-123")
		switch r.URL.Path {
		case "/notfound":
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, `{"error":{"code":"ResourceNotFound"}}`)
		case "/forbidden":
			w.WriteHeader(http.StatusForbidden)
			fmt.Fprint(w, `{"error":{"code":"AuthorizationFailed"}}`)
		default:
			fmt.Fprintf(w, `{"method":"%s"}`, r.Method)
		}
	}))
	t.Cleanup(srv.Close)

	return srv, func() []recordedRequest {
		mu.Lock()
		defer mu.Unlock()
		return append([]recordedRequest(nil), requests...)
	}
}

func testMultiAuthToken(token string) lib.AzureMultiAuthToken {
	return lib.AzureMultiAuthToken{
		TenantName: "TEST",
		TokenData:  lib.AzureTokenData{Token: token, ExpiresOn: time.Now().Add(time.Hour)},
	}
}

func lastRequest(t *testing.T, requests []recordedRequest) recordedRequest {
	t.Helper()
	if len(requests) == 0 {
		t.Fatal("the API did not receive a request")
	}
	return requests[len(requests)-1]
}

// The Http* helpers return response headers as JSON
func decodeHeaders(t *testing.T, headerJson []byte) http.Header {
	t.Helper()
	var headers http.Header
	if err := json.Unmarshal(headerJson, &headers); err != nil {
		t.Fatalf("expected response headers as JSON, got %q: %v", headerJson, err)
	}
	return headers
}

//
// HttpGet
//

func TestHttpGet(t *testing.T) {
	srv, requests := newRecordingApi(t)

	body, err := HttpGet(srv.URL+"/subscriptions?api-version=2022-12-01", testMultiAuthToken("get-token"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if string(body) != `{"method":"GET"}` {
		t.Fatalf("unexpected body %q", body)
	}

	req := lastRequest(t, requests())
	if req.Method != http.MethodGet {
		t.Fatalf("unexpected method %q", req.Method)
	}
	if req.Path != "/subscriptions" || req.Query != "api-version=2022-12-01" {
		t.Fatalf("unexpected path or query: %q, %q", req.Path, req.Query)
	}
	if req.Authorization != "Bearer get-token" {
		t.Fatalf("unexpected Authorization header %q", req.Authorization)
	}
	if req.ContentType != "application/json" {
		t.Fatalf("unexpected Content-Type header %q", req.ContentType)
	}
}

func TestHttpGetNotFoundReturnsError(t *testing.T) {
	srv, _ := newRecordingApi(t)

	body, err := HttpGet(srv.URL+"/notfound", testMultiAuthToken("get-token"))
	if err == nil {
		t.Fatal("expected an error for a 404 response")
	}
	if body != nil {
		t.Fatalf("expected no body with the error, got %q", body)
	}
	// The error is JSON carrying the status line and the API's response
	var parsed struct {
		Status   string `json:"status"`
		Response struct {
			Error struct {
				Code string `json:"code"`
			} `json:"error"`
		} `json:"response"`
	}
	if jsonErr := json.Unmarshal([]byte(err.Error()), &parsed); jsonErr != nil {
		t.Fatalf("expected the error to be JSON, got %q: %v", err.Error(), jsonErr)
	}
	if !strings.HasPrefix(parsed.Status, "404") || parsed.Response.Error.Code != "ResourceNotFound" {
		t.Fatalf("unexpected error content %q", err.Error())
	}
}

func TestHttpGetForbiddenReturnsBody(t *testing.T) {
	srv, _ := newRecordingApi(t)

	// A 403 is printed, but the response body is still handed back without an error
	body, err := HttpGet(srv.URL+"/forbidden", testMultiAuthToken("get-token"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(string(body), "AuthorizationFailed") {
		t.Fatalf("unexpected body %q", body)
	}
}

func TestHttpGetConnectionFailureReturnsError(t *testing.T) {
	srv, _ := newRecordingApi(t)
	url := srv.URL + "/anything"
	srv.Close()

	body, err := HttpGet(url, testMultiAuthToken("get-token"))
	if err == nil {
		t.Fatalf("expected an error when the API cannot be reached, got body %q", body)
	}
}

func TestHttpGetUsesRefreshedTokenForExpiredRegisteredToken(t *testing.T) {
	srv, requests := newRecordingApi(t)
	name := uniqueTokenName("http-get-expired")
	expired, refreshes := registerTestToken(name, time.Now().Add(-time.Minute))

	if _, err := HttpGet(srv.URL+"/resources", lib.AzureMultiAuthToken{TokenData: expired}); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if got := lastRequest(t, requests()).Authorization; got != "Bearer "+name+"-1" {
		t.Fatalf("expected the refreshed token to be sent, got %q", got)
	}
	if n := refreshes.Load(); n != 1 {
		t.Fatalf("expected a single refresh, got %d", n)
	}
}

//
// HttpPost
//

func TestHttpPost(t *testing.T) {
	srv, requests := newRecordingApi(t)
	const payload = `{"query":"Resources | take 1"}`

	body, headers, err := HttpPost(srv.URL+"/providers/Microsoft.ResourceGraph/resources", payload, testMultiAuthToken("post-token"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if string(body) != `{"method":"POST"}` {
		t.Fatalf("unexpected body %q", body)
	}
	if got := decodeHeaders(t, headers).Get("X-Ms-Request-Id"); got != "req-123" {
		t.Fatalf("expected response headers to be returned, got request id %q", got)
	}

	req := lastRequest(t, requests())
	if req.Method != http.MethodPost {
		t.Fatalf("unexpected method %q", req.Method)
	}
	if req.Body != payload {
		t.Fatalf("got request body %q, want %q", req.Body, payload)
	}
	if req.Authorization != "Bearer post-token" {
		t.Fatalf("unexpected Authorization header %q", req.Authorization)
	}
	if req.ContentType != "application/json" {
		t.Fatalf("unexpected Content-Type header %q", req.ContentType)
	}
}

func TestHttpPostErrorStatuses(t *testing.T) {
	srv, _ := newRecordingApi(t)

	for path, wantStatus := range map[string]string{"/notfound": "404", "/forbidden": "403"} {
		t.Run(path, func(t *testing.T) {
			body, headers, err := HttpPost(srv.URL+path, `{}`, testMultiAuthToken("post-token"))
			if err == nil {
				t.Fatalf("expected an error for a %s response", wantStatus)
			}
			if !strings.Contains(err.Error(), wantStatus) {
				t.Fatalf("expected the error to carry the status, got %q", err.Error())
			}
			if body != nil || headers != nil {
				t.Fatalf("expected no body or headers with the error, got %q and %q", body, headers)
			}
		})
	}
}

func TestHttpPostResendsBodyWhenRetryingAfter401(t *testing.T) {
	const payload = `{"important":"payload"}`
	name := uniqueTokenName("http-post-rejected")
	var (
		mu     sync.Mutex
		bodies []string
		auths  []string
	)
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		mu.Lock()
		bodies = append(bodies, string(body))
		auths = append(auths, r.Header.Get("Authorization"))
		mu.Unlock()

		if r.Header.Get("Authorization") == "Bearer "+name {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		fmt.Fprint(w, `{"ok":true}`)
	}))
	defer srv.Close()

	// Not expired as far as the client can tell, but the API rejects it
	rejected, refreshes := registerTestToken(name, time.Now().Add(time.Hour))

	body, _, err := HttpPost(srv.URL+"/things", payload, lib.AzureMultiAuthToken{TokenData: rejected})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if string(body) != `{"ok":true}` {
		t.Fatalf("expected the response from the retried request, got %q", body)
	}

	mu.Lock()
	defer mu.Unlock()
	if len(bodies) != 2 {
		t.Fatalf("expected the request to be sent twice, got %d", len(bodies))
	}
	if bodies[0] != payload || bodies[1] != payload {
		t.Fatalf("expected the full body on both attempts, got %q", bodies)
	}
	if auths[1] != "Bearer "+name+"-1" {
		t.Fatalf("expected the retry to use the new token, got %q", auths[1])
	}
	if n := refreshes.Load(); n != 1 {
		t.Fatalf("expected a single refresh, got %d", n)
	}
}

//
// HttpPut
//

func TestHttpPut(t *testing.T) {
	srv, requests := newRecordingApi(t)
	const payload = `{"properties":{"enabled":true}}`

	body, headers, err := HttpPut(srv.URL+"/resource/one", payload, testMultiAuthToken("put-token"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if string(body) != `{"method":"PUT"}` {
		t.Fatalf("unexpected body %q", body)
	}
	if got := decodeHeaders(t, headers).Get("X-Ms-Request-Id"); got != "req-123" {
		t.Fatalf("expected response headers to be returned, got request id %q", got)
	}

	req := lastRequest(t, requests())
	if req.Method != http.MethodPut || req.Body != payload {
		t.Fatalf("unexpected request: %s with body %q", req.Method, req.Body)
	}
	if req.Authorization != "Bearer put-token" || req.ContentType != "application/json" {
		t.Fatalf("unexpected headers: %q, %q", req.Authorization, req.ContentType)
	}
}

func TestHttpPutForbiddenReturnsBody(t *testing.T) {
	srv, _ := newRecordingApi(t)

	body, _, err := HttpPut(srv.URL+"/forbidden", `{}`, testMultiAuthToken("put-token"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(string(body), "AuthorizationFailed") {
		t.Fatalf("unexpected body %q", body)
	}
}

//
// HttpPatch
//

func TestHttpPatch(t *testing.T) {
	srv, requests := newRecordingApi(t)
	const payload = `{"tags":{"env":"prod"}}`

	body, headers, err := HttpPatch(srv.URL+"/resource/one", payload, testMultiAuthToken("patch-token"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if string(body) != `{"method":"PATCH"}` {
		t.Fatalf("unexpected body %q", body)
	}
	if got := decodeHeaders(t, headers).Get("X-Ms-Request-Id"); got != "req-123" {
		t.Fatalf("expected response headers to be returned, got request id %q", got)
	}

	req := lastRequest(t, requests())
	if req.Method != http.MethodPatch || req.Body != payload {
		t.Fatalf("unexpected request: %s with body %q", req.Method, req.Body)
	}
	if req.Authorization != "Bearer patch-token" || req.ContentType != "application/json" {
		t.Fatalf("unexpected headers: %q, %q", req.Authorization, req.ContentType)
	}
}

func TestHttpPatchForbiddenReturnsBody(t *testing.T) {
	srv, _ := newRecordingApi(t)

	body, _, err := HttpPatch(srv.URL+"/forbidden", `{}`, testMultiAuthToken("patch-token"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(string(body), "AuthorizationFailed") {
		t.Fatalf("unexpected body %q", body)
	}
}

func TestHttpPatchInvalidUrlReturnsError(t *testing.T) {
	if _, _, err := HttpPatch("http://bad host/", `{}`, testMultiAuthToken("patch-token")); err == nil {
		t.Fatal("expected an error for an invalid URL")
	}

	srv, _ := newRecordingApi(t)
	url := srv.URL + "/anything"
	srv.Close()
	if _, _, err := HttpPatch(url, `{}`, testMultiAuthToken("patch-token")); err == nil {
		t.Fatal("expected an error when the API cannot be reached")
	}
}
