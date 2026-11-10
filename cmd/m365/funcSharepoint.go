package m365

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/jercle/cloudini/lib"
)

func UploadFileToSharepoint(dataDirectory string, siteId string, driveId string, folderPath string, fileName string, costExportMonth string, token string) {

	urlString := "https://graph.microsoft.com/v1.0/sites/" +
		siteId +
		"/drives/" +
		driveId +
		"/root:/" +
		folderPath +
		"/" +
		fileName +
		":/content"

	file, err := os.ReadFile(dataDirectory + "/" + costExportMonth + "/" + fileName)
	lib.CheckFatalError(err)

	reader := bytes.NewReader(file)
	req, err := http.NewRequest("PUT", urlString, reader)

	req.Header.Add("Authorization", "Bearer "+token)
	if err != nil {
		log.Fatal(err)
	}
	lib.CheckFatalError(err)

	res, err := lib.HTTPClient.Do(req)
	lib.CheckFatalError(err)
	defer res.Body.Close()

	if res.StatusCode < 200 || res.StatusCode > 299 {
		resBody, _ := io.ReadAll(res.Body)
		lib.CheckFatalError(fmt.Errorf("upload of %s failed: %s: %s", fileName, res.Status, string(resBody)))
	}

	fmt.Println(fmt.Sprintf("Uploaded to: %s/%s", folderPath, fileName))
}
