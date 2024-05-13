/*
Copyright © 2024 Evan Colwell ercolwell@gmail.com
*/
package main

import (
	"time"

	"github.com/jercle/cloudini/cmd"
	_ "github.com/jercle/cloudini/cmd/ado"
	_ "github.com/jercle/cloudini/cmd/azure"
	_ "github.com/jercle/cloudini/cmd/config"
	_ "github.com/jercle/cloudini/cmd/jira"
	_ "github.com/jercle/cloudini/cmd/utils"
	_ "github.com/jercle/cloudini/cmd/web"

	"github.com/carlmjohnson/versioninfo"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	// if installing from source, i.e. not a binary built by `Goreleaser`, this version will not be updated. In that case, use the automagic detection in `versioninfo`
	if version == "dev" {
		version = versioninfo.Version
		commit = versioninfo.Revision
		date = versioninfo.LastCommit.Format(time.RFC3339)
	} else {
		// Goreleaser doesn't prefix with a `v`, which we expect
		version = "v" + version
	}

	cmd.SetVersionInfo(version, commit, date)
	cmd.Execute()

}
