package azure

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"sync"

	"github.com/jercle/azg/lib"
)

type VnetResponse struct {
	Etag       string `json:"etag"`
	ID         string `json:"id"`
	Location   string `json:"location"`
	Name       string `json:"name"`
	Properties struct {
		AddressSpace struct {
			AddressPrefixes []string `json:"addressPrefixes"`
		} `json:"addressSpace"`
		DhcpOptions struct {
			DnsServers []any `json:"dnsServers"`
		} `json:"dhcpOptions"`
		EnableDdosProtection bool   `json:"enableDdosProtection"`
		ProvisioningState    string `json:"provisioningState"`
		ResourceGuid         string `json:"resourceGuid"`
		Subnets              []struct {
			Etag       string `json:"etag"`
			ID         string `json:"id"`
			Name       string `json:"name"`
			Properties struct {
				AddressPrefix    string `json:"addressPrefix"`
				Delegations      []any  `json:"delegations"`
				IpConfigurations []struct {
					ID string `json:"id"`
				} `json:"ipConfigurations"`
				NetworkSecurityGroup *struct {
					ID string `json:"id"`
				} `json:"networkSecurityGroup,omitempty"`
				PrivateEndpointNetworkPolicies string `json:"privateEndpointNetworkPolicies"`
				PrivateEndpoints               []struct {
					ID string `json:"id"`
				} `json:"privateEndpoints,omitempty"`
				PrivateLinkServiceNetworkPolicies string `json:"privateLinkServiceNetworkPolicies"`
				ProvisioningState                 string `json:"provisioningState"`
				Purpose                           string `json:"purpose,omitempty"`
				RouteTable                        *struct {
					ID string `json:"id"`
				} `json:"routeTable,omitempty"`
				ServiceEndpoints []struct {
					Locations         []string `json:"locations"`
					ProvisioningState string   `json:"provisioningState"`
					Service           string   `json:"service"`
				} `json:"serviceEndpoints"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"subnets"`
		VirtualNetworkPeerings []struct {
			Etag       string `json:"etag"`
			ID         string `json:"id"`
			Name       string `json:"name"`
			Properties struct {
				AllowForwardedTraffic     bool   `json:"allowForwardedTraffic"`
				AllowGatewayTransit       bool   `json:"allowGatewayTransit"`
				AllowVirtualNetworkAccess bool   `json:"allowVirtualNetworkAccess"`
				DoNotVerifyRemoteGateways bool   `json:"doNotVerifyRemoteGateways"`
				PeerCompleteVnets         bool   `json:"peerCompleteVnets"`
				PeeringState              string `json:"peeringState"`
				PeeringSyncLevel          string `json:"peeringSyncLevel"`
				ProvisioningState         string `json:"provisioningState"`
				RemoteAddressSpace        struct {
					AddressPrefixes []string `json:"addressPrefixes"`
				} `json:"remoteAddressSpace"`
				RemoteGateways []struct {
					ID string `json:"id"`
				} `json:"remoteGateways,omitempty"`
				RemoteVirtualNetwork struct {
					ID string `json:"id"`
				} `json:"remoteVirtualNetwork"`
				RemoteVirtualNetworkAddressSpace struct {
					AddressPrefixes []string `json:"addressPrefixes"`
				} `json:"remoteVirtualNetworkAddressSpace"`
				ResourceGuid     string `json:"resourceGuid"`
				RouteServiceVips struct {
					Af36ba888c9943f4A5e38fa90652cc96 string `json:"af36ba88-8c99-43f4-a5e3-8fa90652cc96,omitempty"`
				} `json:"routeServiceVips"`
				UseRemoteGateways bool `json:"useRemoteGateways"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"virtualNetworkPeerings"`
	} `json:"properties"`
	Tags map[string]string `json:"tags"`
	Type string            `json:"type"`
}

type VnetListResponse struct {
	Value []VnetResponse
}

type Vnet struct {
	ID                     string                 `json:"id"`
	Location               string                 `json:"location"`
	Name                   string                 `json:"name"`
	ResourceGroup          string                 `json:"resourceGroup"`
	SubscriptionID         string                 `json:"subscriptionId`
	AddressSpace           []string               `json:"addressSpace"`
	Subnets                []SubnetResponse       `json:"subnets"`
	ProvisioningState      string                 `json:"provisioningState"`
	VirtualNetworkPeerings []ProcessedVnetPeering `json:"virtualNetworkPeerings"`
	Tags                   map[string]string      `json:"tags"`
	Type                   string                 `json:"type"`
}

type VirtualNetworkPeering struct {
	Etag       string `json:"etag"`
	ID         string `json:"id"`
	Name       string `json:"name"`
	Properties struct {
		AllowForwardedTraffic     bool   `json:"allowForwardedTraffic"`
		AllowGatewayTransit       bool   `json:"allowGatewayTransit"`
		AllowVirtualNetworkAccess bool   `json:"allowVirtualNetworkAccess"`
		DoNotVerifyRemoteGateways bool   `json:"doNotVerifyRemoteGateways"`
		PeerCompleteVnets         bool   `json:"peerCompleteVnets"`
		PeeringState              string `json:"peeringState"`
		PeeringSyncLevel          string `json:"peeringSyncLevel"`
		ProvisioningState         string `json:"provisioningState"`
		RemoteAddressSpace        struct {
			AddressPrefixes []string `json:"addressPrefixes"`
		} `json:"remoteAddressSpace"`
		RemoteGateways []struct {
			ID string `json:"id"`
		} `json:"remoteGateways,omitempty"`
		RemoteVirtualNetwork struct {
			ID string `json:"id"`
		} `json:"remoteVirtualNetwork"`
		RemoteVirtualNetworkAddressSpace struct {
			AddressPrefixes []string `json:"addressPrefixes"`
		} `json:"remoteVirtualNetworkAddressSpace"`
		ResourceGuid     string `json:"resourceGuid"`
		RouteServiceVips struct {
			Af36ba888c9943f4A5e38fa90652cc96 string `json:"af36ba88-8c99-43f4-a5e3-8fa90652cc96,omitempty"`
		} `json:"routeServiceVips"`
		UseRemoteGateways bool `json:"useRemoteGateways"`
	} `json:"properties"`
	Type string `json:"type"`
}

type ProcessedVnetPeering struct {
	Name                      string   `json:"name"`
	RemoteAddressSpace        []string `json:"remoteAddressSpace"`
	RemoteGateways            []string `json:"remoteGateways"`
	RemoteVirtualNetwork      string   `json:"remoteVirtualNetwork"`
	UseRemoteGateways         bool     `json:"useRemoteGateways"`
	AllowForwardedTraffic     bool     `json:"allowForwardedTraffic"`
	AllowGatewayTransit       bool     `json:"allowGatewayTransit"`
	AllowVirtualNetworkAccess bool     `json:"allowVirtualNetworkAccess"`
	PeeringState              string   `json:"peeringState"`
	ProvisioningState         string   `json:"provisioningState"`
	PeeringSyncLevel          string   `json:"peeringSyncLevel"`
}

type SubnetResponse struct {
	Etag       string `json:"etag"`
	ID         string `json:"id"`
	Name       string `json:"name"`
	Properties struct {
		AddressPrefix string `json:"addressPrefix"`
		Delegations   []struct {
			Etag       string `json:"etag"`
			ID         string `json:"id"`
			Name       string `json:"name"`
			Properties struct {
				Actions           []string `json:"actions"`
				ProvisioningState string   `json:"provisioningState"`
				ServiceName       string   `json:"serviceName"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"delegations"`
		IpConfigurations []struct {
			ID string `json:"id"`
		} `json:"ipConfigurations"`
		NetworkSecurityGroup struct {
			ID string `json:"id"`
		} `json:"networkSecurityGroup"`
		PrivateEndpointNetworkPolicies    string `json:"privateEndpointNetworkPolicies"`
		PrivateLinkServiceNetworkPolicies string `json:"privateLinkServiceNetworkPolicies"`
		ProvisioningState                 string `json:"provisioningState"`
		Purpose                           string `json:"purpose"`
		ServiceEndpoints                  []any  `json:"serviceEndpoints"`
	} `json:"properties"`
	Type string `json:"type"`
}

type ListSubnetsResponse struct {
	Value []SubnetResponse `json:"value"`
}

type ProcessedSubnet struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	VnetName   string `json:"vnetName"`
	Properties struct {
		AddressPrefix string `json:"addressPrefix"`
		Delegations   []struct {
			ID         string `json:"id"`
			Name       string `json:"name"`
			Properties struct {
				Actions           []string `json:"actions"`
				ProvisioningState string   `json:"provisioningState"`
				ServiceName       string   `json:"serviceName"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"delegations"`
		IpConfigurations []struct {
			ID string `json:"id"`
		} `json:"ipConfigurations"`
		NetworkSecurityGroup struct {
			ID string `json:"id"`
		} `json:"networkSecurityGroup"`
		PrivateEndpointNetworkPolicies    string `json:"privateEndpointNetworkPolicies"`
		PrivateLinkServiceNetworkPolicies string `json:"privateLinkServiceNetworkPolicies"`
		ProvisioningState                 string `json:"provisioningState"`
		Purpose                           string `json:"purpose"`
		ServiceEndpoints                  []any  `json:"serviceEndpoints"`
	} `json:"properties"`
	Type string `json:"type"`
}

type SubnetIPConfigResponse struct {
	Etag       string `json:"etag"`
	ID         string `json:"id"`
	Location   string `json:"location"`
	Name       string `json:"name"`
	Properties struct {
		AddressSpace struct {
			AddressPrefixes []string `json:"addressPrefixes"`
		} `json:"addressSpace"`
		DhcpOptions struct {
			DnsServers []any `json:"dnsServers"`
		} `json:"dhcpOptions"`
		EnableDdosProtection bool   `json:"enableDdosProtection"`
		ProvisioningState    string `json:"provisioningState"`
		ResourceGuid         string `json:"resourceGuid"`
		Subnets              []struct {
			Etag       string `json:"etag"`
			ID         string `json:"id"`
			Name       string `json:"name"`
			Properties struct {
				AddressPrefix    string `json:"addressPrefix"`
				Delegations      []any  `json:"delegations"`
				IpConfigurations []struct {
					Etag       string `json:"etag"`
					ID         string `json:"id"`
					Name       string `json:"name"`
					Resource   IPAddressItem
					Properties struct {
						Primary                         bool   `json:"primary,omitempty"`
						PrivateIpAddress                string `json:"privateIPAddress,omitempty"`
						PrivateIpAddressVersion         string `json:"privateIPAddressVersion,omitempty"`
						PrivateIpAllocationMethod       string `json:"privateIPAllocationMethod"`
						PrivateLinkConnectionProperties *struct {
							Fqdns              []string `json:"fqdns"`
							GroupID            string   `json:"groupId"`
							RequiredMemberName string   `json:"requiredMemberName"`
						} `json:"privateLinkConnectionProperties,omitempty"`
						ProvisioningState string `json:"provisioningState"`
						PublicIpAddress   *struct {
							ID string `json:"id"`
						} `json:"publicIPAddress,omitempty"`
						Subnet struct {
							ID string `json:"id"`
						} `json:"subnet"`
					} `json:"properties"`
					Type string `json:"type"`
				} `json:"ipConfigurations"`
				NetworkSecurityGroup *struct {
					ID string `json:"id"`
				} `json:"networkSecurityGroup,omitempty"`
				PrivateEndpointNetworkPolicies string `json:"privateEndpointNetworkPolicies"`
				PrivateEndpoints               []struct {
					ID string `json:"id"`
				} `json:"privateEndpoints,omitempty"`
				PrivateLinkServiceNetworkPolicies string `json:"privateLinkServiceNetworkPolicies"`
				ProvisioningState                 string `json:"provisioningState"`
				Purpose                           string `json:"purpose,omitempty"`
				RouteTable                        *struct {
					ID string `json:"id"`
				} `json:"routeTable,omitempty"`
				ServiceEndpoints []struct {
					Locations         []string `json:"locations"`
					ProvisioningState string   `json:"provisioningState"`
					Service           string   `json:"service"`
				} `json:"serviceEndpoints"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"subnets"`
		VirtualNetworkPeerings []struct {
			Etag       string `json:"etag"`
			ID         string `json:"id"`
			Name       string `json:"name"`
			Properties struct {
				AllowForwardedTraffic     bool   `json:"allowForwardedTraffic"`
				AllowGatewayTransit       bool   `json:"allowGatewayTransit"`
				AllowVirtualNetworkAccess bool   `json:"allowVirtualNetworkAccess"`
				DoNotVerifyRemoteGateways bool   `json:"doNotVerifyRemoteGateways"`
				PeerCompleteVnets         bool   `json:"peerCompleteVnets"`
				PeeringState              string `json:"peeringState"`
				PeeringSyncLevel          string `json:"peeringSyncLevel"`
				ProvisioningState         string `json:"provisioningState"`
				RemoteAddressSpace        struct {
					AddressPrefixes []string `json:"addressPrefixes"`
				} `json:"remoteAddressSpace"`
				RemoteGateways []struct {
					ID string `json:"id"`
				} `json:"remoteGateways,omitempty"`
				RemoteVirtualNetwork struct {
					ID string `json:"id"`
				} `json:"remoteVirtualNetwork"`
				RemoteVirtualNetworkAddressSpace struct {
					AddressPrefixes []string `json:"addressPrefixes"`
				} `json:"remoteVirtualNetworkAddressSpace"`
				ResourceGuid     string `json:"resourceGuid"`
				RouteServiceVips struct {
					Af36ba888c9943f4A5e38fa90652cc96 string `json:"af36ba88-8c99-43f4-a5e3-8fa90652cc96,omitempty"`
				} `json:"routeServiceVips"`
				UseRemoteGateways bool `json:"useRemoteGateways"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"virtualNetworkPeerings"`
	} `json:"properties"`
	Tags map[string]string `json:"tags"`
	Type string            `json:"type"`
}

type PublicIpAddress struct {
	Etag       string `json:"etag"`
	ID         string `json:"id"`
	Location   string `json:"location"`
	Name       string `json:"name"`
	Properties struct {
		DdosSettings struct {
			ProtectionMode string `json:"protectionMode"`
		} `json:"ddosSettings"`
		IdleTimeoutInMinutes int    `json:"idleTimeoutInMinutes"`
		IpAddress            string `json:"ipAddress"`
		IpConfiguration      struct {
			ID string `json:"id"`
		} `json:"ipConfiguration"`
		IpTags                   []any  `json:"ipTags"`
		ProvisioningState        string `json:"provisioningState"`
		PublicIpAddressVersion   string `json:"publicIPAddressVersion"`
		PublicIpAllocationMethod string `json:"publicIPAllocationMethod"`
		ResourceGuid             string `json:"resourceGuid"`
	} `json:"properties"`
	Sku struct {
		Name string `json:"name"`
		Tier string `json:"tier"`
	} `json:"sku"`
	Tags map[string]string `json:"tags"`
	Type string            `json:"type"`
}

type IPAddressItem struct {
	ResourceName string `json:"name"`
	ResourceID   string `json:"id"`
	ResourceType string `json:"type"`
	IpAddress    string
	Vnet         string
	Subnet       string
	Tags         map[string]string
}

type IPAddressList struct {
	PrivateAddresses []IPAddressItem
	PublicAddresses  []IPAddressItem
}

type AllTenantIPs map[string]IPAddressList

func ListAllTenantIpAddresses(token MultiAuthToken) IPAddressList {

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

func ListAllSubscriptionVnetsWithChan(subscriptionId string, mat MultiAuthToken, out chan<- Vnet, wg *lib.WaitGroupCount) {
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

func ListAllVnetIPAddressesWithChan(mat MultiAuthToken, vnet Vnet, publicIps chan<- IPAddressItem, privateIps chan<- IPAddressItem, wg *lib.WaitGroupCount) {
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
				result := HttpGet(confUrl, mat)
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
					result := HttpGet(pubAddressUrl, mat)

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

func ListAllVnetIPAddresses(mat MultiAuthToken, vnet Vnet) IPAddressList {

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
				result := HttpGet(confUrl, mat)
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
					result := HttpGet(pubAddressUrl, mat)

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

func ListAllSubscriptionVnets(subscriptionId string, mat MultiAuthToken) []Vnet {
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

func ListVnetSubnets(subscriptionId string, resourceGroupName string, virtualNetworkName string, mat MultiAuthToken) []SubnetResponse {
	var listSubnetResponse ListSubnetsResponse

	urlString := "https://management.azure.com/subscriptions/" +
		subscriptionId +
		"/resourceGroups/" +
		resourceGroupName +
		"/providers/Microsoft.Network/virtualNetworks/" +
		virtualNetworkName +
		"/subnets?api-version=2023-09-01"

	response := HttpGet(urlString, mat)
	json.Unmarshal(response, &listSubnetResponse)

	return listSubnetResponse.Value
}
