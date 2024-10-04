package lib

import (
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	"golang.org/x/mod/semver"
)

func (tokens *AllTenantTokens) SaveToFile() {

	byteData, err := json.Marshal(tokens)
	CheckFatalError(err)
	if _, err := os.Stat(TokenCacheFile); err != nil {
		os.Create(TokenCacheFile)
	}
	encodedData := b64.StdEncoding.EncodeToString(byteData)
	os.WriteFile(TokenCacheFile, []byte(encodedData), os.ModePerm)
	fmt.Println(encodedData)
}

func (tokens *AllTenantTokens) CheckExpiry() {
	fmt.Println(tokens)
}

func (tokens AllTenantTokens) SelectTenant(tenantName string) (*MultiAuthToken, error) {
	// var tenantToken MultiAuthToken
	// fmt.Println(tenantName)
	var tenantToken *MultiAuthToken

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

func (config AzureConfig) GetDefaultTenant() (*CldConfigTenantAuth, *error) {
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
		return nil, &err
	}
}
