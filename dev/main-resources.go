package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/jercle/cloudini/cmd/azure"
	"github.com/jercle/cloudini/lib"
)

func main() {
	config := lib.GetCldConfig(nil)
	_ = config
	// tokens, err := azure.GetAllTenantSPTokens(lib.MultiAuthTokenRequestOptions{})
	// lib.CheckFatalError(err)
	// token, err := azure.GetTenantSPToken(lib.MultiAuthTokenRequestOptions{})
	tokenReq, err := azure.GetAllTenantSPTokens(lib.MultiAuthTokenRequestOptions{
		// Scope:         "graph",
		GetWriteToken: true,
	}, nil)
	lib.CheckFatalError(err)
	_ = tokenReq

	subscriptionId := ""
	resourceGroupName := ""
	token, err := tokenReq.SelectTenant("RED")
	lib.CheckFatalError(err)
	// _ = tokens
	// _ = token

	allResources := ListAllResGrpResources(subscriptionId, resourceGroupName, token)
	_ = allResources

	// DeleteResources(allResources, token)

	// var allResourcesDetailed []interface{}

	// jsonStr, _ := json.MarshalIndent(allResources, "", "  ")
	// fmt.Println(string(jsonStr))

	// for _, res := range allResources {
	// 	// jsonStr, _ := json.MarshalIndent(res, "", "  ")
	// 	// fmt.Println(string(jsonStr))
	// 	resDetails := GetResourceDetails(res.ID, token)

	// 	allResourcesDetailed = append(allResourcesDetailed, resDetails)

	// 	// jsonStr, _ := json.MarshalIndent(allResourcesDetailed, "", "  ")
	// 	// fmt.Println(string(jsonStr))

	// 	// os.Exit(0)
	// }

	// jsonStr, _ := json.MarshalIndent(allResourcesDetailed, "", "  ")
	// fmt.Println(string(jsonStr))

	// os.Exit(0)
	// allSubs := ListAllAuthenticatedSubscriptions(&tokenReq)

	// SaveAllResourcesTypesToFile("resourceLists", &tokenReq)

	// jsonStr, _ := json.MarshalIndent(allResources, "", "  ")
	// fmt.Println(string(jsonStr))
	// SaveAllResourcesToFile("resourceLists", &tokenReq, false)

	// resourceListPath := "resourceLists"
	// ConvertAzureResourcesToTF(resourceListPath, &tokenReq)
}

func TestResourceGraph(token *lib.MultiAuthToken) {

	// urlString := "https://management.azure.com/providers/Microsoft.ResourceGraph/resources?api-version=2022-10-01"

	// jsonStr := ""
	// jsonBody := []byte(`{"ids": ["e9f4bce2-7308-461a-91ce-3213f50f54f1"]}`)
	// res, _, err := azure.HttpPost(urlString, bodyReader, *token)
}

func GetResourceDetails(resourceId string, token *lib.MultiAuthToken) interface{} {
	var resourceDetails interface{}
	_ = resourceDetails

	urlString := "https://management.azure.com" + resourceId + "?api-version=2021-04-01"
	// fmt.Println(urlString)
	// os.Exit(0)
	res, err := azure.HttpGet(urlString, *token)
	lib.CheckFatalError(err)

	err = json.Unmarshal(res, &resourceDetails)

	// fmt.Println(string(res))
	return resourceDetails
}

func DeleteResources(resourceList []azure.ListRspResource, token *lib.MultiAuthToken) {
	var allResourceData []azure.ListRspResource
	_ = allResourceData
	for _, resource := range resourceList {
		fmt.Println(resource.Name)
	}
}

type ResourceDetailed struct {
	Etag       string `json:"etag"`
	ID         string `json:"id"`
	Location   string `json:"location"`
	ManagedBy  string `json:"managedBy,omitempty"`
	Name       string `json:"name"`
	Properties struct {
		CustomDnsConfigs []any `json:"customDnsConfigs"`
		DnsSettings      *struct {
			AppliedDnsServers        []any  `json:"appliedDnsServers"`
			DnsServers               []any  `json:"dnsServers"`
			InternalDomainNameSuffix string `json:"internalDomainNameSuffix"`
		} `json:"dnsSettings,omitempty"`
		EnableAcceleratedNetworking bool  `json:"enableAcceleratedNetworking"`
		EnableIpForwarding          bool  `json:"enableIPForwarding"`
		HostedWorkloads             []any `json:"hostedWorkloads"`
		IpConfigurations            []struct {
			Etag       string `json:"etag"`
			ID         string `json:"id"`
			Name       string `json:"name"`
			Properties struct {
				Primary                         bool   `json:"primary"`
				PrivateIpAddress                string `json:"privateIPAddress"`
				PrivateIpAddressVersion         string `json:"privateIPAddressVersion"`
				PrivateIpAllocationMethod       string `json:"privateIPAllocationMethod"`
				PrivateLinkConnectionProperties struct {
					Fqdns              []any  `json:"fqdns"`
					GroupID            string `json:"groupId"`
					RequiredMemberName string `json:"requiredMemberName"`
				} `json:"privateLinkConnectionProperties"`
				ProvisioningState string `json:"provisioningState"`
				Subnet            struct {
					ID string `json:"id"`
				} `json:"subnet"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"ipConfigurations"`
		MacAddress                          string `json:"macAddress"`
		ManualPrivateLinkServiceConnections []any  `json:"manualPrivateLinkServiceConnections"`
		NetworkInterfaces                   []struct {
			ID string `json:"id"`
		} `json:"networkInterfaces,omitempty"`
		NicType         string `json:"nicType,omitempty"`
		PrivateEndpoint *struct {
			ID string `json:"id"`
		} `json:"privateEndpoint,omitempty"`
		PrivateLinkServiceConnections []struct {
			Etag       string `json:"etag"`
			ID         string `json:"id"`
			Name       string `json:"name"`
			Properties struct {
				GroupIds                          []string `json:"groupIds"`
				PrivateLinkServiceConnectionState struct {
					ActionsRequired string `json:"actionsRequired"`
					Description     string `json:"description"`
					Status          string `json:"status"`
				} `json:"privateLinkServiceConnectionState"`
				PrivateLinkServiceID string `json:"privateLinkServiceId"`
				ProvisioningState    string `json:"provisioningState"`
				RequestMessage       string `json:"requestMessage"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"privateLinkServiceConnections,omitempty"`
		ProvisioningState string `json:"provisioningState"`
		ResourceGuid      string `json:"resourceGuid"`
		Subnet            *struct {
			ID string `json:"id"`
		} `json:"subnet,omitempty"`
		TapConfigurations       []any `json:"tapConfigurations"`
		VnetEncryptionSupported bool  `json:"vnetEncryptionSupported"`
	} `json:"properties"`
	Tags struct{} `json:"tags"`
	Type string   `json:"type"`
}

func ListAllResGrpResources(subscriptionId string, resourceGroupName string, token *lib.MultiAuthToken) []azure.ListRspResource {
	var (
		unmarshRes   azure.ListAllResourcesResponse
		allResources []azure.ListRspResource
		nextLink     string
	)

	urlString := "https://management.azure.com/subscriptions/" +
		subscriptionId +
		"/resourceGroups/" +
		resourceGroupName +
		"/resources?api-version=2021-04-01&$expand=createdTime,managedBy,kind,properties&$select=name,managedBy,createdTime,kind,properties"

	res, err := azure.HttpGet(urlString, *token)
	lib.CheckFatalError(err)

	// jsonStr, _ := json.MarshalIndent(res, "", "  ")
	fmt.Println(string(res))

	// fmt.Println(string(res))
	json.Unmarshal(res, &unmarshRes)
	allResources = append(allResources, unmarshRes.Value...)

	nextLink = unmarshRes.NextLink

	for nextLink != "" {
		var currentSet azure.ListAllResourcesResponse
		// fmt.Println("Getting next set")
		res, _ := azure.HttpGet(nextLink, *token)
		// fmt.Println(string(res))
		json.Unmarshal(res, &currentSet)
		nextLink = currentSet.NextLink
		// fmt.Println(nextLink)
		allResources = append(allResources, currentSet.Value...)
	}

	return allResources
}

func ConvertAzureResourcesToTF(fileDirectory string, tokens *lib.AllTenantTokens) {
	allSubs := ListAllAuthenticatedSubscriptions(tokens)
	resourceTypeMap := make(map[string]string)
	resTypeMap, err := os.ReadFile("azureResourceTypeMapping.json")
	lib.CheckFatalError(err)
	json.Unmarshal(resTypeMap, &resourceTypeMap)

	jsonStr, _ := json.MarshalIndent(resourceTypeMap, "", "  ")
	fmt.Println(string(jsonStr))
	os.Exit(0)

	for tenant, subs := range allSubs {
		for subName, _ := range subs {
			var subResources []azure.ListRspResource
			file, err := os.ReadFile(fileDirectory + "/" + tenant + "-" + subName + ".json")
			lib.CheckFatalError(err)

			json.Unmarshal(file, &subResources)

			jsonStr, _ := json.MarshalIndent(subResources, "", "  ")
			fmt.Println(string(jsonStr))
			os.Exit(0)
		}
	}
}

// Gets all configured tenants, and their subscriptions, then saves json
// of all resources in each subscription to the given saveDirectory parameter

// Example: SAVEDIR/TENANTNAME-SUBNAME-RESOURCECOUNT
func SaveAllResourcesToFile(saveDirectory string, tokens *lib.AllTenantTokens, includeCountsInFilename bool) {
	allSubs := ListAllAuthenticatedSubscriptions(tokens)

	for tenant, subs := range allSubs {
		token, err := tokens.SelectTenant(tenant)
		lib.CheckFatalError(err)
		for subName, subId := range subs {
			subResources := ListAllSubscriptionResources(subName, subId, token)
			var fileName string
			if includeCountsInFilename {
				fileName = tenant +
					"-" +
					subName +
					"-" +
					strconv.Itoa(len(subResources))
			} else {
				fileName = tenant +
					"-" +
					subName
			}

			jsonStr, _ := json.MarshalIndent(subResources, "", "  ")

			SaveToFile(jsonStr, saveDirectory, fileName)
		}
	}
}

func SaveAllResourcesTypesToFile(saveDirectory string, tokens *lib.AllTenantTokens) {
	allSubs := ListAllAuthenticatedSubscriptions(tokens)
	resourceTypes := make(map[string]string)
	fileName := "resourceTypes"

	for tenant, subs := range allSubs {
		token, err := tokens.SelectTenant(tenant)
		lib.CheckFatalError(err)
		for subName, subId := range subs {
			subResources := ListAllSubscriptionResources(subName, subId, token)
			for _, resource := range subResources {
				resourceTypes[resource.Type] = ""
			}
		}
	}

	jsonStr, _ := json.MarshalIndent(resourceTypes, "", "  ")
	SaveToFile(jsonStr, saveDirectory, fileName)
}

/*
Lists all subscriptions available with provided auth token

Example:

	{
		"TENANTNAME": {
			"SUBNAME": [
				{RESOURCE}...
			]
		}
	}
*/
func ListAllSubscriptionResources(subName string, subscriptionId string, token *lib.MultiAuthToken) []azure.ListRspResource {
	var (
		unmarshRes      azure.ListAllResourcesResponse
		allSubResources []azure.ListRspResource
		nextLink        string
	)

	urlString := "https://management.azure.com/subscriptions/" +
		subscriptionId +
		"/resources?api-version=2021-04-01"

	res, err := azure.HttpGet(urlString, *token)
	lib.CheckFatalError(err)

	// fmt.Println(string(res))
	json.Unmarshal(res, &unmarshRes)
	allSubResources = append(allSubResources, unmarshRes.Value...)

	nextLink = unmarshRes.NextLink

	for nextLink != "" {
		var currentSet azure.ListAllResourcesResponse
		// fmt.Println("Getting next set")
		res, _ := azure.HttpGet(nextLink, *token)
		// fmt.Println(string(res))
		json.Unmarshal(res, &currentSet)
		nextLink = currentSet.NextLink
		// fmt.Println(nextLink)
		allSubResources = append(allSubResources, currentSet.Value...)
	}

	return allSubResources
}

/*
Returns a json string of subscriptions available to given tokens

Example:

	{
		"SUBNAME": "SUBID"
	}
*/
func ListAllAuthenticatedSubscriptions(tokens *lib.AllTenantTokens) TenantList {
	// allSubscriptions := make(map[string]string)
	allTenantSubs := TenantList{}

	for _, token := range *tokens {
		subs, err := azure.ListSubscriptions(token)
		lib.CheckFatalError(err)
		allTenantSubs[token.TenantName] = TenantSubscriptionList{}

		for _, sub := range subs {
			allTenantSubs[token.TenantName][sub.DisplayName] = sub.SubscriptionID
		}
	}
	return allTenantSubs
}

type TenantSubscriptionList map[string]string

type TenantList map[string]TenantSubscriptionList

// Saves []byte to file with given directory and fileName
//
// If directory does not exist, will be created.
func SaveToFile(data []byte, directory string, fileName string) {
	fullFilePath := directory + "/" + fileName + ".json"
	fmt.Println("Saving file: " + fullFilePath)

	if _, err := os.Stat(directory); err != nil {
		os.MkdirAll(directory, os.ModePerm)
	}

	err := os.WriteFile(fullFilePath, data, os.ModePerm)
	lib.CheckFatalError(err)
}
