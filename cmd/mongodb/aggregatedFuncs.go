package mongodb

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/briandowns/spinner"
	"github.com/jercle/cloudini/cmd/ad"
	"github.com/jercle/cloudini/cmd/ado"
	"github.com/jercle/cloudini/cmd/azure"
	"github.com/jercle/cloudini/cmd/citrix"
	"github.com/jercle/cloudini/cmd/m365"
	"github.com/jercle/cloudini/cmd/web"
	"github.com/jercle/cloudini/lib"
	"go.mongodb.org/mongo-driver/bson"
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

	fmt.Println("Updating build and image version data in database...")
	s.Start()
	GetBuildDataAndUpdateImageVesionData(imageGalleryImagesColl)
	s.Stop()
}

//
//

func UpdateAllAzureResourceIPAddresses(ipAddressesColl *mongo.Collection, ipAddressBlocksColl *mongo.Collection, tokenReq lib.AllTenantTokens) {
	s := spinner.New(spinner.CharSets[43], 100*time.Millisecond)

	opts := lib.GetAllResourcesForAllConfiguredTenantsOptions{
		SuppressSteps: true,
		// SelectedIPAddressQueries: &[]string{"GetIPAddressesQueryVirtualMachines"},
	}

	config := lib.GetCldConfig(nil)
	cidrsToCheck := config.Azure.Ipam.CidrBlocksToCheck

	fmt.Println("Fetching all resource IPs...")
	// s.Start()
	resources := azure.GetAllVMIpAddrForAllConfiguredTenants(&opts, tokenReq)

	// s.Stop()

	fmt.Println("Updating all resource IPs in database...")
	s.Start()
	UpsertAzureIPAddresses(resources, ipAddressesColl)
	s.Stop()

	fmt.Println("Getting IP Address Blocks...")
	s.Start()
	var vnets []azure.IPAddressesAllResourceTypes
	for _, res := range resources {
		if res.Type != "microsoft.network/virtualnetworks" {
			continue
		}
		vnets = append(vnets, res)
	}
	ipAddressBlocks := azure.GetIpAddressBlocksForCidrFromVNets(cidrsToCheck, vnets)
	s.Stop()

	fmt.Println("Updating IP Address Blocks in database...")
	s.Start()
	UpdateIpamAddressBlocks(ipAddressBlocks, ipAddressBlocksColl)
	s.Stop()

}

//
//

func UpdateAllAzureResourcesVcpuCountsCostData(opts UpdateAllAzureResourcesAndVcpuCountsOptions, tokenReq lib.AllTenantTokens) lib.AggregatedCostData {
	if opts.AzResGrpsListColl == nil {
		fmt.Println("AzResGrpsListColl == nil")
		os.Exit(1)
	}
	if opts.AzResResourceListColl == nil {
		fmt.Println("AzResResourceListColl == nil")
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
	// resourceSKUsStr, _ := json.Marshal(resourceSKUs)
	// os.WriteFile("resourceSKUs2.json", resourceSKUsStr, 0644)
	// fmt.Println("saved resourceSKUs2.json")
	// os.Exit(0)
	fmt.Println("Updating Azure Resource SKUs in database...")
	s.Start()
	UpsertResourceSKUs(resourceSKUs, opts.AzResSKUColl)
	s.Stop()

	fmt.Println("Getting full list of SKUs from database...")
	s.Start()
	resourceSKUs = GetResourceSKUs(opts.AzResSKUColl)
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

	fmt.Println("")
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
		opts.AzResResourceListColl,
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
	opts.ConfiguredTenantName = config.CertificateManagement.StorageAccountTenantName
	opts.ContainerName = config.CertificateManagement.ContainerName
	opts.DownloadPath = cachePath + "/cert-sync"
	opts.StorageAccountName = config.CertificateManagement.StorageAccountName
	opts.OverwriteExisting = true
	// opts.GetWriteToken = true

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

func UpdateB2CUsers(coll *mongo.Collection) {
	s := spinner.New(spinner.CharSets[43], 100*time.Millisecond)
	fmt.Println("Getting users from all configured B2C tenants")

	s.Start()
	users := azure.GetAllB2CTenantUsers()
	s.Stop()

	fmt.Println("Updating " + strconv.Itoa(len(users)) + " B2C users in database")
	s.Start()
	coll.DeleteMany(context.TODO(), bson.D{{}})
	UpsertB2CUsers(users, coll)
	s.Stop()
}

//
//

func UpdateM365Data(m365MailboxStatisticsColl *mongo.Collection, m365LicenseCountsColl *mongo.Collection) {
	s := spinner.New(spinner.CharSets[43], 100*time.Millisecond)

	fmt.Println("Getting mailbox stats")
	s.Start()
	data := m365.GetMailboxStorageUsedAllConfiguredTenants()
	s.Stop()

	fmt.Println("Upserting mailbox stats")
	s.Start()
	UpsertMailboxStatistics(data, m365MailboxStatisticsColl)
	s.Stop()

	fmt.Println("Getting license counts")
	// s.Start()
	licenseCounts := m365.GetM365LicenseCountsForAllConfiguredTenants(nil)
	// s.Stop()

	fmt.Println("Upserting license counts")
	s.Start()
	InsertM365LicenseCounts(licenseCounts, m365LicenseCountsColl)
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

func UpdateWebsiteCertsPullingFromDatabase(c *mongo.Client) {
	s := spinner.New(spinner.CharSets[43], 100*time.Millisecond)
	fmt.Println("Getting configuration")

	s.Start()

	coll := c.Database("appConfig").Collection("adminSettings")

	filter := bson.D{
		{"_id", "certManagement"},
	}

	config := coll.FindOne(context.TODO(), filter)
	s.Stop()

	var cfg CertManagementConfig
	err := config.Decode(&cfg)
	lib.CheckFatalError(err)
	// lib.JsonMarshalAndPrint(cfg)

	for _, cert := range cfg.UrlsToWatch {
		// fmt.Println(cert.URL)

		cert := web.GetWebsiteCertificate(cert.URL, nil)
		lib.JsonMarshalAndPrint(cert)
	}
	// fmt.Println(config)
	// lib.JsonMarshalAndPrint(config)

	// fmt.Println("Getting mailbox stats")
	// s.Start()
	// data := m365.GetMailboxStorageUsedAllConfiguredTenants()
	// s.Stop()

	// fmt.Println("Upserting mailbox stats")
	// s.Start()
	// UpsertMailboxStatistics(data, coll)
	// s.Stop()

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

func UpdateSupportAlerts(coll *mongo.Collection) {
	s := spinner.New(spinner.CharSets[43], 100*time.Millisecond)

	config := lib.GetCldConfig(nil)
	tenants := config.Azure.MultiTenantAuth.Tenants

	saConf := config.Azure.SupportAlerts

	saToken, err := azure.GetTenantSPToken(lib.AzureMultiAuthTokenRequestOptions{
		TenantName: saConf.DefaultTenant,
	}, nil)
	lib.CheckFatalError(err)
	// fmt.Println(saConf.WorkbookId)
	// os.Exit(0)
	workbookId := saConf.TenantWorkbookIds[saConf.DefaultTenant]
	supportAlertsQuery := azure.GetLogAnalyticsWorkbookQuery(workbookId, saToken)
	// fmt.Println(supportAlertsQuery)
	// os.Exit(0)

	fmt.Println("Getting support alert data")
	s.Start()
	var allAlerts []azure.AzureAlertProcessed
	for tName, tData := range tenants {
		if tData.GetWorkbookAlerts {
			token, err := azure.GetTenantSPToken(lib.AzureMultiAuthTokenRequestOptions{
				TenantName: tName,
			}, nil)
			lib.CheckFatalError(err)
			data := azure.GetAzureWorkbookAlerts(supportAlertsQuery, token)
			allAlerts = append(allAlerts, data...)
		}
	}
	s.Stop()

	fmt.Println("Upserting support alert data")
	s.Start()
	coll.Drop(context.TODO())
	UpsertSupportAlerts(allAlerts, coll)
	s.Stop()
}

//
//

func UpdateAllAzureResources(opts UpdateAllAzureResourcesAndVcpuCountsOptions, tokenReq lib.AllTenantTokens) {

	_, _, cachePath := lib.InitConfig(nil)
	s := spinner.New(spinner.CharSets[43], 100*time.Millisecond)

	resSkuOpts := lib.GetAllResourcesForAllConfiguredTenantsOptions{
		SubscriptionId: opts.SkuListSubscription,
		AzureAuth:      opts.SkuListAuth,
		Location:       opts.Location,
		SuppressSteps:  true,
	}

	fmt.Println("Fetching all Azure Resources...")
	s.Start()
	startTime := time.Now()
	allResources, allResourcesSlice := azure.GetAllResourcesForAllConfiguredTenants(&resSkuOpts, tokenReq)
	s.Stop()
	fmt.Println(strconv.Itoa(len(allResourcesSlice)) + " resources found")
	elapsed := time.Since(startTime)
	fmt.Println(elapsed)
	allResourcesSliceStr, _ := json.MarshalIndent(allResourcesSlice, "", "  ")
	os.WriteFile(cachePath+"/allResourcesSlice.json", allResourcesSliceStr, 0644)
	allResourcesStr, _ := json.MarshalIndent(allResources, "", "  ")
	os.WriteFile(cachePath+"/allResources.json", allResourcesStr, 0644)
	// os.Exit(0)

	fmt.Println("Updating Azure Resources in database...")
	s.Start()
	startTime = time.Now()
	UpsertMultipleResources(allResourcesSlice, opts.AzResResourceListColl)
	s.Stop()
	elapsed = time.Since(startTime)
	fmt.Println(elapsed)

	fmt.Println("")
	fmt.Println("Updating 'existsInAzure' value for all resources in database...")
	s.Start()
	startTime = time.Now()
	UpdateResourcesNotExistInAzure(allResourcesSlice, opts.AzResResourceListColl)
	elapsed = time.Since(startTime)
	s.Stop()
	fmt.Println(elapsed)
}

//
//

func UpdateAWSMonitoringData(coll *mongo.Collection) {
	s := spinner.New(spinner.CharSets[43], 100*time.Millisecond)

	config := lib.GetCldConfig(nil)
	laQuery := config.AWS.LogIngestCountQuery
	tenants := config.Azure.MultiTenantAuth.Tenants

	fmt.Println("Getting and upserting AWS monitoring data")
	startTime := time.Now()

	s.Start()
	for tName, tData := range tenants {
		if tData.AWSIngestWorkspaceID != "" {
			token, err := azure.GetTenantSPToken(lib.AzureMultiAuthTokenRequestOptions{
				TenantName: tName,
				Scope:      "loganalytics",
			}, nil)
			lib.CheckFatalError(err)

			results := azure.RunLogAnalyticsQuery(tData.AWSIngestWorkspaceID, laQuery, *token)
			// lib.JsonMarshalAndPrint(results)
			ingestCounts := ConvertLAResultToAWSIngestCounts(results, tData.AWSIngestRef)
			// _ = ingestCounts
			// lib.JsonMarshalAndPrint(ingestCounts)
			// _ = results
			// os.Exit(0)
			// fmt.Println(tName)
			UpsertAWSMontoringData(ingestCounts, coll)
		}
	}
	s.Stop()
	elapsed := time.Since(startTime)
	fmt.Println(elapsed)
}

//
//

func ConvertLAResultToAWSIngestCounts(data azure.LogAnalyticsQueryResponse, awsIngestRef string) (ingestCounts AWSIngestCounts) {
	ingestCountRows := data.Tables[0].Rows

	// lib.JsonMarshalAndPrint(data)
	// os.Exit(0)
	for _, row := range ingestCountRows {
		// zeroIfNull := float64(0)
		var percentageOfTotalLogs float64
		if row["PercentageOfTotalLogs"] == nil {
			percentageOfTotalLogs = 0
		} else {
			percentageOfTotalLogs = row["PercentageOfTotalLogs"].(float64)
		}
		count := AWSIngestCount{
			LogType:               row["LogType"].(string),
			PercentageOfTotalLogs: percentageOfTotalLogs,
			Count:                 row["Count"].(float64),
		}
		ingestCounts.Counts = append(ingestCounts.Counts, count)
		ingestCounts.TotalLogs = row["TotalLogs"].(float64)
		ingestCounts.TotalSQSMessages = row["TotalSQSMessages"].(float64)
	}
	ingestCounts.Environment = awsIngestRef
	ingestCounts.Monitor = "ingestCountsLast24hr"
	ingestCounts.ID = "ingestCountsLast24hr" + awsIngestRef
	ingestCounts.LastDBSync = time.Now()
	return
}

//
//

func UpsertP2SVpnConnectionDetails(p2sVpnGatewayResourceId string, tenantName string, storageAccountName string, containerName string) {
	// blobSAS := azure.GenerateP2SVpnConnectionHealthDetailed(p2sVpnGatewayResourceId, tenantName, storageAccountName, containerName)

	blobSAS := "https://stapcsharedautomon.blob.core.windows.net/p2svpn/20251218-17.16.10.json?sv=2025-07-05&spr=https&st=2025-12-18T06%3A27%3A41Z&se=2025-12-19T06%3A27%3A41Z&sr=b&sp=r&sig=EbCKVSkW5v3yT94OJxzBuQydP8s4Vbd9yS7RJWevfXs%3D"

	connections := azure.GetP2SVpnConnectionDetailsFromBlobSAS(blobSAS)

	// lib.JsonMarshalAndPrint(connections)

	var (
		wg  sync.WaitGroup
		mut sync.Mutex
	)
	// var connectionsProcessed []azure.AzureP2SConnectionHealth

	connectionsProcessing := make(map[string]azure.AzureP2SConnectionHealth)

	for _, conn := range connections {
		wg.Go(func() {
			deviceSerial := conn.UserName
			_, deviceDetail, err := azure.GetUserFromDeviceSerial(tenantName, deviceSerial)
			lib.CheckFatalError(err)

			conn.UserPrincipalName = deviceDetail.UserPrincipalName
			conn.ManagedDeviceADDeviceId = deviceDetail.AzureAdDeviceID
			conn.ManagedDeviceIntuneId = deviceDetail.ID
			conn.ManagedDeviceName = deviceDetail.DeviceName
			conn.ManagedDeviceLastSyncDateTime = deviceDetail.LastSyncDateTime
			conn.ManagedDeviceSerial = deviceSerial
			conn.UserName = ""
			mut.Lock()
			// connectionsProcessed = append(connectionsProcessed, conn)
			connectionsProcessing[conn.VpnConnectionID] = conn
			mut.Unlock()
		})
	}
	wg.Wait()
	// lib.JsonMarshalAndPrint(connectionsProcessed)
	lib.JsonMarshalAndPrint(connectionsProcessing)
	// TODO: Update all json.marshalindents in this file

	var connectionsProcessed []azure.AzureP2SConnectionHealth

	for _, v := range connectionsProcessing {
		connectionsProcessed = append(connectionsProcessed, v)
	}
	jsonStr, _ := json.MarshalIndent(connectionsProcessed, "", "  ")
	os.WriteFile("/home/jercle/git/cld/cmd/mongodb/aggregatedFuncs-p2svpn-array.json", jsonStr, 0644)

}

//
//

// func UpdateAllAzureResources(opts UpdateAllAzureResourcesAndVcpuCountsOptions, tokenReq lib.AllTenantTokens) {

// 	_, _, cachePath := lib.InitConfig(nil)
// 	s := spinner.New(spinner.CharSets[43], 100*time.Millisecond)

// 	resSkuOpts := lib.GetAllResourcesForAllConfiguredTenantsOptions{
// 		SubscriptionId: opts.SkuListSubscription,
// 		AzureAuth:      opts.SkuListAuth,
// 		Location:       opts.Location,
// 		SuppressSteps:  true,
// 	}

// 	fmt.Println("Fetching all Azure Resources...")
// 	s.Start()
// 	startTime := time.Now()
// 	allResources, allResourcesSlice := azure.GetAllResourcesForAllConfiguredTenants(&resSkuOpts, tokenReq)
// 	s.Stop()
// 	fmt.Println(strconv.Itoa(len(allResourcesSlice)) + " resources found")
// 	elapsed := time.Since(startTime)
// 	fmt.Println(elapsed)
// 	allResourcesSliceStr, _ := json.MarshalIndent(allResourcesSlice, "", "  ")
// 	os.WriteFile(cachePath+"/allResourcesSlice.json", allResourcesSliceStr, 0644)
// 	allResourcesStr, _ := json.MarshalIndent(allResources, "", "  ")
// 	os.WriteFile(cachePath+"/allResources.json", allResourcesStr, 0644)
// 	// os.Exit(0)

// 	fmt.Println("Updating Azure Resources in database...")
// 	s.Start()
// 	startTime = time.Now()
// 	UpsertMultipleResources(allResourcesSlice, opts.AzResResourceListColl)
// 	s.Stop()
// 	elapsed = time.Since(startTime)
// 	fmt.Println(elapsed)

// 	fmt.Println("")
// 	fmt.Println("Updating 'existsInAzure' value for all resources in database...")
// 	s.Start()
// 	startTime = time.Now()
// 	UpdateResourcesNotExistInAzure(allResourcesSlice, opts.AzResResourceListColl)
// 	elapsed = time.Since(startTime)
// 	s.Stop()
// 	fmt.Println(elapsed)
// }
