package m365

import (
	"time"
)

type MailboxUsageDetailOldCsv struct {
	UserPrincipalName            string    `csv:"User Principal Name" json:"userPrincipalName,omitempty,omitzero" bson:"userPrincipalName,omitempty,omitzero"`
	CreatedDate                  string    `csv:"Created Date" json:"createdDate,omitempty,omitzero" bson:"createdDate,omitempty,omitzero"`
	LastActivityDate             string    `csv:"Last Activity Date" json:"lastActivityDate,omitempty,omitzero" bson:"lastActivityDate,omitempty,omitzero"`
	ItemCount                    int       `csv:"Item Count" json:"itemCount" bson:"itemCount"`
	StorageUsedBytes             int       `csv:"Storage Used (Byte)" json:"storageUsedBytes" bson:"storageUsedBytes"`
	ProhibitSendQuotaByte        int       `csv:"Prohibit Send Quota (Byte)" json:"prohibitSendQuota" bson:"prohibitSendQuota"`
	ProhibitSendReceiveQuotaByte int       `csv:"Prohibit Send/Receive Quota (Byte)" json:"prohibitSendRcvQuota" bson:"prohibitSendRcvQuota"`
	DeletedItemCount             int       `csv:"Deleted Item Count" json:"deletedItemCount" bson:"deletedItemCount"`
	DeletedItemSizeBytes         int       `csv:"Deleted Item Size (Byte)" json:"deletedItemSizeBytes" bson:"deletedItemSizeBytes"`
	HasArchive                   bool      `csv:"Has Archive" json:"hasArchive" bson:"hasArchive"`
	RecipientType                string    `csv:"Recipient Type" json:"recipientType,omitempty,omitzero" bson:"recipientType,omitempty,omitzero"`
	TenantName                   string    `json:"tenantName,omitempty,omitzero" bson:"tenantName,omitempty,omitzero"`
	LastAzureSync                time.Time `json:"lastAzureSync,omitempty,omitzero" bson:"lastAzureSync,omitempty,omitzero"`
}

type MailboxUsageDetail struct {
	CreatedDate      string `json:"createdDate,omitempty,omitzero" bson:"createdDate,omitempty,omitzero"`
	DeletedDate      string `json:"deletedDate,omitempty,omitzero" bson:"deletedDate,omitempty,omitzero"`
	DeletedItemCount int    `json:"deletedItemCount" bson:"deletedItemCount"`
	// DeletedItemQuota                int    `json:"deletedItemQuota" bson:"deletedItemQuota"`
	DeletedItemSizeInBytes int `json:"deletedItemSizeInBytes" bson:"deletedItemSizeInBytes"`
	// DisplayName                     string `json:"displayName,omitempty,omitzero" bson:"displayName,omitempty,omitzero"`
	HasArchive bool `json:"hasArchive,omitempty,omitzero" bson:"hasArchive,omitempty,omitzero"`
	IsDeleted  bool `json:"isDeleted,omitempty,omitzero" bson:"isDeleted,omitempty,omitzero"`
	// IssueWarningQuotaInBytes        int    `json:"issueWarningQuotaInBytes" bson:"issueWarningQuotaInBytes"`
	ItemCount                       int    `json:"itemCount" bson:"itemCount"`
	LastActivityDate                string `json:"lastActivityDate,omitempty,omitzero" bson:"lastActivityDate,omitempty,omitzero"`
	ProhibitSendQuotaInBytes        int    `json:"prohibitSendQuotaInBytes" bson:"prohibitSendQuotaInBytes"`
	ProhibitSendReceiveQuotaInBytes int    `json:"prohibitSendReceiveQuotaInBytes" bson:"prohibitSendReceiveQuotaInBytes"`
	RecipientType                   string `json:"recipientType,omitempty,omitzero" bson:"recipientType,omitempty,omitzero"`
	// ReportPeriod                    string `json:"reportPeriod,omitempty,omitzero" bson:"reportPeriod,omitempty,omitzero"`
	ReportRefreshDate  string    `json:"reportRefreshDate,omitempty,omitzero" bson:"reportRefreshDate,omitempty,omitzero"`
	StorageUsedInBytes int       `json:"storageUsedInBytes" bson:"storageUsedInBytes"`
	UserPrincipalName  string    `json:"userPrincipalName,omitempty,omitzero" bson:"userPrincipalName,omitempty,omitzero"`
	TenantName         string    `json:"tenantName,omitempty,omitzero" bson:"tenantName,omitempty,omitzero"`
	LastAzureSync      time.Time `json:"lastAzureSync,omitempty,omitzero" bson:"lastAzureSync,omitempty,omitzero"`
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
	PermissionsMap       map[string][]string        `json:"PermissionsMap,omitempty,omitzero" bson:"PermissionsMap,omitempty,omitzero"`
	UsersByPermission    map[string][]string        `json:"UsersByPermission,omitempty,omitzero" bson:"UsersByPermission,omitempty,omitzero"`
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
	GrantSendOnBehalfTo  []string                      `json:"GrantSendOnBehalfTo,omitempty,omitzero" bson:"GrantSendOnBehalfTo,omitempty,omitzero"`
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
