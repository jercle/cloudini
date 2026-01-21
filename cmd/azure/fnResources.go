package azure

import (
	"context"
	"encoding/json/jsontext"
	json "encoding/json/v2"

	// "encoding/json"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v7"
	"github.com/rmasci/ipsubnet"

	"github.com/jercle/cloudini/lib"
)

func GetAllResGrpsForAllConfiguredTenants(opts *lib.GetAllResourcesForAllConfiguredTenantsOptions, tokens lib.AllTenantTokens) (allResGrps []ResourceGroup) {
	var (
		wg    sync.WaitGroup
		mutex sync.Mutex
	)

	azConfig := lib.GetCldConfig(nil)
	azConfigs := azConfig.Azure.MultiTenantAuth.Tenants

	// resSkus := GetAzureResourceSKUsForSubscription(*opts)
	// vcpuSkus := GetSkusWithVcpus(resSkus)

	for _, token := range tokens {
		if azConfigs[token.TenantName].IsB2C {
			continue
		}
		wg.Add(1)
		go func() {
			defer wg.Done()
			if !opts.SuppressSteps {
				fmt.Println(token.TenantName + ": Fetching resources")
			}
			tenantResGrps := GetAllTenantResourceGroups("", &token)

			mutex.Lock()
			allResGrps = append(allResGrps, tenantResGrps...)
			mutex.Unlock()
		}()
	}

	wg.Wait()

	options := *opts

	outputFilePath := options.OutputFilePath

	if outputFilePath != "" {
		jsonListStr, _ := json.Marshal(allResGrps, jsontext.WithIndent("  "))

		currentDate := time.Now().Format("20060102")

		arrayFileName := outputFilePath + "/allRes-GraphResGrps-FlatArray-" + currentDate + ".json"

		err := os.WriteFile(arrayFileName, jsonListStr, 0644)
		lib.CheckFatalError(err)
		fmt.Println("Saved to " + arrayFileName)
	}

	return allResGrps
}

func GetAllResourcesForAllConfiguredTenants(opts *lib.GetAllResourcesForAllConfiguredTenantsOptions, tokens lib.AllTenantTokens) (allResources map[string]TenantResourceList, allResourcesSlice []lib.AzureResourceDetails) {
	var (
		wg    sync.WaitGroup
		mutex sync.Mutex
	)

	cldConfig := lib.GetCldConfig(nil)
	azConfigs := cldConfig.Azure.MultiTenantAuth.Tenants

	allResources = make(map[string]TenantResourceList)

	resSkus := GetAzureResourceSKUsForSubscription(*opts)
	vcpuSkus := GetSkusWithVcpus(resSkus)

	for _, token := range tokens {
		if azConfigs[token.TenantName].IsB2C {
			continue
		}
		wg.Add(1)
		go func() {
			defer wg.Done()
			if !opts.SuppressSteps {
				fmt.Println(token.TenantName + ": Fetching resources")
			}
			// allResources[token.TenantName] = make(map[string]SubscriptionResourceList)
			tenantResources := GetAllTenantResources("", &token)
			if !opts.SuppressSteps {
				fmt.Println(token.TenantName + ": Fetch complete")
			}
			// var processedTenantResources TenantResourceList
			for sub, subResources := range tenantResources.Subscriptions {

				var currSubResources SubscriptionResourceList
				for _, resource := range subResources.Resources {
					// if resource.Type == "microsoft.network/virtualnetworks/subnets" {
					// 	lib.JsonMarshalAndPrint(resource)
					// 	os.Exit(0)
					// }
					if resource.Type == "microsoft.compute/virtualmachines" {
						mappedDetails := MapVmSizeDetails(resource.Properties.HardwareProfile.VmSize, vcpuSkus)
						mappedDetailsStr, _ := json.Marshal(mappedDetails)
						json.Unmarshal(mappedDetailsStr, &resource.Properties.HardwareProfile.VmSizeSku)
					}

					currRes := resource
					currRes.TenantName = token.TenantName
					currRes.SubscriptionName = sub
					currRes.LastAzureSync = time.Now()
					currRes.ExistsInAzure = true
					currSubResources.Resources = append(currSubResources.Resources, currRes)
					mutex.Lock()
					allResourcesSlice = append(allResourcesSlice, currRes)
					mutex.Unlock()

				}
				// fmt.Println(currSubResources.Resources[0])
				// fmt.Println(subResources.Resources[0])
				currSubResources.ResourceCount = tenantResources.Subscriptions[sub].ResourceCount
				tenantResources.Subscriptions[sub] = currSubResources
				// os.Exit(0)
			}
			mutex.Lock()
			allResources[token.TenantName] = tenantResources
			mutex.Unlock()
			if !opts.SuppressSteps {
				fmt.Println(token.TenantName + ": Processing complete")
			}
		}()
	}

	wg.Wait()

	options := *opts

	outputFilePath := options.OutputFilePath

	if outputFilePath != "" {
		jsonStr, _ := json.Marshal(allResources, jsontext.WithIndent("  "))
		jsonListStr, _ := json.Marshal(allResourcesSlice, jsontext.WithIndent("  "))

		currentDate := time.Now().Format("20060102")

		mapFileName := outputFilePath + "/allRes-GraphResources-ByTenantAndSubscription-" + currentDate + ".json"
		arrayFileName := outputFilePath + "/allRes-GraphResources-FlatArray-" + currentDate + ".json"

		err := os.WriteFile(mapFileName, jsonStr, 0644)
		lib.CheckFatalError(err)
		err = os.WriteFile(arrayFileName, jsonListStr, 0644)
		lib.CheckFatalError(err)
		fmt.Println("Saved to " + mapFileName + " and " + arrayFileName)
	}

	fmt.Println(len(allResourcesSlice))

	// fmt.Println(len(allResourcesSlice))

	return allResources, allResourcesSlice
}

//
//

func GetAzureResourceSKUsForSubscription(opts lib.GetAllResourcesForAllConfiguredTenantsOptions) (processedSkus []lib.AzureResourceSku) {
	cred, err := azidentity.NewClientSecretCredential(opts.AzureAuth.TenantID, opts.AzureAuth.Writer.ClientID, opts.AzureAuth.Writer.ClientSecret, nil)
	lib.CheckFatalError(err)
	var skus []lib.AzureResourceSku

	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory(opts.SubscriptionId, cred, nil)
	lib.CheckFatalError(err)

	pager := clientFactory.NewResourceSKUsClient().NewListPager(&armcompute.ResourceSKUsClientListOptions{
		Filter:                   to.Ptr("location eq '" + opts.Location + "'"),
		IncludeExtendedLocations: nil,
	})

	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			var val lib.AzureResourceSku
			jsonStr, _ := json.Marshal(v)
			json.Unmarshal(jsonStr, &val)
			skus = append(skus, val)
		}
	}

	for _, sku := range skus {
		curr := sku
		curr.LastAzureSync = time.Now()

		for _, cap := range sku.Capabilities {
			if cap.Name == "vCPUs" {
				vmvcpus, err := strconv.Atoi(cap.Value)
				lib.CheckFatalError(err)
				curr.VMvCPUs = vmvcpus
			}
			if cap.Name == "Cores" {
				cores, err := strconv.Atoi(cap.Value)
				lib.CheckFatalError(err)
				curr.VMCores = cores
			}
			if cap.Name == "vCPUsPerCore" {
				vCPUsPerCore, err := strconv.Atoi(cap.Value)
				lib.CheckFatalError(err)
				curr.VMvCPUsPerCore = vCPUsPerCore
			}
		}
		processedSkus = append(processedSkus, curr)
	}

	return
}

//
//

func GetSkusWithVcpus(resSkus []lib.AzureResourceSku) (vcpuSkus []lib.AzureVirtualMachineSku) {
	for _, sku := range resSkus {
		// currRes := lib.StructToMap(sku)
		var curr lib.AzureVirtualMachineSku
		jsonStr, _ := json.Marshal(sku)
		err := json.Unmarshal(jsonStr, &curr)
		lib.CheckFatalError(err)

		vcpuCapabilityExists := false

		for _, cap := range sku.Capabilities {
			if cap.Name == "vCPUs" {
				curr.VCPUs = cap.Value
				vcpuCapabilityExists = true
			}
			if cap.Name == "vCPUsPerCore" {
				curr.VCPUsPerCore = cap.Value
				vcpuCapabilityExists = true
			}
			if cap.Name == "Cores" {
				curr.Cores = cap.Value
				vcpuCapabilityExists = true
			}
			if cap.Name == "vCPUsAvailable" {
				curr.VCPUsAvailable = cap.Value
				vcpuCapabilityExists = true
			}
		}

		if vcpuCapabilityExists {
			vcpuSkus = append(vcpuSkus, curr)
		}
	}
	return
}

//
//

func GetAllTenantResources(outputFile string, token *lib.AzureMultiAuthToken) TenantResourceList {
	subscriptions, err := ListSubscriptions(*token)
	lib.CheckFatalError(err)

	var allResources []lib.AzureResourceDetails

	var allTenantResources TenantResourceList
	allTenantResources.Subscriptions = make(map[string]SubscriptionResourceList)

	subIds := []string{}
	subIdsByNameMap := make(map[string]string)

	for _, sub := range subscriptions {
		subIds = append(subIds, sub.SubscriptionID)
		subIdsByNameMap[sub.SubscriptionID] = sub.DisplayName
		allTenantResources.Subscriptions[sub.DisplayName] = SubscriptionResourceList{}
	}

	urlString := "https://management.azure.com/providers/Microsoft.ResourceGraph/resources?api-version=2022-10-01"

	graphQuery := "Resources | extend cpu = properties.hardwareProfile"

	subIdsStr := ""
	for i, sub := range subIds {
		if i < len(subIds)-1 {
			subIdsStr += "		\"" + sub + "\",\n"
		} else {
			subIdsStr += "		\"" + sub + "\""
		}
	}

	jsonBody := `{"query": "` + graphQuery + `"}`

	res, _, err := HttpPost(urlString, jsonBody, *token)
	if err != nil {
		lib.JsonMarshalAndPrint(token.TenantName)
		lib.CheckFatalError(err)
	}
	// lib.CheckFatalError(err)

	var response ResourceGraphResponse
	err = json.Unmarshal(res, &response)

	if err != nil {
		_, _, cachePath := lib.InitConfig(nil)
		errSplit := strings.Split(err.Error(), "/")
		lib.PrintSliceStringsWithIndexes(errSplit)
		errIndex, subErr := strconv.Atoi(errSplit[2])
		lib.CheckFatalError(subErr)
		var resourcesInterface ResourceGraphResponseDataInterface
		subErr = json.Unmarshal(res, &resourcesInterface)
		errObjStr, _ := json.Marshal(resourcesInterface.Data[errIndex], jsontext.WithIndent("  "))

		os.WriteFile(cachePath+"/allResResponse-ErroredObj.json", errObjStr, 0644)

		// fmt.Println(err)
		lib.CheckFatalError(err)
	}

	for _, res := range response.Data {
		currRes := res
		currRes.ID = strings.ToLower(res.ID)
		currRes.Type = strings.ToLower(res.Type)
		if currRes.Type == "microsoft.network/virtualnetworks" {
			vnet, subnets := BuildVnetAndSubnet(currRes)
			allResources = append(allResources, vnet)
			allResources = append(allResources, subnets...)
		} else {
			allResources = append(allResources, currRes)
		}
	}

	hasSkipToken := false
	skipToken := ""

	if response.SkipToken != "" {
		hasSkipToken = true
		skipToken = response.SkipToken
	}

	for hasSkipToken {
		var whileRes ResourceGraphResponse
		jsonBody := `{
			"subscriptions": [
		` + subIdsStr + `
			],
			"query": "` + graphQuery + `",
			"options": {
				"$skipToken": "` + skipToken + `"
			}
		}`

		res, _, err := HttpPost(urlString, jsonBody, *token)
		lib.CheckFatalError(err)
		err = json.Unmarshal(res, &whileRes)
		lib.CheckFatalError(err)

		for _, res := range whileRes.Data {
			currRes := res
			currRes.ID = strings.ToLower(res.ID)
			if currRes.Type == "microsoft.network/virtualnetworks" {
				vnet, subnets := BuildVnetAndSubnet(currRes)
				allResources = append(allResources, vnet)
				allResources = append(allResources, subnets...)
			} else {
				allResources = append(allResources, currRes)
			}
		}

		if whileRes.SkipToken != "" {
			hasSkipToken = true
			skipToken = whileRes.SkipToken
		} else {
			hasSkipToken = false
			skipToken = ""
		}
	}

	for _, res := range allResources {
		subName, ok := subIdsByNameMap[res.SubscriptionID]
		if !ok {
			fmt.Println("Subscription not found in list of authenticated subs")
			fmt.Println(res)
			os.Exit(0)
		}
		subResList := allTenantResources.Subscriptions[subName]
		subResList.Resources = append(subResList.Resources, res)
		allTenantResources.Subscriptions[subName] = subResList
	}

	for sub, _ := range allTenantResources.Subscriptions {
		subResList := allTenantResources.Subscriptions[sub]
		subResList.ResourceCount = len(subResList.Resources)
		allTenantResources.Subscriptions[sub] = subResList
	}

	if outputFile != "" {
		jsonStr, _ := json.Marshal(allTenantResources, jsontext.WithIndent("  "))

		err = os.WriteFile(outputFile, jsonStr, 0644)
		lib.CheckFatalError(err)
		fmt.Println("Saved to " + outputFile)
	}

	// var allTenantResources TenantResourceList
	for _, resources := range allTenantResources.Subscriptions {
		allTenantResources.ResourceCount += resources.ResourceCount
	}
	// fmt.Println(allTenantResources.ResourceCount)
	// os.Exit(0)
	// allTenantResources.ResourceCount = len(allTenantResourcesBySub)
	// allTenantResources.resources

	return allTenantResources
}

//
//

func BuildVnetAndSubnet(res lib.AzureResourceDetails) (vnet lib.AzureResourceDetails, subnets []lib.AzureResourceDetails) {
	vnet = res

	snets := vnet.Properties.Other["subnets"].([]interface{})

	for _, snet := range snets {
		// snet := snetRaw.(lib.AzureResourceDetails)
		var snetResource lib.AzureResourceDetails
		jsonStr, _ := json.Marshal(snet)
		err := json.Unmarshal(jsonStr, &snetResource)
		lib.CheckFatalError(err)

		snetResource.TenantID = res.TenantID
		snetResource.TenantName = res.TenantName
		snetResource.SubscriptionID = res.SubscriptionID
		snetResource.SubscriptionName = res.SubscriptionName
		snetResource.ResourceGroup = res.ResourceGroup
		snetResource.ID = strings.ToLower(snetResource.ID)
		snetResource.Type = strings.ToLower(snetResource.Type)
		snetResource.Properties.ParentVnet = res.ID

		snetResource.Properties.IpAddressesUsed = len(snetResource.Properties.IpConfigurations)
		vnet.Properties.IpAddressesUsed += snetResource.Properties.IpAddressesUsed

		cidrRangeRaw, ok := snetResource.Properties.Other["addressPrefix"].(interface{})
		if ok {
			cidrRange, isOk := cidrRangeRaw.(string)
			if isOk {
				snetResource.Properties.IpCidrBlock = cidrRange
				cidrNotation, err := strconv.Atoi(strings.Split(cidrRange, "/")[1])
				lib.CheckFatalError(err)
				cidrIp := strings.Split(cidrRange, "/")[0]
				snetCalc := ipsubnet.SubnetCalculator(cidrIp, int64(cidrNotation))

				snetResource.Properties.IpNumberAddresses = int(snetCalc.GetNumberIPAddresses())
				snetResource.Properties.IpNumberAddressableHosts = int(snetCalc.GetNumberAddressableHosts())
				snetResource.Properties.IpNumberAddressableHostsRemaining = snetResource.Properties.IpNumberAddressableHosts - snetResource.Properties.IpAddressesUsed
				snetResource.Properties.IpPercentAddressableHostsUsed = float64(snetResource.Properties.IpAddressesUsed) / float64(snetResource.Properties.IpNumberAddressableHosts) * 100
				snetResource.Properties.IpRange = snetCalc.GetIPAddressRange()
				// fmt.Println("GetHostPortion:", snetCalc.GetHostPortion())
				// fmt.Println("GetNetworkPortion:", snetCalc.GetNetworkPortion())
				// fmt.Println("GetNetworkSize:", snetCalc.GetNetworkSize())
				// fmt.Println("GetIPAddressRange:", snetCalc.GetIPAddressRange())
				// fmt.Println("GetIPAddress:", snetCalc.GetIPAddress())
				delete(snetResource.Properties.Other, "addressPrefix")
			}
		} else {
			cidrRanges, rangesOk := snetResource.Properties.Other["addressPrefixes"].([]interface{})[0].(string)
			if rangesOk {
				snetResource.Properties.IpCidrBlock = cidrRanges
				cidrNotation, err := strconv.Atoi(strings.Split(cidrRanges, "/")[1])
				lib.CheckFatalError(err)
				cidrIp := strings.Split(cidrRanges, "/")[0]
				snet := ipsubnet.SubnetCalculator(cidrIp, int64(cidrNotation))

				snetResource.Properties.IpNumberAddresses = int(snet.GetNumberIPAddresses())
				snetResource.Properties.IpNumberAddressableHosts = int(snet.GetNumberAddressableHosts())
				snetResource.Properties.IpNumberAddressableHostsRemaining = snetResource.Properties.IpNumberAddressableHosts - snetResource.Properties.IpAddressesUsed
				snetResource.Properties.IpPercentAddressableHostsUsed = float64(snetResource.Properties.IpAddressesUsed) / float64(snetResource.Properties.IpNumberAddressableHosts) * 100
				delete(snetResource.Properties.Other, "addressPrefixes")
			}

		}

		subnets = append(subnets, snetResource)
		vnet.Properties.SubnetIds = append(vnet.Properties.SubnetIds, snetResource.ID)
	}

	cidrRange, ok := vnet.Properties.Other["addressSpace"].(map[string]interface{})["addressPrefixes"].([]interface{})[0].(string)
	if ok {
		vnet.Properties.IpCidrBlock = cidrRange
		cidrNotation, err := strconv.Atoi(strings.Split(cidrRange, "/")[1])
		lib.CheckFatalError(err)
		cidrIp := strings.Split(cidrRange, "/")[0]
		snet := ipsubnet.SubnetCalculator(cidrIp, int64(cidrNotation))

		vnet.Properties.IpNumberAddresses = int(snet.GetNumberIPAddresses())
		vnet.Properties.IpNumberAddressableHosts = int(snet.GetNumberAddressableHosts())
		vnet.Properties.IpNumberAddressableHostsRemaining = vnet.Properties.IpNumberAddressableHosts - vnet.Properties.IpAddressesUsed
		vnet.Properties.IpPercentAddressableHostsUsed = float64(vnet.Properties.IpAddressesUsed) / float64(vnet.Properties.IpNumberAddressableHosts) * 100
		delete(vnet.Properties.Other, "addressSpace")
	}
	delete(vnet.Properties.Other, "subnets")
	return
}

//
//

func GetAllTenantResourceGroups(outputFile string, token *lib.AzureMultiAuthToken) (allResGrps []ResourceGroup) {
	urlString := "https://management.azure.com/providers/Microsoft.ResourceGraph/resources?api-version=2022-10-01"
	graphQuery := "resourcecontainers | where type == 'microsoft.resources/subscriptions/resourcegroups'"

	jsonBody := `{
	"query": "` + graphQuery + `"
}`

	res, _, err := HttpPost(urlString, jsonBody, *token)
	lib.CheckFatalError(err)

	// fmt.Println(string(res))
	// os.Exit(0)

	var response ListAllResourceGroupsResponse
	err = json.Unmarshal(res, &response)
	lib.CheckFatalError(err)

	for _, res := range response.Data {
		currRes := res
		currRes.ID = strings.ToLower(res.ID)
		allResGrps = append(allResGrps, currRes)
	}

	hasSkipToken := false
	skipToken := ""

	if response.SkipToken != "" {
		hasSkipToken = true
		skipToken = response.SkipToken
	}

	for hasSkipToken {
		var whileRes ListAllResourceGroupsResponse
		jsonBody := `{
			"query": "` + graphQuery + `",
			"options": {
				"$skipToken": "` + skipToken + `"
			}
		}`

		res, _, err := HttpPost(urlString, jsonBody, *token)
		lib.CheckFatalError(err)
		err = json.Unmarshal(res, &whileRes)
		lib.CheckFatalError(err)

		for _, res := range whileRes.Data {
			currRes := res
			currRes.ID = strings.ToLower(res.ID)
			allResGrps = append(allResGrps, currRes)
		}

		if whileRes.SkipToken != "" {
			hasSkipToken = true
			skipToken = whileRes.SkipToken
		} else {
			hasSkipToken = false
			skipToken = ""
		}
	}

	if outputFile != "" {
		jsonStr, _ := json.Marshal(allResGrps, jsontext.WithIndent("  "))

		err = os.WriteFile(outputFile, jsonStr, 0644)
		lib.CheckFatalError(err)
		fmt.Println("Saved to " + outputFile)
	}

	for i, rg := range allResGrps {
		curr := rg
		curr.TenantName = token.TenantName
		allResGrps[i] = curr
	}

	return allResGrps
}

//
//

func MapVmSizeDetails(vmSize string, resourcesSkus []lib.AzureVirtualMachineSku) (mappedSku lib.AzureVirtualMachineSku) {
	for _, sku := range resourcesSkus {
		if vmSize == sku.Name {
			mappedSku = sku
		}
	}
	return
}

//
//

func GetVcpuCountForAllConfiguredTenants(
	allResources map[string]TenantResourceList,
	opts *lib.GetAllResourcesForAllConfiguredTenantsOptions,
	cfg lib.CldConfigTenants,
) (
	vmResources map[string][]lib.AzureResourceDetails,
	vmResByType map[string]map[string]lib.AzureResourceDetails,
	processedVms map[string][]lib.AzureResourceDetails,
	processedVmsSlice []lib.AzureResourceDetails,
	vCpuCountByTenant lib.VCpuCountByTenant,
	vCpuCountByTenantWithResources lib.VCpuCountByTenant,
) {
	vmResources = make(map[string][]lib.AzureResourceDetails)
	_ = vmResources

	var sqlServers []string

	for _, tData := range allResources {
		for _, sData := range tData.Subscriptions {
			for _, subRes := range sData.Resources {
				if subRes.Type == "microsoft.sqlvirtualmachine/sqlvirtualmachines" {
					sqlServers = append(sqlServers, strings.ToLower(subRes.Properties.VirtualMachineResourceID))
				}
			}
		}
	}

	for tName, tData := range allResources {
		_ = tName
		vmResources[tName] = []lib.AzureResourceDetails{}
		for sName, sData := range tData.Subscriptions {
			_ = sName
			for _, subRes := range sData.Resources {
				currRes := subRes
				currRes.ID = strings.ToLower(currRes.ID)
				if strings.Contains(currRes.Name, "dtw10m") {
					continue
				}
				if currRes.Type == "microsoft.sqlvirtualmachine/sqlvirtualmachines" {
					continue
				}
				if !strings.Contains(strings.ToLower(subRes.Type), "virtualmachines") &&
					!strings.Contains(strings.ToLower(subRes.Type), "hostgroups") {
					continue
				}
				if strings.ToLower(subRes.Type) == "microsoft.compute/virtualmachines/extensions" {
					continue
				}
				if subRes.Properties != nil {
					if subRes.Properties.StorageProfile != nil {
						if subRes.Properties.StorageProfile.OSDisk != nil {
							if strings.ToLower(subRes.Properties.StorageProfile.OSDisk.OSType) == "linux" {
								continue
							}
						}
					}
				}
				if subRes.Properties.VirtualMachineProfile != nil {
					if subRes.Properties.VirtualMachineProfile.StorageProfile != nil {
						if subRes.Properties.VirtualMachineProfile.StorageProfile.OSDisk != nil {
							if strings.ToLower(subRes.Properties.VirtualMachineProfile.StorageProfile.OSDisk.OSType) == "linux" {
								continue
							}
						}
					}
				}
				if subRes.Properties != nil {
					if subRes.Properties.StorageProfile != nil {
						if subRes.Properties.StorageProfile.ImageReference != nil {
							if slices.Contains(sqlServers, subRes.ID) || strings.Contains(strings.ToLower(subRes.Properties.StorageProfile.ImageReference.Publisher), "sql") {
								currRes.IsSqlRelated = true
							}
						}
					}
				}
				if subRes.Properties != nil {
					if subRes.Properties.StorageProfile != nil {
						if subRes.Properties.StorageProfile.ImageReference != nil {
							if !slices.Contains(sqlServers, subRes.ID) && !currRes.IsSqlRelated && strings.ToLower(subRes.Properties.StorageProfile.ImageReference.Publisher) == "windows-10" {
								continue
							}
							if slices.Contains(sqlServers, subRes.ID) && !currRes.IsSqlRelated && strings.ToLower(subRes.Properties.StorageProfile.ImageReference.Publisher) == "windows-10" {
								currRes.WindowsType = "desktop"
							} else {
								currRes.WindowsType = "server"
							}
						}
					}
				}
				if subRes.Properties != nil {
					if subRes.Properties.Extended != nil {
						if subRes.Properties.Extended.InstanceView != nil {
							if !slices.Contains(sqlServers, subRes.ID) && !currRes.IsSqlRelated && strings.Contains(strings.ToLower(subRes.Properties.Extended.InstanceView.OSName), "windows 10") {
								continue
							}
							if slices.Contains(sqlServers, subRes.ID) && !currRes.IsSqlRelated && currRes.WindowsType != "desktop" && strings.Contains(strings.ToLower(subRes.Properties.Extended.InstanceView.OSName), "windows 10") {
								currRes.WindowsType = "desktop"
							} else {
								currRes.WindowsType = "server"
							}
						}
					}
				}
				vmResources[tName] = append(vmResources[tName], currRes)
			}
		}
	}

	vmResByType = make(map[string]map[string]lib.AzureResourceDetails)

	for tName, tData := range vmResources {
		_ = tName
		for _, res := range tData {
			if vmResByType[res.Type] == nil {
				vmResByType[res.Type] = make(map[string]lib.AzureResourceDetails)
			}
			vmResByType[res.Type][res.ID] = res
		}
	}

	processedVms = make(map[string][]lib.AzureResourceDetails)

	for tName, tData := range vmResources {
		processedVms[tName] = []lib.AzureResourceDetails{}
		for _, vm := range tData {
			if vm.Type == "microsoft.sqlvirtualmachine/sqlvirtualmachines" {
				continue
			}
			processedVms[tName] = append(processedVms[tName], vmResByType["microsoft.compute/virtualmachines"][vm.ID])
			processedVmsSlice = append(processedVmsSlice, vm)
		}
	}

	vCpuCountByTenantWithResources = make(lib.VCpuCountByTenant)
	vCpuCountByTenant = make(lib.VCpuCountByTenant)

	for _, vm := range processedVmsSlice {
		currTenant := vCpuCountByTenant[vm.TenantName]
		currTenantWithResources := vCpuCountByTenantWithResources[vm.TenantName]

		id := vm.ID
		vcpus, err := strconv.Atoi(vm.Properties.HardwareProfile.VmSizeSku.VCPUs)
		if err != nil {
			lib.JsonMarshalAndPrint(vm)
			lib.CheckFatalError(err)
		}

		if vm.WindowsType != "desktop" {
			if strings.Contains(strings.ToLower(vm.Properties.Extended.InstanceView.PowerState.Code), "deallocated") {
				currTenant.VmCoreCountDeallocated += vcpus
				currTenantWithResources.VmCoreCountDeallocated += vcpus
			} else {
				currTenant.VmCoreCount += vcpus
				currTenantWithResources.VmCoreCount += vcpus
			}
		}

		if vm.IsSqlRelated {
			if strings.Contains(strings.ToLower(vm.Properties.Extended.InstanceView.PowerState.Code), "deallocated") {
				currTenant.VmCoreCountSqlDeallocated += vcpus
				currTenantWithResources.VmCoreCountSqlDeallocated += vcpus
			} else {
				currTenant.VmCoreCountSql += vcpus
				currTenantWithResources.VmCoreCountSql += vcpus
			}
		}

		vCpuCountByTenant[vm.TenantName] = currTenant

		if vm.WindowsType != "desktop" {
			if strings.Contains(strings.ToLower(vm.Properties.Extended.InstanceView.PowerState.Code), "deallocated") {
				currTenantWithResources.VmResourcesDeallocated = append(currTenantWithResources.VmResourcesDeallocated, id)
			} else {
				currTenantWithResources.VmResources = append(currTenantWithResources.VmResources, id)
			}
		}

		if vm.IsSqlRelated {
			if strings.Contains(strings.ToLower(vm.Properties.Extended.InstanceView.PowerState.Code), "deallocated") {
				currTenantWithResources.VmResourcesSqlDeallocated = append(currTenantWithResources.VmResourcesSqlDeallocated, id)
			} else {
				currTenantWithResources.VmResourcesSql = append(currTenantWithResources.VmResourcesSql, id)
			}
		}

		vCpuCountByTenantWithResources[vm.TenantName] = currTenantWithResources
	}
	return
}

//
//

//
//

//
//

func ListManagementGroups(token *lib.AzureMultiAuthToken) ([]ManagementGroup, error) {
	var response ListManagementGroupsResponse
	// var err error
	urlString := "https://management.azure.com/providers/Microsoft.Management/managementGroups?api-version=2020-05-01"
	res, err := HttpGet(urlString, *token)
	// lib.CheckFatalError(err)
	if err != nil {
		return nil, err
	}
	// fmt.Println(string(res))

	json.Unmarshal(res, &response)

	return response.Value, nil
}

//
