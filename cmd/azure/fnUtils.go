package azure

import (
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
