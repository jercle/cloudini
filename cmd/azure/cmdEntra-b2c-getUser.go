package azure

import (
	"encoding/json"
	"fmt"

	"github.com/jercle/cloudini/lib"
	"github.com/spf13/cobra"
)

// var SetActive bool
// var ShowActive bool
// var Fetch bool

// var addTenant string
// var tenantName string
var (
	getUserByObjectID    string
	getUserByUPN         string
	configuredTenantName string
	getAllUsers          bool
)

// configCmd represents the subs command
var entraB2cGetUserCmd = &cobra.Command{
	Use:     "get-user",
	Aliases: []string{"gu"},
	Short:   "Get Azure B2C Tenant user/s",
	// 	Long: `A longer description that spans multiple lines and likely contains examples
	// and usage of using your command. For example:

	// Cobra is a CLI library for Go that empowers applications.
	// This application is a tool to generate the needed files
	// to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		tokenReq, err := GetAllTenantSPTokens(lib.AzureMultiAuthTokenRequestOptions{
			Scope:         "graph",
			GetWriteToken: true,
		}, nil)
		lib.CheckFatalError(err)
		var token *lib.AzureMultiAuthToken
		if configuredTenantName != "" {
			token, err = tokenReq.SelectTenant(configuredTenantName)
			lib.CheckFatalError(err)
		}

		if getUserByObjectID != "" {
			user := GetB2CUserByObjectId(getUserByObjectID, token)
			jsonStr, _ := json.MarshalIndent(user, "", "  ")
			fmt.Println(string(jsonStr))
		}
		if getUserByUPN != "" {
			user := GetB2CUserByUPN(getUserByUPN, token)
			jsonStr, _ := json.MarshalIndent(user, "", "  ")
			fmt.Println(string(jsonStr))
		}
		if getAllUsers {
			users := GetAllB2CTenantUsers()
			jsonStr, _ := json.MarshalIndent(users, "", "  ")
			fmt.Println(string(jsonStr))
		}
	},
}

func init() {
	entraB2cCmd.AddCommand(entraB2cGetUserCmd)

	entraB2cGetUserCmd.Flags().StringVarP(&getUserByObjectID, "getUserByObjectID", "o", "", "Get user object by ObjectId")
	entraB2cGetUserCmd.Flags().StringVarP(&getUserByUPN, "getUserByUPN", "u", "", "Get user object by User Principal Name.")
	entraB2cGetUserCmd.Flags().StringVarP(&configuredTenantName, "configuredTenantName", "n", "", "Tenant name of tenant configured in cldConfig")
	entraB2cGetUserCmd.Flags().BoolVarP(&getAllUsers, "getAll", "a", false, "Gets all users")

	entraB2cGetUserCmd.MarkFlagsMutuallyExclusive("getUserByObjectID", "getUserByUPN", "getAll")

	// entraB2cGetUserCmd.MarkFlagRequired("configuredTenantName")

	entraB2cGetUserCmd.MarkFlagsRequiredTogether("getUserByObjectID", "configuredTenantName")
	entraB2cGetUserCmd.MarkFlagsRequiredTogether("getUserByUPN", "configuredTenantName")
}
