// Get Azure ACRs

package main

import (
	"fmt"
	"time"

	"github.com/jercle/azg/cmd/azure"
	"github.com/jercle/azg/lib"
)

func main() {
	startTime := time.Now()
	// tenantId := os.Getenv("AZURE_TENANT_ID")
	// subscriptionId := os.Getenv("AZURE_SUBSCRIPTION_ID")
	// clientId := os.Getenv("AZURE_CLIENT_ID")
	// clientSecret := os.Getenv("AZURE_CLIENT_SECRET")
	// token, err := azure.GetServicePrincipalToken(tenantId, lib.CldConfigClientAuthDetails{ClientID: clientId, ClientSecret: clientSecret})
	// lib.CheckFatalError(err)
	// _ = subscriptionId

	// tokens, err := azure.GetAllTenantSPTokens(azure.AzureRequestOptions{})
	// lib.CheckFatalError(err)

	// fmt.Print(tokens)

	// config := lib.GetCldConfig(lib.CldConfigOptions{})
	// config := lib.InitConfig(lib.CldConfigOptions{})
	// fmt.Println(config)

	tokens, err := azure.GetAllTenantSPTokens(azure.MultiAuthTokenRequestOptions{})
	lib.CheckFatalError(err)

	fmt.Println(tokens)

	// fmt.Println(listSubscriptionACRs(tenantId, subscriptionId, token))
	// fmt.Println(listAllTenantACRsBySub("REDDTQ", tenantId, token))
	elapsed := time.Since(startTime)
	_ = elapsed
	// fmt.Println(elapsed)
}
