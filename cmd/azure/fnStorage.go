package azure

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/container"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/sas"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/service"
	"github.com/jercle/cloudini/lib"
)

func StorageBlobHttpGet(options StorageAccountUploadBlobOptions, mat lib.AzureMultiAuthToken) ([]byte, error) {
	urlString := "https://" + options.StorageAccountName + ".blob.core.windows.net/" + options.ContainerName + "/"
	if options.BlobPrefix != "" {
		urlString += options.BlobPrefix + "/"
	}
	urlString += options.BlobFileName

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

func StorageBlobHttpGetFromSAS(blobSAS string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, blobSAS, nil)
	if err != nil {
		return nil, err
	}

	// req.Header.Add("Content-Type", "application/xml")
	// req.Header.Add("Authorization", "Bearer "+mat.TokenData.Token)
	// req.Header.Add("x-ms-version", "2023-11-03")

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

//
//

func UploadBlobFromString(fileData string, options StorageAccountUploadBlobOptions) (azblob.UploadFileResponse, error) {

	var (
		cred *azidentity.ClientSecretCredential
		err  error
	)

	config := lib.GetCldConfig(nil)
	tenant := config.Azure.MultiTenantAuth.Tenants[options.ConfiguredTenantName]

	cred, err = azidentity.NewClientSecretCredential(tenant.TenantID, tenant.Writer.ClientID, tenant.Writer.ClientSecret, nil)
	lib.CheckFatalError(err)

	serviceURL := "https://" + options.StorageAccountName + ".blob.core.windows.net"
	client, err := azblob.NewClient(serviceURL, cred, nil)
	lib.CheckFatalError(err)

	// Upload the file to the specified container with the specified blob name
	blobNameAndPrefix := ""
	if options.BlobPrefix != "" {
		blobNameAndPrefix = options.BlobPrefix + "/" + options.BlobFileName
	} else {
		blobNameAndPrefix = options.BlobFileName
	}

	response, err := client.UploadBuffer(context.TODO(), options.ContainerName, blobNameAndPrefix, []byte(fileData), nil)
	lib.CheckFatalError(err)
	return response, err
}

//
//

func UploadBlobFromFile(fileName string, options StorageAccountUploadBlobOptions) (azblob.UploadFileResponse, error) {

	var (
		cred *azidentity.ClientSecretCredential
		err  error
	)

	// fmt.Println(options.BlobPrefix + options.BlobFileName)
	// os.Exit(0)

	config := lib.GetCldConfig(nil)
	tenant := config.Azure.MultiTenantAuth.Tenants[options.ConfiguredTenantName]

	cred, err = azidentity.NewClientSecretCredential(tenant.TenantID, tenant.Writer.ClientID, tenant.Writer.ClientSecret, nil)
	lib.CheckFatalError(err)

	serviceURL := "https://" + options.StorageAccountName + ".blob.core.windows.net"
	client, err := azblob.NewClient(serviceURL, cred, nil)
	lib.CheckFatalError(err)

	file, err := os.OpenFile(fileName, os.O_RDONLY, 0)
	lib.CheckFatalError(err)

	defer file.Close()

	// Upload the file to the specified container with the specified blob name
	blobNameAndPrefix := ""
	if options.BlobPrefix != "" {
		blobNameAndPrefix = options.BlobPrefix + "/" + options.BlobFileName
	} else {
		blobNameAndPrefix = options.BlobFileName
	}

	response, err := client.UploadFile(context.TODO(), options.ContainerName, blobNameAndPrefix, file, nil)
	lib.CheckFatalError(err)
	return response, err
}

//
//

func BulkUploadBlob(basePath string, options StorageAccountUploadBlobOptions) (responses StorageAccountBulkUploadBlobResponse) {
	var (
		wg    sync.WaitGroup
		mutex sync.Mutex
	)

	paths := lib.GetFullFilePaths(basePath)

	fmt.Println("Uploading", strconv.Itoa(len(paths)), "files to blob storage")

	for _, file := range paths {
		// fmt.Println(file)
		wg.Add(1)
		go func() {
			defer wg.Done()
			fileName := filepath.Base(file)
			// fmt.Println(fileName)
			opts := options
			opts.BlobFileName = fileName
			_, err := UploadBlobFromFile(file, opts)
			mutex.Lock()
			if err != nil {
				responses.Errored = append(responses.Errored, file)
			} else {
				responses.Uploaded = append(responses.Uploaded, file)
			}
			mutex.Unlock()
			// lib.CheckFatalError(err)
			// lib.JsonMarshalAndPrint(rsp)
			// _ = rsp
			// os.Exit(0)
		}()
	}
	wg.Wait()

	// lib.JsonMarshalAndPrint(responses)
	return
}

//
//

func GetBlobSAS(storageAccountName string, containerName string, tenantId string, clientId string, clientSecret string) string {
	cred, err := azidentity.NewClientSecretCredential(tenantId, clientId, clientSecret, nil)
	lib.CheckFatalError(err)

	svcClient, err := service.NewClient(
		fmt.Sprintf(fmt.Sprintf("https://%s.blob.core.windows.net/", storageAccountName)),
		cred,
		&service.ClientOptions{},
	)
	lib.CheckFatalError(err)

	// Set current and past time and create key
	now := time.Now().UTC().Add(-10 * time.Second)
	expiry := now.Add(48 * time.Hour)
	info := service.KeyInfo{
		Start:  to.Ptr(now.UTC().Format(sas.TimeFormat)),
		Expiry: to.Ptr(expiry.UTC().Format(sas.TimeFormat)),
	}

	udc, err := svcClient.GetUserDelegationCredential(context.Background(), info, nil)
	lib.CheckFatalError(err)

	sasQueryParams, err := sas.BlobSignatureValues{
		Protocol:      sas.ProtocolHTTPS,
		StartTime:     time.Now().UTC().Add(time.Second * -10),
		ExpiryTime:    time.Now().UTC().Add(15 * time.Minute),
		Permissions:   to.Ptr(sas.BlobPermissions{Add: true, Write: true, Create: true, Read: true}).String(),
		ContainerName: containerName,
	}.SignWithUserDelegation(udc)
	lib.CheckFatalError(err)

	// sasURL := fmt.Sprintf("https://%s.blob.core.windows.net/?%s", storageAccountName, sasQueryParams.Encode())

	// sasQueryParams :=

	return sasQueryParams.Encode()
}

//
//

type StorageAccountUploadBlobOptions struct {
	StorageAccountName   string
	ContainerName        string
	ConfiguredTenantName string
	BlobFileName         string
	BlobPrefix           string
}

type StorageAccountBulkUploadBlobResponse struct {
	Errored  []string `json:"errored"`
	Uploaded []string `json:"uploaded"`
}
