package lib

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"

	"golang.org/x/mod/semver"
)

// func (tokens *AllTenantTokens) SaveToFile() {

// 	byteData, err := json.Marshal(tokens)
// 	CheckFatalError(err)
// 	if _, err := os.Stat(TokenCacheFile); err != nil {
// 		os.Create(TokenCacheFile)
// 	}
// 	encodedData := b64.StdEncoding.EncodeToString(byteData)
// 	os.WriteFile(TokenCacheFile, []byte(encodedData), os.ModePerm)
// 	fmt.Println(encodedData)
// }

// func (tokens *AllTenantTokens) CheckExpiry() {
// 	fmt.Println(tokens)
// }

func (tokens AllTenantTokens) SelectTenant(tenantName string) (*AzureMultiAuthToken, error) {
	// var tenantToken AzureMultiAuthToken
	// fmt.Println(tenantName)
	var tenantToken *AzureMultiAuthToken

	for _, token := range tokens {
		if token.TenantName == tenantName {
			tenantToken = &token
			break
		}
	}

	if tenantToken != nil {
		return tenantToken, nil
	} else {
		return nil, fmt.Errorf("Token not found for supplied tenant name")
	}
}

func (subs *SubsReqResBody) UpdateTenantName(tenantName string) {
	var localSubs SubsReqResBody
	localSubs.Count = subs.Count
	for _, sub := range subs.Value {
		sub.TenantName = tenantName
		localSubs.Value = append(localSubs.Value, sub)
	}
	*subs = localSubs
}

func (list *GalleryImageVersionList) Latest() (GalleryImageVersion, string) {
	l := *list

	if !list.Sorted {
		l = l.Sort()
	}
	latest := l.Versions[len(l.Versions)-1]

	return latest, latest.Name
}

func (list *GalleryImageVersionList) Sort() GalleryImageVersionList {
	l := *list
	var (
		versionList       []string
		processedVersions []GalleryImageVersion
		sortedVersions    []GalleryImageVersion
	)

	for _, version := range l.Versions {
		currentVersion := ""
		processedVersion := version
		if !strings.HasPrefix(version.Name, "v") {
			currentVersion = "v" + version.Name
			processedVersion.SuffixAdded = true
		} else {
			currentVersion = version.Name
			processedVersion.SuffixAdded = false
		}

		processedVersions = append(processedVersions, processedVersion)
		versionList = append(versionList, currentVersion)

	}

	semver.Sort(versionList)

	for _, version := range versionList {
		for _, usv := range processedVersions {
			currVers := usv
			compareVersion := version
			if currVers.SuffixAdded {
				compareVersion = strings.TrimPrefix(compareVersion, "v")
			}
			if currVers.Name == compareVersion {
				sortedVersions = append(sortedVersions, currVers)
			}
		}
	}

	l.Sorted = true
	l.Versions = sortedVersions

	return l
}

func (list *GalleryImageVersionList) ListVersions() []string {
	var sortedVersionNumbers []string

	l := *list
	if !l.Sorted {
		l = l.Sort()
	}

	for _, v := range l.Versions {
		sortedVersionNumbers = append(sortedVersionNumbers, v.Name)
	}

	return sortedVersionNumbers
}

func (list *GalleryImageVersionList) CheckVersionExists(compareVersion string) bool {
	l := *list
	if !l.Sorted {
		l = l.Sort()
	}
	versionExists := false
Check:
	for _, version := range l.Versions {
		if version.Name == compareVersion {
			versionExists = true
			break Check
		}
	}
	return versionExists
}
func (imgVersion *GalleryImageVersion) IncrementPatchVersion() string {
	version := imgVersion.Name
	var v string

	if !strings.HasPrefix(version, "v") {
		v = "v" + version
	} else {
		v = version
	}
	isValid := semver.IsValid(v)
	if !isValid {
		CheckFatalError(fmt.Errorf("Provide valid semantic version"))
	}

	vnums := strings.Split(version, ".")
	patchVersion, err := strconv.Atoi(vnums[2])
	CheckFatalError(err)

	patchVersion++

	vnums[2] = strconv.Itoa(patchVersion)

	return strings.Join(vnums, ".")
}

func (config AzureConfig) GetDefaultTenant() (*CldConfigTenantAuth, error) {
	var (
		tenant *CldConfigTenantAuth
		err    error
	)
	for _, tConf := range config.MultiTenantAuth.Tenants {
		if tConf.Default {
			tenant = &tConf
		}
	}
	if tenant != nil {
		return tenant, nil
	} else {
		err = fmt.Errorf("No default Azure tenant configured")
		return nil, err
	}
}

// func (tenants *TransformedCostItemsByTenantMap) AddPreTaxCost(tci TransformedCostItem) {
// 	t := *tenants

// 	thing := t[tci.Datafile]
// 	thing.PreTaxCost += tci.PreTaxCost

//		fmt.Println(t)
//	}
func (tenants *TransformedCostItemsByTenant) AddPreTaxCost(tci TransformedCostItem) {
	cfg := GetCldConfig(nil)
	t := *tenants

	tenantName := ""
	custTntName := MapAzureSubscriptionToCustomTenantName(tci.SubscriptionId, *cfg.Azure)
	if custTntName != "" {
		tenantName = custTntName
	} else {
		tenantName = tci.Datafile
	}

	tenant := t[tenantName]
	tenant.PreTaxCost += tci.PreTaxCost
	t[tenantName] = tenant
	*tenants = t
}

func (t *TransformedCostItemsByTenant) AppendTenantData(tci TransformedCostItem) {
	tenants := *t
	// jsonStr, _ := json.MarshalIndent(tenants, "", "  ")
	// fmt.Println(string(jsonStr))
	// os.Exit(0)
	entry, exists := tenants[tci.Tenant]
	if !exists {
		tenant := TransformedTenantData{}
		tenant.PreTaxCost += tci.PreTaxCost
		tenant.ResGroups = append(tenant.ResGroups, tci)
		tenants[tci.Tenant] = tenant
	} else {
		tenant := entry
		tenant.PreTaxCost += tci.PreTaxCost
		tenant.ResGroups = append(tenant.ResGroups, tci)
		tenants[tci.Tenant] = tenant
	}
	*t = tenants
}

func (e *FieldMismatch) Error() string {
	return "CSV line fields mismatch. Expected " + strconv.Itoa(e.Expected) + " found " + strconv.Itoa(e.Found)
}

func (e *UnsupportedType) Error() string {
	return "Unsupported type: " + e.Type
}

func (t *TransformedCostItemsByTenant) SumCosts() {
	tenants := *t

	for key, val := range tenants {
		_ = key
		_ = val
		// fmt.Println(val)

		for _, res := range val.ResGroups {
			jsonStr, _ := json.MarshalIndent(res, "", "  ")
			fmt.Println(string(jsonStr))
			os.Exit(0)
		}
	}
}

func (blob *BlobItem) Download(cred *azidentity.ClientSecretCredential, fileName string) {
	var (
		ctx = context.Background()
	)
	serviceURL := "https://" + blob.StorageAccountName + ".blob.core.windows.net"
	client, err := azblob.NewClient(serviceURL, cred, nil)
	CheckFatalError(err)

	path := filepath.Dir(fileName)
	if _, err := os.Stat(path); err != nil {
		os.MkdirAll(path, os.ModePerm)
	}

	file, err := os.Create(fileName)
	CheckFatalError(err)
	defer file.Close()
	_, err = client.DownloadFile(ctx, blob.ContainerName, blob.Name, file, nil)
}

// TODO: Method to get data without saving file
// func (blob *BlobItem) FetchData(cred *azidentity.ClientSecretCredential, fileName string) {
// 	var (
// 		ctx = context.Background()
// 	)
// 	serviceURL := "https://" + blob.StorageAccountName + ".blob.core.windows.net"
// 	client, err := azblob.NewClient(serviceURL, cred, nil)
// 	CheckFatalError(err)
// 	// file, err := os.Create(fileName)
// 	// CheckFatalError(err)
// 	// defer file.Close()
// 	data := []byte{}
// 	client.DownloadBuffer(ctx, blob.ContainerName, blob.Name, data)
// 	_, err = client.DownloadFile(ctx, blob.ContainerName, blob.Name, file, nil)
// }

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
