package azure

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/container"
	"github.com/jercle/cloudini/lib"
)

func StorageBlobHttpGet(urlString string, mat lib.AzureMultiAuthToken) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, urlString, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/xml")
	req.Header.Add("Authorization", "Bearer "+mat.TokenData.Token)
	req.Header.Add("x-ms-version", "2023-11-03")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	responseBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return responseBody, nil
}

//
//

func ListStorageContainerBlobs(options lib.StorageAccountRequestOptions, cldConfigOpts *lib.CldConfigOptions) lib.BlobList {
	var (
		cred     *azidentity.ClientSecretCredential
		err      error
		ctx      = context.Background()
		BlobList lib.BlobList
	)

	config := lib.GetCldConfig(cldConfigOpts)
	tenant := config.Azure.MultiTenantAuth.Tenants[options.ConfiguredTenantName]
	// lib.PrintSrcLoc(tenant.TenantName)

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
			var blobItem lib.BlobItem
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

//
//

func DownloadAllBlobsInContainer(options lib.StorageAccountRequestOptions) (numFilesDownloaded int) {
	blobList := ListStorageContainerBlobs(options, nil)

	cred, err := GetTenantAzCred(options.ConfiguredTenantName, false, nil)
	lib.CheckFatalError(err)

	for _, blob := range blobList {
		filePath := options.DownloadPath + "/" + blob.Name
		_, notExist := os.Stat(filePath)

		if notExist != nil || options.OverwriteExisting {
			blob.Download(cred, options.DownloadPath+"/"+blob.Name)
			numFilesDownloaded++
		}
	}

	if options.ShowDownloadedCount {
		fmt.Println("Downloaded " + strconv.Itoa(numFilesDownloaded) + " files")
	}

	return numFilesDownloaded
}
