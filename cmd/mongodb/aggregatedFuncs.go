package mongodb

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/briandowns/spinner"
	"github.com/jercle/cloudini/cmd/ad"
	"github.com/jercle/cloudini/cmd/ado"
	"github.com/jercle/cloudini/cmd/azure"
	"github.com/jercle/cloudini/cmd/citrix"
	"github.com/jercle/cloudini/cmd/m365"
	"github.com/jercle/cloudini/lib"
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
	DeleteAllDocumentsInCollection(imageGalleryImagesColl)
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
		s.Stop()
		fmt.Println("Updating Citrix Machine Catalogs in database...")
		s.Start()
		UpsertCitrixMachineCatalogs(machineCatalogs, machineCatalogsColl)
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
	resources := azure.GetAllVMIpAddrForAllConfiguredTenants(&opts, tokenReq)
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

	fmt.Println("Getting and upserting all tenant and subscription details...")
	s.Start()
	UpsertTenantAndSubs(opts.AzResTenantsColl, &tokenReq)
	s.Stop()

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
	startTime := time.Now()
	allResources, allResourcesSlice := azure.GetAllResourcesForAllConfiguredTenants(&resSkuOpts, tokenReq)
	s.Stop()
	elapsed := time.Since(startTime)
	fmt.Println(elapsed)
	allResourcesSliceStr, _ := json.MarshalIndent(allResourcesSlice, "", "  ")
	os.WriteFile(cachePath+"/allResourcesSlice.json", allResourcesSliceStr, 0644)
	allResourcesStr, _ := json.MarshalIndent(allResources, "", "  ")
	os.WriteFile(cachePath+"/allResources.json", allResourcesStr, 0644)
	// os.Exit(0)

	fmt.Println("Getting then updating all storage accounts with minimum TLS versions...")
	s.Start()
	startTime = time.Now()
	stgAccountsOptions := lib.GetAllResourcesForAllConfiguredTenantsOptions{
		GetAllStorageAccountsInTlsCheck: true,
		SuppressSteps:                   true,
	}
	stgAccounts := azure.CheckStorageAccountTlsVersionsForAllConfiguredTenants(&stgAccountsOptions, tokenReq)
	UpsertStorageAccountMinTlsVersions(stgAccounts, opts.AzStorageAcctMinTlsVersions)
	s.Stop()
	elapsed = time.Since(startTime)
	fmt.Println(elapsed)

	fmt.Println("Updating Azure Resources in database...")
	s.Start()
	startTime = time.Now()
	UpsertMultipleResources(allResourcesSlice, opts.AzResResourceListColl)
	s.Stop()
	elapsed = time.Since(startTime)
	fmt.Println(elapsed)

	fmt.Println("Updating 'existsInAzure' value for all resources in database...")
	s.Start()
	startTime = time.Now()
	UpdateResourcesNotExistInAzure(allResourcesSlice, opts.AzResResourceListColl)
	elapsed = time.Since(startTime)
	fmt.Println(elapsed)
	s.Stop()

	fmt.Println("Fetching all Azure Resource Groups...")
	s.Start()
	startTime = time.Now()
	allResGrps := azure.GetAllResGrpsForAllConfiguredTenants(&resSkuOpts, tokenReq)
	s.Stop()
	elapsed = time.Since(startTime)
	fmt.Println(elapsed)
	fmt.Println("Updating Azure Resource Groups in database...")
	s.Start()
	startTime = time.Now()
	UpsertMultipleResGrps(allResGrps, opts.AzResGrpsListColl)
	s.Stop()
	elapsed = time.Since(startTime)
	fmt.Println(elapsed)

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
	// fmt.Println("Processing all resources and cost meters to create relations")
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

//
//

func UpdateAllCertInfo(certsCaCertInfo *mongo.Collection, serverCertsInfoColl *mongo.Collection) {
	s := spinner.New(spinner.CharSets[43], 100*time.Millisecond)

	// azure.DownloadAllBlobsInContainer()
	// lib.GetServerCertInfoFromFile()
	_, _, cachePath := lib.InitConfig(nil)
	config := lib.GetCldConfig(nil)
	// dlPath := cachePath + "/aib-logs"
	var opts lib.StorageAccountRequestOptions
	opts.ConfiguredTenantName = "REDDTQ"
	opts.ContainerName = "cert-sync"
	opts.DownloadPath = cachePath + "/cert-sync"
	opts.StorageAccountName = config.AzureDevOps.Packer.Logs.StorageAcct
	opts.OverwriteExisting = true
	opts.GetWriteToken = true

	fmt.Println("Fetching certs from storage")
	s.Start()
	azure.DownloadAllBlobsInContainer(opts)
	s.Stop()

	fmt.Println("Getting cert info from downloaded files")
	startTime := time.Now()
	s.Start()
	caCertInfo, serverCertInfo := lib.GetCertInfoFromFiles(cachePath+"/cert-sync", cachePath+"/cert-sync-processed")
	s.Stop()
	elapsed := time.Since(startTime)
	fmt.Println(elapsed)

	fmt.Println("Relating server certs to CA requests")
	startTime = time.Now()
	s.Start()
	caCertInfoRelated, serverCertInfoRelated := lib.RelateCertAuthCertsToServerCerts(caCertInfo, serverCertInfo)
	s.Stop()
	elapsed = time.Since(startTime)
	fmt.Println(elapsed)

	certsCount := make(map[string]int)
	var multi []lib.ServerCertInfo
	for _, cert := range serverCertInfoRelated {
		certsCount[cert.ID]++

		if cert.ID == "78b5e4cd429a487ca2e7f4341ca26525" {
			multi = append(multi, cert)
		}
		// os.Exit(0)
	}
	// lib.JsonMarshalAndPrint(certsCount)
	// lib.JsonMarshalAndPrint(multi)
	// os.Exit(0)

	fmt.Println("Clearing collections")
	// clearOpts := options.DeleteOptions{}
	err := serverCertsInfoColl.Drop(context.TODO())
	lib.CheckFatalError(err)
	err = certsCaCertInfo.Drop(context.TODO())
	lib.CheckFatalError(err)
	// _, err := serverCertsInfoColl.DeleteMany(context.TODO(), bson.D{{}}, nil)
	// lib.CheckFatalError(err)
	// lib.JsonMarshalAndPrint(delResult)
	// os.Exit(0)

	fmt.Println("Upserting CA cert info")
	s.Start()
	startTime = time.Now()
	// caCertUpdates :=
	UpsertCACertificates(caCertInfoRelated, certsCaCertInfo)
	s.Stop()
	elapsed = time.Since(startTime)
	fmt.Println(elapsed)
	// os.Exit(0)

	fmt.Println("Upserting Server cert info")
	s.Start()
	startTime = time.Now()
	// serverCertUpdates :=
	UpsertServerCertificates(serverCertInfoRelated, serverCertsInfoColl)
	s.Stop()
	elapsed = time.Since(startTime)
	fmt.Println(elapsed)
	// jsonStr, _ := json.MarshalIndent(serverCertUpdates, "", "  ")
	// fmt.Println(string(jsonStr))
	// lib.MarshalAndPrintJson(caCertUpdates)
	// lib.MarshalAndPrintJson(serverCertUpdates)
	// os.RemoveAll(cachePath + "/cert-sync")
	// os.RemoveAll(cachePath + "/cert-sync-processed")
	// ado.DownloadPackerHostLogs(&dlPath)
	// buildData := lib.GetDataFromMultiplePackerLogFiles(dlPath)
	// UpdateImageDataWithBuildHostLogs(buildData, imageGalleryImagesColl)
}

//
//

func UpdateADUsers(coll *mongo.Collection) {

	config := lib.GetCldConfig(nil)

	adConf := config.ActiveDirectory

	users := ad.GetAllADUsersForAllConfiguredDomains(*adConf)

	UpsertADUsers(users, coll)
	// jsonStr, _ := json.MarshalIndent(serverCertUpdates, "", "  ")
	// fmt.Println(string(jsonStr))
	// lib.MarshalAndPrintJson(caCertUpdates)
	// lib.MarshalAndPrintJson(serverCertUpdates)
	// os.RemoveAll(cachePath + "/cert-sync")
	// os.RemoveAll(cachePath + "/cert-sync-processed")
	// ado.DownloadPackerHostLogs(&dlPath)
	// buildData := lib.GetDataFromMultiplePackerLogFiles(dlPath)
	// UpdateImageDataWithBuildHostLogs(buildData, imageGalleryImagesColl)
}

//
//

func UpdateM365Data(coll *mongo.Collection) {
	s := spinner.New(spinner.CharSets[43], 100*time.Millisecond)

	fmt.Println("Getting mailbox stats")
	s.Start()
	data := m365.GetMailboxStorageUsedAllConfiguredTenants()
	s.Stop()

	fmt.Println("Upserting mailbox stats")
	s.Start()
	UpsertMailboxStatistics(data, coll)
	s.Stop()

	// jsonStr, _ := json.MarshalIndent(serverCertUpdates, "", "  ")
	// fmt.Println(string(jsonStr))
	// lib.MarshalAndPrintJson(caCertUpdates)
	// lib.MarshalAndPrintJson(serverCertUpdates)
	// os.RemoveAll(cachePath + "/cert-sync")
	// os.RemoveAll(cachePath + "/cert-sync-processed")
	// ado.DownloadPackerHostLogs(&dlPath)
	// buildData := lib.GetDataFromMultiplePackerLogFiles(dlPath)
	// UpdateImageDataWithBuildHostLogs(buildData, imageGalleryImagesColl)
}
