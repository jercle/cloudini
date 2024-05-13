/*
Copyright © 2024 Evan Colwell ercolwell@gmail.com
*/
package main

import (
	"github.com/jercle/cloudini/cmd"
	_ "github.com/jercle/cloudini/cmd/ado"
	_ "github.com/jercle/cloudini/cmd/azure"
	_ "github.com/jercle/cloudini/cmd/config"
	_ "github.com/jercle/cloudini/cmd/jira"
	_ "github.com/jercle/cloudini/cmd/utils"
	_ "github.com/jercle/cloudini/cmd/web"
)

func main() {
	// defer lib.TimeTrack(time.Now(), "main")

	var (
		version = "dev"
		commit  = "none"
		date    = "unknown"
	)

	cmd.SetVersionInfo(version, commit, date)

	cmd.Execute()
}
