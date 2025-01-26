package main

import (
	"github.com/jercle/cloudini/cmd"
	_ "github.com/jercle/cloudini/cmd/ado"
	_ "github.com/jercle/cloudini/cmd/azure"
	_ "github.com/jercle/cloudini/cmd/citrix"
	_ "github.com/jercle/cloudini/cmd/config"
	_ "github.com/jercle/cloudini/cmd/jira"
	_ "github.com/jercle/cloudini/cmd/m365"
	_ "github.com/jercle/cloudini/cmd/mongodb"
	_ "github.com/jercle/cloudini/cmd/utils"
	_ "github.com/jercle/cloudini/cmd/web"
)

func main() {
	// defer lib.TimeTrack(time.Now(), "main")

	cmd.Execute()
}
