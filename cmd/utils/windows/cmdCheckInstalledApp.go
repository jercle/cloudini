//go:build windows

package windows

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
)

var appName string
var fullData bool

// checkInstalledAppCmd represents the checkInstalledApp command
var checkInstalledAppCmd = &cobra.Command{
	Use:     "checkInstalledApp",
	Aliases: []string{"installed"},
	Short:   "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		apps := getInstalledApps()
		sort.Slice(apps, func(i, j int) bool {
			return strings.ToLower(apps[i].DisplayName) < strings.ToLower(apps[j].DisplayName)
		})

		if appName == "" {
			if fullData {
				jsonApps, err := json.MarshalIndent(apps, "", "  ")
				if err != nil {
					log.Error(err, err)
				}
				fmt.Println(string(jsonApps))
			} else {
				for _, app := range apps {
					fmt.Println(app.DisplayName)
				}
			}

		}

	},
}

func init() {
	checkInstalledAppCmd.Flags().StringVarP(&appName, "appName", "n", "", "-n [Application Name]")
	checkInstalledAppCmd.Flags().BoolVarP(&fullData, "fullData", "f", false, "Return full data instead of just app names")

	// checkInstalledAppCmd.MarkFlagRequired("appName")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// checkInstalledAppCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// checkInstalledAppCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
