package m365

import "time"

type MailboxUsageDetail struct {
	// ReportRefreshDate            string `csv:"Report Refresh Date"`
	UserPrincipalName string `csv:"User Principal Name" json:"userPrincipalName,omitempty,omitzero" bson:"userPrincipalName,omitempty,omitzero"`
	// DisplayName                  string `csv:"Display Name"`
	// IsDeleted                    string `csv:"Is Deleted"`
	// DeletedDate      string `csv:"Deleted Date"`
	CreatedDate      string `csv:"Created Date" json:"createdDate,omitempty,omitzero" bson:"createdDate,omitempty,omitzero"`
	LastActivityDate string `csv:"Last Activity Date" json:"lastActivityDate,omitempty,omitzero" bson:"lastActivityDate,omitempty,omitzero"`
	// ItemCount                    string `csv:"Item Count"`
	StorageUsedBytes string `csv:"Storage Used (Byte)" json:"storageUsedBytes,omitempty,omitzero" bson:"storageUsedBytes,omitempty,omitzero"`
	// IssueWarningQuotaByte        string `csv:"Issue Warning Quota (Byte)"`
	// ProhibitSendQuotaByte        string `csv:"Prohibit Send Quota (Byte)"`
	// ProhibitSendReceiveQuotaByte string `csv:"Prohibit Send/Receive Quota (Byte)"`
	// DeletedItemCount string `csv:"Deleted Item Count"`
	DeletedItemSizeBytes string `csv:"Deleted Item Size (Byte)" json:"deletedItemSizeBytes,omitempty,omitzero" bson:"deletedItemSizeBytes,omitempty,omitzero"`
	// DeletedItemQuotaByte string `csv:"Deleted Item Quota (Byte)"`
	// HasArchive                   string `csv:"Has Archive"`
	RecipientType string `csv:"Recipient Type" json:"recipientType,omitempty,omitzero" bson:"recipientType,omitempty,omitzero"`
	// ReportPeriod                 string `csv:"Report Period"`
	TenantName    string    `json:"tenantName,omitempty,omitzero" bson:"tenantName,omitempty,omitzero"`
	LastAzureSync time.Time `json:"lastAzureSync,omitempty,omitzero" bson:"lastAzureSync,omitempty,omitzero"`
}

//
//

type EXOMailBox struct {
	Alias                string                     `json:"Alias,omitempty,omitzero" bson:"Alias,omitempty,omitzero"`
	DisplayName          string                     `json:"DisplayName,omitempty,omitzero" bson:"DisplayName,omitempty,omitzero"`
	EmailAddresses       []string                   `json:"EmailAddresses,omitempty,omitzero" bson:"EmailAddresses,omitempty,omitzero"`
	RecipientTypeDetails string                     `json:"RecipientTypeDetails,omitempty,omitzero" bson:"RecipientTypeDetails,omitempty,omitzero"`
	UserPrincipalName    string                     `json:"UserPrincipalName,omitempty,omitzero" bson:"UserPrincipalName,omitempty,omitzero"`
	Permissions          []EXOMailboxUserPermission `json:"Permissions,omitempty,omitzero" bson:"Permissions,omitempty,omitzero"`
	TenantName           string                     `json:"TenantName,omitempty,omitzero" bson:"TenantName,omitempty,omitzero"`
}

//
//

type EXOMailBoxRaw struct {
	Alias                string                        `json:"Alias,omitempty,omitzero" bson:"Alias,omitempty,omitzero"`
	DisplayName          string                        `json:"DisplayName,omitempty,omitzero" bson:"DisplayName,omitempty,omitzero"`
	EmailAddresses       []string                      `json:"EmailAddresses,omitempty,omitzero" bson:"EmailAddresses,omitempty,omitzero"`
	RecipientTypeDetails string                        `json:"RecipientTypeDetails,omitempty,omitzero" bson:"RecipientTypeDetails,omitempty,omitzero"`
	UserPrincipalName    string                        `json:"UserPrincipalName,omitempty,omitzero" bson:"UserPrincipalName,omitempty,omitzero"`
	Permissions          []EXOMailboxUserPermissionRaw `json:"Permissions,omitempty,omitzero" bson:"Permissions,omitempty,omitzero"`
	TenantName           string                        `json:"TenantName,omitempty,omitzero" bson:"TenantName,omitempty,omitzero"`
}

//
//

type EXOMailboxUserPermissionRaw struct {
	User       string `json:"User,omitempty,omitzero" bson:"User,omitempty,omitzero"`
	IsOwner    bool   `json:"IsOwner,omitempty,omitzero" bson:"IsOwner,omitempty,omitzero"`
	Permission string `json:"Permission,omitempty,omitzero" bson:"Permission,omitempty,omitzero"`
}

//
//

type EXOMailboxUserPermission struct {
	User         string   `json:"User,omitempty,omitzero" bson:"User,omitempty,omitzero"`
	IsOwner      bool     `json:"IsOwner,omitempty,omitzero" bson:"IsOwner,omitempty,omitzero"`
	AccessRights []string `json:"AccessRights,omitempty,omitzero" bson:"AccessRights,omitempty,omitzero"`
}
