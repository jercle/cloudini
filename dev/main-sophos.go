package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/jercle/cloudini/lib"
)

func main() {
	startTime := time.Now()
	config := lib.GetCldConfig(nil)
	_ = config
	// tokens, err := azure.GetAllTenantSPTokens(lib.MultiAuthTokenRequestOptions{})
	// lib.CheckFatalError(err)
	// token, err := azure.GetTenantSPToken(lib.MultiAuthTokenRequestOptions{})
	// tokenReq, err := azure.GetAllTenantSPTokens(lib.MultiAuthTokenRequestOptions{})
	// lib.CheckFatalError(err)
	// token, err := tokenReq.SelectTenant("REDDTQ")
	// lib.CheckFatalError(err)
	// _ = tokens
	// _ = token
	// sophosKey := config.SophosConfig.ApiKey
	// serverList := config.SophosConfig.ApiKey
	serverList := config.SophosConfig.Servers["REDDTQ"]
	domain := config.Domains["REDDTQ"]

	sophosPort := "4444"
	sophosUser := config.SophosConfig.ApiUser
	sophosKey := config.SophosConfig.ApiKey

	urlString := "https://" +
		serverList["001"] +
		"." +
		domain +
		":" +
		sophosPort +
		"/webconsole/APIController"

	xmlGetAllPolicies := `
<Request>
    <Login>
        <Username>` + sophosUser + `</Username>
        <Password>` + sophosKey + `</Password>
    </Login>
    <Get>
        <WebFilterPolicy>
        </WebFilterPolicy>
    </Get>
</Request>
`

	jsonStr, _ := json.MarshalIndent(serverList, "", "  ")
	fmt.Println(string(jsonStr))

	// urlString := ""

	res, resHeader, err := HttpPost(urlString, xmlGetAllPolicies, "application/xml", "")
	lib.CheckFatalError(err)

	fmt.Println(string(res))
	fmt.Println(string(resHeader))

	elapsed := time.Since(startTime)
	_ = elapsed
	// fmt.Println(elapsed)
}

func HttpPost(urlString string, body string, contentType string, authToken string) ([]byte, []byte, error) {
	var (
		reqContentType string
	)

	bodyReader := bytes.NewReader([]byte(body))
	// jsonBody := []byte(`{"ids": [e9f4bce2-7308-461a-91ce-3213f50f54f1"]}`)
	// res, _, err := azure.HttpPost(urlString, bodyReader, *token)
	req, err := http.NewRequest(http.MethodPost, urlString, bodyReader)
	lib.CheckFatalError(err)

	if contentType == "" {
		reqContentType = "application/json"
	} else {
		reqContentType = contentType
	}
	req.Header.Add("Content-Type", reqContentType)

	if authToken != "" {
		req.Header.Add("Authorization", authToken)
	}

	envProxy := os.Getenv("http_proxy")
	proxyURL, _ := url.Parse(envProxy)
	proxy := http.ProxyURL(proxyURL)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{},
		Proxy:           proxy,
	}
	client := &http.Client{
		Transport: tr,
	}

	// res, err := http.DefaultClient.Do(req)
	res, err := client.Do(req)
	lib.CheckFatalError(err)
	// if err != nil {
	// 	return nil, err
	// }

	responseBody, err := io.ReadAll(res.Body)

	// jsonStr, _ := json.MarshalIndent(res.Header, "", "  ")
	// fmt.Println(string(jsonStr))
	// fmt.Println(responseBody)
	lib.CheckFatalError(err)
	if res.StatusCode == 404 {
		fmt.Println(string(responseBody))
		lib.CheckFatalError(fmt.Errorf(res.Status))
	}

	resHeader, _ := json.MarshalIndent(res.Header, "", "  ")

	// fmt.Println()
	// if err != nil {
	// 	return nil, err
	// }

	return responseBody, resHeader, nil
}
