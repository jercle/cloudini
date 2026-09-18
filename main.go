package main

import (
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
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	// defer lib.TimeTrack(time.Now(), "main")
	// info, _ := debug.ReadBuildInfo()
	// fmt.Println(info)
	// lib.JsonMarshalAndPrint(info.Main.Sum)
	// cmd.Execute()
	cmd.RootCmd.Execute()
	// fang.Execute()
	// if err := fang.Execute(context.TODO(), cmd.RootCmd); err != nil {
	// 	os.Exit(1)
	// }
}
