package azure

import (
	"encoding/json"
	"log"

	"github.com/jercle/cloudini/lib"
)

func GetGalleryImageVersions(subscriptionId string, resourceGroup string, galleryName string, galleryImageName string, mat lib.MultiAuthToken) lib.GalleryImageVersionList {

	var (
		listResponse  lib.ListGalleryImageVersionsResponse
		imageVersions lib.GalleryImageVersionList
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

	if len(listResponse.Value) == 0 {
		log.Fatalln("No versions exist")
	}

	for _, val := range listResponse.Value {
		str, _ := json.Marshal(val)

		var imgVer lib.GalleryImageVersion

		json.Unmarshal(str, &imgVer)
		imageVersions = append(imageVersions, imgVer)
	}

	// fmt.Println(string(res))
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
