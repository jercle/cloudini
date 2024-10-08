package azure

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/jercle/cloudini/lib"
)

func HttpGet(urlString string, mat lib.MultiAuthToken) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, urlString, nil)
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

func HttpPost(urlString string, body string, mat lib.MultiAuthToken) ([]byte, []byte, error) {
	bodyReader := bytes.NewReader([]byte(body))

	req, err := http.NewRequest(http.MethodPost, urlString, bodyReader)
	lib.CheckFatalError(err)
	// if err != nil {
	// 	return nil, err
	// }

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+mat.TokenData.Token)

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

func HttpPut(urlString string, mat lib.MultiAuthToken) ([]byte, []byte, error) {
	req, err := http.NewRequest(http.MethodPost, urlString, nil)
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

	resHeader, _ := json.MarshalIndent(res.Header, "", "  ")

	// fmt.Println()
	// if err != nil {
	// 	return nil, err
	// }

	return responseBody, resHeader, nil
}
