package azure

import (
	"io"
	"net/http"

	"github.com/jercle/azg/lib"
)

func azureHttpGet(urlString string, mat MultiAuthToken) []byte {
	req, err := http.NewRequest(http.MethodGet, urlString, nil)
	lib.CheckFatalError(err)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+mat.TokenData.Token)

	res, err := http.DefaultClient.Do(req)
	lib.CheckFatalError(err)

	responseBody, err := io.ReadAll(res.Body)
	lib.CheckFatalError(err)

	return responseBody
}
