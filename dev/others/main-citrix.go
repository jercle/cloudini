// Citrix Cloud stuff

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/jercle/azg/lib"
)

type TokenRequestResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   string `json:"expires_in"`
	TokenType   string `json:"token_type"`
}

type CitrixResourceLocationResponse struct {
	Items []CitrixResourceLocationResponseItem `json:"items"`
}

type CitrixResourceLocationResponseItem struct {
	ID           string `json:"id"`
	InternalOnly bool   `json:"internalOnly"`
	Name         string `json:"name"`
	ReadOnly     bool   `json:"readOnly"`
	TimeZone     string `json:"timeZone"`
}

func main() {
	token := GetCitrixCloudToken()

}

func GetCitrixCloudResourceLocations(token TokenRequestResponse) {
	customerId := os.Getenv("CUSTOMER_ID")
	urlString := "https://registry.citrixworkspacesapi.net/" + customerId + "/resourcelocations"

	req, err := http.NewRequest(http.MethodGet, urlString, nil)

	lib.CheckFatalError(err)
	req.Header.Add("Authorization", "CwsAuth Bearer="+token.AccessToken)

	res, err := http.DefaultClient.Do(req)
	lib.CheckFatalError(err)

	responseBody, err := io.ReadAll(res.Body)
	lib.CheckFatalError(err)

	fmt.Println(string(responseBody))
}

func GetCitrixCloudToken() TokenRequestResponse {
	tokenUrl := "https://api-ap-s.cloud.com/cctrustoauth2/root/tokens/clients"

	tokenReqBody := url.Values{}
	tokenReqBody.Set("grant_type", "client_credentials")
	tokenReqBody.Set("client_id", os.Getenv("CLIENT_ID"))
	tokenReqBody.Set("client_secret", os.Getenv("CLIENT_SECRET"))

	req, err := http.NewRequest(http.MethodPost, tokenUrl, strings.NewReader(tokenReqBody.Encode()))

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	lib.CheckFatalError(err)

	res, err := http.DefaultClient.Do(req)
	lib.CheckFatalError(err)

	responseBody, err := io.ReadAll(res.Body)
	lib.CheckFatalError(err)

	var tokenResp TokenRequestResponse

	json.Unmarshal(responseBody, &tokenResp)

	return tokenResp
}
