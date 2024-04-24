package main

import (
	"fmt"

	"github.com/jercle/cloudini/cmd/web"
)

func main() {
	urlString := "https://SUBDOMAIN.azurewebsites.net/api/users/me"
	token := ""

	fmt.Println(string(web.SimpleGetRequestWithToken(urlString, token)))
}
