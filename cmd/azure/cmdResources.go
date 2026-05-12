package azure

import (
	"github.com/jercle/cloudini/lib"
	"github.com/spf13/cobra"
)

var (
	getRoleDefByName string
	// tenantName       string
	getRoleDefById string
	// onlyShowId       bool
	// onlyShowName     bool
)

// configCmd represents the subs command
var resourcesCmd = &cobra.Command{
	Use:     "resources",
	Aliases: []string{"res"},
	Short:   "Information and control of Azure Resources.",
	// 	Long: `A longer description that spans multiple lines and likely contains examples
	// and usage of using your command. For example:

	// Cobra is a CLI library for Go that empowers applications.
	// This application is a tool to generate the needed files
	// to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		config := lib.GetCldConfig(nil)

		// if addTenant != "" {
		// 	var newTenant lib.CldConfigTenantAuth
		// 	newTenant.TenantID = addTenant
		// 	newTenant.TenantName = tenantName
		// 	config.Azure.Tenants = append(config.Azure.Tenants, newTenant)
		// 	// config.SaveToFile()

		// }

		if getRoleDefByName != "" {
			tName := ""
			if tenantName != "" {
				tName = tenantName
			} else {
				activeSub, err := GetActiveCliSub()
				lib.CheckFatalError(err)
				for _, tConf := range config.Azure.MultiTenantAuth.Tenants {
					if tConf.TenantID == activeSub.TenantID {
						tName = tConf.TenantName
					}
				}
			}
			opts := lib.AzureMultiAuthTokenRequestOptions{
				TenantName: tName,
			}
			token, err := GetTenantSPToken(opts, nil)
			lib.CheckFatalError(err)

			_ = token

			// roleDefs, err := ListRoleDefinitions(*token)
			// lib.CheckFatalError(err)

			// var roleDef RoleDefinition

			// for _, rd := range roleDefs {
			// 	if rd.DisplayName == getRoleDefByName {
			// 		roleDef = rd
			// 	}
			// }

			// if onlyShowId {
			// 	fmt.Println(roleDef.ID)
			// } else {
			// 	jsonStr, _ := json.MarshalIndent(roleDef, "", "  ")
			// 	fmt.Println(string(jsonStr))
			// }
		}
	},
}

func init() {
	azCmd.AddCommand(resourcesCmd)
	resourcesCmd.Flags().StringVar(&getRoleDefByName, "getRoleDefByName", "", "Get  Role Definition details by Role Def name")
	resourcesCmd.Flags().StringVar(&getRoleDefById, "getRoleDefById", "", "Get  Role Definition details by Role Def ID")
	resourcesCmd.Flags().StringVarP(&tenantName, "tenantName", "n", "", "Tenant name to use configured auth. Defaults to Tenant of current active Az CLI subscription")
	resourcesCmd.Flags().BoolVar(&onlyShowId, "onlyId", false, "Flag to only print the Role Def ID for 'getRoleDefByName'")
	resourcesCmd.Flags().BoolVar(&onlyShowName, "onlyName", false, "Flag to only print the Role Def Name for 'getRoleDefById'")
	resourcesCmd.MarkFlagsMutuallyExclusive("onlyId", "getRoleDefById")
	resourcesCmd.MarkFlagsMutuallyExclusive("onlyName", "getRoleDefByName")
}
