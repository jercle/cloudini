package azure

import (
	"encoding/json/jsontext"
	json "encoding/json/v2"
	"errors"
	"fmt"
	"math/big"
	"net"
	"net/netip"
	"os"
	"slices"
	"strings"
	"sync"
	"time"

	"github.com/jercle/cloudini/lib"
)

func GetIPAddressesAll(queries map[string]string, outputFile string, token *lib.AzureMultiAuthToken) (allIpAddresses []IPAddressesAllResourceTypes) {

	// fmt.Println("Running queries")
	queryResults := GetIPAddressesRunQueries(queries, token)

	fmt.Println(token.TenantName + " - Processing results")
	for _, resource := range queryResults {
		// 	var resource IPAddressesAllResourceTypes
		// 	jsonStr, _ := json.Marshal(r)
		// 	err := json.Unmarshal(jsonStr, &resource)
		// 	if err != nil {
		// 		lib.JsonMarshalAndPrint(r)
		// 		lib.CheckFatalError(err)
		// 	}

		typeLower := strings.ToLower(resource.Type)

		if len(resource.PublicIps) == 0 && len(resource.PrivateIps) == 0 && len(resource.Cidrs) == 0 {
			continue
		}

		switch typeLower {
		case "microsoft.containerservice/managedclusters":
			ipAddresses := GetManagedClusterIPAddresses(resource, token)
			allIpAddresses = append(allIpAddresses, ipAddresses...)
		case "microsoft.network/virtualnetworks":
			for _, s := range resource.Subnets {
				var snet IPAddressesAllResourceTypes
				jsonStr, _ := json.Marshal(s)
				err := json.Unmarshal(jsonStr, &snet)
				lib.CheckFatalError(err)
				snet.SubscriptionID = resource.SubscriptionID
				snet.SubscriptionName = resource.SubscriptionName
				snet.TenantID = resource.TenantID
				snet.TenantName = resource.TenantName
				snet.ResourceGroup = resource.ResourceGroup
				snet.LastAzureSync = resource.LastAzureSync
				allIpAddresses = append(allIpAddresses, snet)
			}
			resource.Subnets = nil
			allIpAddresses = append(allIpAddresses, resource)
		default:
			allIpAddresses = append(allIpAddresses, resource)
		}
	}

	// fmt.Println(len(allIpAddresses))
	// lib.JsonMarshalAndPrint(allIpAddresses)
	// os.Exit(0)

	if outputFile != "" {
		jsonStr, _ := json.Marshal(queryResults)
		os.WriteFile(outputFile, jsonStr, 0644)
	}
	return
}

//
//

func GetIPAddressesRunQueries(queries map[string]string, token *lib.AzureMultiAuthToken) (allIPAddresses []IPAddressesAllResourceTypes) {
	var (
		wg    sync.WaitGroup
		mutex sync.Mutex
	)

	azTenant := lib.GetCldConfig(nil).Azure.MultiTenantAuth.Tenants[token.TenantName]

	fmt.Println(token.TenantName + " - Getting map of available subscriptions")
	subscriptions, err := ListSubscriptions(*token)
	lib.CheckFatalError(err)
	subIdsByNameMap := make(map[string]string)

	for _, sub := range subscriptions {
		subIdsByNameMap[sub.SubscriptionID] = sub.DisplayName
	}

	urlString := "https://management.azure.com/providers/Microsoft.ResourceGraph/resources?api-version=2022-10-01"

	for k, q := range queries {
		if azTenant.IsB2C {
			continue
		}
		wg.Add(1)
		go func() {
			// startTime := time.Now()
			fmt.Println(token.TenantName + " - Running query: " + k)
			defer wg.Done()
			jsonBody := `{
				"query": "` + q + `"
		}`

			res, _, err := HttpPost(urlString, jsonBody, *token)
			lib.CheckFatalError(err)

			// fmt.Println(string(res))
			// os.WriteFile("/home/jercle/git/cld/dev/main-ips-"+k+".json", res, 0644)

			// os.Exit(0)

			var (
				resData      ResourceGraphIPAddressesResponse
				queryResults []IPAddressesAllResourceTypes
			)

			err = json.Unmarshal(res, &resData)
			lib.CheckFatalError(err)

			// // mutex.Lock()
			// // queryResults = append(queryResults, resData.Data...)
			// // mutex.Unlock()

			for _, r := range resData.Data {
				currRes := r
				currRes.TenantName = token.TenantName
				currRes.SubscriptionName = subIdsByNameMap[currRes.SubscriptionID]
				currRes.ID = strings.ToLower(r.ID)
				currRes.LastAzureSync = time.Now()
				queryResults = append(queryResults, currRes)
			}
			hasSkipToken := false
			skipToken := ""

			if resData.SkipToken != "" {
				hasSkipToken = true
				skipToken = resData.SkipToken
			}

			for hasSkipToken {
				var whileRes ResourceGraphIPAddressesResponse
				jsonBody := `{
						"query": "` + q + `",
						"options": {
							"$skipToken": "` + skipToken + `"
						}
					}`

				res, _, err := HttpPost(urlString, jsonBody, *token)
				lib.CheckFatalError(err)
				err = json.Unmarshal(res, &whileRes)
				lib.CheckFatalError(err)
				for _, r := range whileRes.Data {
					currRes := r
					currRes.TenantName = token.TenantName
					currRes.SubscriptionName = subIdsByNameMap[currRes.SubscriptionID]
					currRes.ID = strings.ToLower(r.ID)
					currRes.LastAzureSync = time.Now()
					queryResults = append(queryResults, currRes)
				}

				if whileRes.SkipToken != "" {
					hasSkipToken = true
					skipToken = whileRes.SkipToken
				} else {
					hasSkipToken = false
					skipToken = ""
				}
			}
			mutex.Lock()
			allIPAddresses = append(allIPAddresses, queryResults...)
			mutex.Unlock()
			// elapsed := time.Since(startTime)
			// fmt.Println(token.TenantName + " - Query complete: " + k + " after " + elapsed.String())
		}()
	}

	wg.Wait()
	return
}

//
//

func GetIPAddressesQueries(selectedQueries *[]string) map[string]string {
	const GetIPAddressesQueryNetworkInterfaces = `Resources
| where type =~ 'microsoft.network/networkinterfaces'
| mv-expand ipconfig = properties.ipConfigurations
| extend publicIpId = tostring(ipconfig.properties.publicIPAddress.id), tags = iff(isnull(tags), dynamic({}), tags)
| extend bareMetalServer= properties.bareMetalServer, privateEndpoint = properties.privateEndpoint, privateLinkService = properties.privateLinkService, virtualMachine = properties.virtualMachine, name, id, properties
| extend isAttached = isnotnull(bareMetalServer) or isnotnull(privateEndpoint) or isnotnull(privateLinkService) or isnotnull(virtualMachine)
| extend attachedTo = dynamic_to_json(coalesce(bareMetalServer.id, privateLinkService.id, privateEndpoint.id, virtualMachine.id))
| join kind=leftouter  (resources | project id, publicIp = properties.ipAddress) on $left.publicIpId == $right.['id']
| project name, resourceGroup, subscriptionId, tenantId, id, privateIp = ipconfig.properties.privateIPAddress, publicIpId = tostring(ipconfig.properties.publicIPAddress.id), publicIp, type, tags, isAttached, attachedTo
| summarize privateIps = make_list(privateIp), publicIps = make_list(publicIp)  by id, name, resourceGroup, subscriptionId, tenantId, type, tags = dynamic_to_json(tags), isAttached, attachedTo`

	const GetIPAddressesQueryVirtualMachines = `Resources
| where type =~ 'microsoft.compute/virtualmachines'
| project id, vmId = tolower(tostring(id)), vmName = name, type, tenantId, subscriptionId, tags = iff(isnull(tags), dynamic({}), tags)
| join (
    Resources
    | where type =~ 'microsoft.network/networkinterfaces'
    | mv-expand ipconfig = properties.ipConfigurations
    | project vmId = tolower(tostring(properties.virtualMachine.id)), nicId = id, privateIp = ipconfig.properties.privateIPAddress, publicIpId = tostring(ipconfig.properties.publicIPAddress.id)
    | join kind=leftouter (
        Resources
        | where type =~ 'microsoft.network/publicipaddresses'
        | project publicIpId = id, publicIp = properties.ipAddress
        )
        on publicIpId
    | project-away publicIpId, publicIpId1
    | summarize associatedNics = make_list(nicId), privateIps = make_list(privateIp), publicIps = make_list(publicIp) by vmId
    )
    on vmId
| project id, name = vmName, type, privateIps, publicIps, tenantId, subscriptionId, associatedNics, tags = dynamic_to_json(tags)`

	const GetIPAddressesQueryLoadBalancers = `Resources
| where type =~ 'microsoft.network/loadbalancers'
| mv-expand feIpConfig = properties.frontendIPConfigurations
| project lbId = id, lbName = name, type, privateIp = feIpConfig.properties.privateIPAddress, publicIpId = tostring(feIpConfig.properties.publicIPAddress.id), tenantId, subscriptionId, tags
| join kind=leftouter (
    Resources
    | where type =~ 'microsoft.network/publicipaddresses'
    | project publicIpId = id, publicIp = properties.ipAddress
    )
    on publicIpId
| project-away publicIpId, publicIpId1
| summarize privateIps = make_list(privateIp), publicIps = make_list(publicIp) by id = lbId, name = lbName, type, tenantId, subscriptionId, tags = dynamic_to_json(tags)`

	const GetIPAddressesQueryManagedEnvironments = `Resources
| where type =~ 'microsoft.app/managedenvironments'
| project id, name, type, tenantId, resourceGroup, subscriptionId, publicNetworkAccess = properties.publicNetworkAccess, privateIps = pack_array(properties.staticIp), tags = dynamic_to_json(tags)`

	const GetIPAddressesQueryBastionHosts = `Resources
| where type =~ 'microsoft.network/bastionhosts'
| mv-expand ipconfig = properties.ipConfigurations
| project id, name, type, tenantId, resourceGroup, subscriptionId, publicIpId = tostring(ipconfig.properties.publicIPAddress.id), tags = dynamic_to_json(tags)
| join kind=leftouter  (resources | project id, publicIp = properties.ipAddress) on $left.publicIpId == $right.['id']
| summarize publicIps = make_list(publicIp), publicIpIds = make_list(publicIpId)  by id, name, resourceGroup, subscriptionId, tenantId, type, tags`

	const GetIPAddressesQueryPrivateEndpoints = `resources
| where type =~ 'microsoft.network/privateendpoints'
| project id, peId = tolower(tostring(id)), name, type, tenantId, subscriptionId, tags = iff(isnull(tags), dynamic({}), tags)
| join (
    Resources
    | where type =~ 'microsoft.network/networkinterfaces'
    | mv-expand ipconfig = properties.ipConfigurations
    | project peId  = tolower(tostring(properties.privateEndpoint.id)), privateIp = ipconfig.properties.privateIPAddress, nicId = id
    )
    on peId
| summarize associatedNics = make_list(nicId), privateIps = make_list(privateIp) by id, name, type, tenantId, subscriptionId, tags = dynamic_to_json(tags)`

	const GetIPAddressesQueryPublicIPs = `
resources
| where type =~ 'microsoft.network/publicipaddresses'
| extend ipConfig = properties.ipConfiguration.id
| extend isAttached = isnotnull(ipConfig)
| project id, name, type, tenantId, subscriptionId, resourceGroup, tags = dynamic_to_json(tags), publicIps = iff(isnotnull(properties.ipAddress), pack_array(properties.ipAddress), dynamic([])) , isAttached`
	const GetIPAddressesQueryWebSites = `Resources
| where type =~ 'microsoft.web/sites'
| extend possibleInboundIps = split(properties.possibleInboundIpAddresses, ',')
| extend possibleOutboundIps = split(properties.possibleOutboundIpAddresses, ',')
| extend inboundIps = split(properties.inboundIpAddress, ',')
| extend outboundIps = split(properties.outboundIpAddresses, ',')
| extend privateIps = array_concat(possibleInboundIps, possibleOutboundIps)
| project id, name, type, tenantId, subscriptionId, resourceGroup, tags = dynamic_to_json(tags), privateIps, possibleInboundIps, possibleOutboundIps, inboundIps, outboundIps`

	const GetIPAddressesQueryManagedClusters = `Resources
| where type =~ 'microsoft.containerservice/managedclusters'
| mv-expand agentPools = properties.agentPoolProfiles
| project id, name, type, tenantId, subscriptionId, resourceGroup, tags = dynamic_to_json(tags)`

	const GetIPAddressesQueryFirewalls = `Resources
| where type =~ 'microsoft.network/azurefirewalls'
| mv-expand hubIPAddresses = properties.hubIPAddresses
| mv-expand publicIp = hubIPAddresses.publicIPs.addresses
| extend privateIp = hubIPAddresses.privateIPAddress
| extend publicIp = publicIp.address
| project-away hubIPAddresses
| mv-expand ipConfig = properties.ipConfigurations
| extend pubIpId = tostring(ipConfig.properties.publicIPAddress.id), privateIp = iff(isnull(privateIp), ipConfig.properties.privateIPAddress, privateIp)
| join kind=leftouter  (resources | project pubIpId = id, publicIp = properties.ipAddress) on pubIpId
| extend publicIp = iff(isnull(publicIp), publicIp1, publicIp)
| summarize privateIps = make_list(privateIp), publicIps = make_list(publicIp) by id, name, resourceGroup, tenantId, subscriptionId, type, tags = dynamic_to_json(tags)`

	const GetIPAddressesQueryP2SVPNGateways = `Resources
| where type =~ 'microsoft.network/p2svpngateways'
| mv-expand p2sConnConfig = properties.p2SConnectionConfigurations
| mv-expand vpnClientAddressPoolPrefix = p2sConnConfig.properties.vpnClientAddressPool.addressPrefixes
| summarize cidrs = make_list(vpnClientAddressPoolPrefix) by id, name, type, tenantId, subscriptionId, resourceGroup, tags = dynamic_to_json(tags)`

	const GetIPAddressesQueryVirtualHubs = `Resources
| where type =~ 'microsoft.network/virtualhubs'
| project id, name, type, tenantId, subscriptionId, resourceGroup, tags = dynamic_to_json(tags), privateIps = properties.virtualRouterIps, cidrs = pack_array(properties.addressPrefix)`

	const GetIPAddressesQueryVirtualNetworks = `Resources
| where type =~ 'microsoft.network/virtualnetworks'
| mv-expand subnetObj = properties.subnets
| extend snetId = subnetObj.id, snetName = subnetObj.name, snetCidrSingle = iff(isnull(subnetObj.properties.addressPrefix), dynamic([]), pack_array(subnetObj.properties.addressPrefix)), snetCidrArr = subnetObj.properties.addressPrefixes, snetType = subnetObj.type
| extend snetCidrs = array_concat(snetCidrSingle, snetCidrArr)
| extend subnet = pack_dictionary('id', snetId, 'name', snetName, 'cidrs', snetCidrs, 'type', snetType)
| summarize cidrs = make_set(properties.addressSpace.addressPrefixes), subnets = make_list(subnet) by id, name, type, tenantId, subscriptionId, resourceGroup, tags = dynamic_to_json(tags)`

	const GetIPAddressesQueryIPGroups = `Resources
| where type =~ 'microsoft.network/ipgroups'
| project id, name, type, tenantId, subscriptionId, resourceGroup, tags = dynamic_to_json(tags), cidrs = properties.ipAddresses`

	// const GetIPAddressesQuery = ``
	queries := make(map[string]string)

	queries["GetIPAddressesQueryBastionHosts"] = GetIPAddressesQueryBastionHosts
	queries["GetIPAddressesQueryFirewalls"] = GetIPAddressesQueryFirewalls
	queries["GetIPAddressesQueryIPGroups"] = GetIPAddressesQueryIPGroups
	queries["GetIPAddressesQueryLoadBalancers"] = GetIPAddressesQueryLoadBalancers
	queries["GetIPAddressesQueryManagedClusters"] = GetIPAddressesQueryManagedClusters
	queries["GetIPAddressesQueryManagedEnvironments"] = GetIPAddressesQueryManagedEnvironments
	queries["GetIPAddressesQueryNetworkInterfaces"] = GetIPAddressesQueryNetworkInterfaces
	queries["GetIPAddressesQueryP2SVPNGateways"] = GetIPAddressesQueryP2SVPNGateways
	queries["GetIPAddressesQueryPrivateEndpoints"] = GetIPAddressesQueryPrivateEndpoints
	queries["GetIPAddressesQueryPublicIPs"] = GetIPAddressesQueryPublicIPs
	queries["GetIPAddressesQueryVirtualHubs"] = GetIPAddressesQueryVirtualHubs
	queries["GetIPAddressesQueryVirtualMachines"] = GetIPAddressesQueryVirtualMachines
	queries["GetIPAddressesQueryVirtualNetworks"] = GetIPAddressesQueryVirtualNetworks
	queries["GetIPAddressesQueryWebSites"] = GetIPAddressesQueryWebSites

	if selectedQueries != nil {
		return lib.SelectMapStringFieldsFromArrayOfKeys(queries, *selectedQueries)
	} else {
		return queries
	}
}

//
//

func GetManagedClusterIPAddresses(clusterIpObject IPAddressesAllResourceTypes, token *lib.AzureMultiAuthToken) (ipAddresses []IPAddressesAllResourceTypes) {
	baseUrl := "https://management.azure.com"
	agentPools := GetManagedClusterAgentPools(clusterIpObject.ID, token)

	var allIPs []string
	var associatedIDs []string

	for _, ap := range agentPools {
		urlString := baseUrl + ap.ID + "/machines?api-version=2025-05-01"
		res, err := HttpGet(urlString, *token)
		if err != nil {
			errStr := err.Error()
			if strings.Contains(errStr, "ParentResourceNotFound") {
				continue
				lib.CheckFatalError(err)
			}
		}

		var resData ResourceManagerResponse
		err = json.Unmarshal(res, &resData)

		var apIPs []string
		var machineIDs []string

		for _, m := range resData.Value {
			jsonStr, _ := json.Marshal(m)
			var machine AzureResourceManagedClusterAgentPoolMachine
			err := json.Unmarshal(jsonStr, &machine)
			lib.CheckFatalError(err)
			var machineIPs []string
			for _, ip := range machine.Properties.Network.IpAddresses {
				parsedIp := net.ParseIP(ip.Ip)
				if parsedIp != nil {
					ipStr := parsedIp.String()
					machineIPs = append(machineIPs, ipStr)
					apIPs = append(apIPs, ipStr)
					allIPs = append(allIPs, ipStr)
				}
			}
			machineIDs = append(machineIDs, machine.ID)
			associatedIDs = append(associatedIDs, machine.ID)

			machineIpObject := IPAddressesAllResourceTypes{
				ID:                    machine.ID,
				PrivateIps:            machineIPs,
				Name:                  machine.Name,
				ResourceGroup:         clusterIpObject.ResourceGroup,
				SubscriptionID:        clusterIpObject.SubscriptionID,
				TenantID:              clusterIpObject.TenantID,
				Type:                  machine.Type,
				AssociatedResourceIDs: []string{machine.Properties.ResourceID, clusterIpObject.ID},
			}
			ipAddresses = append(ipAddresses, machineIpObject)
		}

		apIpObject := IPAddressesAllResourceTypes{
			ID:                    ap.ID,
			PrivateIps:            apIPs,
			Name:                  ap.Name,
			ResourceGroup:         clusterIpObject.ResourceGroup,
			SubscriptionID:        clusterIpObject.SubscriptionID,
			TenantID:              clusterIpObject.TenantID,
			AssociatedResourceIDs: machineIDs,
			Type:                  ap.Type,
		}
		ipAddresses = append(ipAddresses, apIpObject)
		if !slices.Contains(associatedIDs, ap.ID) {
			associatedIDs = append(associatedIDs, ap.ID)
		}
	}

	ipAddrObject := clusterIpObject
	ipAddrObject.PrivateIps = allIPs
	ipAddrObject.AssociatedResourceIDs = append(ipAddrObject.AssociatedResourceIDs, associatedIDs...)
	ipAddresses = append(ipAddresses, ipAddrObject)

	return
}

//
//

func GetManagedClusterAgentPools(clusterId string, token *lib.AzureMultiAuthToken) (agentPools []AzureResourceManagedClusterAgentPool) {
	baseUrl := "https://management.azure.com"
	urlString := baseUrl + clusterId + "/agentPools?api-version=2025-05-01"

	res, err := HttpGet(urlString, *token)
	lib.CheckFatalError(err)

	var resData ResourceManagerResponse
	err = json.Unmarshal(res, &resData)
	lib.CheckFatalError(err)

	var nextLink string
	nextLink = resData.NextLink

	for _, ap := range resData.Value {
		jsonStr, _ := json.Marshal(ap)
		var agentPool AzureResourceManagedClusterAgentPool
		err := json.Unmarshal(jsonStr, &agentPool)
		lib.CheckFatalError(err)
		agentPools = append(agentPools, agentPool)
	}

	for nextLink != "" {
		var currentSet ResourceManagerResponse

		res, err := HttpGet(nextLink, *token)
		lib.CheckFatalError(err)

		err = json.Unmarshal(res, &currentSet)
		lib.CheckFatalError(err)

		nextLink = currentSet.NextLink
		for _, ap := range currentSet.Value {
			jsonStr, _ := json.Marshal(ap)
			var agentPool AzureResourceManagedClusterAgentPool
			err := json.Unmarshal(jsonStr, &agentPool)
			lib.CheckFatalError(err)
			agentPools = append(agentPools, agentPool)
		}
	}
	return
}

//
//

//
//

func GetAllVMIpAddrForAllConfiguredTenants(opts *lib.GetAllResourcesForAllConfiguredTenantsOptions, tokens lib.AllTenantTokens) (allResourceIPs []IPAddressesAllResourceTypes) {
	var (
		wg    sync.WaitGroup
		mutex sync.Mutex
	)

	queries := GetIPAddressesQueries(opts.SelectedIPAddressQueries)

	for _, token := range tokens {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if !opts.SuppressSteps {
				fmt.Println(token.TenantName + ": Fetching resources")
			}
			// allResources[token.TenantName] = make(map[string]SubscriptionResourceList)
			// var r AzureResourceIPConfig
			// tenantResourceIPsOld := GetAllTenantIpAddresses("", &token)
			tenantResourceIPs := GetIPAddressesAll(queries, "", &token)
			if !opts.SuppressSteps {
				fmt.Println(token.TenantName + ": Fetch complete")
			}
			// var processedTenantResources TenantResourceList

			mutex.Lock()
			allResourceIPs = append(allResourceIPs, tenantResourceIPs...)
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
		jsonStr, _ := json.Marshal(allResourceIPs, jsontext.WithIndent("  "))

		currentDate := time.Now().Format("20060102")

		arrayFileName := outputFilePath + "/allRes-GraphResources-AllTenantIPs-" + currentDate + ".json"

		err := os.WriteFile(arrayFileName, jsonStr, 0644)
		lib.CheckFatalError(err)
		fmt.Println("Saved to " + arrayFileName + " and " + arrayFileName)
	}

	// fmt.Println(len(allResourcesSlice))

	return allResourceIPs
}

//
//

func GetAllTenantIpAddresses(outputFile string, token *lib.AzureMultiAuthToken) []AzureResourceIPConfig {
	subscriptions, err := ListSubscriptions(*token)
	lib.CheckFatalError(err)
	var allTenantResourceIPs []AzureResourceIPConfig
	subIdsByNameMap := make(map[string]string)

	for _, sub := range subscriptions {
		subIdsByNameMap[sub.SubscriptionID] = sub.DisplayName
	}

	urlString := "https://management.azure.com/providers/Microsoft.ResourceGraph/resources?api-version=2022-10-01"

	// graphQuery := `Resources
	//   | where type =~ 'microsoft.compute/virtualmachines'
	//   | project id, vmId = tolower(tostring(id)), vmName = name, type, tenantId, subscriptionId
	//   | join (Resources
	//       | where type =~ 'microsoft.network/networkinterfaces'
	//       | mv-expand ipconfig=properties.ipConfigurations
	//       | project vmId = tolower(tostring(properties.virtualMachine.id)), nicId = id, privateIp = ipconfig.properties.privateIPAddress, publicIpId = tostring(ipconfig.properties.publicIPAddress.id)
	//       | join kind=leftouter (Resources
	//           | where type =~ 'microsoft.network/publicipaddresses'
	//           | project publicIpId = id, publicIp = properties.ipAddress
	//       ) on publicIpId
	//       | project-away publicIpId, publicIpId1
	//       | summarize vmNics = make_list(nicId), privateIps = make_list(privateIp), publicIps = make_list(publicIp) by vmId
	//   ) on vmId
	//   | project-away vmId, vmId1
	//   | project id, name = vmName, type, privateIps, publicIps, tenantId, subscriptionId, vmNics
	//   | union (
	//       Resources
	//       | where type =~ 'microsoft.network/loadbalancers'
	//       | project id, lbId = tolower(tostring(id)), lbName = name, properties, type, tenantId, subscriptionId
	//           | mv-expand feIpConfig=properties.frontendIPConfigurations
	//           | project lbId = id, lbName, type, privateIp = feIpConfig.properties.privateIPAddress, publicIpId = tostring(feIpConfig.properties.publicIPAddress.id), tenantId, subscriptionId
	//           | join kind=leftouter (Resources
	//               | where type =~ 'microsoft.network/publicipaddresses'
	//               | project publicIpId = id, publicIp = properties.ipAddress
	//           ) on publicIpId
	//           | project-away publicIpId, publicIpId1
	//           | summarize privateIps = make_list(privateIp), publicIps = make_list(publicIp) by lbId, lbName, type, tenantId, subscriptionId
	//           | project id = lbId, name = lbName, type, privateIps, publicIps, tenantId, subscriptionId
	//   )`

	graphQuery := `Resources
    | where type =~ 'microsoft.network/networkinterfaces'
    | where properties has 'virtualmachine'
    | mv-expand ipconfig=properties.ipConfigurations
    | project attachedId = tolower(tostring(properties.virtualMachine.id)), nicId = id, name, snetId = tolower(tostring(ipconfig.properties.subnet.id)),privateIp = ipconfig.properties.privateIPAddress, publicIpId = tostring(ipconfig.properties.publicIPAddress.id)
    | join kind=leftouter (Resources
        | where type =~ 'microsoft.network/publicipaddresses'
        | project publicIpId = id, publicIp = properties.ipAddress
    ) on publicIpId
    | project-away publicIpId, publicIpId1
    | summarize privateIps = make_list(privateIp), publicIps = make_list(publicIp), snetIds = make_set(snetId) by nicId, name, attachedId
    | join (Resources
        | where type =~ 'microsoft.compute/virtualmachines'
        | project attachedId = tolower(tostring(id)), attachedName = name, type, tenantId, subscriptionId) on attachedId
    | project id = tolower(tostring(attachedId)), name = attachedName, type, privateIps, snetIds, publicIps, nicId, nicName = name, tenantId, subscriptionId
    | summarize privateIps = make_list(privateIps), publicIps = make_list(publicIps), snetIds = make_set(snetIds), nicIds = make_set(nicId) by id, name, subscriptionId, type`

	jsonBody := `{
	"query": "` + graphQuery + `"
}`

	res, _, err := HttpPost(urlString, jsonBody, *token)
	lib.CheckFatalError(err)

	var response ResourceGraphGetIpsResponse
	err = json.Unmarshal(res, &response)
	lib.CheckFatalError(err)

	for _, res := range response.Data {
		currRes := res
		currRes.TenantName = token.TenantName
		currRes.TenantId = token.TenantId
		currRes.SubscriptionName = subIdsByNameMap[currRes.SubscriptionId]
		currRes.ID = strings.ToLower(res.ID)
		currRes.LastAzureSync = time.Now()
		allTenantResourceIPs = append(allTenantResourceIPs, currRes)
		// lib.JsonMarshalAndPrint(currRes)
		// lib.JsonMarshalAndPrint(subIdsByNameMap)
		// fmt.Println(subIdsByNameMap[currRes.SubscriptionId])
		// os.Exit(0)
	}

	// allResources = append(allResources, response.Data...)

	hasSkipToken := false
	skipToken := ""

	if response.SkipToken != "" {
		hasSkipToken = true
		skipToken = response.SkipToken
	}

	for hasSkipToken {
		var whileRes ResourceGraphGetIpsResponse
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

		// allResources = append(allResources, whileRes.Data...)
		for _, res := range whileRes.Data {
			currRes := res
			currRes.ID = strings.ToLower(res.ID)
			allTenantResourceIPs = append(allTenantResourceIPs, currRes)
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
		jsonStr, _ := json.Marshal(allTenantResourceIPs, jsontext.WithIndent("  "))

		err = os.WriteFile(outputFile, jsonStr, 0644)
		lib.CheckFatalError(err)
		fmt.Println("Saved to " + outputFile)
	}

	// var allTenantResources TenantResourceList

	// fmt.Println(allTenantResources.ResourceCount)
	// os.Exit(0)
	// allTenantResources.ResourceCount = len(allTenantResourcesBySub)
	// allTenantResources.resources

	return allTenantResourceIPs
}

//
//

func IpRangeToCidr(start, end string) ([]string, error) {
	ips, err := netip.ParseAddr(start)
	if err != nil {
		return nil, err
	}
	ipe, err := netip.ParseAddr(end)
	if err != nil {
		return nil, err
	}

	isV4 := ips.Is4()
	if isV4 != ipe.Is4() {
		return nil, errors.New("start and end types are different")
	}
	if ips.Compare(ipe) > 0 {
		return nil, errors.New("start > end")
	}

	var (
		ipsInt = new(big.Int).SetBytes(ips.AsSlice())
		ipeInt = new(big.Int).SetBytes(ipe.AsSlice())
		nextIp = new(big.Int)
		maxBit = new(big.Int)
		cmpSh  = new(big.Int)
		bits   = new(big.Int)
		mask   = new(big.Int)
		one    = big.NewInt(1)
		buf    []byte
		cidr   []string
		bitSh  uint
	)
	if isV4 {
		maxBit.SetUint64(32)
		buf = make([]byte, 4)
	} else {
		maxBit.SetUint64(128)
		buf = make([]byte, 16)
	}

	for {
		bits.SetUint64(1)
		mask.SetUint64(1)
		for bits.Cmp(maxBit) < 0 {
			nextIp.Or(ipsInt, mask)

			bitSh = uint(bits.Uint64())
			cmpSh.Lsh(cmpSh.Rsh(ipsInt, bitSh), bitSh)
			if (nextIp.Cmp(ipeInt) > 0) || (cmpSh.Cmp(ipsInt) != 0) {
				bits.Sub(bits, one)
				mask.Rsh(mask, 1)
				break
			}
			bits.Add(bits, one)
			mask.Add(mask.Lsh(mask, 1), one)
		}

		addr, _ := netip.AddrFromSlice(ipsInt.FillBytes(buf))
		cidr = append(cidr, addr.String()+"/"+bits.Sub(maxBit, bits).String())

		if nextIp.Or(ipsInt, mask); nextIp.Cmp(ipeInt) >= 0 {
			break
		}
		ipsInt.Add(nextIp, one)
	}
	return cidr, nil
}

//
//

func GetIpAddressBlocksForCidrFromVNets(cidrsToCheck []lib.IpamCidrBlockToCheck, vnets []IPAddressesAllResourceTypes) (allAddressBlocks []IpAddressBlocksByBlockTag) {
	for _, cidr := range cidrsToCheck {
		lastFoundVnet := ""
		var currentRange IpAddressBlock

		ipsInCidr, err := GetIPsFromCIDR(cidr.CidrBlock)
		lib.CheckFatalError(err)

		var ipAddressBlocks []IpAddressBlock

		for i, ip := range ipsInCidr {
			if i == 0 {
				// currentRange = append(currentRange, ip.String())
				currentRange.FirstIp = ip.String()
			}

			found := false
			foundVnet := ""
			for _, vnet := range vnets {
				_, vnetIpNet, err := net.ParseCIDR(vnet.Cidrs[0])
				lib.CheckFatalError(err)
				if vnetIpNet.Contains(ip) {
					found = true
					foundVnet = vnet.Name
					break
				}
			}

			if found {
				if i == 0 {
					lastFoundVnet = foundVnet
				}
				if lastFoundVnet != foundVnet {
					cidrBlock, err := IpRangeToCidr(currentRange.FirstIp, currentRange.LastIp)
					lib.CheckFatalError(err)
					currentRange.CidrBlocks = cidrBlock
					ipAddressBlocks = append(ipAddressBlocks, currentRange)
					currentRange = IpAddressBlock{
						FirstIp:         ip.String(),
						VNetName:        foundVnet,
						AllocatedToVnet: true,
						// IpAddresses:     []net.IP{ip},
					}
					lastFoundVnet = foundVnet
				} else {
					currentRange.VNetName = foundVnet
					currentRange.LastIp = ip.String()
					currentRange.AllocatedToVnet = true
					lastFoundVnet = foundVnet

					if i == len(ipsInCidr)-1 {
						cidrBlock, err := IpRangeToCidr(currentRange.FirstIp, currentRange.LastIp)
						lib.CheckFatalError(err)
						currentRange.CidrBlocks = cidrBlock
						ipAddressBlocks = append(ipAddressBlocks, currentRange)
					}
				}
			} else {
				if lastFoundVnet != "" {
					cidrBlock, err := IpRangeToCidr(currentRange.FirstIp, currentRange.LastIp)
					lib.CheckFatalError(err)
					currentRange.CidrBlocks = cidrBlock
					ipAddressBlocks = append(ipAddressBlocks, currentRange)
					currentRange = IpAddressBlock{
						FirstIp:         ip.String(),
						VNetName:        "",
						AllocatedToVnet: false,
						// IpAddresses:     []net.IP{ip},
					}
					lastFoundVnet = ""
				} else {
					// currentRange.IpAddresses = append(currentRange.IpAddresses, ip)
					currentRange.VNetName = ""
					currentRange.LastIp = ip.String()
					currentRange.AllocatedToVnet = false

					lastFoundVnet = ""

					if i == len(ipsInCidr)-1 {
						cidrBlock, err := IpRangeToCidr(currentRange.FirstIp, currentRange.LastIp)
						lib.CheckFatalError(err)
						currentRange.CidrBlocks = cidrBlock
						ipAddressBlocks = append(ipAddressBlocks, currentRange)
					}
				}
			}
		}

		byBlockTag := IpAddressBlocksByBlockTag{
			BlockTag:      cidr.BlockTag,
			AddressBlocks: ipAddressBlocks,
		}

		allAddressBlocks = append(allAddressBlocks, byBlockTag)
	}

	return
}

//
//

// getIPsFromCIDR parses a CIDR string and returns a slice of net.IP representing all IPs in the range.
func GetIPsFromCIDR(cidr string) ([]net.IP, error) {
	_, ipNet, err := net.ParseCIDR(cidr)
	if err != nil {
		return nil, fmt.Errorf("invalid CIDR: %w", err)
	}

	var ips []net.IP
	for ip := ipNet.IP.Mask(ipNet.Mask); ipNet.Contains(ip); IncrementIp(ip) {
		newIP := make(net.IP, len(ip))
		copy(newIP, ip)
		ips = append(ips, newIP)
	}
	return ips, nil
}

//
//

// IncrementIp increments an IP address by one.
func IncrementIp(ip net.IP) {
	for i := len(ip) - 1; i >= 0; i-- {
		ip[i]++
		if ip[i] > 0 {
			break
		}
	}
}

//
//
