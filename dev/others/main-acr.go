// acr

package main

import (
	"fmt"
	"os"
	"time"

	"github.com/jercle/cloudini/cmd/azure"
	"github.com/jercle/cloudini/lib"
)

func main() {
	// var (
	// 	acrName  string
	// 	repoName string
	// )
	acrName := "acrapcdtqautomon"
	repoName := "ubuntu-packer"

	startTime := time.Now()
	config := lib.GetCldConfig(nil)
	_ = config
	// tokens, err := azure.GetAllTenantSPTokens(lib.MultiAuthTokenRequestOptions{})
	// lib.CheckFatalError(err)
	token, err := azure.GetTenantSPToken(lib.MultiAuthTokenRequestOptions{TenantName: "REDDTQ", AzureContainerRepositoryName: acrName, Scope: "acr"})
	lib.CheckFatalError(err)
	// _ = tokens
	_ = token

	fmt.Println(token)
	os.Exit(0)
	urlString := "https://" +
		acrName +
		".azurecr.io/acr/v1/" +
		repoName +
		"/_tags"
	res, err := azure.HttpGet(urlString, *token)
	lib.CheckFatalError(err)

	fmt.Println(string(res))

	elapsed := time.Since(startTime)
	_ = elapsed
	// fmt.Println(elapsed)
}
