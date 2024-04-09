package main

import (
	"fmt"
	"time"

	"github.com/jercle/azg/cmd/azure"
	"github.com/jercle/azg/lib"
)

func main() {
	startTime := time.Now()
	config := lib.GetCldConfig(nil)
	_ = config
	// tokens, err := azure.GetAllTenantSPTokens(lib.MultiAuthTokenRequestOptions{})
	// lib.CheckFatalError(err)
	token, err := azure.GetTenantSPToken(lib.MultiAuthTokenRequestOptions{})
	lib.CheckFatalError(err)
	// _ = tokens
	_ = token
	firewallPolicyName := ""
	// azureFirewallName := ""
	subscriptionId := ""
	resourceGroupName := ""
	fwpRuleCollectionName := ""
	urlString := "https://management.azure.com/subscriptions/" + subscriptionId + "/resourceGroups/" + resourceGroupName + "/providers/Microsoft.Network/firewallPolicies/" + firewallPolicyName + "/ruleCollectionGroups/" + fwpRuleCollectionName + "?api-version=2023-09-01"
	// urlString := "https://management.azure.com/subscriptions/" + subscriptionId + "/resourceGroups/" + resourceGroupName + "/providers/Microsoft.Network/azureFirewalls/" + azureFirewallName + "?api-version=2023-09-01"

	resp, err := azure.HttpGet(urlString, *token)
	lib.CheckFatalError(err)

	fmt.Println(string(resp))

	elapsed := time.Since(startTime)
	_ = elapsed
	// fmt.Println(elapsed)
}
