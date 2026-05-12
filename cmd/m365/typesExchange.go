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
