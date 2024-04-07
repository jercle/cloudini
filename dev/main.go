package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/container"
	"github.com/jercle/azg/cmd/azure"
	"github.com/jercle/azg/lib"
)

type StorageAccountBlobListResponse struct {
	ClientRequestID any       `json:"ClientRequestID"`
	ContainerName   string    `json:"ContainerName"`
	ContentType     string    `json:"ContentType"`
	Date            time.Time `json:"Date"`
	Marker          any       `json:"Marker"`
	MaxResults      any       `json:"MaxResults"`
	NextMarker      string    `json:"NextMarker"`
	Prefix          any       `json:"Prefix"`
	RequestID       string    `json:"RequestID"`
	Segment         struct {
		BlobItems []BlobItemResponse `json:"BlobItems"`
	} `json:"Segment"`
	ServiceEndpoint string `json:"ServiceEndpoint"`
	Version         string `json:"Version"`
}

type BlobItemResponse struct {
	BlobTags         any    `json:"BlobTags"`
	Deleted          any    `json:"Deleted"`
	HasVersionsOnly  any    `json:"HasVersionsOnly"`
	IsCurrentVersion any    `json:"IsCurrentVersion"`
	Metadata         any    `json:"Metadata"`
	Name             string `json:"Name"`
	OrMetadata       any    `json:"OrMetadata"`
	Properties       struct {
		AccessTier                  string    `json:"AccessTier"`
		AccessTierChangeTime        any       `json:"AccessTierChangeTime"`
		AccessTierInferred          bool      `json:"AccessTierInferred"`
		ArchiveStatus               any       `json:"ArchiveStatus"`
		BlobSequenceNumber          any       `json:"BlobSequenceNumber"`
		BlobType                    string    `json:"BlobType"`
		CacheControl                string    `json:"CacheControl"`
		ContentDisposition          string    `json:"ContentDisposition"`
		ContentEncoding             string    `json:"ContentEncoding"`
		ContentLanguage             string    `json:"ContentLanguage"`
		ContentLength               float64   `json:"ContentLength"`
		ContentMd5                  string    `json:"ContentMD5"`
		ContentType                 string    `json:"ContentType"`
		CopyCompletionTime          any       `json:"CopyCompletionTime"`
		CopyID                      any       `json:"CopyID"`
		CopyProgress                any       `json:"CopyProgress"`
		CopySource                  any       `json:"CopySource"`
		CopyStatus                  any       `json:"CopyStatus"`
		CopyStatusDescription       any       `json:"CopyStatusDescription"`
		CreationTime                time.Time `json:"CreationTime"`
		CustomerProvidedKeySha256   any       `json:"CustomerProvidedKeySHA256"`
		DeletedTime                 any       `json:"DeletedTime"`
		DestinationSnapshot         any       `json:"DestinationSnapshot"`
		ETag                        string    `json:"ETag"`
		EncryptionScope             any       `json:"EncryptionScope"`
		ExpiresOn                   any       `json:"ExpiresOn"`
		ImmutabilityPolicyExpiresOn any       `json:"ImmutabilityPolicyExpiresOn"`
		ImmutabilityPolicyMode      any       `json:"ImmutabilityPolicyMode"`
		IncrementalCopy             any       `json:"IncrementalCopy"`
		IsSealed                    any       `json:"IsSealed"`
		LastAccessedOn              any       `json:"LastAccessedOn"`
		LastModified                time.Time `json:"LastModified"`
		LeaseDuration               any       `json:"LeaseDuration"`
		LeaseState                  string    `json:"LeaseState"`
		LeaseStatus                 string    `json:"LeaseStatus"`
		LegalHold                   any       `json:"LegalHold"`
		RehydratePriority           any       `json:"RehydratePriority"`
		RemainingRetentionDays      any       `json:"RemainingRetentionDays"`
		ServerEncrypted             bool      `json:"ServerEncrypted"`
		TagCount                    any       `json:"TagCount"`
	} `json:"Properties"`
	Snapshot  any `json:"Snapshot"`
	VersionID any `json:"VersionID"`
}

type BlobItem struct {
	Name             string `json:"Name"`
	BlobTags         any    `json:"BlobTags"`
	Deleted          any    `json:"Deleted"`
	IsCurrentVersion any    `json:"IsCurrentVersion"`
	Metadata         any    `json:"Metadata"`
	OrMetadata       any    `json:"OrMetadata"`
	Properties       struct {
		AccessTier             string    `json:"AccessTier"`
		AccessTierChangeTime   any       `json:"AccessTierChangeTime"`
		AccessTierInferred     bool      `json:"AccessTierInferred"`
		BlobType               string    `json:"BlobType"`
		ContentMd5             string    `json:"ContentMD5"`
		ContentType            string    `json:"ContentType"`
		CreationTime           time.Time `json:"CreationTime"`
		DeletedTime            any       `json:"DeletedTime"`
		LastAccessedOn         any       `json:"LastAccessedOn"`
		LastModified           time.Time `json:"LastModified"`
		RemainingRetentionDays any       `json:"RemainingRetentionDays"`
		ServerEncrypted        bool      `json:"ServerEncrypted"`
	} `json:"Properties"`
	Snapshot  any `json:"Snapshot"`
	VersionID any `json:"VersionID"`
}

type StorageAccountRequestOptions struct {
	StorageAccountName   string
	ContainerName        string
	ConfiguredTenantName string
	GetWriteToken        bool
}

func main() {

	startTime := time.Now()
	config := lib.GetCldConfig(nil)

	_ = config
	// tokens, err := azure.GetAllTenantSPTokens(lib.MultiAuthTokenRequestOptions{})
	// lib.CheckFatalError(err)
	token, err := azure.GetTenantSPToken("Stack Cats", lib.MultiAuthTokenRequestOptions{
		Scope: "storage",
	})
	lib.CheckFatalError(err)
	// _ = tokens
	_ = token

	options := StorageAccountRequestOptions{
		StorageAccountName:   "stkcatmgtstorage",
		ContainerName:        "cost-exports",
		ConfiguredTenantName: "Stack Cats",
	}

	blobList := ListStorageContainerBlobs(options)

	jsonBytes, _ := json.MarshalIndent(blobList, "", "  ")

	// urlString := config.Azure.MultiTenantAuth.Tenants["Stack Cats"].CostExportsLocation + "?restype=container&comp=list&api-version=2023-11-03"

	// response := web.SimpleGetRequestWithToken(urlString, token.TokenData.Token)

	fmt.Println(string(jsonBytes))
	elapsed := time.Since(startTime)
	_ = elapsed
	// fmt.Println(elapsed)
}

func ListStorageContainerBlobs(options StorageAccountRequestOptions) []BlobItem {
	var (
		cred     *azidentity.ClientSecretCredential
		err      error
		ctx      = context.Background()
		BlobList []BlobItem
	)

	config := lib.GetCldConfig(nil)

	tenant := config.Azure.MultiTenantAuth.Tenants[options.ConfiguredTenantName]
	if options.GetWriteToken {
		cred, err = azidentity.NewClientSecretCredential(tenant.TenantID, tenant.Writer.ClientID, tenant.Writer.ClientSecret, nil)
		lib.CheckFatalError(err)
	} else {
		cred, err = azidentity.NewClientSecretCredential(tenant.TenantID, tenant.Reader.ClientID, tenant.Reader.ClientSecret, nil)
		lib.CheckFatalError(err)
	}

	serviceURL := "https://" + options.StorageAccountName + ".blob.core.windows.net"

	client, err := azblob.NewClient(serviceURL, cred, nil)

	pager := client.NewListBlobsFlatPager(options.ContainerName, &azblob.ListBlobsFlatOptions{
		Include: container.ListBlobsInclude{Deleted: false, Versions: false},
	})

	for pager.More() {
		resp, err := pager.NextPage(ctx)
		lib.CheckFatalError(err)
		// jsonBytes, _ := json.MarshalIndent(resp, "", "  ")
		// fmt.Println(string(jsonBytes))
		for _, blob := range resp.Segment.BlobItems {
			var blobItem BlobItem
			jsonBytes, _ := json.Marshal(blob)
			json.Unmarshal(jsonBytes, &blobItem)
			BlobList = append(BlobList, blobItem)
			// fmt.Println(*_blob.Name)
			// BlobList = append(BlobList, *_blob
			// )
		}
	}

	return BlobList
}
