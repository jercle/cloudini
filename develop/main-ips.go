package main

import (
	"encoding/json/jsontext"
	json "encoding/json/v2"
	"fmt"
	"net"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/jercle/cloudini/cmd/azure"
	"github.com/jercle/cloudini/lib"
)

func main() {
	// usrHomeDir, err := os.UserHomeDir()
	// lib.CheckFatalError(err)
	startTime := time.Now()
	config := lib.GetCldConfig(nil)
	_ = config
	// tokens, err := azure.GetAllTenantSPTokens(lib.AzureMultiAuthTokenRequestOptions{})
	// lib.CheckFatalError(err)
	// token, err := azure.GetTenantSPToken(lib.AzureMultiAuthTokenRequestOptions{})
	tokenReq, err := azure.GetAllTenantSPTokens(lib.AzureMultiAuthTokenRequestOptions{
		// GetWriteToken: true,
	}, nil)
	lib.CheckFatalError(err)
	token, err := tokenReq.SelectTenant("")
	lib.CheckFatalError(err)
	// _ = tokens
	_ = token

	// GetAllIPsNetworkInterfaces("", token)

	// file, err := os.ReadFile("./main-ips-all-full.json")
	// lib.CheckFatalError(err)
	// fmt.Println(string(file))
	// var allIPs IPAddressesAllTypes
	// err = json.Unmarshal(file, &allIPs)
	// lib.JsonMarshalAndPrint(allIPs)

	// for _, ipObj := range allIPs {
	// 	if ipObj.VmNics != "" {
	// 		lib.JsonMarshalAndPrint(ipObj)
	// 		os.Exit(0)
	// 	}
	// }

	GetIPAddressesAll(token)
	os.Exit(0)
	elapsed := time.Since(startTime)
	_ = elapsed
	// fmt.Println(elapsed)
}

func GetIPAddressesAll(token *lib.AzureMultiAuthToken) (allIpAddresses []IPAddressesAllTypes) {
	// var (
	// 	wg    sync.WaitGroup
	// 	mutex sync.Mutex
	// )
	queries := []string{
		// GetIPAddressesQueryNetworkInterfaces,
		// GetIPAddressesQueryVirtualMachines,
		// GetIPAddressesQueryLoadBalancers,
		// GetIPAddressesQueryManagedEnvironments,
		// GetIPAddressesQueryBastionHosts,
		// GetIPAddressesQueryPrivateEndpoints,
		// GetIPAddressesQueryPublicIPs,
		// GetIPAddressesQueryWebSites,
		// GetIPAddressesQueryManagedClusters,
		GetIPAddressesQueryFirewalls,
	}

	queryResults := GetIPAddressesRunQueries(queries, token)

	for _, r := range queryResults {
		var resource IPAddressesAllTypes
		jsonStr, _ := json.Marshal(r)
		err := json.Unmarshal(jsonStr, &resource)
		lib.CheckFatalError(err)

		typeLower := strings.ToLower(resource.Type)

		switch typeLower {
		case "microsoft.containerservice/managedclusters":
			ipAddresses := GetManagedClusterIPAddresses(resource, token)
			allIpAddresses = append(allIpAddresses, ipAddresses...)
		default:
			allIpAddresses = append(allIpAddresses, resource)
		}
	}

	fmt.Println(len(allIpAddresses))
	// lib.JsonMarshalAndPrint(allIpAddresses)
	// os.Exit(0)

	// jsonStr, _ := json.Marshal(queryResults)
	// os.WriteFile("main-ips-all-full.json", jsonStr, 0644)
	// lib.JsonMarshalAndPrint(queryResults)
	// fmt.Println(len(queryResults))
	return
}

func GetIPAddressesRunQueries(queries []string, token *lib.AzureMultiAuthToken) (queryResults []interface{}) {
	urlString := "https://management.azure.com/providers/Microsoft.ResourceGraph/resources?api-version=2022-10-01"

	for _, q := range queries {
		jsonBody := `{
    "query": "` + q + `"
}`

		res, _, err := azure.HttpPost(urlString, jsonBody, *token)
		lib.CheckFatalError(err)

		// fmt.Println(string(res))
		// os.WriteFile("main-ips-fw.json", res, 0644)
		// os.Exit(0)

		var resData ResourceGraphResponse

		err = json.Unmarshal(res, &resData)
		lib.CheckFatalError(err)

		queryResults = append(queryResults, resData.Data...)

		hasSkipToken := false
		skipToken := ""

		if resData.SkipToken != "" {
			hasSkipToken = true
			skipToken = resData.SkipToken
		}

		for hasSkipToken {
			var whileRes ResourceGraphResponse
			jsonBody := `{
				"query": "` + q + `",
				"options": {
					"$skipToken": "` + skipToken + `"
				}
			}`

			res, _, err := azure.HttpPost(urlString, jsonBody, *token)
			lib.CheckFatalError(err)
			err = json.Unmarshal(res, &whileRes)
			lib.CheckFatalError(err)

			queryResults = append(queryResults, whileRes.Data...)

			if whileRes.SkipToken != "" {
				hasSkipToken = true
				skipToken = whileRes.SkipToken
			} else {
				hasSkipToken = false
				skipToken = ""
			}
		}
	}
	return
}

func GetAzureResourceTypes(token *lib.AzureMultiAuthToken) {
	graphQuery := "resources | distinct type"
	jsonBody := `{
    "query": "` + graphQuery + `"
}`
	urlString := "https://management.azure.com/providers/Microsoft.ResourceGraph/resources?api-version=2022-10-01"
	res, _, err := azure.HttpPost(urlString, jsonBody, *token)
	lib.CheckFatalError(err)

	var resData ResourceGraphResponse
	// os.WriteFile("main-ips-0-prewhile.json", res, 0644)

	err = json.Unmarshal(res, &resData)
	lib.CheckFatalError(err)

	var types []string

	for _, t := range resData.Data {
		jsonStr, _ := json.Marshal(t)
		var item struct {
			Type string `json:"type"`
		}
		err := json.Unmarshal(jsonStr, &item)
		lib.CheckFatalError(err)

		types = append(types, item.Type)
	}

	jsonStr, _ := json.Marshal(types, jsontext.WithIndent("  "))
	os.WriteFile("main-ips-resourceTypes.json", jsonStr, 0644)
	lib.JsonMarshalAndPrint(types)
	fmt.Println(len(types))
}

func CheckArrayLength(fileName string) {
	file, err := os.ReadFile(fileName)
	lib.CheckFatalError(err)

	var fileData []interface{}
	err = json.Unmarshal(file, &fileData)

	fmt.Println(fileName + ": " + strconv.Itoa(len(fileData)))
}

type ResourceGraphResponse struct {
	Count           int64         `json:"count,omitempty,omitzero" bson:"count,omitempty,omitzero"`
	Data            []interface{} `json:"data,omitempty,omitzero" bson:"data,omitempty,omitzero"`
	Facets          []any         `json:"facets,omitempty,omitzero" bson:"facets,omitempty,omitzero"`
	ResultTruncated string        `json:"resultTruncated,omitempty,omitzero" bson:"resultTruncated,omitempty,omitzero"`
	TotalRecords    int64         `json:"totalRecords,omitempty,omitzero" bson:"totalRecords,omitempty,omitzero"`
	SkipToken       string        `json:"$skipToken,omitempty,omitzero" bson:"$skipToken,omitempty,omitzero"`
}

type IPAddressesVirtualMachine struct {
	AssociatedNics []string          `json:"associatedNics"`
	ID             string            `json:"id"`
	Name           string            `json:"name"`
	PrivateIps     []net.IP          `json:"privateIps"`
	PublicIps      []net.IP          `json:"publicIps"`
	SubscriptionID string            `json:"subscriptionId"`
	Tags           map[string]string `json:"tags"`
	TenantID       string            `json:"tenantId"`
	Type           string            `json:"type"`
}

type IPAddressesNetworkInterface struct {
	AttachedTo     string            `json:"attachedTo"`
	ID             string            `json:"id"`
	IsAttached     float64           `json:"isAttached"`
	Name           string            `json:"name"`
	PrivateIps     []net.IP          `json:"privateIps"`
	PublicIps      []net.IP          `json:"publicIps"`
	ResourceGroup  string            `json:"resourceGroup"`
	SubscriptionID string            `json:"subscriptionId"`
	Tags           map[string]string `json:"tags"`
	TenantID       string            `json:"tenantId"`
	Type           string            `json:"type"`
}

func (t *IPAddressesAllTypes) UnmarshalJSON(data []byte) error {
	if string(data) == "null" || string(data) == `""` {
		return nil
	}

	var nicJson struct {
		AttachedTo     string   `json:"attachedTo"`
		ID             string   `json:"id"`
		IsAttached     float64  `json:"isAttached"`
		Name           string   `json:"name"`
		PrivateIps     []net.IP `json:"privateIps"`
		PublicIps      []net.IP `json:"publicIps"`
		ResourceGroup  string   `json:"resourceGroup"`
		SubscriptionID string   `json:"subscriptionId"`
		Tags           string   `json:"tags"`
		TenantID       string   `json:"tenantId"`
		Type           string   `json:"type"`
	}

	if err := json.Unmarshal(data, &nicJson); err != nil {
		return err
	}

	var tags map[string]string

	if err := json.Unmarshal([]byte(nicJson.Tags), &tags); err != nil {
		return err
	}

	*t = IPAddressesAllTypes{
		AttachedTo:    nicJson.AttachedTo,
		ID:            nicJson.ID,
		IsAttached:    nicJson.IsAttached,
		Name:          nicJson.Name,
		PrivateIps:    nicJson.PrivateIps,
		PublicIps:     nicJson.PublicIps,
		ResourceGroup: nicJson.ResourceGroup,
		TenantID:      nicJson.TenantID,
		Type:          nicJson.Type,
		Tags:          tags,
	}

	return nil
}

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
| project id, name = vmName, type, privateIps, publicIps, tenantId, subscriptionId, associatedNics, tags`

const GetIPAddressesQueryLoadBalancers = `Resources
| where type =~ 'microsoft.network/loadbalancers'
| mv-expand feIpConfig = properties.frontendIPConfigurations
| project lbId = id, lbName = name, type, privateIp = feIpConfig.properties.privateIPAddress, publicIpId = tostring(feIpConfig.properties.publicIPAddress.id), tenantId, subscriptionId, tags = dynamic_to_json(tags)
| join kind=leftouter (
    Resources
    | where type =~ 'microsoft.network/publicipaddresses'
    | project publicIpId = id, publicIp = properties.ipAddress
    )
    on publicIpId
| project-away publicIpId, publicIpId1
| summarize privateIps = make_list(privateIp), publicIps = make_list(publicIp) by id = lbId, name = lbName, type, tenantId, subscriptionId, tags`

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

const GetIPAddressesQueryPublicIPs = `resources
| where type =~ 'microsoft.network/publicipaddresses'
| extend ipConfig = properties.ipConfiguration.id
| extend isAttached = isnotnull(ipConfig)
| project id, name, type, tenantId, subscriptionId, resourceGroup, tags, publicIps = pack_array(properties.ipAddress), ipConfig, isAttached`

const GetIPAddressesQueryWebSites = `resources
| where type =~ 'microsoft.web/sites'
| extend possibleInboundIps = split(properties.possibleInboundIpAddresses, ',')
| extend possibleOutboundIps = split(properties.possibleOutboundIpAddresses, ',')
| extend inboundIps = split(properties.inboundIpAddress, ',')
| extend outboundIps = split(properties.outboundIpAddresses, ',')
| extend privateIps = array_concat(possibleInboundIps, possibleOutboundIps)
| project id, name, type, tenantId, subscriptionId, resourceGroup, tags, privateIps, possibleInboundIps, possibleOutboundIps, inboundIps, outboundIps`

const GetIPAddressesQueryManagedClusters = `resources
| where type =~ 'microsoft.containerservice/managedclusters'
| mv-expand agentPools = properties.agentPoolProfiles
| project id, name, type, tenantId, subscriptionId, resourceGroup, tags`

const GetIPAddressesQueryFirewalls = `resources
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

type ResourceManagerResponse struct {
	NextLink string        `json:"nextLink"`
	Value    []interface{} `json:"value"`
}

type AzureResourceManagedClusterAgentPool struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Properties struct {
		AvailabilityZones          []string `json:"availabilityZones"`
		Count                      float64  `json:"count"`
		CurrentOrchestratorVersion string   `json:"currentOrchestratorVersion"`
		EnableAutoScaling          bool     `json:"enableAutoScaling"`
		EnableEncryptionAtHost     bool     `json:"enableEncryptionAtHost"`
		EnableFips                 bool     `json:"enableFIPS"`
		EnableNodePublicIp         bool     `json:"enableNodePublicIP"`
		EnableUltraSsd             bool     `json:"enableUltraSSD"`
		KubeletDiskType            string   `json:"kubeletDiskType"`
		MaxCount                   float64  `json:"maxCount"`
		MaxPods                    float64  `json:"maxPods"`
		MinCount                   float64  `json:"minCount"`
		Mode                       string   `json:"mode"`
		NodeImageVersion           string   `json:"nodeImageVersion"`
		OrchestratorVersion        string   `json:"orchestratorVersion"`
		OSDiskSizeGb               float64  `json:"osDiskSizeGB"`
		OSDiskType                 string   `json:"osDiskType"`
		OSSku                      string   `json:"osSKU"`
		OSType                     string   `json:"osType"`
		PowerState                 struct {
			Code string `json:"code"`
		} `json:"powerState"`
		ProvisioningState string `json:"provisioningState"`
		ScaleDownMode     string `json:"scaleDownMode"`
		SecurityProfile   struct {
			EnableSecureBoot bool `json:"enableSecureBoot"`
			EnableVtpm       bool `json:"enableVTPM"`
		} `json:"securityProfile"`
		Type            string   `json:"type"`
		UpgradeSettings struct{} `json:"upgradeSettings"`
		VmSize          string   `json:"vmSize"`
		VnetSubnetID    string   `json:"vnetSubnetID"`
		WorkloadRuntime string   `json:"workloadRuntime"`
	} `json:"properties"`
	Type string `json:"type"`
}

type AzureResourceManagedClusterAgentPoolMachine struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Properties struct {
		Network struct {
			IpAddresses []struct {
				Family string `json:"family"`
				Ip     string `json:"ip"`
			} `json:"ipAddresses"`
		} `json:"network"`
		ResourceID string `json:"resourceId"`
	} `json:"properties"`
	Type  string   `json:"type"`
	Zones []string `json:"zones"`
}

type IPAddressesAllTypes struct {
	AssociatedNics        []string          `json:"associatedNics,omitempty"`
	AttachedTo            string            `json:"attachedTo,omitempty"`
	ID                    string            `json:"id"`
	InboundIps            []net.IP          `json:"inboundIps,omitempty"`
	IpConfig              *string           `json:"ipConfig,omitempty"`
	IsAttached            float64           `json:"isAttached,omitempty"`
	Name                  string            `json:"name"`
	OutboundIps           []net.IP          `json:"outboundIps,omitempty"`
	PossibleInboundIps    []net.IP          `json:"possibleInboundIps,omitempty"`
	PossibleOutboundIps   []net.IP          `json:"possibleOutboundIps,omitempty"`
	PrivateIps            []net.IP          `json:"privateIps,omitempty"`
	PublicIpIds           []string          `json:"publicIpIds,omitempty"`
	PublicIps             []net.IP          `json:"publicIps,omitempty"`
	PublicNetworkAccess   string            `json:"publicNetworkAccess,omitempty"`
	ResourceGroup         string            `json:"resourceGroup,omitempty"`
	SubscriptionID        string            `json:"subscriptionId"`
	Tags                  map[string]string `json:"tags,omitempty"`
	TenantID              string            `json:"tenantId"`
	Type                  string            `json:"type,omitempty"`
	VaultURI              string            `json:"vaultUri,omitempty"`
	AssociatedResourceIDs []string          `json:"associatedResourceIDs,omitempty"`
}

func GetManagedClusterIPAddresses(clusterIpObject IPAddressesAllTypes, token *lib.AzureMultiAuthToken) (ipAddresses []IPAddressesAllTypes) {
	baseUrl := "https://management.azure.com"
	agentPools := GetManagedClusterAgentPools(clusterIpObject.ID, token)

	var allIPs []net.IP
	var associatedIDs []string

	for _, ap := range agentPools {
		urlString := baseUrl + ap.ID + "/machines?api-version=2025-05-01"
		res, err := azure.HttpGet(urlString, *token)
		lib.CheckFatalError(err)

		var resData ResourceManagerResponse
		err = json.Unmarshal(res, &resData)

		var apIPs []net.IP
		var machineIDs []string

		for _, m := range resData.Value {
			jsonStr, _ := json.Marshal(m)
			var machine AzureResourceManagedClusterAgentPoolMachine
			err := json.Unmarshal(jsonStr, &machine)
			lib.CheckFatalError(err)
			var machineIPs []net.IP
			for _, ip := range machine.Properties.Network.IpAddresses {
				parsedIp := net.ParseIP(ip.Ip)
				machineIPs = append(machineIPs, parsedIp)
				apIPs = append(apIPs, parsedIp)
				allIPs = append(allIPs, parsedIp)
			}
			machineIDs = append(machineIDs, machine.ID)
			associatedIDs = append(associatedIDs, machine.ID)

			machineIpObject := IPAddressesAllTypes{
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

		apIpObject := IPAddressesAllTypes{
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

func GetManagedClusterAgentPools(clusterId string, token *lib.AzureMultiAuthToken) (agentPools []AzureResourceManagedClusterAgentPool) {
	baseUrl := "https://management.azure.com"
	urlString := baseUrl + clusterId + "/agentPools?api-version=2025-05-01"

	res, err := azure.HttpGet(urlString, *token)
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

		res, err := azure.HttpGet(nextLink, *token)
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

// func GetAllIPsNetworkInterfaces(outputFile string, token *lib.AzureMultiAuthToken) {
// 	urlString := "https://management.azure.com/providers/Microsoft.ResourceGraph/resources?api-version=2022-10-01"

// 	graphQuery := `Resources
// | where type =~ 'microsoft.network/networkinterfaces'
// | mv-expand ipconfig = properties.ipConfigurations
// | extend publicIpId = tostring(ipconfig.properties.publicIPAddress.id), tags = iff(isnull(tags), dynamic({}), tags)
// | extend bareMetalServer= properties.bareMetalServer, privateEndpoint = properties.privateEndpoint, privateLinkService = properties.privateLinkService, virtualMachine = properties.virtualMachine, name, id, properties
// | extend isAttached = isnotnull(bareMetalServer) or isnotnull(privateEndpoint) or isnotnull(privateLinkService) or isnotnull(virtualMachine)
// | extend attachedTo = dynamic_to_json(coalesce(bareMetalServer.id, privateLinkService.id, privateEndpoint.id, virtualMachine.id))
// | join kind=leftouter  (resources | project id, publicIp = properties.ipAddress) on $left.publicIpId == $right.['id']
// | project name, resourceGroup, subscriptionId, tenantId, id, privateIp = ipconfig.properties.privateIPAddress, publicIpId = tostring(ipconfig.properties.publicIPAddress.id), publicIp, type, tags, isAttached, attachedTo
// | summarize privateIps = make_list(privateIp), publicIps = make_list(publicIp)  by id, name, resourceGroup, subscriptionId, tenantId, type, tags= dynamic_to_json(tags), isAttached, attachedTo`

// 	jsonBody := `{
// 	"query": "` + graphQuery + `"
// }`

// 	res, _, err := azure.HttpPost(urlString, jsonBody, *token)
// 	lib.CheckFatalError(err)

// 	// fmt.Println(string(res))
// 	// os.WriteFile("main-ips-"+token.TenantName+".json", res, 0640)
// 	// os.Exit(0)

// 	var resData ResourceGraphResponse
// 	err = json.Unmarshal(res, &resData)
// 	lib.CheckFatalError(err)
// 	var ips []IPAddressesNetworkInterface

// 	for _, nic := range resData.Data {
// 		jsonStr, err := json.Marshal(nic)
// 		lib.CheckFatalError(err)
// 		var nicProcessed IPAddressesNetworkInterface
// 		err = json.Unmarshal(jsonStr, &nicProcessed)
// 		lib.CheckFatalError(err)
// 		lib.JsonMarshalAndPrint(nicProcessed)
// 		os.Exit(0)
// 		lib.CheckFatalError(err)
// 		ips = append(ips, nicProcessed)
// 	}

// }

// func GetAllIPsVirtualMachines(outputFile string, token *lib.AzureMultiAuthToken) (allIPs []IPAddressesVirtualMachine) {
// 	urlString := "https://management.azure.com/providers/Microsoft.ResourceGraph/resources?api-version=2022-10-01"

// 	graphQuery := `Resources
// | where type =~ 'microsoft.compute/virtualmachines'
// | project id, vmId = tolower(tostring(id)), vmName = name, type, tenantId, subscriptionId, tags = iff(isnull(tags), dynamic({}), tags)
// | join (
//     Resources
//     | where type =~ 'microsoft.network/networkinterfaces'
//     | mv-expand ipconfig = properties.ipConfigurations
//     | project vmId = tolower(tostring(properties.virtualMachine.id)), nicId = id, privateIp = ipconfig.properties.privateIPAddress, publicIpId = tostring(ipconfig.properties.publicIPAddress.id)
//     | join kind=leftouter (
//         Resources
//         | where type =~ 'microsoft.network/publicipaddresses'
//         | project publicIpId = id, publicIp = properties.ipAddress
//         )
//         on publicIpId
//     | project-away publicIpId, publicIpId1
//     | summarize associatedNics = make_list(nicId), privateIps = make_list(privateIp), publicIps = make_list(publicIp) by vmId
//     )
//     on vmId
// | project id, name = vmName, type, privateIps, publicIps, tenantId, subscriptionId, associatedNics, tags`

// 	jsonBody := `{
// 	"query": "` + graphQuery + `"
// }`

// 	res, _, err := azure.HttpPost(urlString, jsonBody, *token)
// 	lib.CheckFatalError(err)

// 	var resData ResourceGraphResponse
// 	err = json.Unmarshal(res, &resData)
// 	lib.CheckFatalError(err)

// 	for _, vm := range resData.Data {
// 		jsonStr, err := json.Marshal(vm)
// 		lib.CheckFatalError(err)
// 		var vmProcessed IPAddressesVirtualMachine
// 		err = json.Unmarshal(jsonStr, &vmProcessed)
// 		lib.CheckFatalError(err)
// 		allIPs = append(allIPs, vmProcessed)
// 	}

// 	return
// }

// func GetAllTenantIpAddresses(outputFile string, token *lib.AzureMultiAuthToken) {
// 	// vmIPs := GetAllIPsVirtualMachines(outputFile, token)
// 	// nicIPs := GetAllIPsNetworkInterfaces(outputFile, token)
// }
