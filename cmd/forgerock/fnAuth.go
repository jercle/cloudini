package forgerock

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/jercle/cloudini/lib"
)

func GetTokenForConfiguredTenant(tenant string) (*lib.ForgerockToken, error) {
	tConfig := lib.GetCldConfig(nil).Forgerock.Domains[tenant]
	clientSecret, _ := base64.StdEncoding.DecodeString(tConfig.ClientSecretBase64)
	data := url.Values{}
	data.Set("client_id", tConfig.ClientID)
	data.Set("client_secret", string(clientSecret))
	data.Set("grant_type", "client_credentials")
	data.Set("scope", tConfig.AuthScope)

	urlString := tConfig.UrlBase + "/am/oauth2/realms/root/access_token"

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, urlString, strings.NewReader(data.Encode())) // URL-encoded payload
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	responseBody, err := io.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}
	if res.StatusCode == 404 {
		fmt.Println(string(responseBody))
		lib.CheckFatalError(fmt.Errorf(res.Status))
	}

	var response lib.ForgerockGetTokenResponse
	err = json.Unmarshal(responseBody, &response)
	if err != nil {
		return nil, err
	}

	currentTime := time.Now()
	tokenExpiration := currentTime.Add(time.Second * time.Duration(response.ExpiresIn))

	token := lib.ForgerockToken{
		AccessToken: response.AccessToken,
		TokenType:   response.TokenType,
		ExpiresAt:   tokenExpiration,
		Scope:       response.Scope,
		Tenant:      tenant,
	}

	return &token, nil
}
