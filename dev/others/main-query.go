package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func mainquery() {
	pat := os.Getenv("DEVOPS_PAT")
	org := os.Getenv("DEVOPS_ORG")
	project := os.Getenv("DEVOPS_PROJECT")

	pipelines := GetPipelines(org, project, pat)
	fmt.Println(string(pipelines))
}

func CreateBasicAuthHeaderValue(username string, password string) string {
	auth := username + ":" + password
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

func GetPipelines(org string, project string, pat string) []byte {
	urlString := "https://dev.azure.com/" + org + "/" + project + "/_apis/pipelines?api-version=7.2-preview.1"
	data := azureDevOpsRestGet(urlString, pat)
	return data
}
