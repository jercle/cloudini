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
	// tokens, err := azure.GetAllTenantSPTokens(lib.MultiAuthTokenRequestOptions{})
	// lib.CheckFatalError(err)
	token, err := azure.GetTenantSPToken(lib.MultiAuthTokenRequestOptions{})
	lib.CheckFatalError(err)
	// _ = tokens
	_ = token

	elapsed := time.Since(startTime)
	_ = elapsed
	// fmt.Println(elapsed)
}
