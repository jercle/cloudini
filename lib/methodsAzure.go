package lib

import (
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"os"

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

func (versions *SIGImageVersionList) Latest() SIGImageVersion {
	var latestImage *SIGImageVersion

	for _, version := range *versions {
		if latestImage == nil {
			latestImage = &version
		} else {
			if semver.Compare("v"+version.Name, "v"+latestImage.Name) == 1 {
				latestImage = &version
			}
		}
	}

	return *latestImage
}
