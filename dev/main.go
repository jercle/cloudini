package main

import (
	"encoding/json"
	"fmt"
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
	// lib.CheckFatalError(err)
	// _ = tokens
	// _ = token

	path := "/home/jercle/git/evan-tooling/terraform/apc/dtq/subs/apcdtqshared/out.json"

	byteValue, err := os.ReadFile(path)
	lib.CheckFatalError(err)

	var rules []AzureNSGRule

	err = json.Unmarshal(byteValue, &rules)
	// lib.CheckFatalError(err)
	for _, rule := range rules {
		// fmt.Println(rule.Name)
		importString := "terraform import 'azurerm_network_security_rule.[\"" +
			rule.Name +
			"\"]' '/subscriptions/SUBID/resourceGroups/RES_GRP/providers/Microsoft.Network/networkSecurityGroups/NSG_NAME/securityRules/" +
			rule.Name +
			"'"
		fmt.Println(importString)
	}

	elapsed := time.Since(startTime)
	_ = elapsed
	// fmt.Println(elapsed)
}

type AzureNSGRule struct {
	Access                                 string   `json:"access"`
	Description                            *string  `json:"description"`
	DestinationAddressPrefix               string   `json:"destination_address_prefix"`
	DestinationAddressPrefixes             []any    `json:"destination_address_prefixes"`
	DestinationApplicationSecurityGroupIds []any    `json:"destination_application_security_group_ids"`
	DestinationPortRange                   *string  `json:"destination_port_range"`
	DestinationPortRanges                  []string `json:"destination_port_ranges"`
	Direction                              string   `json:"direction"`
	Name                                   string   `json:"name"`
	Priority                               float64  `json:"priority"`
	Protocol                               string   `json:"protocol"`
	SourceAddressPrefix                    *string  `json:"source_address_prefix"`
	SourceAddressPrefixes                  []string `json:"source_address_prefixes"`
	SourceApplicationSecurityGroupIds      []any    `json:"source_application_security_group_ids"`
	SourcePortRange                        string   `json:"source_port_range"`
	SourcePortRanges                       []any    `json:"source_port_ranges"`
}
