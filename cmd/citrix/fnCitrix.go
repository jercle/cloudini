package citrix

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/jercle/cloudini/lib"
)

func ListMachineCatalogs(creds lib.CitrixCloudAccountConfig, tokenData lib.CitrixTokenData) (machineCatalogs MachineCatalogs) {
	urlString := "https://api.cloud.com/cvad/manage/MachineCatalogs"
	res, err := HttpGet(urlString, creds.CustomerId, creds.SiteId, tokenData)
	// fmt.Println(string(res))
	// os.WriteFile("ctx-res.json", res, 0644)
	lib.CheckFatalError(err)

	var mcatRes GetMachineCatalogsResponse
	json.Unmarshal(res, &mcatRes)
	// lib.JsonMarshalAndPrint(mcatRes)
	// os.Exit(0)

	for _, mcat := range mcatRes.Items {
		curr := mcat
		curr.LastCitrixSync = time.Now()
		machineCatalogs = append(machineCatalogs, curr)
	}

	return
}

func GetMachineCatalogDeliveryGroupAssociations(customerId string, siteId string, machineCatalogId string, tokenData lib.CitrixTokenData) []MchnCatDelGrpAssociation {
	urlString := "https://api.cloud.com/cvad/manage/" +
		"/MachineCatalogs/" +
		machineCatalogId +
		"/DeliveryGroupAssociations"

	res, err := HttpGet(urlString, customerId, siteId, tokenData)
	lib.CheckFatalError(err)

	var MacCatDelGrpAssociations GetMachineCatalogDeliveryGroupAssociationsResponse
	json.Unmarshal(res, &MacCatDelGrpAssociations)

	return MacCatDelGrpAssociations.Items
}

func CacheToken(tokenData lib.CitrixTokenData, cldOpts *lib.CldConfigOptions) {
	_, _, cachePath := lib.InitConfig(cldOpts)
	fmt.Println(cachePath)
	jsonStr, err := json.MarshalIndent(tokenData, "", "  ")
	lib.CheckFatalError(err)

	os.WriteFile(cachePath+"/ctxtok.json", jsonStr, os.ModePerm)
	// fmt.Println(tokenData)
	// fmt.Println(string(jsonStr))

}

func GetUserInfo(customerId string, tokenData lib.CitrixTokenData) (GetUserInfoResponse, error) {
	urlString := "https://api.cloud.com/cvad/manage/Me"

	req, err := http.NewRequest(http.MethodGet, urlString, nil)
	lib.CheckFatalError(err)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Citrix-CustomerId", customerId)
	req.Header.Add("Authorization", " CWSAuth Bearer="+tokenData.AccessToken)

	res, err := http.DefaultClient.Do(req)
	lib.CheckFatalError(err)

	responseBody, err := io.ReadAll(res.Body)

	if res.StatusCode == 404 {
		fmt.Println(string(responseBody))
		lib.CheckFatalError(fmt.Errorf(res.Status))
	}

	var userInfo GetUserInfoResponse
	json.Unmarshal(responseBody, &userInfo)

	return userInfo, nil
}

func HttpGet(urlString string, customerId string, siteId string, tokenData lib.CitrixTokenData) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, urlString, nil)
	lib.CheckFatalError(err)
	// if err != nil {
	// 	return nil, err
	// }

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Citrix-CustomerId", customerId)
	req.Header.Add("Citrix-InstanceId", siteId)
	req.Header.Add("Authorization", " CWSAuth Bearer="+tokenData.AccessToken)

	res, err := http.DefaultClient.Do(req)
	lib.CheckFatalError(err)
	// if err != nil {
	// 	return nil, err
	// }

	responseBody, err := io.ReadAll(res.Body)

	if res.StatusCode == 404 {
		fmt.Println(string(responseBody))
		lib.CheckFatalError(fmt.Errorf(res.Status))
	}

	// fmt.Println()
	// if err != nil {
	// 	return nil, err
	// }

	return responseBody, nil
}

func HttpPost(urlString string, body string, creds lib.CitrixCloudAccountConfig) ([]byte, []byte, error) {
	bodyReader := bytes.NewReader([]byte(body))

	req, err := http.NewRequest(http.MethodPost, urlString, bodyReader)
	lib.CheckFatalError(err)
	// if err != nil {
	// 	return nil, err
	// }

	// req.Header.Add("Content-Type", "application/json")
	// req.Header.Add("Authorization", "Bearer "+mat.TokenData.Token)

	fmt.Println(body)

	res, err := http.DefaultClient.Do(req)
	lib.CheckFatalError(err)
	// if err != nil {
	// 	return nil, err
	// }

	responseBody, err := io.ReadAll(res.Body)

	// jsonStr, _ := json.MarshalIndent(res.Header, "", "  ")
	// fmt.Println(string(jsonStr))
	// fmt.Println(responseBody)
	lib.CheckFatalError(err)
	if res.StatusCode == 404 {
		fmt.Println(string(responseBody))
		lib.CheckFatalError(fmt.Errorf(res.Status))
	}

	resHeader, _ := json.MarshalIndent(res.Header, "", "  ")

	// fmt.Println()
	// if err != nil {
	// 	return nil, err
	// }

	return responseBody, resHeader, nil
}

func GetToken(creds lib.CitrixCloudAccountConfig, cldConfigOpts *lib.CldConfigOptions) (lib.CitrixTokenData, error) {
	cachedToken := lib.GetCachedToken[lib.CitrixTokenData]("citrix-"+creds.CustomerId, cldConfigOpts)

	if cachedToken != nil {
		isExpired := lib.CheckCachedTokenExpired(cachedToken.Expiry)
		if !isExpired {
			return *cachedToken, nil
		}
	}

	regions := make(map[string]string)
	regions["AP"] = "api-ap-s.cloud.com"
	regions["EU"] = "api-eu.cloud.com"
	regions["US"] = "api-us.cloud.com"
	regions["JP"] = "api.citrixcloud.jp"

	regions["TEST"] = "api.cloud.com"

	param := url.Values{}
	param.Set("grant_type", "client_credentials")
	param.Set("client_id", creds.ClientId)
	param.Set("client_secret", creds.ClientSecret)
	payload := bytes.NewBufferString(param.Encode())

	urlString := "https://" +
		regions[creds.Region] +
		"/cctrustoauth2/" +
		creds.CustomerId +
		// "root" +
		"/tokens/clients"

	req, err := http.NewRequest(http.MethodPost, urlString, payload)
	lib.CheckFatalError(err)

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Accept", "application/json")

	res, err := http.DefaultClient.Do(req)
	lib.CheckFatalError(err)

	responseBody, err := io.ReadAll(res.Body)
	lib.CheckFatalError(err)

	var tokenData lib.CitrixTokenData

	json.Unmarshal(responseBody, &tokenData)
	resHeader, _ := json.MarshalIndent(res.Header, "", "  ")

	if res.StatusCode == 404 {
		fmt.Println("Body")
		fmt.Println(string(responseBody))
		fmt.Println("Header")
		fmt.Println(string(resHeader))
		lib.CheckFatalError(fmt.Errorf(res.Status))
	}

	currentTime := time.Now()
	// .Format(time.RFC3339)
	// fmt.Println(currentTime)
	// tokenExpiresIn := tokenData.ExpiresIn
	// fmt.Println(tokenExpiresIn)
	tokenData.Expiry = currentTime.Add(time.Second * 3590)
	// .Format(time.RFC3339)
	// fmt.Println(tokenData.Expiry)
	// parsedExpiry, err := time.Parse(time.RFC3339, tokenExpiry)
	// lib.CheckFatalError(err)
	// fmt.Println(parsedExpiry)

	// jsonStr, _ := json.MarshalIndent(tokenData, "", "  ")
	// fmt.Println(string(jsonStr))
	// os.Exit(0)
	// fmt.Println("New token expiration: " + tokenData.Expiry.Format(time.RFC3339))

	lib.CacheSaveToken(tokenData, "citrix", cldConfigOpts)
	return tokenData, nil
}

func GetMachineCatalaogsForAllConfiguredtenants() (combinedMachineCatalogs MachineCatalogs) {
	config := lib.GetCldConfig(nil)
	citrixEnvs := *config.CitrixCloud.Environments
	if citrixEnvs == nil {
		err := fmt.Errorf("Citrix environments not configured")
		lib.CheckFatalError(err)
	}
	for tName, envCreds := range citrixEnvs {
		tokenData, err := GetToken(envCreds, nil)
		lib.CheckFatalError(err)
		machineCatalogs := ListMachineCatalogs(envCreds, tokenData)
		for _, mc := range machineCatalogs {
			currMc := mc
			currMc.TenantName = tName
			combinedMachineCatalogs = append(combinedMachineCatalogs, currMc)
		}
	}
	return
}
