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
	tokenReq, err := azure.GetAllTenantSPTokens(lib.MultiAuthTokenRequestOptions{}, &lib.CldConfigOptions{
		ConfigFilePath: "/home/jercle/.config/cld/cldConf.json",
	})
	lib.CheckFatalError(err)
	token, err := tokenReq.SelectTenant("RED")
	lib.CheckFatalError(err)
	// _ = tokens
	_ = token

	subscriptionId := ""

	urlString := "https://management.azure.com/subscriptions/" +
		subscriptionId +
		"/providers/Microsoft.Network/IpAllocations?api-version=2024-01-01"
		// "/providers/Microsoft.Network/publicIPAddresses?api-version=2024-01-01"

	res, err := azure.HttpGet(urlString, *token)

	fmt.Println(string(res))
	// jsonStr, _ := json.MarshalIndent(res, "", "  ")
	// fmt.Println(string(jsonStr))

	elapsed := time.Since(startTime)
	_ = elapsed
	// fmt.Println(elapsed)
}
