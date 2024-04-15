package main

import (
	"encoding/json"
	"os"
	"strconv"

	"github.com/jercle/azg/cmd/azure"
	"github.com/jercle/azg/lib"
)

type ListSubResGrpsResponse struct {
	Value []ResourceGroupListResponse `json:"value"`
}

type ResourceGroupListResponse struct {
	ID       string `json:"id"`
	Location string `json:"location"`
	Name     string `json:"name"`
	Tags     *struct {
		CreatedBy                      string `json:"createdBy,omitempty"`
		ImageTemplateName              string `json:"imageTemplateName,omitempty"`
		ImageTemplateResourceGroupName string `json:"imageTemplateResourceGroupName,omitempty"`
	} `json:"tags,omitempty"`
}

type ResourceGroup struct {
	ResourceGroupListResponse
	SubscriptionName string `json:"subscriptionName"`
	TenantName       string `json:"tenantName"`
}

type ListByResourceGroupResponse struct {
	Value []interface{} `json:"value"`
}

// List resources by resource group
// https://management.azure.com/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/resources?api-version=2021-04-01

func main() {
	tokens, err := azure.GetAllTenantSPTokens(lib.MultiAuthTokenRequestOptions{})
	lib.CheckFatalError(err)

	SaveAllResourcesToFile(tokens)

	// jsonStr, _ := json.MarshalIndent(emptyResourceGroups, "", "  ")
	// fmt.Println(string(jsonStr))
	// fmt.Println(len(emptyResourceGroups))

}

func ListAllEmptyResourceGroups(tokens lib.AllTenantTokens) []ResourceGroup {
	var (
		emptyResourceGroups []ResourceGroup
	)

	for _, token := range tokens {
		allSubs, _ := azure.ListSubscriptions(token)
		for _, sub := range allSubs {
			var subResourceGroups ListSubResGrpsResponse
			urlString := "https://management.azure.com/subscriptions/" +
				sub.SubscriptionID +
				"/resourcegroups?api-version=2021-04-01"
			res, err := azure.HttpGet(urlString, token)
			lib.CheckFatalError(err)
			json.Unmarshal(res, &subResourceGroups)

			for _, resGrp := range subResourceGroups.Value {
				var resourceGroup ResourceGroup
				resGrpJson, err := json.Marshal(resGrp)
				json.Unmarshal(resGrpJson, &resourceGroup)
				resourceGroup.SubscriptionName = sub.DisplayName
				resourceGroup.TenantName = token.TenantName

				var resourceList ListByResourceGroupResponse

				urlString := "https://management.azure.com/subscriptions/" +
					sub.SubscriptionID +
					"/resourceGroups/" +
					resGrp.Name +
					"/resources?api-version=2021-04-01"

				res, err := azure.HttpGet(urlString, token)
				lib.CheckFatalError(err)
				json.Unmarshal(res, &resourceList)

				if len(resourceList.Value) == 0 {
					emptyResourceGroups = append(emptyResourceGroups, resourceGroup)
				}
			}
		}
	}

	return emptyResourceGroups
}

func SaveAllResourcesToFile(tokens lib.AllTenantTokens) {
	for _, token := range tokens {
		allSubs, _ := azure.ListSubscriptions(token)
		for _, sub := range allSubs {
			var (
				subResourceGroups ListSubResGrpsResponse
				subResources      []interface{}
			)
			urlString := "https://management.azure.com/subscriptions/" +
				sub.SubscriptionID +
				"/resourcegroups?api-version=2021-04-01"
			res, err := azure.HttpGet(urlString, token)
			lib.CheckFatalError(err)

			json.Unmarshal(res, &subResourceGroups)

			for _, resGrp := range subResourceGroups.Value {
				var resourceGroup ResourceGroup
				resGrpJson, err := json.Marshal(resGrp)
				json.Unmarshal(resGrpJson, &resourceGroup)
				resourceGroup.SubscriptionName = sub.DisplayName
				resourceGroup.TenantName = token.TenantName

				var resourceList ListByResourceGroupResponse

				urlString := "https://management.azure.com/subscriptions/" +
					sub.SubscriptionID +
					"/resourceGroups/" +
					resGrp.Name +
					"/resources?api-version=2021-04-01"

				res, err := azure.HttpGet(urlString, token)
				lib.CheckFatalError(err)
				json.Unmarshal(res, &resourceList)

				subResources = append(subResources, resourceList.Value...)

			}

			baseDir := "outputs/resourceLists/"

			if _, err := os.Stat(baseDir + token.TenantName); err != nil {
				os.MkdirAll(baseDir+token.TenantName, os.ModePerm)
			}

			jsonBytes, _ := json.MarshalIndent(subResources, "", "  ")

			err = os.WriteFile(baseDir+
				token.TenantName+
				"/"+
				sub.DisplayName+
				"-"+
				strconv.Itoa(len(subResources))+
				".json", jsonBytes, os.ModePerm)
			lib.CheckFatalError(err)
		}
	}
}
