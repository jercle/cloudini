package main

import "github.com/jercle/cloudini/lib"

func main() {
	file := "/home/jercle/git/cld/cmd/azure/typesGetIpAddresses.go"
	lib.AddJsonOmitemptyTagsToStructFile(file, true)
	lib.AddBsonTagsFromJsonTagsToStructFile(file, true)
}
