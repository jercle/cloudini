package azure

import (
	json "encoding/json/v2"
	"io"
	"net/http"
	"strings"
	"sync"

	"github.com/jercle/cloudini/lib"
)

func ListAllTenantIpAddresses(token lib.AzureMultiAuthToken) IPAddressList {

	var (
		allIpAddresses IPAddressList
		wg             lib.WaitGroupCount
		mutex          sync.Mutex
	)

	vnets := make(chan Vnet, 100)
	publicIpAddresses := make(chan IPAddressItem, 10000)
	privateIpAddresses := make(chan IPAddressItem, 10000)
	done := make(chan bool, 1)

	allSubs, err := ListSubscriptions(token)
	lib.CheckFatalError(err)

	for _, sub := range allSubs {
		wg.Add(1)
		go ListAllSubscriptionVnetsWithChan(sub.SubscriptionID, token, vnets, &wg)
	}

	go func() {
		// chanLoop:
		for {
			select {
			case vnet := <-vnets:
				wg.Add(1)
				go ListAllVnetIPAddressesWithChan(token, vnet, publicIpAddresses, privateIpAddresses, &wg)
			case publicIp := <-publicIpAddresses:
				mutex.Lock()
				allIpAddresses.PublicAddresses = append(allIpAddresses.PublicAddresses, publicIp)
				mutex.Unlock()
			case privateIp := <-privateIpAddresses:
				mutex.Lock()
				allIpAddresses.PrivateAddresses = append(allIpAddresses.PrivateAddresses, privateIp)
				mutex.Unlock()
			case <-done:
				break
				// break chanLoop
			}
		}
	}()

	wg.Wait()
	done <- true
	return allIpAddresses
}

func ListAllSubscriptionVnetsWithChan(subscriptionId string, mat lib.AzureMultiAuthToken, out chan<- Vnet, wg *lib.WaitGroupCount) {
	var (
		allVnets  []Vnet
		listVnets VnetListResponse
	)

	urlString := "https://management.azure.com/subscriptions/" +
		subscriptionId +
		"/providers/Microsoft.Network/virtualNetworks?api-version=2023-09-01"

	req, err := http.NewRequest(http.MethodGet, urlString, nil)
	lib.CheckFatalError(err)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+mat.TokenData.Token)

	res, err := http.DefaultClient.Do(req)
	lib.CheckFatalError(err)

	responseBody, err := io.ReadAll(res.Body)
	lib.CheckFatalError(err)

	json.Unmarshal(responseBody, &listVnets)
	for _, vnet := range listVnets.Value {
		var (
			currentVnet  Vnet
			vnetPeerings []ProcessedVnetPeering
		)

		currentVnet.Name = vnet.Name
		currentVnet.ID = vnet.ID
		currentVnet.ResourceGroup = strings.Split(strings.Split(vnet.ID, "resourceGroups/")[1], "/")[0]
		currentVnet.AddressSpace = vnet.Properties.AddressSpace.AddressPrefixes
		currentVnet.ProvisioningState = vnet.Properties.ProvisioningState
		currentVnet.Location = vnet.Location
		currentVnet.Type = vnet.Type
		currentVnet.Tags = vnet.Tags
		currentVnet.SubscriptionID = subscriptionId
		for _, peering := range vnet.Properties.VirtualNetworkPeerings {
			var remoteGateways []string
			for _, rgw := range peering.Properties.RemoteGateways {
				remoteGateways = append(remoteGateways, rgw.ID)
			}
			currentPeering := ProcessedVnetPeering{
				Name:                      peering.Name,
				RemoteVirtualNetwork:      peering.Properties.RemoteVirtualNetwork.ID,
				RemoteAddressSpace:        peering.Properties.RemoteAddressSpace.AddressPrefixes,
				AllowForwardedTraffic:     peering.Properties.AllowForwardedTraffic,
				AllowGatewayTransit:       peering.Properties.AllowGatewayTransit,
				AllowVirtualNetworkAccess: peering.Properties.AllowVirtualNetworkAccess,
				UseRemoteGateways:         peering.Properties.UseRemoteGateways,
				PeeringState:              peering.Properties.PeeringState,
				ProvisioningState:         peering.Properties.ProvisioningState,
				PeeringSyncLevel:          peering.Properties.PeeringSyncLevel,
				RemoteGateways:            remoteGateways,
			}
			vnetPeerings = append(vnetPeerings, currentPeering)
		}
		currentVnet.VirtualNetworkPeerings = append(currentVnet.VirtualNetworkPeerings, vnetPeerings...)
		allVnets = append(allVnets, currentVnet)

		out <- currentVnet
	}
	wg.Done()
}

//
//

func ListAllVnetIPAddressesWithChan(mat lib.AzureMultiAuthToken, vnet Vnet, publicIps chan<- IPAddressItem, privateIps chan<- IPAddressItem, wg *lib.WaitGroupCount) {
	urlString := "https://management.azure.com/subscriptions/" +
		vnet.SubscriptionID +
		"/resourceGroups/" +
		vnet.ResourceGroup +
		"/providers/Microsoft.Network/virtualNetworks/" +
		vnet.Name +
		"?api-version=2023-02-01&$expand=subnets/ipConfigurations"

	req, err := http.NewRequest(http.MethodGet, urlString, nil)
	lib.CheckFatalError(err)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+mat.TokenData.Token)

	res, err := http.DefaultClient.Do(req)
	lib.CheckFatalError(err)

	responseBody, err := io.ReadAll(res.Body)
	lib.CheckFatalError(err)

	var vnetIpConfig SubnetIPConfigResponse
	err = json.Unmarshal(responseBody, &vnetIpConfig)
	lib.CheckFatalError(err)

	subnets := vnetIpConfig.Properties.Subnets

	for _, sn := range subnets {
		ipConfigs := sn.Properties.IpConfigurations
		if len(ipConfigs) > 0 {
			for _, conf := range ipConfigs {
				confId := strings.Split(conf.ID, "ipConfigurations")[0]
				confUrl := "https://management.azure.com" + confId + "?api-version=2023-02-01"
				var resourceResp IPAddressItem
				result, err := HttpGet(confUrl, mat)
				lib.CheckFatalError(err)
				json.Unmarshal(result, &resourceResp)

				ipAddressItem := IPAddressItem{
					ResourceName: resourceResp.ResourceName,
					ResourceID:   resourceResp.ResourceID,
					ResourceType: resourceResp.ResourceType,
					Subnet:       sn.Name,
					Vnet:         vnet.Name,
					Tags:         resourceResp.Tags,
				}

				if conf.Properties.PrivateIpAddress != "" {
					// Is a private IP
					ipAddressItem.IpAddress = conf.Properties.PrivateIpAddress
					privateIps <- ipAddressItem
				}

				if conf.Properties.PublicIpAddress != nil {
					// Is a public IP
					pubAddressUrl := "https://management.azure.com" + conf.Properties.PublicIpAddress.ID + "?api-version=2023-02-01"
					result, err := HttpGet(pubAddressUrl, mat)
					lib.CheckFatalError(err)

					var publicIp PublicIpAddress
					json.Unmarshal(result, &publicIp)
					ipAddressItem.IpAddress = publicIp.Properties.IpAddress
					ipAddressItem.ResourceName = publicIp.Name
					ipAddressItem.ResourceID = publicIp.ID
					ipAddressItem.ResourceType = publicIp.Type
					ipAddressItem.Tags = publicIp.Tags
					publicIps <- ipAddressItem
				}
			}
		}
	}
	wg.Done()
}

//
//

func ListAllVnetIPAddresses(mat lib.AzureMultiAuthToken, vnet Vnet) IPAddressList {

	var allVnetIps IPAddressList

	urlString := "https://management.azure.com/subscriptions/" +
		vnet.SubscriptionID +
		"/resourceGroups/" +
		vnet.ResourceGroup +
		"/providers/Microsoft.Network/virtualNetworks/" +
		vnet.Name +
		"?api-version=2023-02-01&$expand=subnets/ipConfigurations"

	req, err := http.NewRequest(http.MethodGet, urlString, nil)
	lib.CheckFatalError(err)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+mat.TokenData.Token)

	res, err := http.DefaultClient.Do(req)
	lib.CheckFatalError(err)

	responseBody, err := io.ReadAll(res.Body)
	lib.CheckFatalError(err)

	var vnetIpConfig SubnetIPConfigResponse
	err = json.Unmarshal(responseBody, &vnetIpConfig)
	lib.CheckFatalError(err)

	subnets := vnetIpConfig.Properties.Subnets

	for _, sn := range subnets {
		ipConfigs := sn.Properties.IpConfigurations
		if len(ipConfigs) > 0 {
			for _, conf := range ipConfigs {
				confId := strings.Split(conf.ID, "ipConfigurations")[0]
				confUrl := "https://management.azure.com" + confId + "?api-version=2023-02-01"
				var resourceResp IPAddressItem
				result, err := HttpGet(confUrl, mat)
				lib.CheckFatalError(err)
				json.Unmarshal(result, &resourceResp)

				ipAddressItem := IPAddressItem{
					ResourceName: resourceResp.ResourceName,
					ResourceID:   resourceResp.ResourceID,
					ResourceType: resourceResp.ResourceType,
					Subnet:       sn.Name,
					Vnet:         vnet.Name,
					Tags:         resourceResp.Tags,
				}

				if conf.Properties.PrivateIpAddress != "" {
					// Is a private IP
					ipAddressItem.IpAddress = conf.Properties.PrivateIpAddress
					allVnetIps.PrivateAddresses = append(allVnetIps.PrivateAddresses, ipAddressItem)
				}

				if conf.Properties.PublicIpAddress != nil {
					// Is a public IP
					pubAddressUrl := "https://management.azure.com" + conf.Properties.PublicIpAddress.ID + "?api-version=2023-02-01"
					result, err := HttpGet(pubAddressUrl, mat)
					lib.CheckFatalError(err)

					var publicIp PublicIpAddress
					json.Unmarshal(result, &publicIp)
					ipAddressItem.IpAddress = publicIp.Properties.IpAddress
					ipAddressItem.ResourceName = publicIp.Name
					ipAddressItem.ResourceID = publicIp.ID
					ipAddressItem.ResourceType = publicIp.Type
					ipAddressItem.Tags = publicIp.Tags
					allVnetIps.PublicAddresses = append(allVnetIps.PublicAddresses, ipAddressItem)
				}
			}
		}
	}
	return allVnetIps
}

//
//

func ListAllSubscriptionVnets(subscriptionId string, mat lib.AzureMultiAuthToken) []Vnet {
	var (
		allVnets  []Vnet
		listVnets VnetListResponse
	)

	urlString := "https://management.azure.com/subscriptions/" +
		subscriptionId +
		"/providers/Microsoft.Network/virtualNetworks?api-version=2023-09-01"

	req, err := http.NewRequest(http.MethodGet, urlString, nil)
	lib.CheckFatalError(err)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+mat.TokenData.Token)

	res, err := http.DefaultClient.Do(req)
	lib.CheckFatalError(err)

	responseBody, err := io.ReadAll(res.Body)
	lib.CheckFatalError(err)

	json.Unmarshal(responseBody, &listVnets)

	for _, vnet := range listVnets.Value {
		var (
			currentVnet  Vnet
			vnetPeerings []ProcessedVnetPeering
		)

		currentVnet.Name = vnet.Name
		currentVnet.ID = vnet.ID
		currentVnet.ResourceGroup = strings.Split(strings.Split(vnet.ID, "resourceGroups/")[1], "/")[0]
		currentVnet.AddressSpace = vnet.Properties.AddressSpace.AddressPrefixes
		currentVnet.ProvisioningState = vnet.Properties.ProvisioningState
		currentVnet.Location = vnet.Location
		currentVnet.Type = vnet.Type
		currentVnet.Tags = vnet.Tags
		currentVnet.SubscriptionID = subscriptionId
		for _, peering := range vnet.Properties.VirtualNetworkPeerings {
			var remoteGateways []string
			for _, rgw := range peering.Properties.RemoteGateways {
				remoteGateways = append(remoteGateways, rgw.ID)
			}
			currentPeering := ProcessedVnetPeering{
				Name:                      peering.Name,
				RemoteVirtualNetwork:      peering.Properties.RemoteVirtualNetwork.ID,
				RemoteAddressSpace:        peering.Properties.RemoteAddressSpace.AddressPrefixes,
				AllowForwardedTraffic:     peering.Properties.AllowForwardedTraffic,
				AllowGatewayTransit:       peering.Properties.AllowGatewayTransit,
				AllowVirtualNetworkAccess: peering.Properties.AllowVirtualNetworkAccess,
				UseRemoteGateways:         peering.Properties.UseRemoteGateways,
				PeeringState:              peering.Properties.PeeringState,
				ProvisioningState:         peering.Properties.ProvisioningState,
				PeeringSyncLevel:          peering.Properties.PeeringSyncLevel,
				RemoteGateways:            remoteGateways,
			}
			vnetPeerings = append(vnetPeerings, currentPeering)
		}
		currentVnet.VirtualNetworkPeerings = append(currentVnet.VirtualNetworkPeerings, vnetPeerings...)
		allVnets = append(allVnets, currentVnet)
	}
	return allVnets
}

//
//

func ListVnetSubnets(subscriptionId string, resourceGroupName string, virtualNetworkName string, mat lib.AzureMultiAuthToken) []SubnetResponse {
	var listSubnetResponse ListSubnetsResponse

	urlString := "https://management.azure.com/subscriptions/" +
		subscriptionId +
		"/resourceGroups/" +
		resourceGroupName +
		"/providers/Microsoft.Network/virtualNetworks/" +
		virtualNetworkName +
		"/subnets?api-version=2023-09-01"

	response, err := HttpGet(urlString, mat)
	lib.CheckFatalError(err)

	json.Unmarshal(response, &listSubnetResponse)

	return listSubnetResponse.Value
}

//
//
