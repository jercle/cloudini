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

	var (
		subscriptionId   = "fdeee0c2-5569-40ea-9ad9-81dd325f6e1e"
		resourceGroup    = "rg-apcdtqdesktop-aib"
		galleryName      = "sigapcdtqdesktopaibimages"
		galleryImageName = "imgdef-apc_oftadtq"
	)
	// tokens, err := azure.GetAllTenantSPTokens(lib.MultiAuthTokenRequestOptions{})
	// lib.CheckFatalError(err)
	// _ = tokens
	token, err := azure.GetTenantSPToken(lib.MultiAuthTokenRequestOptions{
		TenantName: "REDDTQ",
	})
	lib.CheckFatalError(err)
	_ = token

	versions := azure.GetGalleryImageVersions(subscriptionId, resourceGroup, galleryName, galleryImageName, *token)
	// _ = versions
	// azure.GetGalleryImage(subscriptionId, resourceGroup, galleryName, galleryImageName, *token)
	// fmt.Println(latest)
	if len(versions) == 0 {
		fmt.Println("")
	} else {
		latest, _ := versions.Latest()
		fmt.Println(latest.IncrementPatchVersion())
	}

	elapsed := time.Since(startTime)
	_ = elapsed
	// fmt.Println(elapsed)
}
