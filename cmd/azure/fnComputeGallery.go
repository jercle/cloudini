package azure

import (
	"encoding/json"

	"github.com/jercle/cloudini/lib"
)

func GetGalleryImageVersions(subscriptionId string, resourceGroup string, galleryName string, galleryImageName string, mat lib.MultiAuthToken) lib.GalleryImageVersionList {

	var (
		listResponse  lib.ListGalleryImageVersionsResponse
		imageVersions lib.GalleryImageVersionList
		nextLink      string
	)

	urlString := "https://management.azure.com/subscriptions/" +
		subscriptionId +
		"/resourceGroups/" +
		resourceGroup +
		"/providers/Microsoft.Compute/galleries/" +
		galleryName +
		"/images/" +
		galleryImageName +
		"/versions?api-version=2023-07-03"

	res, err := HttpGet(urlString, mat)
	lib.CheckFatalError(err)

	json.Unmarshal(res, &listResponse)

	for _, val := range listResponse.Value {
		str, _ := json.Marshal(val)

		var imgVer lib.GalleryImageVersion

		json.Unmarshal(str, &imgVer)
		imageVersions.Versions = append(imageVersions.Versions, imgVer)
	}

	nextLink = listResponse.NextLink

	for nextLink != "" {
		var currentSet lib.ListGalleryImageVersionsResponse

		res, _ := HttpGet(nextLink, mat)
		json.Unmarshal(res, &currentSet)
		nextLink = currentSet.NextLink

		for _, val := range currentSet.Value {
			str, _ := json.Marshal(val)

			var imgVer lib.GalleryImageVersion

			json.Unmarshal(str, &imgVer)
			imageVersions.Versions = append(imageVersions.Versions, imgVer)
		}
	}
	return imageVersions
}

func GetGalleryImage(subscriptionId string, resourceGroup string, galleryName string, galleryImageName string, mat lib.MultiAuthToken) lib.GalleryImageResponse {

	var (
		imageDefinition lib.GalleryImageResponse
	)

	urlString := "https://management.azure.com/subscriptions/" +
		subscriptionId +
		"/resourceGroups/" +
		resourceGroup +
		"/providers/Microsoft.Compute/galleries/" +
		galleryName +
		"/images/" +
		galleryImageName +
		"?api-version=2023-07-03"

	res, err := HttpGet(urlString, mat)
	lib.CheckFatalError(err)

	json.Unmarshal(res, &imageDefinition)

	// if len(listResponse.Value) == 0 {
	// 	log.Fatalln("No versions exist")
	// }

	// fmt.Println(string(res))
	return imageDefinition
}
