package lib

import "time"

type ForgerockGetTokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	Scope       string `json:"scope"`
	TokenType   string `json:"token_type"`
}

type ForgerockToken struct {
	AccessToken string    `json:"accessToken"`
	ExpiresAt   time.Time `json:"expiresAt"`
	Scope       string    `json:"scope"`
	TokenType   string    `json:"tokenType"`
	Tenant      string    `json:"tenant"`
}
