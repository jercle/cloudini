package m365

import (
	"bytes"
	"fmt"
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

	client := &http.Client{}
	_, err = client.Do(req)
	lib.CheckFatalError(err)

	fmt.Println(fmt.Sprintf("Uploaded to: %s/%s", folderPath, fileName))
}
