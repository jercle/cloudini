package main

import (
	"fmt"
	"log"

	"golang.org/x/sys/windows/registry"
)

type ProxyConf struct {
	Enabled     bool
	ProxyServer string
	ProxyPort   string
	ProxyIgnore string
}

func main() {
	k, err := registry.OpenKey(registry.CURRENT_USER, `SOFTWARE\Microsoft\Windows\CurrentVersion\Internet Settings`, registry.QUERY_VALUE)
	if err != nil {
		log.Fatal(err)
	}
	defer k.Close()

	s, _, err := k.GetStringValue("ProxyServer")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Windows system root is %q\n", s)

}
