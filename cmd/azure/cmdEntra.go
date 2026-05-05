package azure

import (
	"encoding/json"
	"fmt"

	"github.com/jercle/cloudini/lib"
	"github.com/spf13/cobra"
)

var (
	getEntraRoleDefByName string
	tenantName            string
	getEntraRoleDefById   string
	onlyShowId            bool
	onlyShowName          bool
)

// configCmd represents the subs command
var entraCmd = &cobra.Command{
	Use:   "entra",
	Short: "Azure Entra / Azure AD",
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

		if getEntraRoleDefByName != "" {
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
				Scope:      "graph",
			}
			token, err := GetTenantSPToken(opts, nil)
			lib.CheckFatalError(err)

			roleDefs, err := ListEntraRoleDefinitions(*token)
			lib.CheckFatalError(err)

			var roleDef EntraRoleDefinition

			for _, rd := range roleDefs {
				if rd.DisplayName == getEntraRoleDefByName {
					roleDef = rd
				}
			}

			if onlyShowId {
				fmt.Println(roleDef.ID)
			} else {
				jsonStr, _ := json.MarshalIndent(roleDef, "", "  ")
				fmt.Println(string(jsonStr))
			}
		}

		if getEntraRoleDefById != "" {
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
				Scope:      "graph",
			}
			token, err := GetTenantSPToken(opts, nil)
			lib.CheckFatalError(err)

			roleDefs, err := ListEntraRoleDefinitions(*token)
			lib.CheckFatalError(err)

			var roleDef EntraRoleDefinition

			for _, rd := range roleDefs {
				if rd.ID == getEntraRoleDefById {
					roleDef = rd
				}
			}

			if onlyShowName {
				fmt.Println(roleDef.ID)
			} else {
				jsonStr, _ := json.MarshalIndent(roleDef, "", "  ")
				fmt.Println(string(jsonStr))
			}
		}
	},
}

func init() {
	azCmd.AddCommand(entraCmd)
	entraCmd.Flags().StringVar(&getEntraRoleDefByName, "getEntraRoleDefByName", "", "Get Entra Role Definition details by Role Def name")
	entraCmd.Flags().StringVar(&getEntraRoleDefById, "getEntraRoleDefById", "", "Get Entra Role Definition details by Role Def ID")
	entraCmd.Flags().StringVarP(&tenantName, "tenantName", "n", "", "Tenant name to use configured auth. Defaults to Tenant of current active Az CLI subscription")
	entraCmd.Flags().BoolVar(&onlyShowId, "onlyId", false, "Flag to only print the Role Def ID for 'getEntraRoleDefByName'")
	entraCmd.Flags().BoolVar(&onlyShowName, "onlyName", false, "Flag to only print the Role Def Name for 'getEntraRoleDefById'")
	entraCmd.MarkFlagsMutuallyExclusive("onlyId", "getEntraRoleDefById")
	entraCmd.MarkFlagsMutuallyExclusive("onlyName", "getEntraRoleDefByName")
}
