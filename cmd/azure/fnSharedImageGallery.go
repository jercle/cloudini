package azure

import (
	"encoding/json"

	"github.com/jercle/cloudini/lib"
)

func GetSIGImageVersions(subscriptionId string, resourceGroupName string, galleryName string, galleryImageName string, token *lib.MultiAuthToken) lib.SIGImageVersionList {
	var (
		listVersionsResponse lib.ListSIGImageVersionsResponse
		imageVersions        lib.SIGImageVersionList
	)

	urlString := "https://management.azure.com/subscriptions/" +
		subscriptionId +
		"/resourceGroups/" +
		resourceGroupName +
		"/providers/Microsoft.Compute/galleries/" +
		galleryName +
		"/images/" +
		galleryImageName +
		// "?api-version=2023-07-03"
		"/versions?api-version=2023-07-03"

	res, err := HttpGet(urlString, *token)
	lib.CheckFatalError(err)

	json.Unmarshal(res, &listVersionsResponse)

	for _, img := range listVersionsResponse.Value {

		var image lib.SIGImageVersion
		jsonBytes, _ := json.Marshal(img)
		err := json.Unmarshal(jsonBytes, &image)
		lib.CheckFatalError(err)

		imageVersions = append(imageVersions, image)
	}

	return imageVersions
}
