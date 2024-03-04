/*
Copyright © 2024 Evan Colwell ercolwell@gmail.com
*/
package main

import (
	"time"

	"github.com/jercle/azg/cmd"
	_ "github.com/jercle/azg/cmd/ado"
	_ "github.com/jercle/azg/cmd/azure"
	_ "github.com/jercle/azg/cmd/jira"
	_ "github.com/jercle/azg/cmd/utils"
	_ "github.com/jercle/azg/cmd/web"
	"github.com/jercle/azg/lib"
)

func main() {
	defer lib.TimeTrack(time.Now(), "main")
	cmd.Execute()

}
