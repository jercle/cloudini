package azure

import (
	"github.com/jercle/cloudini/lib"
	"github.com/spf13/cobra"
)

var (
	// getRoleDefByName string
	// tenantName       string
	// getRoleDefById string
	// onlyShowId       bool
	// onlyShowName     bool
	checkStorageAcctTlsVersions bool
	getAll                      bool
)

// configCmd represents the subs command
var generalCmd = &cobra.Command{
	Use:     "general",
	Aliases: []string{"gen"},
	Short:   "Generalised Azure commands",
	// 	Long: `A longer description that spans multiple lines and likely contains examples
	// and usage of using your command. For example:

	// Cobra is a CLI library for Go that empowers applications.
	// This application is a tool to generate the needed files
	// to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		// config := lib.GetCldConfig(nil)

		if checkStorageAcctTlsVersions {
			tokenReq, err := GetAllTenantSPTokens(lib.AzureMultiAuthTokenRequestOptions{
				// Scope:         "graph",
				GetWriteToken: true,
			}, nil)
			lib.CheckFatalError(err)
			opts := lib.GetAllResourcesForAllConfiguredTenantsOptions{
				SuppressSteps: true,
			}
			var storageAccounts []StorageAccountTlsVersion
			if getAll {
				opts.GetAllStorageAccountsInTlsCheck = true
				storageAccounts = CheckStorageAccountTlsVersionsForAllConfiguredTenants(&opts, tokenReq)
			} else {
				storageAccounts = CheckStorageAccountTlsVersionsForAllConfiguredTenants(&opts, tokenReq)
			}

			lib.JsonMarshalAndPrint(storageAccounts)
		}
	},
}

func init() {
	azCmd.AddCommand(generalCmd)
	generalCmd.Flags().BoolVar(&checkStorageAcctTlsVersions, "checkStorageAcctTlsVersions", false, "Returns Storage Accounts with a minimum TLS version of less than 1.2")
	generalCmd.Flags().BoolVar(&getAll, "getAll", false, "Returns all storage account with minimum TLS details")
	// resourcesCmd.Flags().StringVar(&getRoleDefById, "getRoleDefById", "", "Get  Role Definition details by Role Def ID")
	// resourcesCmd.Flags().StringVarP(&tenantName, "tenantName", "n", "", "Tenant name to use configured auth. Defaults to Tenant of current active Az CLI subscription")
	// resourcesCmd.Flags().BoolVar(&onlyShowId, "onlyId", false, "Flag to only print the Role Def ID for 'getRoleDefByName'")
	// resourcesCmd.Flags().BoolVar(&onlyShowName, "onlyName", false, "Flag to only print the Role Def Name for 'getRoleDefById'")
	// resourcesCmd.MarkFlagsMutuallyExclusive("onlyId", "getRoleDefById")
	// resourcesCmd.MarkFlagsMutuallyExclusive("onlyName", "getRoleDefByName")
}
