// Get vm images
// https://learn.microsoft.com/en-us/rest/api/compute/virtual-machine-images/list?view=rest-compute-2024-03-01&tabs=HTTP

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

	subscriptionId := ""
	resourceGroupName := ""
	firewallPolicyName := ""
	ruleCollectionGroupName := ""

	urlString := "https://management.azure.com/subscriptions/" +
		subscriptionId +
		"/resourceGroups/" +
		resourceGroupName +
		"/providers/Microsoft.Network/firewallPolicies/" +
		firewallPolicyName +
		"/ruleCollectionGroups/" +
		ruleCollectionGroupName +
		"?api-version=2023-09-01"

	// "https://management.azure.com/subscriptions/" +
	// 	subscriptionId +
	// 	"/providers/Microsoft.Compute/locations/" +
	// 	location + "/publishers/" +
	// 	publisherName + "/artifacttypes/vmimage/offers/" +
	// 	offer + "/skus/" +
	// 	skus + "/versions?api-version=2024-03-01&$orderby=name"

	res, err := azure.HttpGet(urlString, *token)
	lib.CheckFatalError(err)

	fmt.Println(string(res))

	// var imageList ListVirtualMachineImagesResponse
	// json.Unmarshal(res, &imageList)

	// for _, image := range imageList {
	// 	fmt.Println(image.Name)
	// }

	elapsed := time.Since(startTime)
	_ = elapsed
	// fmt.Println(elapsed)
}

type ListVirtualMachineImagesResponse []VirtualMachineImage

type VirtualMachineImage struct {
	ID       string `json:"id"`
	Location string `json:"location"`
	Name     string `json:"name"`
}
