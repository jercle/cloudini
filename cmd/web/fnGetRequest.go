// https://www.developer.com/languages/json-files-golang/

package web

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/jercle/cloudini/lib"
	"github.com/TylerBrock/colorjson"
	"github.com/tidwall/pretty"
)

type Request struct {
	Url     string
	Outfile string
}

func Get(opts Request) {

	_, err := url.ParseRequestURI(opts.Url)
	if err != nil {
		log.Fatal("Invalid URL")
	}

	res, error := http.Get(opts.Url)
	if error != nil {
		log.Fatal(error)
	}

	body, _ := io.ReadAll(res.Body)

	var jsonObject map[string]interface{}
	json.Unmarshal([]byte(body), &jsonObject)

	formatter := colorjson.NewFormatter()
	formatter.Indent = 2
	jsonString, _ := formatter.Marshal(jsonObject)
	fmt.Println(string(jsonString))

	if opts.Outfile != "" {
		err = os.WriteFile(opts.Outfile, pretty.Pretty(body), os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
	}
}

// Returns response body
func SimpleGetRequestWithToken(urlString string, token string) []byte {
	req, err := http.NewRequest(http.MethodGet, urlString, nil)
	lib.CheckFatalError(err)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+token)

	res, err := http.DefaultClient.Do(req)
	lib.CheckFatalError(err)

	responseBody, err := io.ReadAll(res.Body)
	lib.CheckFatalError(err)
	defer res.Body.Close()

	return responseBody
}

// Returns response body
func SimpleGetRequest(urlString string, token string) []byte {
	req, err := http.NewRequest(http.MethodGet, urlString, nil)
	lib.CheckFatalError(err)

	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	lib.CheckFatalError(err)

	responseBody, err := io.ReadAll(res.Body)
	lib.CheckFatalError(err)
	defer res.Body.Close()

	return responseBody
}

// func RemoveBOM(resp *http.Response) error {
// 	_, err := exported.Payload(resp, &exported.PayloadOptions{
// 		BytesModifier: func(b []byte) []byte {
// 			// UTF8
// 			return bytes.TrimPrefix(b, []byte("\xef\xbb\xbf"))
// 		},
// 	})
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
