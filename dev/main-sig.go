package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/jercle/cloudini/cmd/azure"
	"github.com/jercle/cloudini/lib"
	"golang.org/x/mod/semver"
)

func main() {
	startTime := time.Now()
	config := lib.GetCldConfig(nil)
	_ = config

	var (
		subscriptionId   = "fdeee0c2-5569-40ea-9ad9-81dd325f6e1e"
		resourceGroup    = "rg-apcdtqdesktop-aib"
		galleryName      = "sigapcdtqdesktopaibimages"
		galleryImageName = "imgdef-base-winupdrun"
	)
	// tokens, err := azure.GetAllTenantSPTokens(lib.MultiAuthTokenRequestOptions{})
	// lib.CheckFatalError(err)
	// _ = tokens
	token, err := azure.GetTenantSPToken(lib.MultiAuthTokenRequestOptions{
		TenantName: "REDDTQ",
	})
	lib.CheckFatalError(err)
	_ = token

	versions := GetGalleryImageVersions(subscriptionId, resourceGroup, galleryName, galleryImageName, *token)

	_, latest := versions.Latest()

	fmt.Println(latest)

	elapsed := time.Since(startTime)
	_ = elapsed
	// fmt.Println(elapsed)
}

func GetGalleryImageVersions(subscriptionId string, resourceGroup string, galleryName string, galleryImageName string, mat lib.MultiAuthToken) GalleryImageVersionList {

	var (
		listResponse  ListGalleryImageVersionsResponse
		imageVersions GalleryImageVersionList
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

	res, err := azure.HttpGet(urlString, mat)
	lib.CheckFatalError(err)

	json.Unmarshal(res, &listResponse)

	for _, val := range listResponse.Value {
		str, _ := json.Marshal(val)

		var imgVer GalleryImageVersion

		json.Unmarshal(str, &imgVer)
		imageVersions = append(imageVersions, imgVer)
	}

	// fmt.Println(string(res))
	return imageVersions
}

func (list *GalleryImageVersionList) Latest() (GalleryImageVersion, string) {
	latestVersion := GalleryImageVersion{}

	for _, version := range *list {
		currentVersion := ""

		if !strings.HasPrefix(version.Name, "v") {
			currentVersion = "v" + version.Name
		} else {
			currentVersion = version.Name
		}

		if semver.Compare(currentVersion, latestVersion.Name) == 1 {
			latestVersion = version
		}
	}

	return latestVersion, latestVersion.Name
}
