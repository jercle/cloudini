package azure

import (
	"fmt"

	"github.com/jercle/cloudini/lib"
	"github.com/spf13/cobra"

	jsonc "github.com/nwidger/jsoncolor"
)

// var SetActive bool
// var ShowActive bool
// var Fetch bool

var (
	addTenant              string
	showConfig             bool
	showTenantsConfig      bool
	showSingleTenantConfig string
)

// configCmd represents the subs command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Commands related to Cloudini's Azure configuration",
	// 	Long: `A longer description that spans multiple lines and likely contains examples
	// and usage of using your command. For example:

	// Cobra is a CLI library for Go that empowers applications.
	// This application is a tool to generate the needed files
	// to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if showConfig {
			config := lib.GetCldConfig(nil)
			jsonStr, _ := jsonc.MarshalIndent(config.Azure, "", "  ")
			fmt.Println(string(jsonStr))
		}

		if showSingleTenantConfig != "" {
			config := lib.GetCldConfig(nil)
			tenantCfg, ok := config.Azure.MultiTenantAuth.Tenants[showSingleTenantConfig]
			if ok {
				jsonStr, err := jsonc.MarshalIndent(tenantCfg, "", "  ")
				lib.CheckFatalError(err)
				fmt.Println(string(jsonStr))
			} else {
				fmt.Println("Tenant '" + showSingleTenantConfig + "' does not exist in configuration")
			}
		}

		// config := lib.GetCldConfig(lib.CldConfigOptions{})

		// if addTenant != "" {
		// 	var newTenant lib.CldConfigTenantAuth
		// 	newTenant.TenantID = addTenant
		// 	newTenant.TenantName = tenantName
		// 	config.Azure.Tenants = append(config.Azure.Tenants, newTenant)
		// 	// config.SaveToFile()

		// }

	},
}

func init() {
	azCmd.AddCommand(configCmd)

	configCmd.Flags().StringVarP(&addTenant, "addTenant", "t", "", "Add tenant configuration")
	configCmd.Flags().BoolVarP(&showConfig, "show", "s", false, "Show current config file Azure configuration")
	configCmd.Flags().BoolVarP(&showTenantsConfig, "showTenantsConfig", "m", false, "Show current config file Azure Tenants configuration")
	configCmd.Flags().StringVarP(&showSingleTenantConfig, "showSingleTenantConfig", "n", "", "Show current config file for single Azure Tenant. Provide configured tenant name")

	configCmd.MarkFlagsMutuallyExclusive("showSingleTenantConfig", "showTenantsConfig", "show")
	// configCmd.Flags().BoolVarP(&SetActive, "setActive", "x", false, "Change active Azure subscription")
	// configCmd.Flags().BoolVarP(&ShowActive, "showActive", "a", false, "Show current active Azure subscription")
	// configCmd.Flags().BoolVarP(&Fetch, "fetch", "f", false, "Fetch all available subscriptions from Azure")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
