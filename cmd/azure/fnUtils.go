package azure

import (
	"bytes"
	"encoding/json/jsontext"
	json "encoding/json/v2"
	"fmt"
	"io"
	"net/http"

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

	// lib.JsonMarshalAndPrint(res.Header)

	responseBody, err := io.ReadAll(res.Body)

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
