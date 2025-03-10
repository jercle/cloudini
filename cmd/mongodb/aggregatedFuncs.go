package mongodb

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/jercle/cloudini/cmd/ado"
	"github.com/jercle/cloudini/cmd/azure"
	"github.com/jercle/cloudini/cmd/citrix"
	"github.com/jercle/cloudini/lib"
	"github.com/briandowns/spinner"
	"go.mongodb.org/mongo-driver/mongo"
)

func UpdateAllGalleryImagesAndUpdateWithUsedByCitrix(imageGalleryImagesColl *mongo.Collection, machineCatalogsColl *mongo.Collection, tokenReq lib.AllTenantTokens) {
	s := spinner.New(spinner.CharSets[43], 100*time.Millisecond)

	fmt.Println("Fetching gallery images...")
	s.Start()
	galleryImagesWithVersions := azure.GetAllImagesAndVersionsForAllGalleries(tokenReq)
	s.Stop()
	fmt.Println("Updating gallery images in database...")
	s.Start()
	UpsertImageGalleryImages(galleryImagesWithVersions, imageGalleryImagesColl)
	s.Stop()

	config := lib.GetCldConfig(nil)
	citrixEnvs := *config.CitrixCloud.Environments
	if citrixEnvs == nil {
		err := fmt.Errorf("Citrix environments not configured")
		lib.CheckFatalError(err)
	}

	for envName, envCreds := range citrixEnvs {
		tokenData, err := citrix.GetToken(envCreds, nil)
		lib.CheckFatalError(err)
		fmt.Println("Fetching Machine Catalogs for " + envName + "...")
		s.Start()
		machineCatalogs := citrix.ListMachineCatalogs(envCreds, tokenData)
		fmt.Println("Updating Citrix Machine Catalogs in database...")
		s.Start()
		UpsertCitrixMachineCatalogs(machineCatalogs, machineCatalogsColl)
		s.Stop()
		mcMasterImageVersions := machineCatalogs.ListImageVersions()
		s.Stop()
		fmt.Println("Updating Azure Images used by Citrix in database...")
		s.Start()
		MarkImageGalleryImagesUsedByCitrix(mcMasterImageVersions, imageGalleryImagesColl)
		s.Stop()
	}

	GetBuildDataAndUpdateImageVesionData(imageGalleryImagesColl)
}

//
//

func UpdateAllAzureResourceIPAddresses(resIPAddressesColl *mongo.Collection, tokenReq lib.AllTenantTokens) {
	s := spinner.New(spinner.CharSets[43], 100*time.Millisecond)

	opts := lib.GetAllResourcesForAllConfiguredTenantsOptions{
		SuppressSteps: true,
	}

	fmt.Println("Fetching all resource IPs...")
	s.Start()
	resources := azure.GetAllIpAddrForAllConfiguredTenants(&opts, tokenReq)
	s.Stop()

	fmt.Println("Updating all resource IPs in database...")
	s.Start()
	UpsertAzureIPAddresses(resources, resIPAddressesColl)
	s.Stop()

}

//
//

func UpdateAllAzureResourcesVcpuCountsCostData(opts UpdateAllAzureResourcesAndVcpuCountsOptions, tokenReq lib.AllTenantTokens) lib.AggregatedCostData {
	if opts.AzResGrpsListColl == nil {
		fmt.Println("AzResGrpsListColl == nil")
		os.Exit(1)
	}

	var costExportMonth string

	if opts.CostDataMonth == "" {
		costExportMonth = time.Now().AddDate(0, 0, -1).Format("200601")
	} else {
		costExportMonth = opts.CostDataMonth
	}

	config := lib.GetCldConfig(nil)
	_, _, cachePath := lib.InitConfig(nil)
	s := spinner.New(spinner.CharSets[43], 100*time.Millisecond)

	resSkuOpts := lib.GetAllResourcesForAllConfiguredTenantsOptions{
		SubscriptionId: opts.SkuListSubscription,
		AzureAuth:      opts.SkuListAuth,
		Location:       opts.Location,
		SuppressSteps:  true,
	}

	fmt.Println("Getting Azure Resource SKUs...")
	s.Start()
	resourceSKUs := azure.GetAzureResourceSKUsForSubscription(resSkuOpts)
	s.Stop()
	fmt.Println("Updating Azure Resource SKUs in database...")
	s.Start()
	UpsertResourceSKUs(resourceSKUs, opts.AzResSKUColl)
	s.Stop()

	fmt.Println("Fetching all Azure Resources...")
	s.Start()
	allResources, allResourcesSlice := azure.GetAllResourcesForAllConfiguredTenants(&resSkuOpts, tokenReq)
	s.Stop()
	allResourcesSliceStr, _ := json.MarshalIndent(allResourcesSlice, "", "  ")
	os.WriteFile(cachePath+"/allResourcesSlice.json", allResourcesSliceStr, 0644)

	fmt.Println("Updating Azure Resources in database...")
	s.Start()
	UpsertMultipleResources(allResourcesSlice, opts.AzResResourceListColl)
	s.Stop()

	fmt.Println("Updating 'existsInAzure' value for all resources in database...")
	s.Start()
	UpdateResourcesNotExistInAzure(allResourcesSlice, opts.AzResResourceListColl)
	s.Stop()

	fmt.Println("Fetching all Azure Resource Groups...")
	s.Start()
	allResGrps := azure.GetAllResGrpsForAllConfiguredTenants(&resSkuOpts, tokenReq)
	s.Stop()
	fmt.Println("Updating Azure Resource Groups in database...")
	s.Start()
	UpsertMultipleResGrps(allResGrps, opts.AzResGrpsListColl)
	s.Stop()

	fmt.Println("Getting vCPU Counts...")
	s.Start()
	_,
		_,
		_,
		_,
		_,
		vCpuCountWithResources := azure.GetVcpuCountForAllConfiguredTenants(allResources, nil, config.Azure.MultiTenantAuth.Tenants)
	s.Stop()
	vCpuCountWithResourcesStr, _ := json.MarshalIndent(vCpuCountWithResources, "", "  ")
	os.WriteFile(cachePath+"/vCpuCountWithResources.json", vCpuCountWithResourcesStr, 0644)
	fmt.Println("Updating vCPU Counts in database...")
	s.Start()
	UpsertVcpuCounts(vCpuCountWithResources, opts.AzResVcpuCountsColl)
	s.Stop()

	tempBlobDir := cachePath + "/costexports"
	costExportsOutfilePath := tempBlobDir + "/" + costExportMonth

	fmt.Println("Getting cost export data for " + costExportMonth + "...")
	s.Start()
	azure.DownloadAllConfiguredTenantCostExportsForMonth(lib.DownloadAllConfiguredTenantCostExportsForMonthOptions{
		BlobPrefix:        opts.CostDataBlobPrefix + "/" + costExportMonth,
		OutfilePath:       costExportsOutfilePath,
		OutfileNamePrefix: "cost-export",
		CostExportMonth:   costExportMonth,
		SuppressSteps:     true,
	}, nil)
	s.Stop()

	fmt.Println("Combining cost export data")
	s.Start()
	combinedCostData := azure.CombineCostExportCSVData(costExportsOutfilePath)
	s.Stop()

	fmt.Println("Transforming cost export data")
	transformedData := azure.TransformCostDataNew(combinedCostData, 1, 2)
	transformedDataStr, _ := json.MarshalIndent(transformedData, "", "  ")
	os.WriteFile(cachePath+"/transformedData.json", transformedDataStr, 0644)

	fmt.Println("Updating cost data in database")
	UpsertMonthlyTenantSubResGrpCosts(transformedData,
		costExportMonth,
		opts.EnvOptCostingTenantsColl,
		opts.EnvOptCostingSubsColl,
		opts.EnvOptCostingResGrpsColl,
		opts.EnvOptCostingResourcesColl,
		opts.EnvOptCostingMetersColl,
		opts.AzResTenantsColl,
	)

	fmt.Println("Deleting cached cost data")
	os.RemoveAll(tempBlobDir)

	return transformedData
}

//
//

func UpdateAzureResourceRelations(transformedData lib.AggregatedCostData, opts UpdateAllAzureResourcesAndVcpuCountsOptions) {
	s := spinner.New(spinner.CharSets[43], 100*time.Millisecond)

	// fmt.Println("Processing cost data into usable array")
	_, costDataSlice := azure.ProcessCostData(transformedData)

	fmt.Println("Getting current and historical resources from database")
	s.Start()
	resourceFromDatabase := GetAllResources(opts.AzResResourceListColl)
	s.Stop()

	_, _, cachePath := lib.InitConfig(nil)
	fmt.Println("Processing all resources and cost meters to create relations")
	costDataSliceStr, _ := json.MarshalIndent(costDataSlice, "", "  ")
	os.WriteFile(cachePath+"/costDataSlice.json", costDataSliceStr, 0644)
	resourceFromDatabaseStr, _ := json.MarshalIndent(resourceFromDatabase, "", "  ")
	os.WriteFile(cachePath+"/resourceFromDatabase.json", resourceFromDatabaseStr, 0644)
	_, processedResourcesSlice := azure.GatherRelatedResourcesAndCostMeters(costDataSlice, resourceFromDatabase, 2, 2)
	fmt.Println("Updating all resources with related cost data in database")

	fmt.Println("Upserting all processed resources to database...")
	s.Start()
	UpsertMultipleResources(processedResourcesSlice, opts.AzResResourceListColl)
	s.Stop()
}

//
//

func UpdateEntraItems(opts UpdateEntraItemsOptions, tokenReq lib.AllTenantTokens) {
	s := spinner.New(spinner.CharSets[43], 100*time.Millisecond)
	_, _, cachePath := lib.InitConfig(nil)

	fmt.Println("Fetching all App Registrations...")
	s.Start()
	allAppRegistrations, appRegExpiringCreds := azure.GetAppRegDataForAllConfiguredTenants("")
	s.Stop()

	allAppRegistrationsStr, _ := json.MarshalIndent(allAppRegistrations, "", "  ")
	os.WriteFile(cachePath+"/allAppRegistrations.json", allAppRegistrationsStr, 0644)
	defer os.Remove(cachePath + "/allAppRegistrations.json")
	appRegExpiringCredsStr, _ := json.MarshalIndent(appRegExpiringCreds, "", "  ")
	os.WriteFile(cachePath+"/appRegExpiringCreds.json", appRegExpiringCredsStr, 0644)
	defer os.Remove(cachePath + "/appRegExpiringCreds.json")

	fmt.Println("Updating App Registrations in database...")
	s.Start()
	UpsertMultipleEntraApps(allAppRegistrations, opts.EntraAppRegColl)
	s.Stop()

	fmt.Println("Updating App Registrations with expired or expiring credentials in database...")
	s.Start()
	DeleteAllDocumentsInCollection(opts.EntraAppRegCredsExpiringColl)
	UpsertMultipleEntraApps(appRegExpiringCreds, opts.EntraAppRegCredsExpiringColl)
	s.Stop()

}

//
//

func UpdateEntraPimItems(opts UpdateEntraPimItemsOptions) {
	s := spinner.New(spinner.CharSets[43], 100*time.Millisecond)

	fmt.Println("Fetching all resource PIM assignments and eligibilities..")
	s.Start()
	assignments, eligibilities := azure.ListAllTenantPIMScheduleInstancesForAllTenants()
	s.Stop()

	fmt.Println("Updating resource PIM assignments in database...")
	s.Start()
	UpsertMultipleRoleAssignmentScheduleInstances(assignments, opts.EntraRoleAssignmentScheduleInstancesColl)
	s.Stop()

	fmt.Println("Updating resource PIM eligibilities in database...")
	s.Start()
	UpsertMultipleRoleEligibilityScheduleInstances(eligibilities, opts.EntraRoleEligibilityScheduleInstancesColl)
	s.Stop()
}

//
//

func GetBuildDataAndUpdateImageVesionData(imageGalleryImagesColl *mongo.Collection) {
	_, _, cachePath := lib.InitConfig(nil)
	dlPath := cachePath + "/aib-logs"

	ado.DownloadPackerHostLogs(&dlPath)
	buildData := lib.GetDataFromMultiplePackerLogFiles(dlPath)
	UpdateImageDataWithBuildHostLogs(buildData, imageGalleryImagesColl)
}
