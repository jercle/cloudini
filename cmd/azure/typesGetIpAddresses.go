package azure

import (
	json "encoding/json/v2"
	"net"
	"time"
)

type ResourceGraphIPAddressesResponse struct {
	Count           int64                         `json:"count,omitempty,omitzero" bson:"count,omitempty,omitzero"`
	Data            []IPAddressesAllResourceTypes `json:"data,omitempty,omitzero" bson:"data,omitempty,omitzero"`
	Facets          []any                         `json:"facets,omitempty,omitzero" bson:"facets,omitempty,omitzero"`
	ResultTruncated string                        `json:"resultTruncated,omitempty,omitzero" bson:"resultTruncated,omitempty,omitzero"`
	TotalRecords    int64                         `json:"totalRecords,omitempty,omitzero" bson:"totalRecords,omitempty,omitzero"`
	SkipToken       string                        `json:"$skipToken,omitempty,omitzero" bson:"$skipToken,omitempty,omitzero"`
}

//
//

type ResourceManagerResponse struct {
	NextLink string        `json:"nextLink,omitempty,omitzero" bson:"nextLink,omitempty,omitzero"`
	Value    []interface{} `json:"value,omitempty,omitzero" bson:"value,omitempty,omitzero"`
}

//
//

type AzureResourceManagedClusterAgentPool struct {
	ID         string `json:"id,omitempty,omitzero" bson:"id,omitempty,omitzero"`
	Name       string `json:"name,omitempty,omitzero" bson:"name,omitempty,omitzero"`
	Properties struct {
		AvailabilityZones          []string `json:"availabilityZones,omitempty,omitzero" bson:"availabilityZones,omitempty,omitzero"`
		Count                      float64  `json:"count,omitempty,omitzero" bson:"count,omitempty,omitzero"`
		CurrentOrchestratorVersion string   `json:"currentOrchestratorVersion,omitempty,omitzero" bson:"currentOrchestratorVersion,omitempty,omitzero"`
		EnableAutoScaling          bool     `json:"enableAutoScaling,omitempty,omitzero" bson:"enableAutoScaling,omitempty,omitzero"`
		EnableEncryptionAtHost     bool     `json:"enableEncryptionAtHost,omitempty,omitzero" bson:"enableEncryptionAtHost,omitempty,omitzero"`
		EnableFips                 bool     `json:"enableFIPS,omitempty,omitzero" bson:"enableFIPS,omitempty,omitzero"`
		EnableNodePublicIp         bool     `json:"enableNodePublicIP,omitempty,omitzero" bson:"enableNodePublicIP,omitempty,omitzero"`
		EnableUltraSsd             bool     `json:"enableUltraSSD,omitempty,omitzero" bson:"enableUltraSSD,omitempty,omitzero"`
		KubeletDiskType            string   `json:"kubeletDiskType,omitempty,omitzero" bson:"kubeletDiskType,omitempty,omitzero"`
		MaxCount                   float64  `json:"maxCount,omitempty,omitzero" bson:"maxCount,omitempty,omitzero"`
		MaxPods                    float64  `json:"maxPods,omitempty,omitzero" bson:"maxPods,omitempty,omitzero"`
		MinCount                   float64  `json:"minCount,omitempty,omitzero" bson:"minCount,omitempty,omitzero"`
		Mode                       string   `json:"mode,omitempty,omitzero" bson:"mode,omitempty,omitzero"`
		NodeImageVersion           string   `json:"nodeImageVersion,omitempty,omitzero" bson:"nodeImageVersion,omitempty,omitzero"`
		OrchestratorVersion        string   `json:"orchestratorVersion,omitempty,omitzero" bson:"orchestratorVersion,omitempty,omitzero"`
		OSDiskSizeGb               float64  `json:"osDiskSizeGB,omitempty,omitzero" bson:"osDiskSizeGB,omitempty,omitzero"`
		OSDiskType                 string   `json:"osDiskType,omitempty,omitzero" bson:"osDiskType,omitempty,omitzero"`
		OSSku                      string   `json:"osSKU,omitempty,omitzero" bson:"osSKU,omitempty,omitzero"`
		OSType                     string   `json:"osType,omitempty,omitzero" bson:"osType,omitempty,omitzero"`
		PowerState                 struct {
			Code string `json:"code,omitempty,omitzero" bson:"code,omitempty,omitzero"`
		} `json:"powerState,omitempty,omitzero" bson:"powerState,omitempty,omitzero"`
		ProvisioningState string `json:"provisioningState,omitempty,omitzero" bson:"provisioningState,omitempty,omitzero"`
		ScaleDownMode     string `json:"scaleDownMode,omitempty,omitzero" bson:"scaleDownMode,omitempty,omitzero"`
		SecurityProfile   struct {
			EnableSecureBoot bool `json:"enableSecureBoot,omitempty,omitzero" bson:"enableSecureBoot,omitempty,omitzero"`
			EnableVtpm       bool `json:"enableVTPM,omitempty,omitzero" bson:"enableVTPM,omitempty,omitzero"`
		} `json:"securityProfile,omitempty,omitzero" bson:"securityProfile,omitempty,omitzero"`
		Type            string   `json:"type,omitempty,omitzero" bson:"type,omitempty,omitzero"`
		UpgradeSettings struct{} `json:"upgradeSettings,omitempty,omitzero" bson:"upgradeSettings,omitempty,omitzero"`
		VmSize          string   `json:"vmSize,omitempty,omitzero" bson:"vmSize,omitempty,omitzero"`
		VnetSubnetID    string   `json:"vnetSubnetID,omitempty,omitzero" bson:"vnetSubnetID,omitempty,omitzero"`
		WorkloadRuntime string   `json:"workloadRuntime,omitempty,omitzero" bson:"workloadRuntime,omitempty,omitzero"`
	} `json:"properties,omitempty,omitzero" bson:"properties,omitempty,omitzero"`
	Type string `json:"type,omitempty,omitzero" bson:"type,omitempty,omitzero"`
}

//
//

type AzureResourceManagedClusterAgentPoolMachine struct {
	ID         string `json:"id,omitempty,omitzero" bson:"id,omitempty,omitzero"`
	Name       string `json:"name,omitempty,omitzero" bson:"name,omitempty,omitzero"`
	Properties struct {
		Network struct {
			IpAddresses []struct {
				Family string `json:"family,omitempty,omitzero" bson:"family,omitempty,omitzero"`
				Ip     string `json:"ip,omitempty,omitzero" bson:"ip,omitempty,omitzero"`
			} `json:"ipAddresses,omitempty,omitzero" bson:"ipAddresses,omitempty,omitzero"`
		} `json:"network,omitempty,omitzero" bson:"network,omitempty,omitzero"`
		ResourceID string `json:"resourceId,omitempty,omitzero" bson:"resourceId,omitempty,omitzero"`
	} `json:"properties,omitempty,omitzero" bson:"properties,omitempty,omitzero"`
	Type  string   `json:"type,omitempty,omitzero" bson:"type,omitempty,omitzero"`
	Zones []string `json:"zones,omitempty,omitzero" bson:"zones,omitempty,omitzero"`
}

//
//

type IPAddressesAllResourceTypes struct {
	AssociatedNics        []string  `json:"associatedNics,omitempty,omitzero" bson:"associatedNics,omitempty,omitzero"`
	AssociatedResourceIDs []string  `json:"associatedResourceIDs,omitempty,omitzero" bson:"associatedResourceIDs,omitempty,omitzero"`
	AttachedTo            string    `json:"attachedTo,omitempty,omitzero" bson:"attachedTo,omitempty,omitzero"`
	Cidrs                 []string  `json:"cidrs,omitempty,omitzero" bson:"cidrs,omitempty,omitzero"`
	ID                    string    `json:"id,omitempty,omitzero" bson:"_id,omitempty,omitzero"`
	InboundIps            []string  `json:"inboundIps,omitempty,omitzero" bson:"inboundIps,omitempty,omitzero"`
	IsAttached            bool      `json:"isAttached,omitempty,omitzero" bson:"isAttached,omitempty,omitzero"`
	LastAzureSync         time.Time `json:"lastAzureSync,omitempty,omitzero" bson:"lastAzureSync,omitempty,omitzero"`
	LastDBSync            time.Time `json:"lastDatabaseSync,omitempty,omitzero" bson:"lastDatabaseSync,omitempty,omitzero"`
	Name                  string    `json:"name,omitempty,omitzero" bson:"name,omitempty,omitzero"`
	OutboundIps           []string  `json:"outboundIps,omitempty,omitzero" bson:"outboundIps,omitempty,omitzero"`
	PossibleInboundIps    []string  `json:"possibleInboundIps,omitempty,omitzero" bson:"possibleInboundIps,omitempty,omitzero"`
	PossibleOutboundIps   []string  `json:"possibleOutboundIps,omitempty,omitzero" bson:"possibleOutboundIps,omitempty,omitzero"`
	PrivateIps            []string  `json:"privateIps,omitempty,omitzero" bson:"privateIps,omitempty,omitzero"`
	PublicIpIds           []string  `json:"publicIpIds,omitempty,omitzero" bson:"publicIpIds,omitempty,omitzero"`
	PublicIps             []string  `json:"publicIps,omitempty,omitzero" bson:"publicIps,omitempty,omitzero"`
	PublicNetworkAccess   string    `json:"publicNetworkAccess,omitempty,omitzero" bson:"publicNetworkAccess,omitempty,omitzero"`
	ResourceGroup         string    `json:"resourceGroup,omitempty,omitzero" bson:"resourceGroup,omitempty,omitzero"`
	SubscriptionID        string    `json:"subscriptionId,omitempty,omitzero" bson:"subscriptionId,omitempty,omitzero"`
	SubscriptionName      string    `json:"subscriptionName,omitempty,omitzero" bson:"subscriptionName,omitempty,omitzero"`
	Subnets               []struct {
		Cidrs []string `json:"cidrs,omitempty,omitzero" bson:"cidrs,omitempty,omitzero"`
		ID    string   `json:"id,omitempty,omitzero" bson:"id,omitempty,omitzero"`
		Name  string   `json:"name,omitempty,omitzero" bson:"name,omitempty,omitzero"`
		Type  string   `json:"type,omitempty,omitzero" bson:"type,omitempty,omitzero"`
	} `json:"subnets,omitempty,omitzero" bson:"subnets,omitempty,omitzero"`
	Tags       map[string]string `json:"tags" bson:"tags"`
	TenantID   string            `json:"tenantId,omitempty,omitzero" bson:"tenantId,omitempty,omitzero"`
	TenantName string            `json:"tenantName,omitempty,omitzero" bson:"tenantName,omitempty,omitzero"`
	Type       string            `json:"type,omitempty,omitzero" bson:"type,omitempty,omitzero"`
}

func (t *IPAddressesAllResourceTypes) UnmarshalJSON(data []byte) error {
	if string(data) == "null" || string(data) == `""` {
		return nil
	}

	var nicJson struct {
		AssociatedNics        []string  `json:"associatedNics,omitempty,omitzero" bson:"associatedNics,omitempty,omitzero"`
		AssociatedResourceIDs []string  `json:"associatedResourceIDs,omitempty,omitzero" bson:"associatedResourceIDs,omitempty,omitzero"`
		AttachedTo            string    `json:"attachedTo,omitempty,omitzero" bson:"attachedTo,omitempty,omitzero"`
		Cidrs                 []string  `json:"cidrs,omitempty,omitzero" bson:"cidrs,omitempty,omitzero"`
		ID                    string    `json:"id,omitempty,omitzero" bson:"id,omitempty,omitzero"`
		InboundIps            []string  `json:"inboundIps,omitempty,omitzero" bson:"inboundIps,omitempty,omitzero"`
		IsAttached            int64     `json:"isAttached,omitempty,omitzero" bson:"isAttached,omitempty,omitzero"`
		LastAzureSync         time.Time `json:"lastAzureSync,omitempty,omitzero" bson:"lastAzureSync,omitempty,omitzero"`
		LastDBSync            time.Time `json:"lastDatabaseSync,omitempty,omitzero" bson:"lastDatabaseSync,omitempty,omitzero"`
		Name                  string    `json:"name,omitempty,omitzero" bson:"name,omitempty,omitzero"`
		OutboundIps           []string  `json:"outboundIps,omitempty,omitzero" bson:"outboundIps,omitempty,omitzero"`
		PossibleInboundIps    []string  `json:"possibleInboundIps,omitempty,omitzero" bson:"possibleInboundIps,omitempty,omitzero"`
		PossibleOutboundIps   []string  `json:"possibleOutboundIps,omitempty,omitzero" bson:"possibleOutboundIps,omitempty,omitzero"`
		PrivateIps            []string  `json:"privateIps,omitempty,omitzero" bson:"privateIps,omitempty,omitzero"`
		PublicIpIds           []string  `json:"publicIpIds,omitempty,omitzero" bson:"publicIpIds,omitempty,omitzero"`
		PublicIps             []string  `json:"publicIps,omitempty,omitzero" bson:"publicIps,omitempty,omitzero"`
		PublicNetworkAccess   string    `json:"publicNetworkAccess,omitempty,omitzero" bson:"publicNetworkAccess,omitempty,omitzero"`
		ResourceGroup         string    `json:"resourceGroup,omitempty,omitzero" bson:"resourceGroup,omitempty,omitzero"`
		SubscriptionID        string    `json:"subscriptionId,omitempty,omitzero" bson:"subscriptionId,omitempty,omitzero"`
		SubscriptionName      string    `json:"subscriptionName,omitempty,omitzero" bson:"subscriptionName,omitempty,omitzero"`
		Subnets               []struct {
			Cidrs []string `json:"cidrs,omitempty,omitzero" bson:"cidrs,omitempty,omitzero"`
			ID    string   `json:"id,omitempty,omitzero" bson:"id,omitempty,omitzero"`
			Name  string   `json:"name,omitempty,omitzero" bson:"name,omitempty,omitzero"`
			Type  string   `json:"type,omitempty,omitzero" bson:"type,omitempty,omitzero"`
		} `json:"subnets,omitempty,omitzero" bson:"subnets,omitempty,omitzero"`
		Tags       string `json:"tags" bson:"tags"`
		TenantID   string `json:"tenantId,omitempty,omitzero" bson:"tenantId,omitempty,omitzero"`
		TenantName string `json:"tenantName,omitempty,omitzero" bson:"tenantName,omitempty,omitzero"`
		Type       string `json:"type,omitempty,omitzero" bson:"type,omitempty,omitzero"`
	}

	if err := json.Unmarshal(data, &nicJson); err != nil {
		return err
	}

	var tags map[string]string

	var isAttached bool
	if nicJson.IsAttached > 0 {
		isAttached = true
	}

	if nicJson.Tags != "" {
		if err := json.Unmarshal([]byte(nicJson.Tags), &tags); err != nil {
			return err
		}
		nicJson.Tags = ""
	}

	var cidrs []string

	for _, cidr := range nicJson.Cidrs {
		_, netip, err := net.ParseCIDR(cidr)
		if err != nil {
			_, netip, err := net.ParseCIDR(cidr + "/32")
			if err != nil {
				continue
			}
			cidrs = append(cidrs, netip.String())
		} else {
			cidrs = append(cidrs, netip.String())
		}
	}

	*t = IPAddressesAllResourceTypes{
		AssociatedNics:        nicJson.AssociatedNics,
		AssociatedResourceIDs: nicJson.AssociatedResourceIDs,
		AttachedTo:            nicJson.AttachedTo,
		Cidrs:                 cidrs,
		ID:                    nicJson.ID,
		InboundIps:            nicJson.InboundIps,
		IsAttached:            isAttached,
		LastAzureSync:         nicJson.LastAzureSync,
		LastDBSync:            nicJson.LastDBSync,
		Name:                  nicJson.Name,
		OutboundIps:           nicJson.OutboundIps,
		PossibleInboundIps:    nicJson.PossibleInboundIps,
		PossibleOutboundIps:   nicJson.PossibleOutboundIps,
		PrivateIps:            nicJson.PrivateIps,
		PublicIpIds:           nicJson.PublicIpIds,
		PublicIps:             nicJson.PublicIps,
		PublicNetworkAccess:   nicJson.PublicNetworkAccess,
		ResourceGroup:         nicJson.ResourceGroup,
		SubscriptionID:        nicJson.SubscriptionID,
		SubscriptionName:      nicJson.SubscriptionName,
		Subnets:               nicJson.Subnets,
		Tags:                  tags,
		TenantID:              nicJson.TenantID,
		TenantName:            nicJson.TenantName,
		Type:                  nicJson.Type,
	}
	// Subnets               []struct {
	// 	Cidrs
	// 	ID
	// 	Name
	// }

	return nil
}

type IPAddressesAllResourceTypesProcessed struct {
	AssociatedNics        []string  `json:"associatedNics,omitempty,omitzero" bson:"associatedNics,omitempty,omitzero"`
	AssociatedResourceIDs []string  `json:"associatedResourceIDs,omitempty,omitzero" bson:"associatedResourceIDs,omitempty,omitzero"`
	AttachedTo            string    `json:"attachedTo,omitempty,omitzero" bson:"attachedTo,omitempty,omitzero"`
	Cidrs                 []string  `json:"cidrs,omitempty,omitzero" bson:"cidrs,omitempty,omitzero"`
	ID                    string    `json:"id,omitempty,omitzero" bson:"id,omitempty,omitzero"`
	InboundIps            []string  `json:"inboundIps,omitempty,omitzero" bson:"inboundIps,omitempty,omitzero"`
	IsAttached            bool      `json:"isAttached,omitempty,omitzero" bson:"isAttached,omitempty,omitzero"`
	LastAzureSync         time.Time `json:"lastAzureSync,omitempty,omitzero" bson:"lastAzureSync,omitempty,omitzero"`
	LastDBSync            time.Time `json:"lastDatabaseSync,omitempty,omitzero" bson:"lastDatabaseSync,omitempty,omitzero"`
	Name                  string    `json:"name,omitempty,omitzero" bson:"name,omitempty,omitzero"`
	OutboundIps           []string  `json:"outboundIps,omitempty,omitzero" bson:"outboundIps,omitempty,omitzero"`
	PossibleInboundIps    []string  `json:"possibleInboundIps,omitempty,omitzero" bson:"possibleInboundIps,omitempty,omitzero"`
	PossibleOutboundIps   []string  `json:"possibleOutboundIps,omitempty,omitzero" bson:"possibleOutboundIps,omitempty,omitzero"`
	PrivateIps            []string  `json:"privateIps,omitempty,omitzero" bson:"privateIps,omitempty,omitzero"`
	PublicIpIds           []string  `json:"publicIpIds,omitempty,omitzero" bson:"publicIpIds,omitempty,omitzero"`
	PublicIps             []string  `json:"publicIps,omitempty,omitzero" bson:"publicIps,omitempty,omitzero"`
	PublicNetworkAccess   string    `json:"publicNetworkAccess,omitempty,omitzero" bson:"publicNetworkAccess,omitempty,omitzero"`
	ResourceGroup         string    `json:"resourceGroup,omitempty,omitzero" bson:"resourceGroup,omitempty,omitzero"`
	SubscriptionID        string    `json:"subscriptionId,omitempty,omitzero" bson:"subscriptionId,omitempty,omitzero"`
	SubscriptionName      string    `json:"subscriptionName,omitempty,omitzero" bson:"subscriptionName,omitempty,omitzero"`
	Subnets               []struct {
		Cidrs []string `json:"cidrs,omitempty,omitzero" bson:"cidrs,omitempty,omitzero"`
		ID    string   `json:"id,omitempty,omitzero" bson:"id,omitempty,omitzero"`
		Name  string   `json:"name,omitempty,omitzero" bson:"name,omitempty,omitzero"`
		Type  string   `json:"type,omitempty,omitzero" bson:"type,omitempty,omitzero"`
	} `json:"subnets,omitempty,omitzero" bson:"subnets,omitempty,omitzero"`
	Tags       map[string]string `json:"tags" bson:"tags"`
	TenantID   string            `json:"tenantId,omitempty,omitzero" bson:"tenantId,omitempty,omitzero"`
	TenantName string            `json:"tenantName,omitempty,omitzero" bson:"tenantName,omitempty,omitzero"`
	Type       string            `json:"type,omitempty,omitzero" bson:"type,omitempty,omitzero"`
}

//
//

type ResourceGraphGetIpsResponse struct {
	Count           float64                 `json:"count,omitempty,omitzero" bson:"count,omitempty,omitzero"`
	Data            []AzureResourceIPConfig `json:"data,omitempty,omitzero" bson:"data,omitempty,omitzero"`
	Facets          []interface{}           `json:"facets,omitempty,omitzero" bson:"facets,omitempty,omitzero"`
	ResultTruncated string                  `json:"resultTruncated,omitempty,omitzero" bson:"resultTruncated,omitempty,omitzero"`
	SkipToken       string                  `json:"$skipToken,omitempty,omitzero" bson:"$skipToken,omitempty,omitzero"`
	TotalRecords    float64                 `json:"totalRecords,omitempty,omitzero" bson:"totalRecords,omitempty,omitzero"`
}

//
//

type AzureResourceIPConfig struct {
	ID               string    `json:"id,omitempty,omitzero" bson:"id,omitempty,omitzero"`
	Name             string    `json:"name,omitempty,omitzero" bson:"name,omitempty,omitzero"`
	Type             string    `json:"type,omitempty,omitzero" bson:"type,omitempty,omitzero"`
	TenantName       string    `json:"tenantName,omitempty,omitzero" bson:"tenantName,omitempty,omitzero"`
	TenantId         string    `json:"tenantId,omitempty,omitzero" bson:"tenantId,omitempty,omitzero"`
	SubscriptionName string    `json:"subscriptionName,omitempty,omitzero" bson:"subscriptionName,omitempty,omitzero"`
	SubscriptionId   string    `json:"subscriptionId,omitempty,omitzero" bson:"subscriptionId,omitempty,omitzero"`
	PrivateIPs       []string  `json:"privateIps,omitempty,omitzero" bson:"privateIps,omitempty,omitzero"`
	PublicIPs        []string  `json:"publicIps,omitempty,omitzero" bson:"publicIps,omitempty,omitzero"`
	SubnetIds        []string  `json:"snetIds,omitempty,omitzero" bson:"snetIds,omitempty,omitzero"`
	AttachedVmId     string    `json:"attachedVmId,omitempty,omitzero" bson:"attachedVmId,omitempty,omitzero"`
	AttachedVmName   string    `json:"attachedVmName,omitempty,omitzero" bson:"attachedVmName,omitempty,omitzero"`
	LastAzureSync    time.Time `json:"lastAzureSync,omitempty,omitzero" bson:"lastAzureSync,omitempty,omitzero"`
	LastDBSync       time.Time `json:"lastDatabaseSync,omitempty,omitzero" bson:"lastDatabaseSync,omitempty,omitzero"`
}

//
//

type IpAddressBlock struct {
	AllocatedToVnet bool   `json:"allocatedToVnet" bson:"allocatedToVnet"`
	VNetName        string `json:"vnetName" bson:"vnetName"`
	// IpAddresses     []net.IP `json:"ipAddresses" bson:"ipAddresses"`
	FirstIp    string    `json:"firstIp" bson:"firstIp"`
	LastIp     string    `json:"lastIp" bson:"lastIp"`
	CidrBlocks []string  `json:"cidrBlocks" bson:"cidrBlocks"`
	LastDBSync time.Time `json:"lastDatabaseSync,omitempty,omitzero" bson:"lastDatabaseSync,omitempty,omitzero"`
}

//
//

type IpAddressBlocksByBlockTag struct {
	AddressBlocks []IpAddressBlock `json:"addressBlocks" bson:"addressBlocks"`
	BlockTag      string           `json:"blockTag" bson:"_id"`
}
