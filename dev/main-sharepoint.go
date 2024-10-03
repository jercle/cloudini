package main

import (
	"bytes"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/jercle/cloudini/cmd/azure"
	"github.com/jercle/cloudini/lib"
)

func main() {
	startTime := time.Now()
	config := lib.GetCldConfig(nil)
	_ = config
	// tokens, err := azure.GetAllTenantSPTokens(lib.MultiAuthTokenRequestOptions{})
	// lib.CheckFatalError(err)
	// token, err := azure.GetTenantSPToken(lib.MultiAuthTokenRequestOptions{})
	tokenReq, err := azure.GetAllTenantSPTokens(lib.MultiAuthTokenRequestOptions{
		Scope:         "graph",
		GetWriteToken: true,
	})
	lib.CheckFatalError(err)
	token, err := tokenReq.SelectTenant("RED")
	lib.CheckFatalError(err)
	// _ = tokens
	_ = token

	// jsonStr, _ := json.MarshalIndent(token, "", "  ")
	// fmt.Println(string(jsonStr))
	// os.Exit(0)

	// See: https://learn.microsoft.com/en-us/graph/api/driveitem-list-children?view=graph-rest-1.0&tabs=http

	// urlString := "https://SITE.sharepoint.com/sites/CloudOperationsSupportTeam/drive/items"

	// GET https://graph.microsoft.com/v1.0/drives/{drive-id}/root:/{path-relative-to-root}:/children

	// urlString := "https://graph.microsoft.com/v1.0/drives/"
	// urlString := "https://graph.microsoft.com/v1.0/sites/SITE.sharepoint.com/sites/TEAM"
	// urlString := "https://graph.microsoft.com/v1.0/sites/getAllSites"

	siteId := ""
	driveId := ""
	// driveItemId := ""
	// _ = driveItemId

	costExportMonth := "202407"
	dataPath := "./cost-exports/" + costExportMonth + "/"
	fileName := "MonthlyCostReport-" + costExportMonth + ".xlsx"
	_ = dataPath
	_ = fileName

	folderPath := ""
	urlString := "https://graph.microsoft.com/v1.0/sites/" +
		siteId +
		"/drives/" +
		driveId +
		// "/root:/General" +
		"/root:/" +
		folderPath +
		"/" +
		fileName +
		":/content"
		// "/items/" +
		// driveItemId +
		// ":/children"

	UploadFileToSharepoint(urlString, dataPath+fileName, token.TokenData.Token)

	elapsed := time.Since(startTime)
	_ = elapsed
}

func UploadFileToSharepoint(url string, filepath string, token string) {
	file, err := os.ReadFile(filepath)
	lib.CheckFatalError(err)

	reader := bytes.NewReader(file)
	req, err := http.NewRequest("PUT", url, reader)

	req.Header.Add("Authorization", "Bearer "+token)
	if err != nil {
		log.Fatal(err)
	}
	lib.CheckFatalError(err)

	client := &http.Client{}
	_, err = client.Do(req)
	lib.CheckFatalError(err)
}
