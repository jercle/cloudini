package citrix

import "time"

type MachineMetricResponse struct {
	Context  string          `json:"@odata.context,omitempty,omitzero" bson:"@odata.context,omitempty,omitzero"`
	NextLink string          `json:"@odata.nextLink,omitempty,omitzero" bson:"@odata.nextLink,omitempty,omitzero"`
	Value    []MachineMetric `json:"value,omitempty,omitzero" bson:"value,omitempty,omitzero"`
}

type MachineMetric struct {
	CollectedDate time.Time `json:"CollectedDate,omitempty,omitzero" bson:"CollectedDate,omitempty,omitzero"`
	Iops          float64   `json:"Iops,omitempty,omitzero" bson:"Iops,omitempty,omitzero"`
	Latency       float64   `json:"Latency,omitempty,omitzero" bson:"Latency,omitempty,omitzero"`
	MachineID     string    `json:"MachineId,omitempty,omitzero" bson:"MachineId,omitempty,omitzero"`
}

type MachineResourceUtilisationResponse struct {
	Context  string                       `json:"@odata.context,omitempty,omitzero" bson:"@odata.context,omitempty,omitzero"`
	NextLink string                       `json:"@odata.nextLink,omitempty,omitzero" bson:"@odata.nextLink,omitempty,omitzero"`
	Value    []MachineResourceUtilisation `json:"value,omitempty,omitzero" bson:"value,omitempty,omitzero"`
}

type MachineResourceUtilisation struct {
	AvgIcaRttInMs               *float64  `json:"AvgIcaRttInMs,omitempty,omitzero" bson:"AvgIcaRttInMs,omitempty,omitzero"`
	CpuRelativeFrequencyPercent any       `json:"CPURelativeFrequencyPercent,omitempty,omitzero" bson:"CPURelativeFrequencyPercent,omitempty,omitzero"`
	CollectedDate               time.Time `json:"CollectedDate,omitempty,omitzero" bson:"CollectedDate,omitempty,omitzero"`
	CreatedDate                 time.Time `json:"CreatedDate,omitempty,omitzero" bson:"CreatedDate,omitempty,omitzero"`
	DesktopGroupID              string    `json:"DesktopGroupId,omitempty,omitzero" bson:"DesktopGroupId,omitempty,omitzero"`
	DiskIops                    float64   `json:"DiskIops,omitempty,omitzero" bson:"DiskIops,omitempty,omitzero"`
	DiskLatency                 float64   `json:"DiskLatency,omitempty,omitzero" bson:"DiskLatency,omitempty,omitzero"`
	DiskUsagePercent            any       `json:"DiskUsagePercent,omitempty,omitzero" bson:"DiskUsagePercent,omitempty,omitzero"`
	IdlenessPercent             any       `json:"IdlenessPercent,omitempty,omitzero" bson:"IdlenessPercent,omitempty,omitzero"`
	LoadIndexPercent            *float64  `json:"LoadIndexPercent,omitempty,omitzero" bson:"LoadIndexPercent,omitempty,omitzero"`
	MachineID                   string    `json:"MachineId,omitempty,omitzero" bson:"MachineId,omitempty,omitzero"`
	ModifiedDate                time.Time `json:"ModifiedDate,omitempty,omitzero" bson:"ModifiedDate,omitempty,omitzero"`
	NetUtilizationPercent       any       `json:"NetUtilizationPercent,omitempty,omitzero" bson:"NetUtilizationPercent,omitempty,omitzero"`
	PagefileUsagePercent        any       `json:"PagefileUsagePercent,omitempty,omitzero" bson:"PagefileUsagePercent,omitempty,omitzero"`
	PercentCpu                  float64   `json:"PercentCpu,omitempty,omitzero" bson:"PercentCpu,omitempty,omitzero"`
	PercentGpu                  any       `json:"PercentGPU,omitempty,omitzero" bson:"PercentGPU,omitempty,omitzero"`
	ProfileDiskUsagePercent     any       `json:"ProfileDiskUsagePercent,omitempty,omitzero" bson:"ProfileDiskUsagePercent,omitempty,omitzero"`
	RamUsagePercent             any       `json:"RAMUsagePercent,omitempty,omitzero" bson:"RAMUsagePercent,omitempty,omitzero"`
	SessionCount                float64   `json:"SessionCount,omitempty,omitzero" bson:"SessionCount,omitempty,omitzero"`
	TotalMemory                 float64   `json:"TotalMemory,omitempty,omitzero" bson:"TotalMemory,omitempty,omitzero"`
	UsedMemory                  float64   `json:"UsedMemory,omitempty,omitzero" bson:"UsedMemory,omitempty,omitzero"`
}

//
//

type MachineLoadIndexesResponse struct {
	Context  string             `json:"@odata.context,omitempty,omitzero" bson:"@odata.context,omitempty,omitzero"`
	NextLink string             `json:"@odata.nextLink,omitempty,omitzero" bson:"@odata.nextLink,omitempty,omitzero"`
	Value    []MachineLoadIndex `json:"value,omitempty,omitzero" bson:"value,omitempty,omitzero"`
}

type MachineLoadIndex struct {
	Cpu                float64   `json:"Cpu,omitempty,omitzero" bson:"Cpu,omitempty,omitzero"`
	CreatedDate        time.Time `json:"CreatedDate,omitempty,omitzero" bson:"CreatedDate,omitempty,omitzero"`
	Disk               any       `json:"Disk,omitempty,omitzero" bson:"Disk,omitempty,omitzero"`
	EffectiveLoadIndex float64   `json:"EffectiveLoadIndex,omitempty,omitzero" bson:"EffectiveLoadIndex,omitempty,omitzero"`
	ID                 float64   `json:"Id,omitempty,omitzero" bson:"_id,omitempty,omitzero"`
	MachineID          string    `json:"MachineId,omitempty,omitzero" bson:"MachineId,omitempty,omitzero"`
	Memory             *float64  `json:"Memory,omitempty,omitzero" bson:"Memory,omitempty,omitzero"`
	ModifiedDate       time.Time `json:"ModifiedDate,omitempty,omitzero" bson:"ModifiedDate,omitempty,omitzero"`
	Network            any       `json:"Network,omitempty,omitzero" bson:"Network,omitempty,omitzero"`
	SessionCount       float64   `json:"SessionCount,omitempty,omitzero" bson:"SessionCount,omitempty,omitzero"`
}

//
//

type MonitorMachinesResponse struct {
	Context  string           `json:"@odata.context,omitempty,omitzero" bson:"@odata.context,omitempty,omitzero"`
	NextLink string           `json:"@odata.nextLink,omitempty,omitzero" bson:"@odata.nextLink,omitempty,omitzero"`
	Value    []MonitorMachine `json:"value,omitempty,omitzero" bson:"value,omitempty,omitzero"`
}

type MonitorMachine struct {
	TenantName               string                      `json:"tenantName,omitempty,omitzero" bson:"tenantName,omitempty,omitzero"`
	AgentVersion             string                      `json:"AgentVersion,omitempty,omitzero" bson:"AgentVersion,omitempty,omitzero"`
	Metrics                  *MachineMetric              `json:"metrics,omitempty,omitzero" bson:"metrics,omitempty,omitzero"`
	ResourceUtilisation      *MachineResourceUtilisation `json:"resourceUtilisation,omitempty,omitzero" bson:"resourceUtilisation,omitempty,omitzero"`
	AssociatedUserNames      string                      `json:"AssociatedUserNames,omitempty,omitzero" bson:"AssociatedUserNames,omitempty,omitzero"`
	CurrentLoadIndex         *MachineLoadIndex           `json:"CurrentLoadIndex,omitempty,omitzero" bson:"CurrentLoadIndex,omitempty,omitzero"`
	CurrentLoadIndexID       float64                     `json:"CurrentLoadIndexId,omitempty,omitzero" bson:"CurrentLoadIndexId,omitempty,omitzero"`
	CurrentPowerState        float64                     `json:"CurrentPowerState,omitempty,omitzero" bson:"CurrentPowerState,omitempty,omitzero"`
	CurrentRegistrationState float64                     `json:"CurrentRegistrationState,omitempty,omitzero" bson:"CurrentRegistrationState,omitempty,omitzero"`
	CurrentSessionCount      float64                     `json:"CurrentSessionCount,omitempty,omitzero" bson:"CurrentSessionCount,omitempty,omitzero"`
	DnsName                  string                      `json:"DnsName,omitempty,omitzero" bson:"DnsName,omitempty,omitzero"`
	IpAddress                string                      `json:"IPAddress,omitempty,omitzero" bson:"IPAddress,omitempty,omitzero"`
	ID                       string                      `json:"Id,omitempty,omitzero" bson:"_id,omitempty,omitzero"`
	IsInMaintenanceMode      bool                        `json:"IsInMaintenanceMode,omitempty,omitzero" bson:"IsInMaintenanceMode,omitempty,omitzero"`
	LifecycleState           float64                     `json:"LifecycleState,omitempty,omitzero" bson:"LifecycleState,omitempty,omitzero"`
	Name                     string                      `json:"Name,omitempty,omitzero" bson:"Name,omitempty,omitzero"`
	PoweredOnDate            time.Time                   `json:"PoweredOnDate,omitempty,omitzero" bson:"PoweredOnDate,omitempty,omitzero"`
	// AssociatedUserFullNames      string     `json:"AssociatedUserFullNames,omitempty,omitzero" bson:"AssociatedUserFullNames,omitempty,omitzero"`
	// AssociatedUserUpNs           string     `json:"AssociatedUserUPNs,omitempty,omitzero" bson:"AssociatedUserUPNs,omitempty,omitzero"`
	// CatalogID                    string     `json:"CatalogId,omitempty,omitzero" bson:"CatalogId,omitempty,omitzero"`
	// ConsecutiveFailuresCount     *float64   `json:"ConsecutiveFailuresCount,omitempty,omitzero" bson:"ConsecutiveFailuresCount,omitempty,omitzero"`
	// ControllerDnsName            string     `json:"ControllerDnsName,omitempty,omitzero" bson:"ControllerDnsName,omitempty,omitzero"`
	// CreatedDate                  time.Time  `json:"CreatedDate,omitempty,omitzero" bson:"CreatedDate,omitempty,omitzero"`
	// DesktopGroupID               string     `json:"DesktopGroupId,omitempty,omitzero" bson:"DesktopGroupId,omitempty,omitzero"`
	// FailureDate                  *time.Time `json:"FailureDate,omitempty,omitzero" bson:"FailureDate,omitempty,omitzero"`
	// FaultState float64 `json:"FaultState,omitempty,omitzero" bson:"FaultState,omitempty,omitzero"`
	// FunctionalLevel              float64    `json:"FunctionalLevel,omitempty,omitzero" bson:"FunctionalLevel,omitempty,omitzero"`
	// Hash                         string     `json:"Hash,omitempty,omitzero" bson:"Hash,omitempty,omitzero"`
	// HostedMachineID              string     `json:"HostedMachineId,omitempty,omitzero" bson:"HostedMachineId,omitempty,omitzero"`
	// HostedMachineName            string     `json:"HostedMachineName,omitempty,omitzero" bson:"HostedMachineName,omitempty,omitzero"`
	// HostingServerName            string     `json:"HostingServerName,omitempty,omitzero" bson:"HostingServerName,omitempty,omitzero"`
	// HypervisorID                 string     `json:"HypervisorId,omitempty,omitzero" bson:"HypervisorId,omitempty,omitzero"`
	// IsAssigned                   bool       `json:"IsAssigned,omitempty,omitzero" bson:"IsAssigned,omitempty,omitzero"`
	// IsPendingUpdate              bool       `json:"IsPendingUpdate,omitempty,omitzero" bson:"IsPendingUpdate,omitempty,omitzero"`
	// IsPreparing                  bool       `json:"IsPreparing,omitempty,omitzero" bson:"IsPreparing,omitempty,omitzero"`
	// LastDeregisteredCode         float64    `json:"LastDeregisteredCode,omitempty,omitzero" bson:"LastDeregisteredCode,omitempty,omitzero"`
	// LastDeregisteredDate         time.Time  `json:"LastDeregisteredDate,omitempty,omitzero" bson:"LastDeregisteredDate,omitempty,omitzero"`
	// LastPowerActionCompletedDate time.Time  `json:"LastPowerActionCompletedDate,omitempty,omitzero" bson:"LastPowerActionCompletedDate,omitempty,omitzero"`
	// LastPowerActionFailureReason float64    `json:"LastPowerActionFailureReason,omitempty,omitzero" bson:"LastPowerActionFailureReason,omitempty,omitzero"`
	// LastPowerActionReason        float64    `json:"LastPowerActionReason,omitempty,omitzero" bson:"LastPowerActionReason,omitempty,omitzero"`
	// LastPowerActionType          float64    `json:"LastPowerActionType,omitempty,omitzero" bson:"LastPowerActionType,omitempty,omitzero"`
	// LastUpgradeState             any        `json:"LastUpgradeState,omitempty,omitzero" bson:"LastUpgradeState,omitempty,omitzero"`
	// LastUpgradeStateChangeDate   any        `json:"LastUpgradeStateChangeDate,omitempty,omitzero" bson:"LastUpgradeStateChangeDate,omitempty,omitzero"`
	// MachineRole                  float64    `json:"MachineRole,omitempty,omitzero" bson:"MachineRole,omitempty,omitzero"`
	// ModifiedDate                 time.Time  `json:"ModifiedDate,omitempty,omitzero" bson:"ModifiedDate,omitempty,omitzero"`
	// OSType                       string     `json:"OSType,omitempty,omitzero" bson:"OSType,omitempty,omitzero"`
	// PowerStateChangeDate         time.Time  `json:"PowerStateChangeDate,omitempty,omitzero" bson:"PowerStateChangeDate,omitempty,omitzero"`
	// RegistrationStateChangeDate  time.Time  `json:"RegistrationStateChangeDate,omitempty,omitzero" bson:"RegistrationStateChangeDate,omitempty,omitzero"`
	// ResumedDate                  any        `json:"ResumedDate,omitempty,omitzero" bson:"ResumedDate,omitempty,omitzero"`
	// Sid                          string     `json:"Sid,omitempty,omitzero" bson:"Sid,omitempty,omitzero"`
	// Tags                         []any      `json:"Tags,omitempty,omitzero" bson:"Tags,omitempty,omitzero"`
	// WindowsConnectionSetting     float64    `json:"WindowsConnectionSetting,omitempty,omitzero" bson:"WindowsConnectionSetting,omitempty,omitzero"`
	// ZombieFailuresCount          *float64   `json:"ZombieFailuresCount,omitempty,omitzero" bson:"ZombieFailuresCount,omitempty,omitzero"`
}
