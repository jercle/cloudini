// fakeDataGen
package main

import (
	"encoding/json"
	"fmt"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/jercle/cloudini/cmd/citrix"
	"github.com/jercle/cloudini/lib"
)

func main() {
	var machineCatalog citrix.MachineCatalog

	err := gofakeit.Struct(&machineCatalog)
	lib.CheckFatalError(err)

	jsonStr, _ := json.MarshalIndent(machineCatalog, "", "  ")
	fmt.Println(string(jsonStr))
}

func generateMachineCatalog() {

}
