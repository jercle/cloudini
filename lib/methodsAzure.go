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
	latestVersion := GalleryImageVersion{}
	var versionList []string
	vAppended := false
	latestVersionNum := ""

	for _, version := range *list {
		currentVersion := ""

		if !strings.HasPrefix(version.Name, "v") {
			currentVersion = "v" + version.Name
			vAppended = true
		} else {
			currentVersion = version.Name
		}

		versionList = append(versionList, currentVersion)

	}

	semver.Sort(versionList)
	latest := versionList[len(versionList)-1]

	if vAppended {
		latestVersionNum = strings.TrimPrefix(latest, "v")
	} else {
		latestVersionNum = latest
	}

	for _, version := range *list {
		if version.Name == latestVersionNum {
			latestVersion = version
		}
	}

	return latestVersion, latestVersion.Name
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
