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

type ListGalleryImageVersionsResponse struct {
	Value []GalleryImageVersionResponse `json:"value"`
}
type GalleryImageVersionResponse struct {
	ID         string `json:"id"`
	Location   string `json:"location"`
	Name       string `json:"name"`
	Properties struct {
		ProvisioningState string `json:"provisioningState"`
		PublishingProfile struct {
			ExcludeFromLatest  bool      `json:"excludeFromLatest"`
			PublishedDate      time.Time `json:"publishedDate"`
			ReplicaCount       float64   `json:"replicaCount"`
			ReplicationMode    string    `json:"replicationMode"`
			StorageAccountType string    `json:"storageAccountType"`
			TargetRegions      []struct {
				Name                 string  `json:"name"`
				RegionalReplicaCount float64 `json:"regionalReplicaCount"`
				StorageAccountType   string  `json:"storageAccountType"`
			} `json:"targetRegions"`
		} `json:"publishingProfile"`
		SafetyProfile struct {
			AllowDeletionOfReplicatedLocations bool `json:"allowDeletionOfReplicatedLocations"`
			ReportedForPolicyViolation         bool `json:"reportedForPolicyViolation"`
		} `json:"safetyProfile"`
		StorageProfile struct {
			OSDiskImage struct {
				HostCaching string   `json:"hostCaching"`
				SizeInGb    float64  `json:"sizeInGB"`
				Source      struct{} `json:"source"`
			} `json:"osDiskImage"`
			Source struct {
				VirtualMachineID string `json:"virtualMachineId"`
			} `json:"source"`
		} `json:"storageProfile"`
	} `json:"properties"`
	Tags struct {
		CostGroup string `json:"cost_group"`
		Env       string `json:"env"`
		ManagedBy string `json:"managed_by"`
	} `json:"tags"`
	Type string `json:"type"`
}

type GalleryImageVersion struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Properties struct {
		ProvisioningState string `json:"provisioningState"`
		PublishingProfile struct {
			ExcludeFromLatest bool `json:"excludeFromLatest"`
		} `json:"publishingProfile"`
	} `json:"properties"`
}

type GalleryImageVersionList []GalleryImageVersion

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
