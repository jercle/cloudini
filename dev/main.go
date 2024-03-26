package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/jercle/azg/cmd/azure"
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
		AddressPrefix    string `json:"addressPrefix"`
		Delegations      []any  `json:"delegations"`
		IpConfigurations []struct {
			ID string `json:"id"`
		} `json:"ipConfigurations"`
		PrivateEndpointNetworkPolicies string `json:"privateEndpointNetworkPolicies"`
		PrivateEndpoints               []struct {
			ID string `json:"id"`
		} `json:"privateEndpoints"`
		PrivateLinkServiceNetworkPolicies string `json:"privateLinkServiceNetworkPolicies"`
		ProvisioningState                 string `json:"provisioningState"`
		Purpose                           string `json:"purpose"`
		RouteTable                        struct {
			ID string `json:"id"`
		} `json:"routeTable"`
		ServiceEndpoints []struct {
			Locations         []string `json:"locations"`
			ProvisioningState string   `json:"provisioningState"`
			Service           string   `json:"service"`
		} `json:"serviceEndpoints"`
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

func main() {
	var (
		tenantId = os.Getenv("AZURE_TENANT_ID")
		// subscriptionId = "fdeee0c2-5569-40ea-9ad9-81dd325f6e1e"
		subscriptionId     = os.Getenv("AZURE_SUBSCRIPTION_ID")
		spDetails          lib.CldConfigClientAuthDetails
		resourceGroupName  = "rg-apcdtqshared-automon"
		virtualNetworkName = "vnet-apcdtqshared-automon"
		// subnetName         = "snet-apcdtqshared-automon-builders"
		subnetName = "snet-apcdtqshared-automon"
	)

	spDetails.ClientID = os.Getenv("AZURE_CLIENT_ID")
	spDetails.ClientSecret = os.Getenv("AZURE_CLIENT_SECRET")

	_ = tenantId
	_ = subscriptionId
	_ = resourceGroupName
	_ = virtualNetworkName
	_ = subnetName

	// Get vNet
	// urlString := "https://management.azure.com/subscriptions/" +
	// 	subscriptionId +
	// 	"/resourceGroups/" +
	// 	resourceGroupName +
	// 	"/providers/Microsoft.Network/virtualNetworks/" +
	// 	virtualNetworkName +
	// 	"?api-version=2023-09-01"

	// Get vNet Usages
	// urlString := "https://management.azure.com/subscriptions/" +
	// 	subscriptionId +
	// 	"/resourceGroups/" +
	// 	resourceGroupName +
	// 	"/providers/Microsoft.Network/virtualNetworks/" +
	// 	virtualNetworkName +
	// 	"/usages?api-version=2023-09-01"

	// Get Subnet
	// urlString := "https://management.azure.com/subscriptions/" +
	// 	subscriptionId +
	// 	"/resourceGroups/" +
	// 	resourceGroupName +
	// 	"/providers/Microsoft.Network/virtualNetworks/" +
	// 	virtualNetworkName +
	// 	"/subnets/" +
	// 	subnetName +
	// 	"?api-version=2023-09-01"

	// Get Subnet IP configurations
	// urlString := "https://management.azure.com/subscriptions/" +
	// subscriptionId +
	// "/resourceGroups/" +
	// resourceGroupName +
	// "/providers/Microsoft.Network/virtualNetworks/" +
	// virtualNetworkName +
	// "?api-version=2023-02-01&$expand=subnets/ipConfigurations"

	// resp := printHttpGetResult(urlString)

	// fmt.Println(string(resp))
	// listAllSubscriptionVnets(tenantId, subscriptionId, spDetails)
	// listAllTenantIpAddresses(tenantId)
	tokens, err := azure.GetAllTenantSPTokens(azure.MultiAuthTokenRequestOptions{})
	lib.CheckFatalError(err)
	redToken, err := tokens.SelectTenant("REDDTQ")
	lib.CheckFatalError(err)

	allIps := listAllTenantIpAddresses(*redToken)

	fmt.Println(allIps)

	jsonStr, err := json.MarshalIndent(allIps, "", "  ")
	lib.CheckFatalError(err)

	fmt.Println(string(jsonStr))
	// config := lib.GetCldConfig(lib.CldConfigOptions{})

	// fmt.Println(config.Azure.MultiTenantAuth.Tenants)

}

func listAllTenantIpAddresses(token azure.MultiAuthToken) IPAddressList {
	var allIpAddresses IPAddressList
	vnets := make(chan Vnet, 25)
	publicIpAddresses := make(chan IPAddressItem, 25)
	privateIpAddresses := make(chan IPAddressItem, 25)

	allSubs, err := azure.ListSubscriptions(token)
	lib.CheckFatalError(err)

	for _, sub := range allSubs {
		go listAllSubscriptionVnets(token.TenantId, sub.SubscriptionID, token.TokenData, vnets)
		// for _, vnet := range vnets {
		// 	vnetIps := getAllVnetIPAddresses(token.TenantId, sub.SubscriptionID, vnet.ResourceGroup, vnet.Name, token.TokenData)
		// 	allIpAddresses.PrivateAddresses = append(allIpAddresses.PrivateAddresses, vnetIps.PrivateAddresses...)
		// 	allIpAddresses.PublicAddresses = append(allIpAddresses.PublicAddresses, vnetIps.PublicAddresses...)
		// }
	}

	for vnet := range vnets {

	}

	return allIpAddresses
}

func listAllSubscriptionVnets(tenantId string, subscriptionId string, token azure.TokenData, vnets chan<- Vnet) {
	var (
		listVnets VnetListResponse
	)

	urlString := "https://management.azure.com/subscriptions/" +
		subscriptionId +
		"/providers/Microsoft.Network/virtualNetworks?api-version=2023-09-01"

	// token, err := azure.GetServicePrincipalToken(tenantId, spDetails)
	// lib.CheckFatalError(err)

	req, err := http.NewRequest(http.MethodGet, urlString, nil)
	lib.CheckFatalError(err)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+token.Token)

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
		// allVnets = append(allVnets, currentVnet)
		vnets <- currentVnet

	}

	// return allVnets
}

func getAllVnetIPAddresses(tenantId string, token azure.TokenData, vnets <-chan Vnet) IPAddressList {
	var (
		allIpAddresses IPAddressList
	)

	for vnet := range vnets {
		urlString := "https://management.azure.com/subscriptions/" +
			vnet.SubscriptionID +
			"/resourceGroups/" +
			vnet.ResourceGroup +
			"/providers/Microsoft.Network/virtualNetworks/" +
			vnet.Name +
			"?api-version=2023-02-01&$expand=subnets/ipConfigurations"

		// token, err := azure.GetServicePrincipalToken(tenantId, spDetails)
		// lib.CheckFatalError(err)

		req, err := http.NewRequest(http.MethodGet, urlString, nil)
		lib.CheckFatalError(err)

		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("Authorization", "Bearer "+token.Token)

		res, err := http.DefaultClient.Do(req)
		lib.CheckFatalError(err)

		responseBody, err := io.ReadAll(res.Body)
		lib.CheckFatalError(err)

		var vnet SubnetIPConfigResponse
		err = json.Unmarshal(responseBody, &vnet)
		lib.CheckFatalError(err)

		subnets := vnet.Properties.Subnets

		for _, sn := range subnets {
			ipConfigs := sn.Properties.IpConfigurations
			for _, conf := range ipConfigs {
				go func() {
					confId := strings.Split(conf.ID, "ipConfigurations")[0]
					confUrl := "https://management.azure.com" + confId + "?api-version=2023-02-01"
					var resourceResp IPAddressItem
					result := printHttpGetResult(confUrl)
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

						allIpAddresses.PrivateAddresses = append(allIpAddresses.PrivateAddresses, ipAddressItem)
					} else {
						// Is a public IP
						pubAddressUrl := "https://management.azure.com" + conf.Properties.PublicIpAddress.ID + "?api-version=2023-02-01"
						result := printHttpGetResult(pubAddressUrl)

						var publicIp PublicIpAddress
						json.Unmarshal(result, &publicIp)
						ipAddressItem.IpAddress = publicIp.Properties.IpAddress
						ipAddressItem.ResourceName = publicIp.Name
						ipAddressItem.ResourceID = publicIp.ID
						ipAddressItem.ResourceType = publicIp.Type
						ipAddressItem.Tags = publicIp.Tags

						allIpAddresses.PublicAddresses = append(allIpAddresses.PublicAddresses, ipAddressItem)

					}
				}()
			}
		}
	}

	return allIpAddresses
}

func printHttpGetResult(urlString string) []byte {
	var (
		tenantId = os.Getenv("AZURE_TENANT_ID")
		// subscriptionId     = os.Getenv("AZURE_SUBSCRIPTION_ID")
		spDetails lib.CldConfigClientAuthDetails
		// resourceGroupName  = "rg-apcdtqshared-automon"
		// virtualNetworkName = "vnet-apcdtqshared-automon"
	)
	spDetails.ClientID = os.Getenv("AZURE_CLIENT_ID")
	spDetails.ClientSecret = os.Getenv("AZURE_CLIENT_SECRET")

	token, err := azure.GetServicePrincipalToken(tenantId, spDetails)
	lib.CheckFatalError(err)

	req, err := http.NewRequest(http.MethodGet, urlString, nil)
	lib.CheckFatalError(err)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+token.Token)

	res, err := http.DefaultClient.Do(req)
	lib.CheckFatalError(err)

	responseBody, err := io.ReadAll(res.Body)
	lib.CheckFatalError(err)

	// fmt.Println(string(responseBody))
	return responseBody
}
