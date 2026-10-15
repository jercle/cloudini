package defender

import "time"

type ArmResponse[T Machine] struct {
	Odata_Context string `json:"@odata.context,omitempty,omitzero" bson:"@odata.context,omitempty,omitzero"`
	NextLink      string `json:"@odata.nextLink,omitempty,omitzero" bson:"@odata.nextLink,omitempty,omitzero"`
	Value         []T    `json:"value,omitempty,omitzero" bson:"value,omitempty,omitzero"`
}

type Machine struct {
	AadDeviceID            *string              `json:"aadDeviceId,omitempty,omitzero" bson:"aadDeviceId,omitempty,omitzero"`
	AgentVersion           string               `json:"agentVersion,omitempty,omitzero" bson:"agentVersion,omitempty,omitzero"`
	ComputerDnsName        string               `json:"computerDnsName,omitempty,omitzero" bson:"computerDnsName,omitempty,omitzero"`
	DeviceValue            string               `json:"deviceValue,omitempty,omitzero" bson:"deviceValue,omitempty,omitzero"`
	ExclusionReason        any                  `json:"exclusionReason,omitempty,omitzero" bson:"exclusionReason,omitempty,omitzero"`
	ExposureLevel          string               `json:"exposureLevel,omitempty,omitzero" bson:"exposureLevel,omitempty,omitzero"`
	FirstSeen              time.Time            `json:"firstSeen,omitempty,omitzero" bson:"firstSeen,omitempty,omitzero"`
	HealthStatus           string               `json:"healthStatus,omitempty,omitzero" bson:"healthStatus,omitempty,omitzero"`
	ID                     string               `json:"id,omitempty,omitzero" bson:"id,omitempty,omitzero"`
	IpAddresses            []MachineIpAddresses `json:"ipAddresses,omitempty,omitzero" bson:"ipAddresses,omitempty,omitzero"`
	IsAadJoined            *bool                `json:"isAadJoined,omitempty,omitzero" bson:"isAadJoined,omitempty,omitzero"`
	IsExcluded             bool                 `json:"isExcluded,omitempty,omitzero" bson:"isExcluded,omitempty,omitzero"`
	IsPotentialDuplication bool                 `json:"isPotentialDuplication,omitempty,omitzero" bson:"isPotentialDuplication,omitempty,omitzero"`
	LastExternalIpAddress  *string              `json:"lastExternalIpAddress,omitempty,omitzero" bson:"lastExternalIpAddress,omitempty,omitzero"`
	LastIpAddress          *string              `json:"lastIpAddress,omitempty,omitzero" bson:"lastIpAddress,omitempty,omitzero"`
	LastSeen               time.Time            `json:"lastSeen,omitempty,omitzero" bson:"lastSeen,omitempty,omitzero"`
	MachineTags            []any                `json:"machineTags,omitempty,omitzero" bson:"machineTags,omitempty,omitzero"`
	ManagedBy              string               `json:"managedBy,omitempty,omitzero" bson:"managedBy,omitempty,omitzero"`
	ManagedByStatus        string               `json:"managedByStatus,omitempty,omitzero" bson:"managedByStatus,omitempty,omitzero"`
	MergedIntoMachineID    any                  `json:"mergedIntoMachineId,omitempty,omitzero" bson:"mergedIntoMachineId,omitempty,omitzero"`
	OnboardingStatus       string               `json:"onboardingStatus,omitempty,omitzero" bson:"onboardingStatus,omitempty,omitzero"`
	OSArchitecture         *string              `json:"osArchitecture,omitempty,omitzero" bson:"osArchitecture,omitempty,omitzero"`
	OSBuild                *float64             `json:"osBuild,omitempty,omitzero" bson:"osBuild,omitempty,omitzero"`
	OSPlatform             string               `json:"osPlatform,omitempty,omitzero" bson:"osPlatform,omitempty,omitzero"`
	OSProcessor            *string              `json:"osProcessor,omitempty,omitzero" bson:"osProcessor,omitempty,omitzero"`
	RbacGroupID            float64              `json:"rbacGroupId,omitempty,omitzero" bson:"rbacGroupId,omitempty,omitzero"`
	RbacGroupName          string               `json:"rbacGroupName,omitempty,omitzero" bson:"rbacGroupName,omitempty,omitzero"`
	RiskScore              string               `json:"riskScore,omitempty,omitzero" bson:"riskScore,omitempty,omitzero"`
	Version                string               `json:"version,omitempty,omitzero" bson:"version,omitempty,omitzero"`
	VmMetadata             MachineVmMetadata    `json:"vmMetadata,omitempty,omitzero" bson:"vmMetadata,omitempty,omitzero"`
	TenantName             string               `json:"tenantName,omitempty,omitzero" bson:"tenantName,omitempty,omitzero"`
	MongoId                string               `json:"_id,omitempty,omitzero" bson:"_id,omitempty,omitzero"`
}

type MachineIpAddresses struct {
	IpAddress         string  `json:"ipAddress,omitempty,omitzero" bson:"ipAddress,omitempty,omitzero"`
	MacAddress        *string `json:"macAddress,omitempty,omitzero" bson:"macAddress,omitempty,omitzero"`
	OperationalStatus string  `json:"operationalStatus,omitempty,omitzero" bson:"operationalStatus,omitempty,omitzero"`
	Type              *string `json:"type,omitempty,omitzero" bson:"type,omitempty,omitzero"`
}

type MachineVmMetadata struct {
	CloudProvider  string  `json:"cloudProvider,omitempty,omitzero" bson:"cloudProvider,omitempty,omitzero"`
	ResourceID     string  `json:"resourceId,omitempty,omitzero" bson:"resourceId,omitempty,omitzero"`
	SubscriptionID *string `json:"subscriptionId,omitempty,omitzero" bson:"subscriptionId,omitempty,omitzero"`
	VmID           string  `json:"vmId,omitempty,omitzero" bson:"vmId,omitempty,omitzero"`
}
