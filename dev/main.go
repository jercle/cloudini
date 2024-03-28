package main

import (
	"encoding/json"
	"fmt"

	"github.com/jercle/azg/cmd/azure"
	"github.com/jercle/azg/lib"
)

func main() {
	token, err := azure.GetTenantSPToken("REDDTQ", azure.MultiAuthTokenRequestOptions{Scope: "graph"})
	lib.CheckFatalError(err)
	// roleDefs, err :=
	defs, err := azure.ListEntraRoleDefinitions(*token)
	jsonStr, _ := json.MarshalIndent(defs, "", "  ")

	fmt.Println(string(jsonStr))

}
