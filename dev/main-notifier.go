package main

import (
	"fmt"
	"time"

	"github.com/jercle/cloudini/cmd/azure"
	"github.com/jercle/cloudini/lib"
)

func main() {
	startTime := time.Now()
	config := lib.GetCldConfig(nil)
	_ = config
	// tokens, err := azure.GetAllTenantSPTokens(lib.MultiAuthTokenRequestOptions{})
	// lib.CheckFatalError(err)
	// token, err := azure.GetTenantSPToken(lib.MultiAuthTokenRequestOptions{})
	// tokenReq, err := azure.GetAllTenantSPTokens(lib.MultiAuthTokenRequestOptions{}, nil)
	// lib.CheckFatalError(err)
	// token, err := tokenReq.SelectTenant("REDDTQ")

	token := azure.GetAzCliToken()
	// lib.CheckFatalError(err)
	// _ = tokens
	_ = token

	// fmt.Println(token)
	// fmt.Println(tok)
	// os.Exit(0)

	urlString := "https://graph.microsoft.com/v1.0/subscriptions"
	body := `{
   "changeType": "created",
   "notificationUrl": "",
   "resource": "me/mailFolders('Inbox')/messages",
   "expirationDateTime":"2016-11-20T18:23:45.9356913Z",
   "clientState": "secretClientValue",
   "latestSupportedTlsVersion": "v1_2"
}`

	// res, err := azure.HttpGet(urlString, token)
	res, resHeader, err := azure.HttpPost(urlString, body, token)
	lib.CheckFatalError(err)

	fmt.Println(string(res))
	fmt.Println(string(resHeader))

	elapsed := time.Since(startTime)
	_ = elapsed
	// fmt.Println(elapsed)
}
