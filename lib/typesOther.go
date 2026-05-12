package lib

import (
	"time"
)

type PackerLogBuildData struct {
	AzDoBuildName          string                                     `json:"azDoBuildName,omitempty" bson:"azDoBuildName,omitempty"`
	AzDoCompleteTime       time.Time                                  `json:"azDoCompleteTime,omitempty" bson:"azDoCompleteTime,omitempty"`
	AzDoDuration           time.Duration                              `json:"azDoDuration,omitempty" bson:"azDoDuration,omitempty"`
	AzDoHost               string                                     `json:"azDoHost,omitempty" bson:"azDoHost,omitempty"`
	AzDoLogFile            string                                     `json:"azDoLogFile,omitempty" bson:"azDoLogFile,omitempty"`
	AzDoStartTime          time.Time                                  `json:"azDoStartTime,omitempty" bson:"azDoStartTime,omitempty"`
	BuildBaseImageVersion  *AzureResourceStorageProfileImageReference `json:"buildBaseImageVersion,omitempty" bson:"buildBaseImageVersion,omitempty"`
	BuildImageCompleteTime time.Time                                  `json:"buildImageCompleteTime,omitempty" bson:"buildImageCompleteTime,omitempty"`
	BuildImageDuration     time.Duration                              `json:"buildImageDuration,omitempty" bson:"buildImageDuration,omitempty"`
	BuildImageEnv          string                                     `json:"buildImageEnv,omitempty" bson:"buildImageEnv,omitempty"`
	BuildImageStartTime    time.Time                                  `json:"buildImageStartTime,omitempty" bson:"buildImageStartTime,omitempty"`
	OutputImgDef           string                                     `json:"outputImgDef,omitempty" bson:"outputImgDef,omitempty"`
	OutputImgGalleryName   string                                     `json:"outputImgGalleryName,omitempty" bson:"outputImgGalleryName,omitempty"`
	OutputImgId            string                                     `json:"outputImgId,omitempty" bson:"outputImgId,omitempty"`
	OutputImgResGrp        string                                     `json:"outputImgResGrp,omitempty" bson:"outputImgResGrp,omitempty"`
	OutputImgVersion       string                                     `json:"outputImgVersion,omitempty" bson:"outputImgVersion,omitempty"`
}

//
//

type PackerPublishImageResponse struct {
	ID         string `json:"id,omitempty" bson:"id,omitempty"`
	Location   string `json:"location,omitempty" bson:"location,omitempty"`
	Name       string `json:"name,omitempty" bson:"name,omitempty"`
	Properties struct {
		ProvisioningState string `json:"provisioningState,omitempty" bson:"provisioningState,omitempty"`
		PublishingProfile struct {
			ExcludeFromLatest  bool      `json:"excludeFromLatest,omitempty" bson:"excludeFromLatest,omitempty"`
			PublishedDate      time.Time `json:"publishedDate,omitempty" bson:"publishedDate,omitempty"`
			ReplicaCount       float64   `json:"replicaCount,omitempty" bson:"replicaCount,omitempty"`
			ReplicationMode    string    `json:"replicationMode,omitempty" bson:"replicationMode,omitempty"`
			StorageAccountType string    `json:"storageAccountType,omitempty" bson:"storageAccountType,omitempty"`
			TargetRegions      []struct {
				Name                 string  `json:"name,omitempty" bson:"name,omitempty"`
				RegionalReplicaCount float64 `json:"regionalReplicaCount,omitempty" bson:"regionalReplicaCount,omitempty"`
				StorageAccountType   string  `json:"storageAccountType,omitempty" bson:"storageAccountType,omitempty"`
			} `json:"targetRegions,omitempty" bson:"targetRegions,omitempty"`
		} `json:"publishingProfile,omitempty" bson:"publishingProfile,omitempty"`
		SafetyProfile struct {
			AllowDeletionOfReplicatedLocations bool `json:"allowDeletionOfReplicatedLocations,omitempty" bson:"allowDeletionOfReplicatedLocations,omitempty"`
			ReportedForPolicyViolation         bool `json:"reportedForPolicyViolation,omitempty" bson:"reportedForPolicyViolation,omitempty"`
		} `json:"safetyProfile,omitempty" bson:"safetyProfile,omitempty"`
		StorageProfile struct {
			Source struct {
				VirtualMachineID string `json:"virtualMachineId,omitempty" bson:"virtualMachineId,omitempty"`
			} `json:"source,omitempty" bson:"source,omitempty"`
		} `json:"storageProfile,omitempty" bson:"storageProfile,omitempty"`
	} `json:"properties,omitempty" bson:"properties,omitempty"`
	Tags map[string]string `json:"tags" bson:"tags"`
	Type string            `json:"type,omitempty" bson:"type,omitempty"`
}

//
//
