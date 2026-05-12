package ado

import (
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"
)

var AzureDevOpsBaseUrl = "https://dev.azure.com/"
var AzureDevopsApiVersion = "api-version=7.2-preview.1"

func CreateBasicAuthHeaderValue(username string, password string) string {
	auth := ":" + password
	return "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))
}

func azureDevOpsRestGet(uri string, pat string) []byte {
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", CreateBasicAuthHeaderValue("", pat))

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	var resData *interface{}
	derr := json.NewDecoder(res.Body).Decode(&resData)
	if derr != nil {
		log.Fatal(derr)
	}

	byteData, err := json.Marshal(resData)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(string(byteData))

	return byteData
}
