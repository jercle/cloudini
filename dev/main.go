package main

import (
	"github.com/jercle/azg/cmd/azure"
	"github.com/jercle/azg/lib"
)

func main() {
	tokens, err := azure.GetAllTenantSPTokens(azure.MultiAuthTokenRequestOptions{})
	lib.CheckFatalError(err)
	redToken, _ := tokens.SelectTenant("REDDTQ")
	_ = redToken

}
