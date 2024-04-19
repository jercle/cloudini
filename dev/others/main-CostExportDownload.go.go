package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/container"
	"github.com/jercle/cloudini/cmd/azure"
	"github.com/jercle/cloudini/lib"
)

type BlobItem struct {
	Name               string `json:"Name"`
	ContainerName      string `json:"containerName"`
	TenantName         string `json:"tenantName"`
	StorageAccountName string `json:"storageAccountName"`
	BlobTags           any    `json:"BlobTags"`
	Deleted            any    `json:"Deleted"`
	IsCurrentVersion   any    `json:"IsCurrentVersion"`
	Metadata           any    `json:"Metadata"`
	OrMetadata         any    `json:"OrMetadata"`
	Properties         struct {
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

type BlobList []BlobItem

type StorageAccountRequestOptions struct {
	StorageAccountName   string
	ContainerName        string
	ConfiguredTenantName string
	GetWriteToken        bool
	BlobName             string
	DownloadFileName     string
}

func main() {

	startTime := time.Now()

	// jsonBytes, _ := json.MarshalIndent(blobList, "", "  ")
	// fmt.Println(string(jsonBytes))

	config := lib.GetCldConfig(nil)

	tenant := config.Azure.MultiTenantAuth.Tenants["REDDTQ"]
	split := strings.Split(tenant.CostExportsLocation, "/")

	options := StorageAccountRequestOptions{
		StorageAccountName:   strings.Split(split[2], ".")[0],
		ContainerName:        split[3],
		ConfiguredTenantName: tenant.TenantName,
	}
	blobs := ListStorageContainerBlobs(options)
	fmt.Println(len(blobs))
	// DownloadAllConfiguredTenantLastMonthCostExports(DownloadAllConfiguredTenantLastMonthCostExportsOptions{
	// 	BlobPrefix:  "monthly-cost-exports/202403",
	// 	OutfilePath: "cost-exports/monthly-cost-exports",
	// })
	elapsed := time.Since(startTime)
	_ = elapsed
}

type DownloadAllConfiguredTenantLastMonthCostExportsOptions struct {

	// Prefix for blob files
	// Example: "monthly-cost-exports/202404"
	// Example: "daily-month-to-date-exports/202404"
	BlobPrefix string

	// Path and filename prefix for file download
	// Tenant Name will be appended to filename
	// Example: "outputs/cost-exports" would become "outputs/cost-exports_TENANTNAME"
	OutfilePath string
}

func DownloadAllConfiguredTenantLastMonthCostExports(opts DownloadAllConfiguredTenantLastMonthCostExportsOptions) {
	var wg sync.WaitGroup
	config := lib.GetCldConfig(nil)

	for _, tenant := range config.Azure.MultiTenantAuth.Tenants {
		if tenant.CostExportsLocation == "" {
			continue
		}
		wg.Add(1)
		go func() {
			defer wg.Done()
			split := strings.Split(tenant.CostExportsLocation, "/")
			if len(split) != 4 {
				lib.CheckFatalError(fmt.Errorf("CostExportsLocation does not seem to be correctly formatted"))
			}

			options := StorageAccountRequestOptions{
				StorageAccountName:   strings.Split(split[2], ".")[0],
				ContainerName:        split[3],
				ConfiguredTenantName: tenant.TenantName,
			}

			blobList := ListStorageContainerBlobs(options)
			blobList.Filter(BlobListFilterOptions{FilterPrefix: opts.BlobPrefix})
			blobList.SortByCreateDate("descending")

			cred := azure.GetTenantAzCred(tenant.TenantName, false)
			fileName := opts.OutfilePath + "_" + tenant.TenantName + ".csv"
			fmt.Println(tenant.TenantName, len(blobList))
			if len(blobList) > 0 {
				blobList[0].Download(cred, fileName)
				fmt.Println("Downloaded", tenant.TenantName)
			} else {
				fmt.Println("No blobs found for ", tenant.TenantName)
			}

		}()
	}
	wg.Wait()
}

func ListStorageContainerBlobs(options StorageAccountRequestOptions) BlobList {
	var (
		cred     *azidentity.ClientSecretCredential
		err      error
		ctx      = context.Background()
		BlobList BlobList
	)

	config := lib.GetCldConfig(nil)
	tenant := config.Azure.MultiTenantAuth.Tenants[options.ConfiguredTenantName]
	lib.PrintSrcLoc(tenant.TenantName)

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
		for _, blob := range resp.Segment.BlobItems {
			var blobItem BlobItem
			jsonBytes, _ := json.MarshalIndent(blob, "", "  ")
			blobItem.TenantName = tenant.TenantName
			blobItem.StorageAccountName = options.StorageAccountName
			blobItem.ContainerName = options.ContainerName
			json.Unmarshal(jsonBytes, &blobItem)
			// fmt.Println(blobItem)
			BlobList = append(BlobList, blobItem)
		}
	}

	return BlobList
}

func (blob *BlobItem) Download(cred *azidentity.ClientSecretCredential, fileName string) {
	var (
		ctx = context.Background()
	)
	serviceURL := "https://" + blob.StorageAccountName + ".blob.core.windows.net"
	client, err := azblob.NewClient(serviceURL, cred, nil)
	lib.CheckFatalError(err)
	file, err := os.Create(fileName)
	lib.CheckFatalError(err)
	defer file.Close()
	_, err = client.DownloadFile(ctx, blob.ContainerName, blob.Name, file, nil)
}

func (bl *BlobList) Filter(opts BlobListFilterOptions) {
	var filteredBlobs []BlobItem
	for _, blob := range *bl {
		if strings.HasPrefix(blob.Name, opts.FilterPrefix) {
			filteredBlobs = append(filteredBlobs, blob)
		}
	}
	*bl = filteredBlobs
}

func (blobList *BlobList) SortByCreateDate(sortOrder string) {
	bl := *blobList
	if sortOrder == "ascending" {
		sort.Slice(bl, func(i, j int) bool {
			return bl[i].Properties.CreationTime.Before(bl[j].Properties.CreationTime)
		})
	} else if sortOrder == "descending" {
		sort.Slice(bl, func(i, j int) bool {
			return bl[j].Properties.CreationTime.Before(bl[i].Properties.CreationTime)
		})
	} else {
		fmt.Println("Sort order must be 'ascending' or 'descending'")
	}

	*blobList = bl
}

type BlobListFilterOptions struct {
	FilterPrefix   string
	FilterSuffix   string
	FilterContains string
}
