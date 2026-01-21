package azure

import "time"

type GetManagedDevicesResponse struct {
	Context string                 `json:"@odata.context,omitempty,omitzero" bson:"@odata.context,omitempty,omitzero"`
	Count   float64                `json:"@odata.count,omitempty,omitzero" bson:"@odata.count,omitempty,omitzero"`
	Value   []ManagedDeviceMinimal `json:"value,omitempty,omitzero" bson:"value,omitempty,omitzero"`
}

//
//

type ManagedDeviceMinimal struct {
	AzureAdDeviceID   string    `json:"azureADDeviceId,omitempty,omitzero" bson:"azureADDeviceId,omitempty,omitzero"`
	DeviceName        string    `json:"deviceName,omitempty,omitzero" bson:"deviceName,omitempty,omitzero"`
	ID                string    `json:"id,omitempty,omitzero" bson:"_id,omitempty,omitzero"`
	LastSyncDateTime  time.Time `json:"lastSyncDateTime,omitempty,omitzero" bson:"lastSyncDateTime,omitempty,omitzero"`
	SerialNumber      string    `json:"serialNumber,omitempty,omitzero" bson:"serialNumber,omitempty,omitzero"`
	UserPrincipalName string    `json:"userPrincipalName,omitempty,omitzero" bson:"userPrincipalName,omitempty,omitzero"`
}
