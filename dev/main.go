package main

import (
	"fmt"
	"time"

	"github.com/jercle/cloudini/cmd/azure"
	"github.com/jercle/cloudini/lib"
)

func main() {
	var (
		subscriptionId    = "fdeee0c2-5569-40ea-9ad9-81dd325f6e1e"
		resourceGroupName = "rg-apcdtqdesktop-aib"
		galleryName       = "sigapcdtqdesktopaibimages"
		galleryImageName  = "imgdef-specialised-win10-multi-session-gen2"
	)
	startTime := time.Now()
	config := lib.GetCldConfig(nil)
	_ = config
	// tokens, err := azure.GetAllTenantSPTokens(lib.MultiAuthTokenRequestOptions{})
	// lib.CheckFatalError(err)
	token, err := azure.GetTenantSPToken(lib.MultiAuthTokenRequestOptions{})
	lib.CheckFatalError(err)
	// _ = tokens
	_ = token

	images := azure.GetSIGImageVersions(subscriptionId, resourceGroupName, galleryName, galleryImageName, token)

	_ = images

	latest := images.Latest()

	fmt.Println(latest)

	elapsed := time.Since(startTime)
	_ = elapsed
	// fmt.Println(elapsed)
}
