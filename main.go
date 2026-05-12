package main

import (
	"context"
	"os"

	"github.com/jercle/cloudini/cmd"
	_ "github.com/jercle/cloudini/cmd/ado"
	_ "github.com/jercle/cloudini/cmd/azure"
	_ "github.com/jercle/cloudini/cmd/citrix"
	_ "github.com/jercle/cloudini/cmd/config"
	_ "github.com/jercle/cloudini/cmd/forgerock"
	_ "github.com/jercle/cloudini/cmd/jira"
	_ "github.com/jercle/cloudini/cmd/m365"
	_ "github.com/jercle/cloudini/cmd/mongodb"
	_ "github.com/jercle/cloudini/cmd/utils"
	_ "github.com/jercle/cloudini/cmd/web"

	"github.com/charmbracelet/fang"
)

func main() {
	// defer lib.TimeTrack(time.Now(), "main")

	// cmd.Execute()

	if err := fang.Execute(context.Background(), cmd.RootCmd); err != nil {
		os.Exit(1)
	}
}
