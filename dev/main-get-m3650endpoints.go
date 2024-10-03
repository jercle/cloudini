// Get vm images
// https://learn.microsoft.com/en-us/rest/api/compute/virtual-machine-images/list?view=rest-compute-2024-03-01&tabs=HTTP

package main

import (
	"encoding/json"
	"fmt"
	"slices"
	"strings"
	"time"

	"github.com/jercle/cloudini/cmd/azure"
	"github.com/jercle/cloudini/lib"
)

type Microsoft365Endpoint struct {
	Category               string   `json:"category"`
	ExpressRoute           bool     `json:"expressRoute"`
	ID                     float64  `json:"id"`
	Ips                    []string `json:"ips,omitempty"`
	Notes                  string   `json:"notes,omitempty"`
	Required               bool     `json:"required"`
	ServiceArea            string   `json:"serviceArea"`
	ServiceAreaDisplayName string   `json:"serviceAreaDisplayName"`
	TcpPorts               string   `json:"tcpPorts,omitempty"`
	UdpPorts               string   `json:"udpPorts,omitempty"`
	Urls                   []string `json:"urls,omitempty"`
	Protocols              []string `json:"protocols,omitempty"`
}

func main() {
	startTime := time.Now()
	config := lib.GetCldConfig(nil)
	_ = config
	// tokens, err := azure.GetAllTenantSPTokens(lib.MultiAuthTokenRequestOptions{})
	// lib.CheckFatalError(err)
	// token, err := azure.GetTenantSPToken(lib.MultiAuthTokenRequestOptions{})
	tokenReq, err := azure.GetAllTenantSPTokens(lib.MultiAuthTokenRequestOptions{})
	lib.CheckFatalError(err)
	token, err := tokenReq.SelectTenant("")
	lib.CheckFatalError(err)
	// _ = tokens
	_ = token

	webserviceUrl := "https://endpoints.office.com"
	clientRequestId := "b10c5ed1-bad1-445f-b386-b919946339a7"
	tenantName := ""

	// Get the latest versioning data from the Office 365 IP Address and URL web service
	// urlString := webserviceUrl +
	// 	"/version/Worldwide?ClientRequestId=" +
	// 	clientRequestId

	// Query the Office 365 IP Address and URL web service for new data
	urlString := webserviceUrl + "/endpoints/Worldwide?NoIPv6=true&ClientRequestId=" +
		clientRequestId +
		"&TenantName=" +
		tenantName

	res, err := azure.HttpGet(urlString, *token)
	lib.CheckFatalError(err)

	var (
		endpointList       []Microsoft365Endpoint
		processedEndpoints []Microsoft365Endpoint
		networkEndpoints   []Microsoft365Endpoint
		appEndpoints       []Microsoft365Endpoint
	)

	// networkEndpoints := []Microsoft365Endpoint{}

	json.Unmarshal(res, &endpointList)

	for _, item := range endpointList {
		if item.TcpPorts != "" {
			item.Protocols = append(item.Protocols, "TCP")
		}

		if item.UdpPorts != "" {
			item.Protocols = append(item.Protocols, "UDP")
		}

		// fmt.Println(len(item.Ips))

		if len(item.Ips) != 0 {
			// fmt.Println(len(item.Ips))
			networkEndpoints = append(networkEndpoints, item)
			// jsonstr, _ := json.MarshalIndent(item, "", "  ")
			// fmt.Println(string(jsonstr))
		}

		if len(item.Urls) != 0 {
			// fmt.Println(len(item.Urls))
			appEndpoints = append(appEndpoints, item)
		}

		processedEndpoints = append(processedEndpoints, item)
	}

	// jsonStr, err := json.MarshalIndent(processedEndpoints, "", "  ")
	// lib.CheckFatalError(err)

	var fwprcAppM365Allow FirewallPolicyRuleCollectionGroup
	fwprcAppM365Allow.ID = "fwrulecoll-group-app-microsoft365-allow"
	fwprcAppM365Allow.Schema = "https://schema.management.azure.com/schemas/2019-04-01/deploymentParameters.json#"
	fwprcAppM365Allow.ContentVersion = "1.0.0.0"
	fwprcAppM365Allow.Parameters.Priority.Value = 32000

	var fwprcAppM365AllowRuleCollection RuleCollection
	fwprcAppM365AllowRuleCollection.Action.Type = "allow"
	fwprcAppM365AllowRuleCollection.RuleCollectionType = "FirewallPolicyFilterRuleCollection"
	fwprcAppM365AllowRuleCollection.Name = "fwrulecoll-app-microsoft365-allow"
	fwprcAppM365AllowRuleCollection.Priority = 32001

	for _, ep := range appEndpoints {
		var _ Microsoft365Endpoint
		// var rule FirewallPolicyRuleCollectionRule
		// json.Unmarshal(endpoint, &)
		if slices.Contains(strings.Split(ep.TcpPorts, ","), "80") {
			fmt.Println(ep)
		}

	}

	// fmt.Println(string(jsonStr))
	// fmt.Println(string(res))

	// var imageList ListVirtualMachineImagesResponse
	// json.Unmarshal(res, &imageList)

	// for _, image := range imageList {
	// 	fmt.Println(image.Name)
	// }

	elapsed := time.Since(startTime)
	_ = elapsed
	// fmt.Println(elapsed)
}

type FirewallPolicyRuleCollectionGroup struct {
	Attachments                 string  `json:"_attachments"`
	Etag                        string  `json:"_etag"`
	Rid                         string  `json:"_rid"`
	Self                        string  `json:"_self"`
	Ts                          float64 `json:"_ts"`
	ContentVersion              string  `json:"contentVersion"`
	ID                          string  `json:"id"`
	MsO365WebServiceDataVersion string  `json:"msO365WebServiceDataVersion"`
	Parameters                  struct {
		CuaID struct {
			Value string `json:"value"`
		} `json:"cuaId"`
		FirewallPolicyName struct {
			Value string `json:"value"`
		} `json:"firewallPolicyName"`
		Priority struct {
			Value float64 `json:"value"`
		} `json:"priority"`
		RuleCollectionGroupName struct {
			Value string `json:"value"`
		} `json:"ruleCollectionGroupName"`

		RuleCollections struct {
			Value []RuleCollection `json:"value"`
		} `json:"ruleCollections"`
	} `json:"parameters"`
	Schema string `json:"schema"`
}

type RuleCollection struct {
	Action struct {
		Type string `json:"type"`
	} `json:"action"`
	Name               string                             `json:"name"`
	Priority           int                                `json:"priority"`
	RuleCollectionType string                             `json:"ruleCollectionType"`
	Rules              []FirewallPolicyRuleCollectionRule `json:"rules"`
}

type FirewallPolicyRuleCollectionRule struct {
	DestinationAddresses []any                       `json:"destinationAddresses"`
	FqdnTags             []any                       `json:"fqdnTags"`
	Name                 string                      `json:"name"`
	Protocols            []FwPolRuleColRuleProtocols `json:"protocols"`
	RuleType             string                      `json:"ruleType"`
	SourceIpGroups       []string                    `json:"sourceIpGroups"`
	TargetFqdns          []string                    `json:"targetFqdns"`
	TargetUrls           []any                       `json:"targetUrls"`
	TerminateTls         bool                        `json:"terminateTLS"`
	WebCategories        []any                       `json:"webCategories"`
}

type FwPolRuleColRuleProtocols struct {
	Port         string `json:"port"`
	ProtocolType string `json:"protocolType"`
}
