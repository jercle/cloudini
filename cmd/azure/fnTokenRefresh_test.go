package azure

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/jercle/cloudini/lib"
)

// Registers a token expiring at the given time. Each refresh issues
// "<name>-1", "<name>-2", ... valid for an hour
func registerTestToken(name string, expiresOn time.Time) (lib.AzureTokenData, *atomic.Int32) {
	var refreshes atomic.Int32
	original := lib.AzureTokenData{Token: name, ExpiresOn: expiresOn}
	registerTokenSource(original, func() (*lib.AzureTokenData, error) {
		n := refreshes.Add(1)
		return &lib.AzureTokenData{
			Token:     fmt.Sprintf("%s-%d", name, n),
			ExpiresOn: time.Now().Add(time.Hour),
		}, nil
	})
	return original, &refreshes
}

func TestFreshTokenDataUnregisteredTokenUnchanged(t *testing.T) {
	original := lib.AzureTokenData{Token: "unregistered", ExpiresOn: time.Now().Add(-time.Hour)}
	if got := FreshTokenData(original); got != original {
		t.Fatalf("expected unregistered token to be returned as is, got %+v", got)
	}
}

func TestFreshTokenDataValidTokenNotRefreshed(t *testing.T) {
	original, refreshes := registerTestToken("valid", time.Now().Add(time.Hour))
	if got := FreshTokenData(original); got.Token != "valid" {
		t.Fatalf("expected valid token to be kept, got %q", got.Token)
	}
	if refreshes.Load() != 0 {
		t.Fatalf("expected no refresh, got %d", refreshes.Load())
	}
}

func TestFreshTokenDataRefreshesOnceForConcurrentStaleCopies(t *testing.T) {
	// Inside the refresh window, so should be replaced before use
	original, refreshes := registerTestToken("expiring", time.Now().Add(time.Minute))

	var wg sync.WaitGroup
	results := make([]string, 50)
	for i := range results {
		wg.Add(1)
		go func() {
			defer wg.Done()
			results[i] = FreshTokenData(original).Token
		}()
	}
	wg.Wait()

	for _, got := range results {
		if got != "expiring-1" {
			t.Fatalf("expected every stale copy to resolve to expiring-1, got %q", got)
		}
	}
	if refreshes.Load() != 1 {
		t.Fatalf("expected a single refresh, got %d", refreshes.Load())
	}

	mat := FreshToken(lib.AzureMultiAuthToken{TenantName: "tenant", TokenData: original})
	if mat.TokenData.Token != "expiring-1" || mat.TenantName != "tenant" {
		t.Fatalf("unexpected token from FreshToken: %+v", mat)
	}
}

func TestFreshTokenDataKeepsTokenWhenRefreshFails(t *testing.T) {
	var attempts atomic.Int32
	original := lib.AzureTokenData{Token: "refresh-fails", ExpiresOn: time.Now().Add(-time.Minute)}
	registerTokenSource(original, func() (*lib.AzureTokenData, error) {
		attempts.Add(1)
		return nil, fmt.Errorf("token endpoint unavailable")
	})

	for range 3 {
		if got := FreshTokenData(original); got != original {
			t.Fatalf("expected existing token to be kept, got %+v", got)
		}
	}
	// Further attempts are held off until the retry interval has passed
	if attempts.Load() != 1 {
		t.Fatalf("expected 1 refresh attempt, got %d", attempts.Load())
	}
}

func TestDoRequestWithTokenRefreshReplacesExpiredToken(t *testing.T) {
	original, _ := registerTestToken("expired", time.Now().Add(-time.Minute))

	var gotAuth string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotAuth = r.Header.Get("Authorization")
	}))
	defer srv.Close()

	req, _ := http.NewRequest(http.MethodGet, srv.URL, nil)
	res, err := doRequestWithTokenRefresh(req, original)
	if err != nil {
		t.Fatal(err)
	}
	res.Body.Close()

	if gotAuth != "Bearer expired-1" {
		t.Fatalf("expected request to use refreshed token, got %q", gotAuth)
	}
}

func TestDoRequestWithTokenRefreshRetriesOn401(t *testing.T) {
	// Not due for refresh by expiry time, but rejected by the API
	original, refreshes := registerTestToken("revoked", time.Now().Add(time.Hour))

	var (
		requests int
		bodies   []string
	)
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requests++
		body, _ := io.ReadAll(r.Body)
		bodies = append(bodies, string(body))
		if r.Header.Get("Authorization") != "Bearer revoked-1" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.Write([]byte("ok"))
	}))
	defer srv.Close()

	req, _ := http.NewRequest(http.MethodPost, srv.URL, strings.NewReader(`{"query": "x"}`))
	res, err := doRequestWithTokenRefresh(req, original)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()
	resBody, _ := io.ReadAll(res.Body)

	if res.StatusCode != http.StatusOK || string(resBody) != "ok" {
		t.Fatalf("expected retried request to succeed, got %d %q", res.StatusCode, resBody)
	}
	if requests != 2 || refreshes.Load() != 1 {
		t.Fatalf("expected 2 requests and 1 refresh, got %d and %d", requests, refreshes.Load())
	}
	if bodies[0] != `{"query": "x"}` || bodies[1] != bodies[0] {
		t.Fatalf("expected body to be resent on retry, got %q", bodies)
	}
}

func TestDoRequestWithTokenRefreshReturns401ForUnregisteredToken(t *testing.T) {
	requests := 0
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requests++
		w.WriteHeader(http.StatusUnauthorized)
	}))
	defer srv.Close()

	req, _ := http.NewRequest(http.MethodGet, srv.URL, nil)
	res, err := doRequestWithTokenRefresh(req, lib.AzureTokenData{Token: "cli-token"})
	if err != nil {
		t.Fatal(err)
	}
	res.Body.Close()

	if res.StatusCode != http.StatusUnauthorized || requests != 1 {
		t.Fatalf("expected a single 401 response, got %d after %d requests", res.StatusCode, requests)
	}
}
