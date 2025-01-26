package lib

import "time"

type CitrixTokenData struct {
	AccessToken string    `json:"access_token"`
	ExpiresIn   string    `json:"expires_in"`
	TokenType   string    `json:"token_type"`
	CustomerId  string    `json:"customer_id"`
	SiteId      string    `json:"siteId"`
	Expiry      time.Time `json:"expiry,omitempty"`
}
