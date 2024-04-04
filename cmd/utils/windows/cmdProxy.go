//go:build windows

/*
Copyright © 2024 Evan Colwell ercolwell@gmail.com
*/
package windows

import (
	"github.com/jercle/azg/lib"
	"github.com/spf13/cobra"
)

var setProxyConfig string
var getProxyConfig bool
var deleteProxyConfig bool

// checkInstalledAppCmd represents the checkInstalledApp command
var proxyCmd = &cobra.Command{
	Use:     "proxy",
	Aliases: []string{"installed"},
	Short:   "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		if getProxyConfig {
			GetProxySettings()
		}

		if setProxyConfig != "" {
			cldConf := lib.GetCldConfig(nil)
			SetProxySettings(cldConf.ProxyConfig[setProxyConfig], false)
			GetProxySettings()
		}

		if deleteProxyConfig {
			RemoveProxyConfig()
		}

	},
}

func init() {
	proxyCmd.Flags().BoolVarP(&getProxyConfig, "getProxyConfig", "g", false, "Shows current proxy configuration")
	proxyCmd.Flags().BoolVarP(&deleteProxyConfig, "deleteProxyConfig", "d", false, "Shows current proxy configuration")
	proxyCmd.Flags().StringVarP(&setProxyConfig, "setProxyConfig", "s", "", "Shows current proxy configuration")

	proxyCmd.MarkFlagsOneRequired("getProxyConfig", "setProxyConfig", "deleteProxyConfig")
	proxyCmd.MarkFlagsMutuallyExclusive("getProxyConfig", "setProxyConfig", "deleteProxyConfig")

	// checkInstalledAppCmd.MarkFlagRequired("appName")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// checkInstalledAppCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// checkInstalledAppCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
