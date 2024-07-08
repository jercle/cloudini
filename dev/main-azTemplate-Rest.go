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
	tokenReq, err := azure.GetAllTenantSPTokens(lib.MultiAuthTokenRequestOptions{})
	lib.CheckFatalError(err)
	token, err := tokenReq.SelectTenant("REDDTQ")
	lib.CheckFatalError(err)
	// _ = tokens
	_ = token

	urlString := ""
	res, err := azure.HttpGet(urlString, *token)
	lib.CheckFatalError(err)

	fmt.Println(string(res))

	elapsed := time.Since(startTime)
	_ = elapsed
	// fmt.Println(elapsed)
}
