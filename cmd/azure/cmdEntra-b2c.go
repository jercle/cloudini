package azure

import (
	"github.com/spf13/cobra"
)

// var SetActive bool
// var ShowActive bool
// var Fetch bool

// var addTenant string
// var tenantName string

// configCmd represents the subs command
var entraB2cCmd = &cobra.Command{
	Use:   "b2c",
	Short: "Azure B2C Tenant",
	// 	Long: `A longer description that spans multiple lines and likely contains examples
	// and usage of using your command. For example:

	// Cobra is a CLI library for Go that empowers applications.
	// This application is a tool to generate the needed files
	// to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

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
	entraCmd.AddCommand(entraB2cCmd)
}
