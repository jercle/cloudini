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

	"github.com/TylerBrock/colorjson"
	"github.com/tidwall/pretty"
)

type Request struct {
	Url     string
	Outfile string
}

func get(opts Request) {

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
