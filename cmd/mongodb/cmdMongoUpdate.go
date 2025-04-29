package mongodb

import (
	"context"
	"fmt"
	"time"

	"github.com/jercle/cloudini/cmd/azure"
	"github.com/jercle/cloudini/lib"
	"github.com/spf13/cobra"
)

var (
	updateAllGalleryImagesAndUpdateWithUsedByCitrix bool
	updateAzureResVcpuCountsCostData                bool
	updateAzureResourceRelations                    bool
	costDataMonth                                   string
	updateEntraItems                                bool
	updateEntraPimItems                             bool
	updateIpAddresses                               bool
	updateAllCertInfo                               bool
	showExecutionTime                               bool

// tenantId       string
// subscriptionId string
// resourceGroup  string
// clientSecret   string
// clientId       string
)

var cmdMongoUpdate = &cobra.Command{
	Use:   "update",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		startTime := time.Now()
		config := lib.GetCldConfig(nil)
		mongoConf := config.MongoDBConfig

		ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		defer cancel()

		c := ClientFromConfig(ctx, cancel)
		defer c.Disconnect(ctx)

		azResImageGalleryImagesColl := c.Database(mongoConf.DbAzRes).Collection(mongoConf.CollAzResImageGalleryImages)
		azResResourceListColl := c.Database(mongoConf.DbAzRes).Collection(mongoConf.CollAzResResourceList)
		azResGrpsListColl := c.Database(mongoConf.DbAzRes).Collection(mongoConf.CollAzResGrpsList)
		azResSKUColl := c.Database(mongoConf.DbAzRes).Collection(mongoConf.CollAzResSKU)
		azResTenantsColl := c.Database(mongoConf.DbAzRes).Collection(mongoConf.CollAzResTenants)
		azResVcpuCountsColl := c.Database(mongoConf.DbAzRes).Collection(mongoConf.CollAzResVcpuCounts)
		azResIpAddresses := c.Database(mongoConf.DbAzRes).Collection(mongoConf.CollAzResIPAddresses)

		citrixMachineCatalogsColl := c.Database(mongoConf.DbCitrix).Collection(mongoConf.CollCitrixMachineCatalogs)

		certsCaCertInfo := c.Database(mongoConf.DbCertificates).Collection(mongoConf.CollCertsCaCertInfo)
		certsServerCertInfo := c.Database(mongoConf.DbCertificates).Collection(mongoConf.CollCertsServerCertInfo)

		entraAppRegColl := c.Database(mongoConf.DbEntra).Collection(mongoConf.CollEntraAppReg)
		entraAppRegCredsExpiringColl := c.Database(mongoConf.DbEntra).Collection(mongoConf.CollEntraAppRegCredsExpiring)
		entraRoleAssignmentScheduleInstancesColl := c.Database(mongoConf.DbEntra).Collection(mongoConf.CollEntraRoleAssignmentScheduleInstances)
		entraRoleEligibilityScheduleInstancesColl := c.Database(mongoConf.DbEntra).Collection(mongoConf.CollEntraRoleEligibilityScheduleInstances)

		// envOptCostingColl := c.Database(mongoConf.DbEnvironmentOptimisation).Collection(mongoConf.CollEnvOptCosting)
		envOptCostingMetersColl := c.Database(mongoConf.DbEnvironmentOptimisation).Collection(mongoConf.CollEnvOptCostingMeters)
		envOptCostingResGrpsColl := c.Database(mongoConf.DbEnvironmentOptimisation).Collection(mongoConf.CollEnvOptCostingResGrps)
		envOptCostingResourcesColl := c.Database(mongoConf.DbEnvironmentOptimisation).Collection(mongoConf.CollEnvOptCostingResources)
		envOptCostingSubsColl := c.Database(mongoConf.DbEnvironmentOptimisation).Collection(mongoConf.CollEnvOptCostingSubs)
		envOptCostingTenantsColl := c.Database(mongoConf.DbEnvironmentOptimisation).Collection(mongoConf.CollEnvOptCostingTenants)

		tokenReq, err := azure.GetAllTenantSPTokens(lib.AzureMultiAuthTokenRequestOptions{}, nil)
		lib.CheckFatalError(err)

		if updateIpAddresses {
			UpdateAllAzureResourceIPAddresses(azResIpAddresses, tokenReq)
		}

		if updateAllGalleryImagesAndUpdateWithUsedByCitrix {
			UpdateAllGalleryImagesAndUpdateWithUsedByCitrix(azResImageGalleryImagesColl, citrixMachineCatalogsColl, tokenReq)
		}

		if updateAllCertInfo {
			UpdateAllCertInfo(certsCaCertInfo, certsServerCertInfo)
		}

		if updateEntraItems {
			appRegOpts := UpdateEntraItemsOptions{
				EntraAppRegColl:              entraAppRegColl,
				EntraAppRegCredsExpiringColl: entraAppRegCredsExpiringColl,
			}
			UpdateEntraItems(appRegOpts, tokenReq)
		}

		if updateEntraPimItems {
			opts := UpdateEntraPimItemsOptions{
				EntraRoleAssignmentScheduleInstancesColl:  entraRoleAssignmentScheduleInstancesColl,
				EntraRoleEligibilityScheduleInstancesColl: entraRoleEligibilityScheduleInstancesColl,
			}
			UpdateEntraPimItems(opts)
		}

		// Longest running, so keep last
		if updateAzureResVcpuCountsCostData {
			opts := UpdateAllAzureResourcesAndVcpuCountsOptions{
				SkuListSubscription:        config.Azure.SkuListSubscription,
				SkuListAuth:                config.Azure.MultiTenantAuth.Tenants[config.Azure.SkuListAuthTenant],
				Location:                   config.Azure.ResourceLocation,
				CostDataMonth:              costDataMonth,
				CostDataBlobPrefix:         config.Azure.CostDataBlobPrefix,
				AzResSKUColl:               azResSKUColl,
				AzResVcpuCountsColl:        azResVcpuCountsColl,
				AzResTenantsColl:           azResTenantsColl,
				AzResResourceListColl:      azResResourceListColl,
				AzResGrpsListColl:          azResGrpsListColl,
				EnvOptCostingTenantsColl:   envOptCostingTenantsColl,
				EnvOptCostingSubsColl:      envOptCostingSubsColl,
				EnvOptCostingResGrpsColl:   envOptCostingResGrpsColl,
				EnvOptCostingResourcesColl: envOptCostingResourcesColl,
				EnvOptCostingMetersColl:    envOptCostingMetersColl,
			}
			transformedData := UpdateAllAzureResourcesVcpuCountsCostData(opts, tokenReq)

			if updateAzureResourceRelations {
				UpdateAzureResourceRelations(transformedData, opts)
			}
		}

		elapsed := time.Since(startTime)
		if showExecutionTime {
			fmt.Println("Execution time: " + elapsed.String())
		}
	},
}

func init() {
	cmdMongo.AddCommand(cmdMongoUpdate)
	cmdMongoUpdate.Flags().BoolVarP(&updateAllGalleryImagesAndUpdateWithUsedByCitrix, "updateAllGalleryImagesAndUpdateWithUsedByCitrix", "g", false, "Gets all gallery images in configured tenants, then checks agains Citrix and updates in database")
	cmdMongoUpdate.Flags().BoolVarP(&updateAzureResVcpuCountsCostData, "updateAzureResVcpuCountsCostData", "c", false, "Gets latest cost data and all resources, transforms and relates them, then updates database")
	// cmdMongoUpdate.Flags().BoolVarP(&updateAzureResourceRelations, "updateAzureResourceRelations", "r", false, "Gets all resources from cost data and database, aggregates and finds relations, then updates database. This can only be used in conjunction with 'updateAzureResVcpuCountsCostData'")
	cmdMongoUpdate.Flags().BoolVarP(&updateEntraItems, "updateEntraItems", "e", false, "Gets all App Registrations from configured Azure tenants and finds expiring credentials, then updates database")
	cmdMongoUpdate.Flags().BoolVarP(&updateIpAddresses, "updateIpAddresses", "i", false, "Gets all App Registrations from configured Azure tenants and finds expiring credentials, then updates database")
	cmdMongoUpdate.Flags().BoolVarP(&updateAllCertInfo, "updateAllCertInfo", "x", false, "Update server certificates and expiries")
	cmdMongoUpdate.Flags().BoolVarP(&updateEntraPimItems, "updateEntraPimItems", "p", false, "Gets all PIM assignments and eligibilities, then updates database")
	cmdMongoUpdate.Flags().BoolVarP(&showExecutionTime, "showExecutionTime", "t", false, "Prints execution time when complete")
	cmdMongoUpdate.Flags().StringVarP(&costDataMonth, "costDataMonth", "m", "", "Which month to get cost data from - defaults to whatever month it was yesterday. Use with 'updateAzureResVcpuCountsCostData' Format: YYYYMM")

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
