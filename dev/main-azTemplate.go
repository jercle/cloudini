package main

import (
	"time"

	"github.com/jercle/azg/cmd/azure"
	"github.com/jercle/azg/lib"
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
