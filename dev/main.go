package main

import (
	"github.com/jercle/azg/lib"
)

func main() {
	// var options lib.MultiAuthTokenRequestOptions
	// options.Scope = "graph"
	// token, err := azure.GetTenantSPToken("REDDTQ", options)
	// lib.CheckFatalError(err)
	// // roleDefs, err :=
	// defs, err := azure.ListEntraRoleDefinitions(*token)
	// jsonStr, _ := json.MarshalIndent(defs, "", "  ")

	// fmt.Println(string(jsonStr))

	// lib.InitConfig(lib.CldConfigOptions{})
	// config := lib.GetCldConfig(lib.CldConfigOptions{})
	config := lib.GetCldConfig(nil)
	_ = config

}
