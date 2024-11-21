/*
Copyright © 2024 Evan Colwell ercolwell@gmail.com
*/
package citrix

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/jercle/cloudini/lib"
)

func ListMachineCatalogs(creds lib.CitrixCloudAccountConfig, tokenData lib.CitrixTokenData) MachineCatalogs {
	urlString := "https://api.cloud.com/cvad/manage/MachineCatalogs"
	res, err := HttpGet(urlString, creds.CustomerId, creds.SiteId, tokenData)
	lib.CheckFatalError(err)

	var machineCatalogs GetMachineCatalogsResponse
	json.Unmarshal(res, &machineCatalogs)

	return machineCatalogs.Items
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
