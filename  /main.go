package main

import (
	"dev.azure.com/APCDevOps/cloudini/cmd"
	_ "dev.azure.com/APCDevOps/cloudini/cmd/ado"
	_ "dev.azure.com/APCDevOps/cloudini/cmd/azure"
	_ "dev.azure.com/APCDevOps/cloudini/cmd/citrix"
	_ "dev.azure.com/APCDevOps/cloudini/cmd/config"
	_ "dev.azure.com/APCDevOps/cloudini/cmd/jira"
	_ "dev.azure.com/APCDevOps/cloudini/cmd/m365"
	_ "dev.azure.com/APCDevOps/cloudini/cmd/mongodb"
	_ "dev.azure.com/APCDevOps/cloudini/cmd/utils"
	_ "dev.azure.com/APCDevOps/cloudini/cmd/web"
)

func main() {
	// defer lib.TimeTrack(time.Now(), "main")

	cmd.Execute()
}
