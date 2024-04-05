//go:build windows

/*
Copyright © 2024 Evan Colwell ercolwell@gmail.com
*/
package windows

import (
	"fmt"

	"github.com/jercle/azg/lib"
	"github.com/spf13/cobra"
)

var setProxyConfig bool
var selectProxyConfig string
var deleteProxyConfig bool

// checkInstalledAppCmd represents the checkInstalledApp command
var proxyCmd = &cobra.Command{
	Use:   "proxy",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		// var proxyConfig lib.ProxyConfig

		if !setProxyConfig && !deleteProxyConfig {
			proxyConfig := GetProxySettings()
			fmt.Println(proxyConfig)
			// if cmd.
		}

		if setProxyConfig {
			cldConf := lib.GetCldConfig(nil)
			if selectProxyConfig == "" {
				SetProxySettings(cldConf.ProxyConfig["default"], false)
			} else {
				SetProxySettings(cldConf.ProxyConfig[selectProxyConfig], false)
			}
			proxyConfig := GetProxySettings()
			fmt.Println(proxyConfig)
		}

		if deleteProxyConfig {
			RemoveProxyConfig()
		}

	},
}

func init() {
	proxyCmd.Flags().BoolVarP(&deleteProxyConfig, "deleteProxyConfig", "d", false, "Removes proxy configuration")
	proxyCmd.Flags().BoolVarP(&setProxyConfig, "setProxyConfig", "s", false, "Set proxy configration")
	proxyCmd.Flags().StringVarP(&selectProxyConfig, "selectProxyConfig", "n", "", "Select proxy configuration from cld config file")

	// proxyCmd.MarkFlagsOneRequired("setProxyConfig", "deleteProxyConfig")
	proxyCmd.MarkFlagsMutuallyExclusive("setProxyConfig", "deleteProxyConfig")
	proxyCmd.MarkFlagsMutuallyExclusive("selectProxyConfig", "deleteProxyConfig")
}
