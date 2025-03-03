//go:build windows

package windows

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"

	"github.com/jercle/cloudini/lib"
	"github.com/spf13/cobra"
)

var (
	setProxyConfig      bool
	selectProxyConfig   string
	deleteProxyConfig   bool
	listConfig          bool
	openInternetOptions bool
)

// checkInstalledAppCmd represents the checkInstalledApp command
var proxyCmd = &cobra.Command{
	Use:   "proxy",
	Short: "A brief description of your command",
	Long: `* Set default proxy configured in cloudini
	* cld.exe u win proxy -s
* Get currently configured proxy in Windows
	* cld.exe u win proxy
* Open 'Internet Options' control panel cmdlet
	* cld.exe u win proxy -o
* Remove proxy configuration
	* cld.exe u win proxy -d
* Set proxy configuration from a proxyConfig object in the cloudini config file
	* cld.exe u win proxy -s -n NAME`,
	Run: func(cmd *cobra.Command, args []string) {

		// var proxyConfig lib.ProxyConfig

		if openInternetOptions {
			command := exec.Command("control", "inetcpl.cpl")
			command.Run()
			os.Exit(0)
		}

		if listConfig {
			proxyConfig := lib.GetCldConfig(nil).ProxyConfig
			jsonBytes, _ := json.MarshalIndent(proxyConfig, "", "  ")
			fmt.Println(string(jsonBytes))
			os.Exit(0)
		}

		if !setProxyConfig && !deleteProxyConfig {
			proxyConfig := GetProxySettings()
			jsonBytes, _ := json.MarshalIndent(proxyConfig, "", "  ")
			fmt.Println(string(jsonBytes))
		}

		if setProxyConfig {
			cldConf := lib.GetCldConfig(nil)
			proxyConfigs := *cldConf.ProxyConfig
			if selectProxyConfig == "" {
				SetProxySettings(proxyConfigs["default"], false)
			} else {
				SetProxySettings(proxyConfigs[selectProxyConfig], false)
			}
			proxyConfig := GetProxySettings()
			jsonBytes, _ := json.MarshalIndent(proxyConfig, "", "  ")
			fmt.Println(string(jsonBytes))
		}

		if deleteProxyConfig {
			RemoveProxyConfig()
			proxyConfig := GetProxySettings()
			jsonBytes, _ := json.MarshalIndent(proxyConfig, "", "  ")
			fmt.Println(string(jsonBytes))
		}

	},
}

func init() {
	proxyCmd.Flags().BoolVarP(&deleteProxyConfig, "deleteProxyConfig", "d", false, "Removes proxy configuration")
	proxyCmd.Flags().BoolVarP(&setProxyConfig, "setProxyConfig", "s", false, "Set proxy configration")
	proxyCmd.Flags().BoolVarP(&listConfig, "listConfig", "l", false, "List proxy settings configured in cld")
	proxyCmd.Flags().BoolVarP(&openInternetOptions, "openInternetOptions", "o", false, "Opens Windows Internet Options cmdlet")
	proxyCmd.Flags().StringVarP(&selectProxyConfig, "selectProxyConfig", "n", "", "Select proxy configuration from cld config file")

	// proxyCmd.MarkFlagsOneRequired("setProxyConfig", "deleteProxyConfig")
	proxyCmd.MarkFlagsMutuallyExclusive("setProxyConfig", "deleteProxyConfig", "listConfig", "openInternetOptions")
	proxyCmd.MarkFlagsMutuallyExclusive("selectProxyConfig", "deleteProxyConfig", "listConfig", "openInternetOptions")
}
