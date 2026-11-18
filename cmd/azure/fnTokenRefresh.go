package azure

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/jercle/cloudini/lib"
)

// Tokens are passed around this codebase by value, so a long running job keeps
// using the copy it was handed at startup. To stop that copy expiring mid-run,
// every token issued by GetServicePrincipalToken (and therefore GetTenantSPToken
// and GetAllTenantSPTokens) and GetServicePrincipalMultiAuthToken is registered
// here along with a function that knows how to request a replacement.
//
// Anything about to send a token calls FreshToken/FreshTokenData (the Http*
// helpers do this for you), which swaps a stale copy for the current token and
// requests a new one shortly before the current one expires.

const (
	// A token is replaced once it is within this window of expiring, so that
	// requests sent with it do not expire in flight.
	tokenRefreshSkew = 5 * time.Minute
	// Minimum gap between refresh attempts for the same token source. Stops a
	// failing token endpoint, or an API returning 401 for reasons other than
	// expiry, from triggering a token request per call.
	tokenRefreshRetryInterval = 30 * time.Second
)

type tokenSource struct {
	mu         sync.Mutex
	current    lib.AzureTokenData
	refresh    func() (*lib.AzureTokenData, error)
	retryAfter time.Time
}

var (
	tokenSourcesMu sync.RWMutex
	// Keyed by token string. Every token a source has issued maps to that
	// source, so stale copies still resolve to the current token.
	tokenSources = map[string]*tokenSource{}
)

// Records how to obtain a replacement for the given token
func registerTokenSource(tokenData lib.AzureTokenData, refresh func() (*lib.AzureTokenData, error)) {
	if tokenData.Token == "" || refresh == nil {
		return
	}
	tokenSourcesMu.Lock()
	defer tokenSourcesMu.Unlock()
	if _, exists := tokenSources[tokenData.Token]; exists {
		return
	}
	tokenSources[tokenData.Token] = &tokenSource{
		current: tokenData,
		refresh: refresh,
	}
}

// Returns the current token for the source that issued tokenStr, requesting a
// new one if it is about to expire. When rejected is true the caller has had
// tokenStr refused with a 401, so a new token is requested if tokenStr is still
// the current one.
//
// The bool is false when tokenStr did not come from a registered source.
func resolveToken(tokenStr string, rejected bool) (lib.AzureTokenData, bool) {
	tokenSourcesMu.RLock()
	src := tokenSources[tokenStr]
	tokenSourcesMu.RUnlock()
	if src == nil {
		return lib.AzureTokenData{}, false
	}

	// Held while refreshing so concurrent callers wait for, and then share,
	// a single new token rather than each requesting their own
	src.mu.Lock()
	defer src.mu.Unlock()

	now := time.Now()
	needsRefresh := !src.current.ExpiresOn.After(now.Add(tokenRefreshSkew))
	if rejected && src.current.Token == tokenStr {
		needsRefresh = true
	}
	if !needsRefresh || now.Before(src.retryAfter) {
		return src.current, true
	}
	src.retryAfter = now.Add(tokenRefreshRetryInterval)

	newToken, err := src.refresh()
	if err != nil || newToken == nil || newToken.Token == "" {
		fmt.Fprintln(os.Stderr, "Unable to refresh Azure token, continuing with existing token:", err)
		return src.current, true
	}

	src.current = *newToken
	tokenSourcesMu.Lock()
	tokenSources[newToken.Token] = src
	tokenSourcesMu.Unlock()

	return src.current, true
}

// Returns up to date token data for the given token data. If the token has
// expired, or is about to, a new one is requested. Token data that was not
// issued by one of the service principal token functions is returned as is.
func FreshTokenData(tokenData lib.AzureTokenData) lib.AzureTokenData {
	current, ok := resolveToken(tokenData.Token, false)
	if !ok {
		return tokenData
	}
	return current
}

// Returns a copy of the given token with up to date token data.
//
// Use this before handing a token to anything that does not go through the
// Http* helpers in this package, eg an external script or SDK
func FreshToken(mat lib.AzureMultiAuthToken) lib.AzureMultiAuthToken {
	mat.TokenData = FreshTokenData(mat.TokenData)
	return mat
}

// Sends req authorised with tokenData, first swapping in a new token if the
// one provided has expired or is about to.
//
// If the API still responds 401, a new token is requested and the request is
// retried once.
func doRequestWithTokenRefresh(req *http.Request, tokenData lib.AzureTokenData) (*http.Response, error) {
	current := FreshTokenData(tokenData)
	req.Header.Set("Authorization", "Bearer "+current.Token)

	res, err := http.DefaultClient.Do(req)
	if err != nil || res.StatusCode != http.StatusUnauthorized {
		return res, err
	}

	renewed, ok := resolveToken(current.Token, true)
	if !ok || renewed.Token == current.Token {
		return res, nil
	}

	retry := req.Clone(req.Context())
	if req.Body != nil && req.Body != http.NoBody {
		if req.GetBody == nil {
			return res, nil
		}
		body, bodyErr := req.GetBody()
		if bodyErr != nil {
			return res, nil
		}
		retry.Body = body
	}
	retry.Header.Set("Authorization", "Bearer "+renewed.Token)

	io.Copy(io.Discard, res.Body)
	res.Body.Close()

	return http.DefaultClient.Do(retry)
}
