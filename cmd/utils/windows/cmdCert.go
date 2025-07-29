//go:build windows

package windows

// import (
// 	"github.com/spf13/cobra"
// )

// var (
// // setProxyConfig      bool
// // selectProxyConfig   string
// // deleteProxyConfig   bool
// // listConfig          bool
// // openInternetOptions bool
// )

// // checkInstalledAppCmd represents the checkInstalledApp command
// var certCmd = &cobra.Command{
// 	Use:   "cert",
// 	Short: "A brief description of your command",
// 	Long: `* Set default proxy configured in cloudini
// 	* cld.exe u win proxy -s
// * Get currently configured proxy in Windows
// 	* cld.exe u win proxy
// * Open 'Internet Options' control panel cmdlet
// 	* cld.exe u win proxy -o
// * Remove proxy configuration
// 	* cld.exe u win proxy -d
// * Set proxy configuration from a proxyConfig object in the cloudini config file
// 	* cld.exe u win proxy -s -n NAME`,
// 	Run: func(cmd *cobra.Command, args []string) {

// 		// var proxyConfig lib.ProxyConfig

// 	},
// }

// func init() {
// 	proxyCmd.Flags().BoolVarP(&deleteProxyConfig, "deleteProxyConfig", "d", false, "Removes proxy configuration")
// 	proxyCmd.Flags().BoolVarP(&setProxyConfig, "setProxyConfig", "s", false, "Set proxy configration")
// 	proxyCmd.Flags().BoolVarP(&listConfig, "listConfig", "l", false, "List proxy settings configured in cld")
// 	proxyCmd.Flags().BoolVarP(&openInternetOptions, "openInternetOptions", "o", false, "Opens Windows Internet Options cmdlet")
// 	proxyCmd.Flags().StringVarP(&selectProxyConfig, "selectProxyConfig", "n", "", "Select proxy configuration from cld config file")

// 	// proxyCmd.MarkFlagsOneRequired("setProxyConfig", "deleteProxyConfig")
// 	proxyCmd.MarkFlagsMutuallyExclusive("setProxyConfig", "deleteProxyConfig", "listConfig", "openInternetOptions")
// 	proxyCmd.MarkFlagsMutuallyExclusive("selectProxyConfig", "deleteProxyConfig", "listConfig", "openInternetOptions")
// }
