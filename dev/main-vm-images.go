// Get vm images
// https://learn.microsoft.com/en-us/rest/api/compute/virtual-machine-images/list?view=rest-compute-2024-03-01&tabs=HTTP

package main

import (
	"encoding/json"
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
	token, err := azure.GetTenantSPToken(lib.MultiAuthTokenRequestOptions{})
	lib.CheckFatalError(err)
	// _ = tokens
	_ = token

	subscriptionId := ""
	location := "australiaeast"
	publisherName := "MicrosoftWindowsDesktop"
	offer := "office-365"
	skus := "win10-22h2-avd-m365-g2"

	urlString := "https://management.azure.com/subscriptions/" +
		subscriptionId +
		"/providers/Microsoft.Compute/locations/" +
		location + "/publishers/" +
		publisherName + "/artifacttypes/vmimage/offers/" +
		offer + "/skus/" +
		skus + "/versions?api-version=2024-03-01&$orderby=name"

	res, err := azure.HttpGet(urlString, *token)
	lib.CheckFatalError(err)

	// fmt.Println(string(res))

	var imageList ListVirtualMachineImagesResponse
	json.Unmarshal(res, &imageList)

	for _, image := range imageList {
		fmt.Println(image.Name)
	}

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
