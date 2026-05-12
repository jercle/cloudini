package azure

import (
	"bytes"
	"encoding/json/jsontext"
	json "encoding/json/v2"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/jercle/cloudini/lib"
)

func HttpGet(urlString string, mat lib.AzureMultiAuthToken) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, urlString, nil)
	lib.CheckFatalError(err)
	// if err != nil {
	// 	return nil, err
	// }

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+mat.TokenData.Token)

	res, err := http.DefaultClient.Do(req)
	// lib.CheckFatalError(err)

	if err != nil {
		return nil, err
	}

	// fmt.Println(res.Status)
	// fmt.Println(res.StatusCode)

	// lib.JsonMarshalAndPrint(res.Header)

	responseBody, err := io.ReadAll(res.Body)

	if res.StatusCode == 403 {
		fmt.Println(res.Status, urlString)
	}

	if res.StatusCode == 404 {
		// fmt.Println(string(responseBody))
		// lib.CheckFatalError()
		jsonErr := `{"status": "` + res.Status + `", "response": ` + string(responseBody) + `}`
		err = fmt.Errorf(jsonErr)
		return nil, err
	}

	// fmt.Println()
	// if err != nil {
	// 	return nil, err
	// }

	return responseBody, nil
}

//
//

func HttpGetErrLogToCache(urlString string, mat lib.AzureMultiAuthToken) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, urlString, nil)
	lib.CheckFatalError(err)
	// if err != nil {
	// 	return nil, err
	// }

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+mat.TokenData.Token)

	res, err := http.DefaultClient.Do(req)
	// lib.CheckFatalError(err)
	if err != nil {
		return nil, err
	}

	// fmt.Println(res.Status)
	// fmt.Println(res.StatusCode)

	// lib.JsonMarshalAndPrint(res.Header)

	responseBody, err := io.ReadAll(res.Body)

	// if res.StatusCode == 404 {
	// 	// fmt.Println(string(responseBody))
	// 	// lib.CheckFatalError()
	// 	jsonErr := `{"status": "` + res.Status + `", "response": ` + string(responseBody) + `}`
	// 	err = fmt.Errorf(jsonErr)
	// 	return nil, err
	// }

	if res.StatusCode != 200 && res.StatusCode != 204 {
		if res.StatusCode == 403 {
			_, _, cachePath := lib.InitConfig(nil)
			fmt.Println(res.Status, urlString)
			f, err := os.OpenFile(cachePath+"/403.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				log.Fatal(err)
			}
			defer f.Close()
			date := time.Now().Format(time.RFC3339)
			if _, err := f.WriteString(date + " - PATCH - " + res.Status + " - " + urlString + "\n"); err != nil {
				log.Fatal(err)
			}
		}
		resStr, _ := json.Marshal(res)
		err = fmt.Errorf(urlString, string(resStr))
		return nil, err
	}

	// fmt.Println()
	// if err != nil {
	// 	return nil, err
	// }

	return responseBody, nil
}

//
//

func HttpPost(urlString string, body string, mat lib.AzureMultiAuthToken) ([]byte, []byte, error) {
	bodyReader := bytes.NewReader([]byte(body))

	req, err := http.NewRequest(http.MethodPost, urlString, bodyReader)
	lib.CheckFatalError(err)
	// if err != nil {
	// 	return nil, err
	// }

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+mat.TokenData.Token)

	// fmt.Println(body)

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
		// lib.CheckFatalError(fmt.Errorf(res.Status))
		return nil, nil, fmt.Errorf(res.Status)
	}
	if res.StatusCode == 403 {
		fmt.Println(string(responseBody))
		// lib.CheckFatalError(fmt.Errorf(res.Status))
		return nil, nil, fmt.Errorf(res.Status)
	}

	resHeader, _ := json.Marshal(res.Header, jsontext.WithIndent("  "))

	// fmt.Println()
	// if err != nil {
	// 	return nil, err
	// }

	return responseBody, resHeader, nil
}

//
//

func HttpPut(urlString string, body string, mat lib.AzureMultiAuthToken) ([]byte, []byte, error) {
	bodyReader := bytes.NewReader([]byte(body))

	req, err := http.NewRequest(http.MethodPut, urlString, bodyReader)
	lib.CheckFatalError(err)
	// if err != nil {
	// 	return nil, err
	// }

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+mat.TokenData.Token)

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

	resHeader, _ := json.Marshal(res.Header, jsontext.WithIndent("  "))

	// fmt.Println()
	// if err != nil {
	// 	return nil, err
	// }

	return responseBody, resHeader, nil
}

//
//

func HttpPatch(urlString string, body string, mat lib.AzureMultiAuthToken) ([]byte, []byte, error) {
	bodyReader := bytes.NewReader([]byte(body))

	req, err := http.NewRequest(http.MethodPatch, urlString, bodyReader)
	lib.CheckFatalError(err)
	// if err != nil {
	// 	return nil, err
	// }

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+mat.TokenData.Token)

	// fmt.Println(body)

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
	} else if res.StatusCode == 403 {
		fmt.Println(res.Status, urlString)
	}

	resHeader, _ := json.Marshal(res.Header, jsontext.WithIndent("  "))

	// fmt.Println()
	// if err != nil {
	// 	return nil, err
	// }

	return responseBody, resHeader, nil
}

func HttpPatchErrLogToCache(urlString string, body string, mat lib.AzureMultiAuthToken) ([]byte, []byte, error) {
	bodyReader := bytes.NewReader([]byte(body))

	req, err := http.NewRequest(http.MethodPatch, urlString, bodyReader)
	lib.CheckFatalError(err)
	// if err != nil {
	// 	return nil, err
	// }

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+mat.TokenData.Token)

	// fmt.Println(body)

	res, err := http.DefaultClient.Do(req)
	lib.CheckFatalError(err)
	// if err != nil {
	// 	return nil, err
	// }

	responseBody, err := io.ReadAll(res.Body)

	resHeader, _ := json.Marshal(res.Header, jsontext.WithIndent("  "))
	// lib.CheckFatalError(err)
	if err != nil {
		return responseBody, []byte(res.Status), nil
	}
	// jsonStr, _ := json.MarshalIndent(res.Header, "", "  ")
	// fmt.Println(string(jsonStr))
	// fmt.Println(responseBody)
	switch res.StatusCode {
	case 404:
		fmt.Println(string(responseBody))
		lib.CheckFatalError(fmt.Errorf(res.Status))
		return responseBody, resHeader, err
	case 403:
		date := time.Now().Format(time.RFC3339)
		_, _, cachePath := lib.InitConfig(nil)
		// fmt.Println(res.Status, urlString)
		f, e := os.OpenFile(cachePath+"/403.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if e != nil {
			// log.Fatal(e)
			log.Println(date + " - PATCH - " + res.Status + " - " + urlString + "\n")
		}
		defer f.Close()
		if _, e := f.WriteString(date + " - PATCH - " + res.Status + " - " + urlString + "\n"); e != nil {
			log.Println(e)
			// log.Fatal(e)
		}
		return responseBody, resHeader, err
	}

	// fmt.Println()
	// if err != nil {
	// 	return nil, err
	// }

	return responseBody, resHeader, nil
}

//
//

func GetAzureResourceTypes(token *lib.AzureMultiAuthToken) (types []string) {
	graphQuery := "resources | distinct type"
	jsonBody := `{
    "query": "` + graphQuery + `"
}`
	urlString := "https://management.azure.com/providers/Microsoft.ResourceGraph/resources?api-version=2022-10-01"
	res, _, err := HttpPost(urlString, jsonBody, *token)
	lib.CheckFatalError(err)

	var resData ResourceGraphIPAddressesResponse
	// os.WriteFile("main-ips-0-prewhile.json", res, 0644)

	err = json.Unmarshal(res, &resData)
	lib.CheckFatalError(err)

	for _, t := range resData.Data {
		jsonStr, _ := json.Marshal(t)
		var item struct {
			Type string `json:"type"`
		}
		err := json.Unmarshal(jsonStr, &item)
		lib.CheckFatalError(err)

		types = append(types, item.Type)
	}

	// jsonStr, _ := json.Marshal(types, jsontext.WithIndent("  "))
	// os.WriteFile("main-ips-resourceTypes.json", jsonStr, 0644)
	// lib.JsonMarshalAndPrint(types)
	// fmt.Println(len(types))
	return
}
