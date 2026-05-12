package citrix

import "time"

type GetAllMachinesResponse struct {
	ContinuationToken string          `json:"ContinuationToken,omitempty,omitzero" bson:"ContinuationToken,omitempty,omitzero"`
	Items             []CitrixMachine `json:"Items,omitempty,omitzero" bson:"Items,omitempty,omitzero"`
}

type CitrixMachine struct {
	AgentVersion       string `json:"AgentVersion,omitempty,omitzero" bson:"AgentVersion,omitempty,omitzero"`
	AllocationType     string `json:"AllocationType,omitempty,omitzero" bson:"AllocationType,omitempty,omitzero"`
	ApplicationsInUse  []any  `json:"ApplicationsInUse,omitempty,omitzero" bson:"ApplicationsInUse,omitempty,omitzero"`
	AzureResourceGroup string `json:"AzureResourceGroup,omitempty,omitzero" bson:"AzureResourceGroup,omitempty,omitzero"`
	AssignedUsers      []struct {
		CanonicalName           any     `json:"CanonicalName,omitempty,omitzero" bson:"CanonicalName,omitempty,omitzero"`
		City                    any     `json:"City,omitempty,omitzero" bson:"City,omitempty,omitzero"`
		Claims                  any     `json:"Claims,omitempty,omitzero" bson:"Claims,omitempty,omitzero"`
		CommonName              any     `json:"CommonName,omitempty,omitzero" bson:"CommonName,omitempty,omitzero"`
		Country                 any     `json:"Country,omitempty,omitzero" bson:"Country,omitempty,omitzero"`
		DaysUntilPasswordExpiry any     `json:"DaysUntilPasswordExpiry,omitempty,omitzero" bson:"DaysUntilPasswordExpiry,omitempty,omitzero"`
		DenyOnlySids            any     `json:"DenyOnlySids,omitempty,omitzero" bson:"DenyOnlySids,omitempty,omitzero"`
		Directory               any     `json:"Directory,omitempty,omitzero" bson:"Directory,omitempty,omitzero"`
		DirectoryServer         any     `json:"DirectoryServer,omitempty,omitzero" bson:"DirectoryServer,omitempty,omitzero"`
		DisplayName             string  `json:"DisplayName,omitempty,omitzero" bson:"DisplayName,omitempty,omitzero"`
		DistinguishedName       any     `json:"DistinguishedName,omitempty,omitzero" bson:"DistinguishedName,omitempty,omitzero"`
		Domain                  any     `json:"Domain,omitempty,omitzero" bson:"Domain,omitempty,omitzero"`
		Enabled                 any     `json:"Enabled,omitempty,omitzero" bson:"Enabled,omitempty,omitzero"`
		Forest                  any     `json:"Forest,omitempty,omitzero" bson:"Forest,omitempty,omitzero"`
		GroupSids               any     `json:"GroupSids,omitempty,omitzero" bson:"GroupSids,omitempty,omitzero"`
		Guid                    any     `json:"Guid,omitempty,omitzero" bson:"Guid,omitempty,omitzero"`
		HomePhone               any     `json:"HomePhone,omitempty,omitzero" bson:"HomePhone,omitempty,omitzero"`
		IsBuiltIn               any     `json:"IsBuiltIn,omitempty,omitzero" bson:"IsBuiltIn,omitempty,omitzero"`
		IsGroup                 any     `json:"IsGroup,omitempty,omitzero" bson:"IsGroup,omitempty,omitzero"`
		Locked                  any     `json:"Locked,omitempty,omitzero" bson:"Locked,omitempty,omitzero"`
		Mail                    any     `json:"Mail,omitempty,omitzero" bson:"Mail,omitempty,omitzero"`
		Mobile                  any     `json:"Mobile,omitempty,omitzero" bson:"Mobile,omitempty,omitzero"`
		Name                    any     `json:"Name,omitempty,omitzero" bson:"Name,omitempty,omitzero"`
		Oid                     any     `json:"Oid,omitempty,omitzero" bson:"Oid,omitempty,omitzero"`
		PasswordCanExpire       any     `json:"PasswordCanExpire,omitempty,omitzero" bson:"PasswordCanExpire,omitempty,omitzero"`
		PossibleLookupFailure   bool    `json:"PossibleLookupFailure,omitempty,omitzero" bson:"PossibleLookupFailure,omitempty,omitzero"`
		PrincipalName           string  `json:"PrincipalName,omitempty,omitzero" bson:"PrincipalName,omitempty,omitzero"`
		PropertiesFetched       float64 `json:"PropertiesFetched,omitempty,omitzero" bson:"PropertiesFetched,omitempty,omitzero"`
		SamAccountName          any     `json:"SamAccountName,omitempty,omitzero" bson:"SamAccountName,omitempty,omitzero"`
		SamName                 string  `json:"SamName,omitempty,omitzero" bson:"SamName,omitempty,omitzero"`
		Sid                     string  `json:"Sid,omitempty,omitzero" bson:"Sid,omitempty,omitzero"`
		State                   any     `json:"State,omitempty,omitzero" bson:"State,omitempty,omitzero"`
		StreetAddress           any     `json:"StreetAddress,omitempty,omitzero" bson:"StreetAddress,omitempty,omitzero"`
		TelephoneNumber         any     `json:"TelephoneNumber,omitempty,omitzero" bson:"TelephoneNumber,omitempty,omitzero"`
		UserIdentity            any     `json:"UserIdentity,omitempty,omitzero" bson:"UserIdentity,omitempty,omitzero"`
	} `json:"AssignedUsers,omitempty,omitzero" bson:"AssignedUsers,omitempty,omitzero"`
	AssociatedUsers []struct {
		CanonicalName           any     `json:"CanonicalName,omitempty,omitzero" bson:"CanonicalName,omitempty,omitzero"`
		City                    any     `json:"City,omitempty,omitzero" bson:"City,omitempty,omitzero"`
		Claims                  any     `json:"Claims,omitempty,omitzero" bson:"Claims,omitempty,omitzero"`
		CommonName              any     `json:"CommonName,omitempty,omitzero" bson:"CommonName,omitempty,omitzero"`
		Country                 any     `json:"Country,omitempty,omitzero" bson:"Country,omitempty,omitzero"`
		DaysUntilPasswordExpiry any     `json:"DaysUntilPasswordExpiry,omitempty,omitzero" bson:"DaysUntilPasswordExpiry,omitempty,omitzero"`
		DenyOnlySids            any     `json:"DenyOnlySids,omitempty,omitzero" bson:"DenyOnlySids,omitempty,omitzero"`
		Directory               any     `json:"Directory,omitempty,omitzero" bson:"Directory,omitempty,omitzero"`
		DirectoryServer         any     `json:"DirectoryServer,omitempty,omitzero" bson:"DirectoryServer,omitempty,omitzero"`
		DisplayName             string  `json:"DisplayName,omitempty,omitzero" bson:"DisplayName,omitempty,omitzero"`
		DistinguishedName       any     `json:"DistinguishedName,omitempty,omitzero" bson:"DistinguishedName,omitempty,omitzero"`
		Domain                  any     `json:"Domain,omitempty,omitzero" bson:"Domain,omitempty,omitzero"`
		Enabled                 any     `json:"Enabled,omitempty,omitzero" bson:"Enabled,omitempty,omitzero"`
		Forest                  any     `json:"Forest,omitempty,omitzero" bson:"Forest,omitempty,omitzero"`
		GroupSids               any     `json:"GroupSids,omitempty,omitzero" bson:"GroupSids,omitempty,omitzero"`
		Guid                    any     `json:"Guid,omitempty,omitzero" bson:"Guid,omitempty,omitzero"`
		HomePhone               any     `json:"HomePhone,omitempty,omitzero" bson:"HomePhone,omitempty,omitzero"`
		IsBuiltIn               any     `json:"IsBuiltIn,omitempty,omitzero" bson:"IsBuiltIn,omitempty,omitzero"`
		IsGroup                 any     `json:"IsGroup,omitempty,omitzero" bson:"IsGroup,omitempty,omitzero"`
		Locked                  any     `json:"Locked,omitempty,omitzero" bson:"Locked,omitempty,omitzero"`
		Mail                    any     `json:"Mail,omitempty,omitzero" bson:"Mail,omitempty,omitzero"`
		Mobile                  any     `json:"Mobile,omitempty,omitzero" bson:"Mobile,omitempty,omitzero"`
		Name                    any     `json:"Name,omitempty,omitzero" bson:"Name,omitempty,omitzero"`
		Oid                     any     `json:"Oid,omitempty,omitzero" bson:"Oid,omitempty,omitzero"`
		PasswordCanExpire       any     `json:"PasswordCanExpire,omitempty,omitzero" bson:"PasswordCanExpire,omitempty,omitzero"`
		PossibleLookupFailure   bool    `json:"PossibleLookupFailure,omitempty,omitzero" bson:"PossibleLookupFailure,omitempty,omitzero"`
		PrincipalName           string  `json:"PrincipalName,omitempty,omitzero" bson:"PrincipalName,omitempty,omitzero"`
		PropertiesFetched       float64 `json:"PropertiesFetched,omitempty,omitzero" bson:"PropertiesFetched,omitempty,omitzero"`
		SamAccountName          any     `json:"SamAccountName,omitempty,omitzero" bson:"SamAccountName,omitempty,omitzero"`
		SamName                 string  `json:"SamName,omitempty,omitzero" bson:"SamName,omitempty,omitzero"`
		Sid                     string  `json:"Sid,omitempty,omitzero" bson:"Sid,omitempty,omitzero"`
		State                   any     `json:"State,omitempty,omitzero" bson:"State,omitempty,omitzero"`
		StreetAddress           any     `json:"StreetAddress,omitempty,omitzero" bson:"StreetAddress,omitempty,omitzero"`
		TelephoneNumber         any     `json:"TelephoneNumber,omitempty,omitzero" bson:"TelephoneNumber,omitempty,omitzero"`
		UserIdentity            any     `json:"UserIdentity,omitempty,omitzero" bson:"UserIdentity,omitempty,omitzero"`
	} `json:"AssociatedUsers,omitempty,omitzero" bson:"AssociatedUsers,omitempty,omitzero"`
	AzureAdJoinedMode       string  `json:"AzureAdJoinedMode,omitempty,omitzero" bson:"AzureAdJoinedMode,omitempty,omitzero"`
	CloudPcProvisioningType string  `json:"CloudPCProvisioningType,omitempty,omitzero" bson:"CloudPCProvisioningType,omitempty,omitzero"`
	ConnectorID             *string `json:"ConnectorId,omitempty,omitzero" bson:"ConnectorId,omitempty,omitzero"`
	ConnectorName           any     `json:"ConnectorName,omitempty,omitzero" bson:"ConnectorName,omitempty,omitzero"`
	ContainerMetadata       any     `json:"ContainerMetadata,omitempty,omitzero" bson:"ContainerMetadata,omitempty,omitzero"`
	ContainerScopes         []struct {
		ScopeType string `json:"ScopeType,omitempty,omitzero" bson:"ScopeType,omitempty,omitzero"`
		Scopes    []any  `json:"Scopes,omitempty,omitzero" bson:"Scopes,omitempty,omitzero"`
	} `json:"ContainerScopes,omitempty,omitzero" bson:"ContainerScopes,omitempty,omitzero"`
	ControllerDnsName any `json:"ControllerDnsName,omitempty,omitzero" bson:"ControllerDnsName,omitempty,omitzero"`
	CriticalIssues    []struct {
		Detail  any    `json:"Detail,omitempty,omitzero" bson:"Detail,omitempty,omitzero"`
		Message any    `json:"Message,omitempty,omitzero" bson:"Message,omitempty,omitzero"`
		Type    string `json:"Type,omitempty,omitzero" bson:"Type,omitempty,omitzero"`
	} `json:"CriticalIssues,omitempty,omitzero" bson:"CriticalIssues,omitempty,omitzero"`
	DeliveryGroup struct {
		ID   string  `json:"Id,omitempty,omitzero" bson:"Id,omitempty,omitzero"`
		Name string  `json:"Name,omitempty,omitzero" bson:"Name,omitempty,omitzero"`
		Uid  float64 `json:"Uid,omitempty,omitzero" bson:"Uid,omitempty,omitzero"`
	} `json:"DeliveryGroup,omitempty,omitzero" bson:"DeliveryGroup,omitempty,omitzero"`
	DeliveryType                    string     `json:"DeliveryType,omitempty,omitzero" bson:"DeliveryType,omitempty,omitzero"`
	Description                     *string    `json:"Description,omitempty,omitzero" bson:"Description,omitempty,omitzero"`
	DesktopConditions               []any      `json:"DesktopConditions,omitempty,omitzero" bson:"DesktopConditions,omitempty,omitzero"`
	DnsName                         string     `json:"DnsName,omitempty,omitzero" bson:"DnsName,omitempty,omitzero"`
	DrainingReason                  any        `json:"DrainingReason,omitempty,omitzero" bson:"DrainingReason,omitempty,omitzero"`
	DrainingUntilShutdown           bool       `json:"DrainingUntilShutdown,omitempty,omitzero" bson:"DrainingUntilShutdown,omitempty,omitzero"`
	FaultState                      string     `json:"FaultState,omitempty,omitzero" bson:"FaultState,omitempty,omitzero"`
	FormattedLastConnectionTime     *time.Time `json:"FormattedLastConnectionTime,omitempty,omitzero" bson:"FormattedLastConnectionTime,omitempty,omitzero"`
	FormattedLastDeregistrationTime time.Time  `json:"FormattedLastDeregistrationTime,omitempty,omitzero" bson:"FormattedLastDeregistrationTime,omitempty,omitzero"`
	FormattedLastErrorTime          any        `json:"FormattedLastErrorTime,omitempty,omitzero" bson:"FormattedLastErrorTime,omitempty,omitzero"`
	FormattedSessionStartTime       any        `json:"FormattedSessionStartTime,omitempty,omitzero" bson:"FormattedSessionStartTime,omitempty,omitzero"`
	FormattedSessionStateChangeTime any        `json:"FormattedSessionStateChangeTime,omitempty,omitzero" bson:"FormattedSessionStateChangeTime,omitempty,omitzero"`
	FunctionalLevel                 string     `json:"FunctionalLevel,omitempty,omitzero" bson:"FunctionalLevel,omitempty,omitzero"`
	Hosting                         *struct {
		FormattedLastHostingUpdateTime time.Time `json:"FormattedLastHostingUpdateTime,omitempty,omitzero" bson:"FormattedLastHostingUpdateTime,omitempty,omitzero"`
		HostedMachineID                string    `json:"HostedMachineId,omitempty,omitzero" bson:"HostedMachineId,omitempty,omitzero"`
		HostedMachineName              *string   `json:"HostedMachineName,omitempty,omitzero" bson:"HostedMachineName,omitempty,omitzero"`
		HostingServerName              any       `json:"HostingServerName,omitempty,omitzero" bson:"HostingServerName,omitempty,omitzero"`
		HypervisorConnection           struct {
			ID   string  `json:"Id,omitempty,omitzero" bson:"Id,omitempty,omitzero"`
			Name string  `json:"Name,omitempty,omitzero" bson:"Name,omitempty,omitzero"`
			Uid  float64 `json:"Uid,omitempty,omitzero" bson:"Uid,omitempty,omitzero"`
		} `json:"HypervisorConnection,omitempty,omitzero" bson:"HypervisorConnection,omitempty,omitzero"`
		ImageOutOfDate        bool   `json:"ImageOutOfDate,omitempty,omitzero" bson:"ImageOutOfDate,omitempty,omitzero"`
		LastHostingUpdateTime string `json:"LastHostingUpdateTime,omitempty,omitzero" bson:"LastHostingUpdateTime,omitempty,omitzero"`
	} `json:"Hosting,omitempty,omitzero" bson:"Hosting,omitempty,omitzero"`
	IpAddress             *string `json:"IPAddress,omitempty,omitzero" bson:"IPAddress,omitempty,omitzero"`
	IconID                string  `json:"IconId,omitempty,omitzero" bson:"IconId,omitempty,omitzero"`
	ID                    string  `json:"Id,omitempty,omitzero" bson:"Id,omitempty,omitzero"`
	InMaintenanceMode     bool    `json:"InMaintenanceMode,omitempty,omitzero" bson:"InMaintenanceMode,omitempty,omitzero"`
	IsAssigned            bool    `json:"IsAssigned,omitempty,omitzero" bson:"IsAssigned,omitempty,omitzero"`
	IsDraining            any     `json:"IsDraining,omitempty,omitzero" bson:"IsDraining,omitempty,omitzero"`
	LastConnectionFailure any     `json:"LastConnectionFailure,omitempty,omitzero" bson:"LastConnectionFailure,omitempty,omitzero"`
	LastConnectionTime    *string `json:"LastConnectionTime,omitempty,omitzero" bson:"LastConnectionTime,omitempty,omitzero"`
	LastConnectionUser    *struct {
		CanonicalName           any     `json:"CanonicalName,omitempty,omitzero" bson:"CanonicalName,omitempty,omitzero"`
		City                    any     `json:"City,omitempty,omitzero" bson:"City,omitempty,omitzero"`
		Claims                  any     `json:"Claims,omitempty,omitzero" bson:"Claims,omitempty,omitzero"`
		CommonName              any     `json:"CommonName,omitempty,omitzero" bson:"CommonName,omitempty,omitzero"`
		Country                 any     `json:"Country,omitempty,omitzero" bson:"Country,omitempty,omitzero"`
		DaysUntilPasswordExpiry any     `json:"DaysUntilPasswordExpiry,omitempty,omitzero" bson:"DaysUntilPasswordExpiry,omitempty,omitzero"`
		DenyOnlySids            any     `json:"DenyOnlySids,omitempty,omitzero" bson:"DenyOnlySids,omitempty,omitzero"`
		Directory               any     `json:"Directory,omitempty,omitzero" bson:"Directory,omitempty,omitzero"`
		DirectoryServer         any     `json:"DirectoryServer,omitempty,omitzero" bson:"DirectoryServer,omitempty,omitzero"`
		DisplayName             string  `json:"DisplayName,omitempty,omitzero" bson:"DisplayName,omitempty,omitzero"`
		DistinguishedName       any     `json:"DistinguishedName,omitempty,omitzero" bson:"DistinguishedName,omitempty,omitzero"`
		Domain                  any     `json:"Domain,omitempty,omitzero" bson:"Domain,omitempty,omitzero"`
		Enabled                 any     `json:"Enabled,omitempty,omitzero" bson:"Enabled,omitempty,omitzero"`
		Forest                  any     `json:"Forest,omitempty,omitzero" bson:"Forest,omitempty,omitzero"`
		GroupSids               any     `json:"GroupSids,omitempty,omitzero" bson:"GroupSids,omitempty,omitzero"`
		Guid                    any     `json:"Guid,omitempty,omitzero" bson:"Guid,omitempty,omitzero"`
		HomePhone               any     `json:"HomePhone,omitempty,omitzero" bson:"HomePhone,omitempty,omitzero"`
		IsBuiltIn               any     `json:"IsBuiltIn,omitempty,omitzero" bson:"IsBuiltIn,omitempty,omitzero"`
		IsGroup                 any     `json:"IsGroup,omitempty,omitzero" bson:"IsGroup,omitempty,omitzero"`
		Locked                  any     `json:"Locked,omitempty,omitzero" bson:"Locked,omitempty,omitzero"`
		Mail                    any     `json:"Mail,omitempty,omitzero" bson:"Mail,omitempty,omitzero"`
		Mobile                  any     `json:"Mobile,omitempty,omitzero" bson:"Mobile,omitempty,omitzero"`
		Name                    any     `json:"Name,omitempty,omitzero" bson:"Name,omitempty,omitzero"`
		Oid                     any     `json:"Oid,omitempty,omitzero" bson:"Oid,omitempty,omitzero"`
		PasswordCanExpire       any     `json:"PasswordCanExpire,omitempty,omitzero" bson:"PasswordCanExpire,omitempty,omitzero"`
		PossibleLookupFailure   bool    `json:"PossibleLookupFailure,omitempty,omitzero" bson:"PossibleLookupFailure,omitempty,omitzero"`
		PrincipalName           any     `json:"PrincipalName,omitempty,omitzero" bson:"PrincipalName,omitempty,omitzero"`
		PropertiesFetched       float64 `json:"PropertiesFetched,omitempty,omitzero" bson:"PropertiesFetched,omitempty,omitzero"`
		SamAccountName          any     `json:"SamAccountName,omitempty,omitzero" bson:"SamAccountName,omitempty,omitzero"`
		SamName                 string  `json:"SamName,omitempty,omitzero" bson:"SamName,omitempty,omitzero"`
		Sid                     any     `json:"Sid,omitempty,omitzero" bson:"Sid,omitempty,omitzero"`
		State                   any     `json:"State,omitempty,omitzero" bson:"State,omitempty,omitzero"`
		StreetAddress           any     `json:"StreetAddress,omitempty,omitzero" bson:"StreetAddress,omitempty,omitzero"`
		TelephoneNumber         any     `json:"TelephoneNumber,omitempty,omitzero" bson:"TelephoneNumber,omitempty,omitzero"`
		UserIdentity            any     `json:"UserIdentity,omitempty,omitzero" bson:"UserIdentity,omitempty,omitzero"`
	} `json:"LastConnectionUser,omitempty,omitzero" bson:"LastConnectionUser,omitempty,omitzero"`
	LastDeregistrationReason string    `json:"LastDeregistrationReason,omitempty,omitzero" bson:"LastDeregistrationReason,omitempty,omitzero"`
	LastDeregistrationTime   string    `json:"LastDeregistrationTime,omitempty,omitzero" bson:"LastDeregistrationTime,omitempty,omitzero"`
	LastErrorReason          any       `json:"LastErrorReason,omitempty,omitzero" bson:"LastErrorReason,omitempty,omitzero"`
	LastErrorTime            any       `json:"LastErrorTime,omitempty,omitzero" bson:"LastErrorTime,omitempty,omitzero"`
	LoadIndex                *float64  `json:"LoadIndex,omitempty,omitzero" bson:"LoadIndex,omitempty,omitzero"`
	LoadIndexNames           []string  `json:"LoadIndexNames,omitempty,omitzero" bson:"LoadIndexNames,omitempty,omitzero"`
	LoadIndexes              []float64 `json:"LoadIndexes,omitempty,omitzero" bson:"LoadIndexes,omitempty,omitzero"`
	MachineCatalog           struct {
		ID   string  `json:"Id,omitempty,omitzero" bson:"Id,omitempty,omitzero"`
		Name string  `json:"Name,omitempty,omitzero" bson:"Name,omitempty,omitzero"`
		Uid  float64 `json:"Uid,omitempty,omitzero" bson:"Uid,omitempty,omitzero"`
	} `json:"MachineCatalog,omitempty,omitzero" bson:"MachineCatalog,omitempty,omitzero"`
	MachineConfigurationOutOfSync any    `json:"MachineConfigurationOutOfSync,omitempty,omitzero" bson:"MachineConfigurationOutOfSync,omitempty,omitzero"`
	MachineType                   string `json:"MachineType,omitempty,omitzero" bson:"MachineType,omitempty,omitzero"`
	MachineUnavailableReason      string `json:"MachineUnavailableReason,omitempty,omitzero" bson:"MachineUnavailableReason,omitempty,omitzero"`
	MaintenanceModeReason         string `json:"MaintenanceModeReason,omitempty,omitzero" bson:"MaintenanceModeReason,omitempty,omitzero"`
	Metadata                      []struct {
		Name  string `json:"Name,omitempty,omitzero" bson:"Name,omitempty,omitzero"`
		Value string `json:"Value,omitempty,omitzero" bson:"Value,omitempty,omitzero"`
	} `json:"Metadata,omitempty,omitzero" bson:"Metadata,omitempty,omitzero"`
	Name              string `json:"Name,omitempty,omitzero" bson:"Name,omitempty,omitzero"`
	NonCriticalIssues []struct {
		Detail  any    `json:"Detail,omitempty,omitzero" bson:"Detail,omitempty,omitzero"`
		Message any    `json:"Message,omitempty,omitzero" bson:"Message,omitempty,omitzero"`
		Type    string `json:"Type,omitempty,omitzero" bson:"Type,omitempty,omitzero"`
	} `json:"NonCriticalIssues,omitempty,omitzero" bson:"NonCriticalIssues,omitempty,omitzero"`
	OSType                                         string   `json:"OSType,omitempty,omitzero" bson:"OSType,omitempty,omitzero"`
	OSVersion                                      string   `json:"OSVersion,omitempty,omitzero" bson:"OSVersion,omitempty,omitzero"`
	PersistUserChanges                             string   `json:"PersistUserChanges,omitempty,omitzero" bson:"PersistUserChanges,omitempty,omitzero"`
	PowerActionPending                             bool     `json:"PowerActionPending,omitempty,omitzero" bson:"PowerActionPending,omitempty,omitzero"`
	PowerState                                     string   `json:"PowerState,omitempty,omitzero" bson:"PowerState,omitempty,omitzero"`
	ProvisioningMaintenanceMode                    string   `json:"ProvisioningMaintenanceMode,omitempty,omitzero" bson:"ProvisioningMaintenanceMode,omitempty,omitzero"`
	ProvisioningType                               string   `json:"ProvisioningType,omitempty,omitzero" bson:"ProvisioningType,omitempty,omitzero"`
	PublishedApplications                          []any    `json:"PublishedApplications,omitempty,omitzero" bson:"PublishedApplications,omitempty,omitzero"`
	PublishedName                                  *string  `json:"PublishedName,omitempty,omitzero" bson:"PublishedName,omitempty,omitzero"`
	RegistrationState                              string   `json:"RegistrationState,omitempty,omitzero" bson:"RegistrationState,omitempty,omitzero"`
	ScheduledReboot                                string   `json:"ScheduledReboot,omitempty,omitzero" bson:"ScheduledReboot,omitempty,omitzero"`
	SessionClientAddress                           any      `json:"SessionClientAddress,omitempty,omitzero" bson:"SessionClientAddress,omitempty,omitzero"`
	SessionClientName                              any      `json:"SessionClientName,omitempty,omitzero" bson:"SessionClientName,omitempty,omitzero"`
	SessionClientVersion                           any      `json:"SessionClientVersion,omitempty,omitzero" bson:"SessionClientVersion,omitempty,omitzero"`
	SessionConnectedViaHostName                    any      `json:"SessionConnectedViaHostName,omitempty,omitzero" bson:"SessionConnectedViaHostName,omitempty,omitzero"`
	SessionConnectedViaIp                          any      `json:"SessionConnectedViaIP,omitempty,omitzero" bson:"SessionConnectedViaIP,omitempty,omitzero"`
	SessionCount                                   float64  `json:"SessionCount,omitempty,omitzero" bson:"SessionCount,omitempty,omitzero"`
	SessionLaunchedViaHostName                     any      `json:"SessionLaunchedViaHostName,omitempty,omitzero" bson:"SessionLaunchedViaHostName,omitempty,omitzero"`
	SessionLaunchedViaIp                           any      `json:"SessionLaunchedViaIP,omitempty,omitzero" bson:"SessionLaunchedViaIP,omitempty,omitzero"`
	SessionProtocol                                any      `json:"SessionProtocol,omitempty,omitzero" bson:"SessionProtocol,omitempty,omitzero"`
	SessionSecureIcaActive                         any      `json:"SessionSecureIcaActive,omitempty,omitzero" bson:"SessionSecureIcaActive,omitempty,omitzero"`
	SessionSmartAccessTags                         []any    `json:"SessionSmartAccessTags,omitempty,omitzero" bson:"SessionSmartAccessTags,omitempty,omitzero"`
	SessionStartTime                               any      `json:"SessionStartTime,omitempty,omitzero" bson:"SessionStartTime,omitempty,omitzero"`
	SessionState                                   any      `json:"SessionState,omitempty,omitzero" bson:"SessionState,omitempty,omitzero"`
	SessionStateChangeTime                         any      `json:"SessionStateChangeTime,omitempty,omitzero" bson:"SessionStateChangeTime,omitempty,omitzero"`
	SessionSupport                                 string   `json:"SessionSupport,omitempty,omitzero" bson:"SessionSupport,omitempty,omitzero"`
	SessionUserName                                any      `json:"SessionUserName,omitempty,omitzero" bson:"SessionUserName,omitempty,omitzero"`
	Sid                                            string   `json:"Sid,omitempty,omitzero" bson:"Sid,omitempty,omitzero"`
	SummaryState                                   string   `json:"SummaryState,omitempty,omitzero" bson:"SummaryState,omitempty,omitzero"`
	SupportedPowerActions                          []string `json:"SupportedPowerActions,omitempty,omitzero" bson:"SupportedPowerActions,omitempty,omitzero"`
	Tags                                           []any    `json:"Tags,omitempty,omitzero" bson:"Tags,omitempty,omitzero"`
	Uid                                            float64  `json:"Uid,omitempty,omitzero" bson:"Uid,omitempty,omitzero"`
	UpgradeDetail                                  any      `json:"UpgradeDetail,omitempty,omitzero" bson:"UpgradeDetail,omitempty,omitzero"`
	UpgradeState                                   any      `json:"UpgradeState,omitempty,omitzero" bson:"UpgradeState,omitempty,omitzero"`
	UpgradeType                                    any      `json:"UpgradeType,omitempty,omitzero" bson:"UpgradeType,omitempty,omitzero"`
	WillShutdownAfterUse                           bool     `json:"WillShutdownAfterUse,omitempty,omitzero" bson:"WillShutdownAfterUse,omitempty,omitzero"`
	WindowsActivationStatus                        any      `json:"WindowsActivationStatus,omitempty,omitzero" bson:"WindowsActivationStatus,omitempty,omitzero"`
	WindowsActivationStatusErrorCode               any      `json:"WindowsActivationStatusErrorCode,omitempty,omitzero" bson:"WindowsActivationStatusErrorCode,omitempty,omitzero"`
	WindowsActivationStatusVirtualMachineError     any      `json:"WindowsActivationStatusVirtualMachineError,omitempty,omitzero" bson:"WindowsActivationStatusVirtualMachineError,omitempty,omitzero"`
	WindowsActivationTypeProvisionedVirtualMachine any      `json:"WindowsActivationTypeProvisionedVirtualMachine,omitempty,omitzero" bson:"WindowsActivationTypeProvisionedVirtualMachine,omitempty,omitzero"`
	WindowsConnectionSetting                       string   `json:"WindowsConnectionSetting,omitempty,omitzero" bson:"WindowsConnectionSetting,omitempty,omitzero"`
	Zone                                           struct {
		ID   string `json:"Id,omitempty,omitzero" bson:"Id,omitempty,omitzero"`
		Name string `json:"Name,omitempty,omitzero" bson:"Name,omitempty,omitzero"`
		Uid  any    `json:"Uid,omitempty,omitzero" bson:"Uid,omitempty,omitzero"`
	} `json:"Zone,omitempty,omitzero" bson:"Zone,omitempty,omitzero"`
}
