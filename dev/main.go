package main

import (
	"time"

	"github.com/jercle/cloudini/cmd/azure"
	"github.com/jercle/cloudini/lib"
)

func main() {
	startTime := time.Now()
	config := lib.GetCldConfig(nil)
	_ = config

	var (
		subscriptionId   = ""
		resourceGroup    = "rg--aib"
		galleryName      = ""
		galleryImageName = "imgdef-"
	)
	// tokens, err := azure.GetAllTenantSPTokens(lib.MultiAuthTokenRequestOptions{})
	// lib.CheckFatalError(err)
	// _ = tokens
	token, err := azure.GetTenantSPToken(lib.MultiAuthTokenRequestOptions{
		TenantName: "REDDTQ",
	})
	lib.CheckFatalError(err)
	_ = token

	// versions := azure.GetGalleryImageVersions(subscriptionId, resourceGroup, galleryName, galleryImageName, *token)
	// _ = versions
	azure.GetGalleryImage(subscriptionId, resourceGroup, galleryName, galleryImageName, *token)
	// latest, _ := versions.Latest()

	// fmt.Println(latest.IncrementPatchVersion())

	elapsed := time.Since(startTime)
	_ = elapsed
	// fmt.Println(elapsed)
}
