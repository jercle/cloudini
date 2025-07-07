package azure

import (
	"github.com/spf13/cobra"
)

// var SetActive bool
// var ShowActive bool
// var Fetch bool

var addTenant string

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
