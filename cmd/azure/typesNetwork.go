package azure

import "time"

type VnetResponse struct {
	Etag       string `json:"etag,omitempty,omitzero" bson:"etag,omitempty,omitzero"`
	ID         string `json:"id,omitempty,omitzero" bson:"id,omitempty,omitzero"`
	Location   string `json:"location,omitempty,omitzero" bson:"location,omitempty,omitzero"`
	Name       string `json:"name,omitempty,omitzero" bson:"name,omitempty,omitzero"`
	Properties struct {
		AddressSpace struct {
			AddressPrefixes []string `json:"addressPrefixes,omitempty,omitzero" bson:"addressPrefixes,omitempty,omitzero"`
		} `json:"addressSpace,omitempty,omitzero" bson:"addressSpace,omitempty,omitzero"`
		DhcpOptions struct {
			DnsServers []any `json:"dnsServers,omitempty,omitzero" bson:"dnsServers,omitempty,omitzero"`
		} `json:"dhcpOptions,omitempty,omitzero" bson:"dhcpOptions,omitempty,omitzero"`
		EnableDdosProtection bool   `json:"enableDdosProtection,omitempty,omitzero" bson:"enableDdosProtection,omitempty,omitzero"`
		ProvisioningState    string `json:"provisioningState,omitempty,omitzero" bson:"provisioningState,omitempty,omitzero"`
		ResourceGuid         string `json:"resourceGuid,omitempty,omitzero" bson:"resourceGuid,omitempty,omitzero"`
		Subnets              []struct {
			Etag       string `json:"etag,omitempty,omitzero" bson:"etag,omitempty,omitzero"`
			ID         string `json:"id,omitempty,omitzero" bson:"id,omitempty,omitzero"`
			Name       string `json:"name,omitempty,omitzero" bson:"name,omitempty,omitzero"`
			Properties struct {
				AddressPrefix    string `json:"addressPrefix,omitempty,omitzero" bson:"addressPrefix,omitempty,omitzero"`
				Delegations      []any  `json:"delegations,omitempty,omitzero" bson:"delegations,omitempty,omitzero"`
				IpConfigurations []struct {
					ID string `json:"id,omitempty,omitzero" bson:"id,omitempty,omitzero"`
				} `json:"ipConfigurations,omitempty,omitzero" bson:"ipConfigurations,omitempty,omitzero"`
				NetworkSecurityGroup *struct {
					ID string `json:"id,omitempty,omitzero" bson:"id,omitempty,omitzero"`
				} `json:"networkSecurityGroup,omitempty" bson:"networkSecurityGroup,omitempty"`
				PrivateEndpointNetworkPolicies string `json:"privateEndpointNetworkPolicies,omitempty,omitzero" bson:"privateEndpointNetworkPolicies,omitempty,omitzero"`
				PrivateEndpoints               []struct {
					ID string `json:"id,omitempty,omitzero" bson:"id,omitempty,omitzero"`
				} `json:"privateEndpoints,omitempty" bson:"privateEndpoints,omitempty"`
				PrivateLinkServiceNetworkPolicies string `json:"privateLinkServiceNetworkPolicies,omitempty,omitzero" bson:"privateLinkServiceNetworkPolicies,omitempty,omitzero"`
				ProvisioningState                 string `json:"provisioningState,omitempty,omitzero" bson:"provisioningState,omitempty,omitzero"`
				Purpose                           string `json:"purpose,omitempty" bson:"purpose,omitempty"`
				RouteTable                        *struct {
					ID string `json:"id,omitempty,omitzero" bson:"id,omitempty,omitzero"`
				} `json:"routeTable,omitempty" bson:"routeTable,omitempty"`
				ServiceEndpoints []struct {
					Locations         []string `json:"locations,omitempty,omitzero" bson:"locations,omitempty,omitzero"`
					ProvisioningState string   `json:"provisioningState,omitempty,omitzero" bson:"provisioningState,omitempty,omitzero"`
					Service           string   `json:"service,omitempty,omitzero" bson:"service,omitempty,omitzero"`
				} `json:"serviceEndpoints,omitempty,omitzero" bson:"serviceEndpoints,omitempty,omitzero"`
			} `json:"properties,omitempty,omitzero" bson:"properties,omitempty,omitzero"`
			Type string `json:"type,omitempty,omitzero" bson:"type,omitempty,omitzero"`
		} `json:"subnets,omitempty,omitzero" bson:"subnets,omitempty,omitzero"`
		VirtualNetworkPeerings []struct {
			Etag       string `json:"etag,omitempty,omitzero" bson:"etag,omitempty,omitzero"`
			ID         string `json:"id,omitempty,omitzero" bson:"id,omitempty,omitzero"`
			Name       string `json:"name,omitempty,omitzero" bson:"name,omitempty,omitzero"`
			Properties struct {
				AllowForwardedTraffic     bool   `json:"allowForwardedTraffic,omitempty,omitzero" bson:"allowForwardedTraffic,omitempty,omitzero"`
				AllowGatewayTransit       bool   `json:"allowGatewayTransit,omitempty,omitzero" bson:"allowGatewayTransit,omitempty,omitzero"`
				AllowVirtualNetworkAccess bool   `json:"allowVirtualNetworkAccess,omitempty,omitzero" bson:"allowVirtualNetworkAccess,omitempty,omitzero"`
				DoNotVerifyRemoteGateways bool   `json:"doNotVerifyRemoteGateways,omitempty,omitzero" bson:"doNotVerifyRemoteGateways,omitempty,omitzero"`
				PeerCompleteVnets         bool   `json:"peerCompleteVnets,omitempty,omitzero" bson:"peerCompleteVnets,omitempty,omitzero"`
				PeeringState              string `json:"peeringState,omitempty,omitzero" bson:"peeringState,omitempty,omitzero"`
				PeeringSyncLevel          string `json:"peeringSyncLevel,omitempty,omitzero" bson:"peeringSyncLevel,omitempty,omitzero"`
				ProvisioningState         string `json:"provisioningState,omitempty,omitzero" bson:"provisioningState,omitempty,omitzero"`
				RemoteAddressSpace        struct {
					AddressPrefixes []string `json:"addressPrefixes,omitempty,omitzero" bson:"addressPrefixes,omitempty,omitzero"`
				} `json:"remoteAddressSpace,omitempty,omitzero" bson:"remoteAddressSpace,omitempty,omitzero"`
				RemoteGateways []struct {
					ID string `json:"id,omitempty,omitzero" bson:"id,omitempty,omitzero"`
				} `json:"remoteGateways,omitempty" bson:"remoteGateways,omitempty"`
				RemoteVirtualNetwork struct {
					ID string `json:"id,omitempty,omitzero" bson:"id,omitempty,omitzero"`
				} `json:"remoteVirtualNetwork,omitempty,omitzero" bson:"remoteVirtualNetwork,omitempty,omitzero"`
				RemoteVirtualNetworkAddressSpace struct {
					AddressPrefixes []string `json:"addressPrefixes,omitempty,omitzero" bson:"addressPrefixes,omitempty,omitzero"`
				} `json:"remoteVirtualNetworkAddressSpace,omitempty,omitzero" bson:"remoteVirtualNetworkAddressSpace,omitempty,omitzero"`
				ResourceGuid     string `json:"resourceGuid,omitempty,omitzero" bson:"resourceGuid,omitempty,omitzero"`
				RouteServiceVips struct {
					Af36ba888c9943f4A5e38fa90652cc96 string `json:"af36ba88-8c99-43f4-a5e3-8fa90652cc96,omitempty" bson:"af36ba88-8c99-43f4-a5e3-8fa90652cc96,omitempty"`
				} `json:"routeServiceVips,omitempty,omitzero" bson:"routeServiceVips,omitempty,omitzero"`
				UseRemoteGateways bool `json:"useRemoteGateways,omitempty,omitzero" bson:"useRemoteGateways,omitempty,omitzero"`
			} `json:"properties,omitempty,omitzero" bson:"properties,omitempty,omitzero"`
			Type string `json:"type,omitempty,omitzero" bson:"type,omitempty,omitzero"`
		} `json:"virtualNetworkPeerings,omitempty,omitzero" bson:"virtualNetworkPeerings,omitempty,omitzero"`
	} `json:"properties,omitempty,omitzero" bson:"properties,omitempty,omitzero"`
	Tags map[string]string `json:"tags,omitempty,omitzero" bson:"tags,omitempty,omitzero"`
	Type string            `json:"type,omitempty,omitzero" bson:"type,omitempty,omitzero"`
}

type VnetListResponse struct {
	Value []VnetResponse
}

type Vnet struct {
	ID                     string                 `json:"id,omitempty,omitzero" bson:"id,omitempty,omitzero"`
	Location               string                 `json:"location,omitempty,omitzero" bson:"location,omitempty,omitzero"`
	Name                   string                 `json:"name,omitempty,omitzero" bson:"name,omitempty,omitzero"`
	ResourceGroup          string                 `json:"resourceGroup,omitempty,omitzero" bson:"resourceGroup,omitempty,omitzero"`
	SubscriptionID         string                 `json:"subscriptionId,omitempty,omitzero" bson:"subscriptionId,omitempty,omitzero"`
	AddressSpace           []string               `json:"addressSpace,omitempty,omitzero" bson:"addressSpace,omitempty,omitzero"`
	Subnets                []SubnetResponse       `json:"subnets,omitempty,omitzero" bson:"subnets,omitempty,omitzero"`
	ProvisioningState      string                 `json:"provisioningState,omitempty,omitzero" bson:"provisioningState,omitempty,omitzero"`
	VirtualNetworkPeerings []ProcessedVnetPeering `json:"virtualNetworkPeerings,omitempty,omitzero" bson:"virtualNetworkPeerings,omitempty,omitzero"`
	Tags                   map[string]string      `json:"tags,omitempty,omitzero" bson:"tags,omitempty,omitzero"`
	Type                   string                 `json:"type,omitempty,omitzero" bson:"type,omitempty,omitzero"`
}

type VirtualNetworkPeering struct {
	Etag       string `json:"etag,omitempty,omitzero" bson:"etag,omitempty,omitzero"`
	ID         string `json:"id,omitempty,omitzero" bson:"id,omitempty,omitzero"`
	Name       string `json:"name,omitempty,omitzero" bson:"name,omitempty,omitzero"`
	Properties struct {
		AllowForwardedTraffic     bool   `json:"allowForwardedTraffic,omitempty,omitzero" bson:"allowForwardedTraffic,omitempty,omitzero"`
		AllowGatewayTransit       bool   `json:"allowGatewayTransit,omitempty,omitzero" bson:"allowGatewayTransit,omitempty,omitzero"`
		AllowVirtualNetworkAccess bool   `json:"allowVirtualNetworkAccess,omitempty,omitzero" bson:"allowVirtualNetworkAccess,omitempty,omitzero"`
		DoNotVerifyRemoteGateways bool   `json:"doNotVerifyRemoteGateways,omitempty,omitzero" bson:"doNotVerifyRemoteGateways,omitempty,omitzero"`
		PeerCompleteVnets         bool   `json:"peerCompleteVnets,omitempty,omitzero" bson:"peerCompleteVnets,omitempty,omitzero"`
		PeeringState              string `json:"peeringState,omitempty,omitzero" bson:"peeringState,omitempty,omitzero"`
		PeeringSyncLevel          string `json:"peeringSyncLevel,omitempty,omitzero" bson:"peeringSyncLevel,omitempty,omitzero"`
		ProvisioningState         string `json:"provisioningState,omitempty,omitzero" bson:"provisioningState,omitempty,omitzero"`
		RemoteAddressSpace        struct {
			AddressPrefixes []string `json:"addressPrefixes,omitempty,omitzero" bson:"addressPrefixes,omitempty,omitzero"`
		} `json:"remoteAddressSpace,omitempty,omitzero" bson:"remoteAddressSpace,omitempty,omitzero"`
		RemoteGateways []struct {
			ID string `json:"id,omitempty,omitzero" bson:"id,omitempty,omitzero"`
		} `json:"remoteGateways,omitempty" bson:"remoteGateways,omitempty"`
		RemoteVirtualNetwork struct {
			ID string `json:"id,omitempty,omitzero" bson:"id,omitempty,omitzero"`
		} `json:"remoteVirtualNetwork,omitempty,omitzero" bson:"remoteVirtualNetwork,omitempty,omitzero"`
		RemoteVirtualNetworkAddressSpace struct {
			AddressPrefixes []string `json:"addressPrefixes,omitempty,omitzero" bson:"addressPrefixes,omitempty,omitzero"`
		} `json:"remoteVirtualNetworkAddressSpace,omitempty,omitzero" bson:"remoteVirtualNetworkAddressSpace,omitempty,omitzero"`
		ResourceGuid     string `json:"resourceGuid,omitempty,omitzero" bson:"resourceGuid,omitempty,omitzero"`
		RouteServiceVips struct {
			Af36ba888c9943f4A5e38fa90652cc96 string `json:"af36ba88-8c99-43f4-a5e3-8fa90652cc96,omitempty" bson:"af36ba88-8c99-43f4-a5e3-8fa90652cc96,omitempty"`
		} `json:"routeServiceVips,omitempty,omitzero" bson:"routeServiceVips,omitempty,omitzero"`
		UseRemoteGateways bool `json:"useRemoteGateways,omitempty,omitzero" bson:"useRemoteGateways,omitempty,omitzero"`
	} `json:"properties,omitempty,omitzero" bson:"properties,omitempty,omitzero"`
	Type string `json:"type,omitempty,omitzero" bson:"type,omitempty,omitzero"`
}

type ProcessedVnetPeering struct {
	Name                      string   `json:"name,omitempty,omitzero" bson:"name,omitempty,omitzero"`
	RemoteAddressSpace        []string `json:"remoteAddressSpace,omitempty,omitzero" bson:"remoteAddressSpace,omitempty,omitzero"`
	RemoteGateways            []string `json:"remoteGateways,omitempty,omitzero" bson:"remoteGateways,omitempty,omitzero"`
	RemoteVirtualNetwork      string   `json:"remoteVirtualNetwork,omitempty,omitzero" bson:"remoteVirtualNetwork,omitempty,omitzero"`
	UseRemoteGateways         bool     `json:"useRemoteGateways,omitempty,omitzero" bson:"useRemoteGateways,omitempty,omitzero"`
	AllowForwardedTraffic     bool     `json:"allowForwardedTraffic,omitempty,omitzero" bson:"allowForwardedTraffic,omitempty,omitzero"`
	AllowGatewayTransit       bool     `json:"allowGatewayTransit,omitempty,omitzero" bson:"allowGatewayTransit,omitempty,omitzero"`
	AllowVirtualNetworkAccess bool     `json:"allowVirtualNetworkAccess,omitempty,omitzero" bson:"allowVirtualNetworkAccess,omitempty,omitzero"`
	PeeringState              string   `json:"peeringState,omitempty,omitzero" bson:"peeringState,omitempty,omitzero"`
	ProvisioningState         string   `json:"provisioningState,omitempty,omitzero" bson:"provisioningState,omitempty,omitzero"`
	PeeringSyncLevel          string   `json:"peeringSyncLevel,omitempty,omitzero" bson:"peeringSyncLevel,omitempty,omitzero"`
}

type SubnetResponse struct {
	Etag       string `json:"etag,omitempty,omitzero" bson:"etag,omitempty,omitzero"`
	ID         string `json:"id,omitempty,omitzero" bson:"id,omitempty,omitzero"`
	Name       string `json:"name,omitempty,omitzero" bson:"name,omitempty,omitzero"`
	Properties struct {
		AddressPrefix string `json:"addressPrefix,omitempty,omitzero" bson:"addressPrefix,omitempty,omitzero"`
		Delegations   []struct {
			Etag       string `json:"etag,omitempty,omitzero" bson:"etag,omitempty,omitzero"`
			ID         string `json:"id,omitempty,omitzero" bson:"id,omitempty,omitzero"`
			Name       string `json:"name,omitempty,omitzero" bson:"name,omitempty,omitzero"`
			Properties struct {
				Actions           []string `json:"actions,omitempty,omitzero" bson:"actions,omitempty,omitzero"`
				ProvisioningState string   `json:"provisioningState,omitempty,omitzero" bson:"provisioningState,omitempty,omitzero"`
				ServiceName       string   `json:"serviceName,omitempty,omitzero" bson:"serviceName,omitempty,omitzero"`
			} `json:"properties,omitempty,omitzero" bson:"properties,omitempty,omitzero"`
			Type string `json:"type,omitempty,omitzero" bson:"type,omitempty,omitzero"`
		} `json:"delegations,omitempty,omitzero" bson:"delegations,omitempty,omitzero"`
		IpConfigurations []struct {
			ID string `json:"id,omitempty,omitzero" bson:"id,omitempty,omitzero"`
		} `json:"ipConfigurations,omitempty,omitzero" bson:"ipConfigurations,omitempty,omitzero"`
		NetworkSecurityGroup struct {
			ID string `json:"id,omitempty,omitzero" bson:"id,omitempty,omitzero"`
		} `json:"networkSecurityGroup,omitempty,omitzero" bson:"networkSecurityGroup,omitempty,omitzero"`
		PrivateEndpointNetworkPolicies    string `json:"privateEndpointNetworkPolicies,omitempty,omitzero" bson:"privateEndpointNetworkPolicies,omitempty,omitzero"`
		PrivateLinkServiceNetworkPolicies string `json:"privateLinkServiceNetworkPolicies,omitempty,omitzero" bson:"privateLinkServiceNetworkPolicies,omitempty,omitzero"`
		ProvisioningState                 string `json:"provisioningState,omitempty,omitzero" bson:"provisioningState,omitempty,omitzero"`
		Purpose                           string `json:"purpose,omitempty,omitzero" bson:"purpose,omitempty,omitzero"`
		ServiceEndpoints                  []any  `json:"serviceEndpoints,omitempty,omitzero" bson:"serviceEndpoints,omitempty,omitzero"`
	} `json:"properties,omitempty,omitzero" bson:"properties,omitempty,omitzero"`
	Type string `json:"type,omitempty,omitzero" bson:"type,omitempty,omitzero"`
}

type ListSubnetsResponse struct {
	Value []SubnetResponse `json:"value,omitempty,omitzero" bson:"value,omitempty,omitzero"`
}

type ProcessedSubnet struct {
	ID         string `json:"id,omitempty,omitzero" bson:"id,omitempty,omitzero"`
	Name       string `json:"name,omitempty,omitzero" bson:"name,omitempty,omitzero"`
	VnetName   string `json:"vnetName,omitempty,omitzero" bson:"vnetName,omitempty,omitzero"`
	Properties struct {
		AddressPrefix string `json:"addressPrefix,omitempty,omitzero" bson:"addressPrefix,omitempty,omitzero"`
		Delegations   []struct {
			ID         string `json:"id,omitempty,omitzero" bson:"id,omitempty,omitzero"`
			Name       string `json:"name,omitempty,omitzero" bson:"name,omitempty,omitzero"`
			Properties struct {
				Actions           []string `json:"actions,omitempty,omitzero" bson:"actions,omitempty,omitzero"`
				ProvisioningState string   `json:"provisioningState,omitempty,omitzero" bson:"provisioningState,omitempty,omitzero"`
				ServiceName       string   `json:"serviceName,omitempty,omitzero" bson:"serviceName,omitempty,omitzero"`
			} `json:"properties,omitempty,omitzero" bson:"properties,omitempty,omitzero"`
			Type string `json:"type,omitempty,omitzero" bson:"type,omitempty,omitzero"`
		} `json:"delegations,omitempty,omitzero" bson:"delegations,omitempty,omitzero"`
		IpConfigurations []struct {
			ID string `json:"id,omitempty,omitzero" bson:"id,omitempty,omitzero"`
		} `json:"ipConfigurations,omitempty,omitzero" bson:"ipConfigurations,omitempty,omitzero"`
		NetworkSecurityGroup struct {
			ID string `json:"id,omitempty,omitzero" bson:"id,omitempty,omitzero"`
		} `json:"networkSecurityGroup,omitempty,omitzero" bson:"networkSecurityGroup,omitempty,omitzero"`
		PrivateEndpointNetworkPolicies    string `json:"privateEndpointNetworkPolicies,omitempty,omitzero" bson:"privateEndpointNetworkPolicies,omitempty,omitzero"`
		PrivateLinkServiceNetworkPolicies string `json:"privateLinkServiceNetworkPolicies,omitempty,omitzero" bson:"privateLinkServiceNetworkPolicies,omitempty,omitzero"`
		ProvisioningState                 string `json:"provisioningState,omitempty,omitzero" bson:"provisioningState,omitempty,omitzero"`
		Purpose                           string `json:"purpose,omitempty,omitzero" bson:"purpose,omitempty,omitzero"`
		ServiceEndpoints                  []any  `json:"serviceEndpoints,omitempty,omitzero" bson:"serviceEndpoints,omitempty,omitzero"`
	} `json:"properties,omitempty,omitzero" bson:"properties,omitempty,omitzero"`
	Type string `json:"type,omitempty,omitzero" bson:"type,omitempty,omitzero"`
}

type SubnetIPConfigResponse struct {
	Etag       string `json:"etag,omitempty,omitzero" bson:"etag,omitempty,omitzero"`
	ID         string `json:"id,omitempty,omitzero" bson:"id,omitempty,omitzero"`
	Location   string `json:"location,omitempty,omitzero" bson:"location,omitempty,omitzero"`
	Name       string `json:"name,omitempty,omitzero" bson:"name,omitempty,omitzero"`
	Properties struct {
		AddressSpace struct {
			AddressPrefixes []string `json:"addressPrefixes,omitempty,omitzero" bson:"addressPrefixes,omitempty,omitzero"`
		} `json:"addressSpace,omitempty,omitzero" bson:"addressSpace,omitempty,omitzero"`
		DhcpOptions struct {
			DnsServers []any `json:"dnsServers,omitempty,omitzero" bson:"dnsServers,omitempty,omitzero"`
		} `json:"dhcpOptions,omitempty,omitzero" bson:"dhcpOptions,omitempty,omitzero"`
		EnableDdosProtection bool   `json:"enableDdosProtection,omitempty,omitzero" bson:"enableDdosProtection,omitempty,omitzero"`
		ProvisioningState    string `json:"provisioningState,omitempty,omitzero" bson:"provisioningState,omitempty,omitzero"`
		ResourceGuid         string `json:"resourceGuid,omitempty,omitzero" bson:"resourceGuid,omitempty,omitzero"`
		Subnets              []struct {
			Etag       string `json:"etag,omitempty,omitzero" bson:"etag,omitempty,omitzero"`
			ID         string `json:"id,omitempty,omitzero" bson:"id,omitempty,omitzero"`
			Name       string `json:"name,omitempty,omitzero" bson:"name,omitempty,omitzero"`
			Properties struct {
				AddressPrefix    string `json:"addressPrefix,omitempty,omitzero" bson:"addressPrefix,omitempty,omitzero"`
				Delegations      []any  `json:"delegations,omitempty,omitzero" bson:"delegations,omitempty,omitzero"`
				IpConfigurations []struct {
					Etag       string `json:"etag,omitempty,omitzero" bson:"etag,omitempty,omitzero"`
					ID         string `json:"id,omitempty,omitzero" bson:"id,omitempty,omitzero"`
					Name       string `json:"name,omitempty,omitzero" bson:"name,omitempty,omitzero"`
					Resource   IPAddressItem
					Properties struct {
						Primary                         bool   `json:"primary,omitempty" bson:"primary,omitempty"`
						PrivateIpAddress                string `json:"privateIPAddress,omitempty" bson:"privateIPAddress,omitempty"`
						PrivateIpAddressVersion         string `json:"privateIPAddressVersion,omitempty" bson:"privateIPAddressVersion,omitempty"`
						PrivateIpAllocationMethod       string `json:"privateIPAllocationMethod,omitempty,omitzero" bson:"privateIPAllocationMethod,omitempty,omitzero"`
						PrivateLinkConnectionProperties *struct {
							Fqdns              []string `json:"fqdns,omitempty,omitzero" bson:"fqdns,omitempty,omitzero"`
							GroupID            string   `json:"groupId,omitempty,omitzero" bson:"groupId,omitempty,omitzero"`
							RequiredMemberName string   `json:"requiredMemberName,omitempty,omitzero" bson:"requiredMemberName,omitempty,omitzero"`
						} `json:"privateLinkConnectionProperties,omitempty" bson:"privateLinkConnectionProperties,omitempty"`
						ProvisioningState string `json:"provisioningState,omitempty,omitzero" bson:"provisioningState,omitempty,omitzero"`
						PublicIpAddress   *struct {
							ID string `json:"id,omitempty,omitzero" bson:"id,omitempty,omitzero"`
						} `json:"publicIPAddress,omitempty" bson:"publicIPAddress,omitempty"`
						Subnet struct {
							ID string `json:"id,omitempty,omitzero" bson:"id,omitempty,omitzero"`
						} `json:"subnet,omitempty,omitzero" bson:"subnet,omitempty,omitzero"`
					} `json:"properties,omitempty,omitzero" bson:"properties,omitempty,omitzero"`
					Type string `json:"type,omitempty,omitzero" bson:"type,omitempty,omitzero"`
				} `json:"ipConfigurations,omitempty,omitzero" bson:"ipConfigurations,omitempty,omitzero"`
				NetworkSecurityGroup *struct {
					ID string `json:"id,omitempty,omitzero" bson:"id,omitempty,omitzero"`
				} `json:"networkSecurityGroup,omitempty" bson:"networkSecurityGroup,omitempty"`
				PrivateEndpointNetworkPolicies string `json:"privateEndpointNetworkPolicies,omitempty,omitzero" bson:"privateEndpointNetworkPolicies,omitempty,omitzero"`
				PrivateEndpoints               []struct {
					ID string `json:"id,omitempty,omitzero" bson:"id,omitempty,omitzero"`
				} `json:"privateEndpoints,omitempty" bson:"privateEndpoints,omitempty"`
				PrivateLinkServiceNetworkPolicies string `json:"privateLinkServiceNetworkPolicies,omitempty,omitzero" bson:"privateLinkServiceNetworkPolicies,omitempty,omitzero"`
				ProvisioningState                 string `json:"provisioningState,omitempty,omitzero" bson:"provisioningState,omitempty,omitzero"`
				Purpose                           string `json:"purpose,omitempty" bson:"purpose,omitempty"`
				RouteTable                        *struct {
					ID string `json:"id,omitempty,omitzero" bson:"id,omitempty,omitzero"`
				} `json:"routeTable,omitempty" bson:"routeTable,omitempty"`
				ServiceEndpoints []struct {
					Locations         []string `json:"locations,omitempty,omitzero" bson:"locations,omitempty,omitzero"`
					ProvisioningState string   `json:"provisioningState,omitempty,omitzero" bson:"provisioningState,omitempty,omitzero"`
					Service           string   `json:"service,omitempty,omitzero" bson:"service,omitempty,omitzero"`
				} `json:"serviceEndpoints,omitempty,omitzero" bson:"serviceEndpoints,omitempty,omitzero"`
			} `json:"properties,omitempty,omitzero" bson:"properties,omitempty,omitzero"`
			Type string `json:"type,omitempty,omitzero" bson:"type,omitempty,omitzero"`
		} `json:"subnets,omitempty,omitzero" bson:"subnets,omitempty,omitzero"`
		VirtualNetworkPeerings []struct {
			Etag       string `json:"etag,omitempty,omitzero" bson:"etag,omitempty,omitzero"`
			ID         string `json:"id,omitempty,omitzero" bson:"id,omitempty,omitzero"`
			Name       string `json:"name,omitempty,omitzero" bson:"name,omitempty,omitzero"`
			Properties struct {
				AllowForwardedTraffic     bool   `json:"allowForwardedTraffic,omitempty,omitzero" bson:"allowForwardedTraffic,omitempty,omitzero"`
				AllowGatewayTransit       bool   `json:"allowGatewayTransit,omitempty,omitzero" bson:"allowGatewayTransit,omitempty,omitzero"`
				AllowVirtualNetworkAccess bool   `json:"allowVirtualNetworkAccess,omitempty,omitzero" bson:"allowVirtualNetworkAccess,omitempty,omitzero"`
				DoNotVerifyRemoteGateways bool   `json:"doNotVerifyRemoteGateways,omitempty,omitzero" bson:"doNotVerifyRemoteGateways,omitempty,omitzero"`
				PeerCompleteVnets         bool   `json:"peerCompleteVnets,omitempty,omitzero" bson:"peerCompleteVnets,omitempty,omitzero"`
				PeeringState              string `json:"peeringState,omitempty,omitzero" bson:"peeringState,omitempty,omitzero"`
				PeeringSyncLevel          string `json:"peeringSyncLevel,omitempty,omitzero" bson:"peeringSyncLevel,omitempty,omitzero"`
				ProvisioningState         string `json:"provisioningState,omitempty,omitzero" bson:"provisioningState,omitempty,omitzero"`
				RemoteAddressSpace        struct {
					AddressPrefixes []string `json:"addressPrefixes,omitempty,omitzero" bson:"addressPrefixes,omitempty,omitzero"`
				} `json:"remoteAddressSpace,omitempty,omitzero" bson:"remoteAddressSpace,omitempty,omitzero"`
				RemoteGateways []struct {
					ID string `json:"id,omitempty,omitzero" bson:"id,omitempty,omitzero"`
				} `json:"remoteGateways,omitempty" bson:"remoteGateways,omitempty"`
				RemoteVirtualNetwork struct {
					ID string `json:"id,omitempty,omitzero" bson:"id,omitempty,omitzero"`
				} `json:"remoteVirtualNetwork,omitempty,omitzero" bson:"remoteVirtualNetwork,omitempty,omitzero"`
				RemoteVirtualNetworkAddressSpace struct {
					AddressPrefixes []string `json:"addressPrefixes,omitempty,omitzero" bson:"addressPrefixes,omitempty,omitzero"`
				} `json:"remoteVirtualNetworkAddressSpace,omitempty,omitzero" bson:"remoteVirtualNetworkAddressSpace,omitempty,omitzero"`
				ResourceGuid     string `json:"resourceGuid,omitempty,omitzero" bson:"resourceGuid,omitempty,omitzero"`
				RouteServiceVips struct {
					Af36ba888c9943f4A5e38fa90652cc96 string `json:"af36ba88-8c99-43f4-a5e3-8fa90652cc96,omitempty" bson:"af36ba88-8c99-43f4-a5e3-8fa90652cc96,omitempty"`
				} `json:"routeServiceVips,omitempty,omitzero" bson:"routeServiceVips,omitempty,omitzero"`
				UseRemoteGateways bool `json:"useRemoteGateways,omitempty,omitzero" bson:"useRemoteGateways,omitempty,omitzero"`
			} `json:"properties,omitempty,omitzero" bson:"properties,omitempty,omitzero"`
			Type string `json:"type,omitempty,omitzero" bson:"type,omitempty,omitzero"`
		} `json:"virtualNetworkPeerings,omitempty,omitzero" bson:"virtualNetworkPeerings,omitempty,omitzero"`
	} `json:"properties,omitempty,omitzero" bson:"properties,omitempty,omitzero"`
	Tags map[string]string `json:"tags,omitempty,omitzero" bson:"tags,omitempty,omitzero"`
	Type string            `json:"type,omitempty,omitzero" bson:"type,omitempty,omitzero"`
}

type PublicIpAddress struct {
	Etag       string `json:"etag,omitempty,omitzero" bson:"etag,omitempty,omitzero"`
	ID         string `json:"id,omitempty,omitzero" bson:"id,omitempty,omitzero"`
	Location   string `json:"location,omitempty,omitzero" bson:"location,omitempty,omitzero"`
	Name       string `json:"name,omitempty,omitzero" bson:"name,omitempty,omitzero"`
	Properties struct {
		DdosSettings struct {
			ProtectionMode string `json:"protectionMode,omitempty,omitzero" bson:"protectionMode,omitempty,omitzero"`
		} `json:"ddosSettings,omitempty,omitzero" bson:"ddosSettings,omitempty,omitzero"`
		IdleTimeoutInMinutes int    `json:"idleTimeoutInMinutes,omitempty,omitzero" bson:"idleTimeoutInMinutes,omitempty,omitzero"`
		IpAddress            string `json:"ipAddress,omitempty,omitzero" bson:"ipAddress,omitempty,omitzero"`
		IpConfiguration      struct {
			ID string `json:"id,omitempty,omitzero" bson:"id,omitempty,omitzero"`
		} `json:"ipConfiguration,omitempty,omitzero" bson:"ipConfiguration,omitempty,omitzero"`
		IpTags                   []any  `json:"ipTags,omitempty,omitzero" bson:"ipTags,omitempty,omitzero"`
		ProvisioningState        string `json:"provisioningState,omitempty,omitzero" bson:"provisioningState,omitempty,omitzero"`
		PublicIpAddressVersion   string `json:"publicIPAddressVersion,omitempty,omitzero" bson:"publicIPAddressVersion,omitempty,omitzero"`
		PublicIpAllocationMethod string `json:"publicIPAllocationMethod,omitempty,omitzero" bson:"publicIPAllocationMethod,omitempty,omitzero"`
		ResourceGuid             string `json:"resourceGuid,omitempty,omitzero" bson:"resourceGuid,omitempty,omitzero"`
	} `json:"properties,omitempty,omitzero" bson:"properties,omitempty,omitzero"`
	Sku struct {
		Name string `json:"name,omitempty,omitzero" bson:"name,omitempty,omitzero"`
		Tier string `json:"tier,omitempty,omitzero" bson:"tier,omitempty,omitzero"`
	} `json:"sku,omitempty,omitzero" bson:"sku,omitempty,omitzero"`
	Tags map[string]string `json:"tags,omitempty,omitzero" bson:"tags,omitempty,omitzero"`
	Type string            `json:"type,omitempty,omitzero" bson:"type,omitempty,omitzero"`
}

type IPAddressItem struct {
	ResourceName string `json:"name,omitempty,omitzero" bson:"name,omitempty,omitzero"`
	ResourceID   string `json:"id,omitempty,omitzero" bson:"id,omitempty,omitzero"`
	ResourceType string `json:"type,omitempty,omitzero" bson:"type,omitempty,omitzero"`
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

//
//

type AzureP2SConnectionDetails struct {
	P2SConnectionConfigurationResourceID string `json:"P2SConnectionConfigurationResourceId,omitempty,omitzero" bson:"P2SConnectionConfigurationResourceId,omitempty,omitzero"`
	UserNameVpnConnectionHealths         []struct {
		UserName             string                     `json:"UserName,omitempty,omitzero" bson:"UserName,omitempty,omitzero"`
		VpnConnectionHealths []AzureP2SConnectionHealth `json:"VpnConnectionHealths,omitempty,omitzero" bson:"VpnConnectionHealths,omitempty,omitzero"`
	} `json:"UserNameVpnConnectionHealths,omitempty,omitzero" bson:"UserNameVpnConnectionHealths,omitempty,omitzero"`
}

//
//

type AzureP2SConnectionHealth struct {
	EgressBytesTransferred        float64   `json:"EgressBytesTransferred,omitempty,omitzero" bson:"EgressBytesTransferred,omitempty,omitzero"`
	EgressPacketsTransferred      float64   `json:"EgressPacketsTransferred,omitempty,omitzero" bson:"EgressPacketsTransferred,omitempty,omitzero"`
	IngressBytesTransferred       float64   `json:"IngressBytesTransferred,omitempty,omitzero" bson:"IngressBytesTransferred,omitempty,omitzero"`
	IngressPacketsTransferred     float64   `json:"IngressPacketsTransferred,omitempty,omitzero" bson:"IngressPacketsTransferred,omitempty,omitzero"`
	MaxBandwidth                  float64   `json:"MaxBandwidth,omitempty,omitzero" bson:"MaxBandwidth,omitempty,omitzero"`
	MaxPacketsPerSecond           float64   `json:"MaxPacketsPerSecond,omitempty,omitzero" bson:"MaxPacketsPerSecond,omitempty,omitzero"`
	PrivateIpAddress              string    `json:"PrivateIpAddress,omitempty,omitzero" bson:"PrivateIpAddress,omitempty,omitzero"`
	PrivateIpv6Address            any       `json:"PrivateIpv6Address,omitempty,omitzero" bson:"PrivateIpv6Address,omitempty,omitzero"`
	PublicIpAddress               string    `json:"PublicIpAddress,omitempty,omitzero" bson:"PublicIpAddress,omitempty,omitzero"`
	UserName                      string    `json:"UserName,omitempty,omitzero" bson:"UserName,omitempty,omitzero"`
	UserPrincipalName             string    `json:"UserPrincipalName,omitempty,omitzero" bson:"UserPrincipalName,omitempty,omitzero"`
	VpnConnectionDuration         float64   `json:"VpnConnectionDuration,omitempty,omitzero" bson:"VpnConnectionDuration,omitempty,omitzero"`
	VpnConnectionID               string    `json:"VpnConnectionId,omitempty,omitzero" bson:"_id,omitempty,omitzero"`
	VpnConnectionTime             string    `json:"VpnConnectionTime,omitempty,omitzero" bson:"VpnConnectionTime,omitempty,omitzero"`
	ManagedDeviceADDeviceId       string    `json:"ManagedDeviceADDeviceId,omitempty,omitzero" bson:"ManagedDeviceADDeviceId,omitempty,omitzero"`
	ManagedDeviceIntuneId         string    `json:"ManagedDeviceIntuneId,omitempty,omitzero" bson:"ManagedDeviceIntuneId,omitempty,omitzero"`
	ManagedDeviceName             string    `json:"ManagedDeviceName,omitempty,omitzero" bson:"ManagedDeviceName,omitempty,omitzero"`
	ManagedDeviceSerial           string    `json:"ManagedDeviceSerial,omitempty,omitzero" bson:"ManagedDeviceSerial,omitempty,omitzero"`
	ManagedDeviceLastSyncDateTime time.Time `json:"ManagedDeviceLastSyncDateTime,omitempty,omitzero" bson:"ManagedDeviceLastSyncDateTime,omitempty,omitzero"`
}
