/*
Copyright © 2024 Evan Colwell ercolwell@gmail.com
*/
package azure

import (
	"io"
	"net/http"

	"github.com/jercle/cloudini/lib"
)

func StorageBlobHttpGet(urlString string, mat lib.AzureMultiAuthToken) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, urlString, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/xml")
	req.Header.Add("Authorization", "Bearer "+mat.TokenData.Token)
	req.Header.Add("x-ms-version", "2023-11-03")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	responseBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return responseBody, nil
}
