package azure

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
	SubscriptionID         string                 `json:"subscriptionId"`
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
