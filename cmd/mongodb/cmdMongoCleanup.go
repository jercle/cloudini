package mongodb

import (
	"context"
	"time"

	"github.com/jercle/cloudini/lib"
	"github.com/spf13/cobra"
)

var (
	resetLifetimeCostValues           bool
	resetRelatedResourcesAndCostItems bool
	unsetResourceField                string

// updateAllAzureResourcesVcpuCountsCostData       bool
// costDataMonth                                   string
// updateEntraItems                                bool

// tenantId       string
// subscriptionId string
// resourceGroup  string
// clientSecret   string
// clientId       string
)

var cmdMongoCleanup = &cobra.Command{
	Use:   "cleanup",
	Short: "Cleans up data within MongoDB generated from Cloudini",
	// 	Long: `A longer description that spans multiple lines and likely contains examples
	// and usage of using your command. For example:

	// Cobra is a CLI library for Go that empowers applications.
	// This application is a tool to generate the needed files
	// to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		config := lib.GetCldConfig(nil)
		mongoConf := config.MongoDBConfig

		ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		defer cancel()

		c := ClientFromConfig(ctx, cancel)
		defer c.Disconnect(ctx)

		// azResImageGalleryImagesColl := c.Database(mongoConf.DbAzRes).Collection(mongoConf.CollAzResImageGalleryImages)
		azResResourceListColl := c.Database(mongoConf.DbAzRes).Collection(mongoConf.CollAzResResourceList)
		// azResGrpsListColl := c.Database(mongoConf.DbAzRes).Collection(mongoConf.CollAzResGrpsList)
		// azResSKUColl := c.Database(mongoConf.DbAzRes).Collection(mongoConf.CollAzResSKU)
		// azResTenantsColl := c.Database(mongoConf.DbAzRes).Collection(mongoConf.CollAzResTenants)
		// azResVcpuCountsColl := c.Database(mongoConf.DbAzRes).Collection(mongoConf.CollAzResVcpuCounts)

		// citrixMachineCatalogsColl := c.Database(mongoConf.DbCitrix).Collection(mongoConf.CollCitrixMachineCatalogs)

		// entraAppRegColl := c.Database(mongoConf.DbEntra).Collection(mongoConf.CollEntraAppReg)
		// entraAppRegCredsExpiringColl := c.Database(mongoConf.DbEntra).Collection(mongoConf.CollEntraAppRegCredsExpiring)
		// envOptCostingColl := c.Database(mongoConf.DbEnvironmentOptimisation).Collection(mongoConf.CollEnvOptCosting)
		envOptCostingMetersColl := c.Database(mongoConf.DbEnvironmentOptimisation).Collection(mongoConf.CollEnvOptCostingMeters)
		envOptCostingResGrpsColl := c.Database(mongoConf.DbEnvironmentOptimisation).Collection(mongoConf.CollEnvOptCostingResGrps)
		envOptCostingResourcesColl := c.Database(mongoConf.DbEnvironmentOptimisation).Collection(mongoConf.CollEnvOptCostingResources)
		envOptCostingSubsColl := c.Database(mongoConf.DbEnvironmentOptimisation).Collection(mongoConf.CollEnvOptCostingSubs)
		envOptCostingTenantsColl := c.Database(mongoConf.DbEnvironmentOptimisation).Collection(mongoConf.CollEnvOptCostingTenants)

		// tokenReq, err := azure.GetAllTenantSPTokens(lib.AzureMultiAuthTokenRequestOptions{}, nil)
		// lib.CheckFatalError(err)

		if resetLifetimeCostValues {
			ResetLifetimeCostValues(envOptCostingMetersColl)
			ResetLifetimeCostValues(envOptCostingResGrpsColl)
			ResetLifetimeCostValues(envOptCostingResourcesColl)
			ResetLifetimeCostValues(envOptCostingSubsColl)
			ResetLifetimeCostValues(envOptCostingTenantsColl)
		}

		if resetRelatedResourcesAndCostItems {
			ResetRelatedResourcesAndCostItems(azResResourceListColl)
		}

		if unsetResourceField != "" {
			UnsetField(unsetResourceField, azResResourceListColl)
		}
	},
}

func init() {
	cmdMongo.AddCommand(cmdMongoCleanup)
	cmdMongoCleanup.Flags().BoolVarP(&resetLifetimeCostValues, "resetLifetimeCostValues", "c", false, "Gets all cost data, and resets it flowing up from Cost Meters to Tenants")
	cmdMongoCleanup.Flags().BoolVarP(&resetRelatedResourcesAndCostItems, "resetRelatedResourcesAndCostItems", "r", false, "Gets all resources, then removes duplicates from relatedCostMeters and relatedResources fields")
	cmdMongo.PersistentFlags().StringVarP(&unsetResourceField, "unsetResourceField", "u", "", "Unsets a field in the allResources collection")
	// cmdMongoUpdate.Flags().BoolVarP(&updateEntraItems, "updateEntraItems", "e", false, "Gets all App Registrations from configured Azure tenants and finds expiring credentials, then updates database")
	// cmdMongoUpdate.Flags().StringVarP(&costDataMonth, "costDataMonth", "m", "", "Which month to get cost data from - defaults to whatever month it was yesterday. Use with 'updateAllAzureResourcesVcpuCountsCostData' Format: YYYYMM")
	// cmdMongo.PersistentFlags().StringVarP(&subscriptionId, "subscriptionId", "s", "", "Subscription ID to run command against. If not supplied, current default Azure CLI subscription is used.")
	// cmdMongo.PersistentFlags().StringVarP(&resourceGroup, "resourceGroup", "r", "", "Resource group to run command against.")
	// cmdMongo.PersistentFlags().StringVar(&clientId, "clientId", "", "Client ID for Service Principal authentication.")
	// cmdMongo.PersistentFlags().StringVar(&clientSecret, "clientSecret", "", "Client Secret for Service Principal authentication.")
	// cmdMongo.PersistentFlags().StringVarP(&tenantId, "tenantId", "t", "", "Tenant ID.")

	// if subscriptionId == "" {
	// 	sub, err := GetActiveSub()
	// 	if err != nil {
	// 		log.Fatalln(err)
	// 	}
	// 	subscriptionId = sub.ID
	// }

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// azCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// azCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
