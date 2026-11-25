// fakeDataGen
package main

import (
	"encoding/json/jsontext"
	json "encoding/json/v2"
	"fmt"

	"github.com/jercle/cloudini/cmd/citrix"
	"github.com/jercle/cloudini/lib"
	"github.com/brianvoe/gofakeit/v7"
)

func main() {
	var machineCatalog citrix.MachineCatalog

	err := gofakeit.Struct(&machineCatalog)
	lib.CheckFatalError(err)

	jsonStr, _ := json.Marshal(machineCatalog, jsontext.WithIndent("  "))
	fmt.Println(string(jsonStr))
}

func generateMachineCatalog() {

}
