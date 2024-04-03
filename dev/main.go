package main

import (
	"fmt"
	"log"

	"github.com/jercle/azg/lib"
	"golang.org/x/sys/windows/registry"
)

func main() {
	GetProxySettings()
}

func GetProxySettings() {
	k, err := registry.OpenKey(registry.CURRENT_USER, `SOFTWARE\Microsoft\Windows\CurrentVersion\Internet Settings`, registry.QUERY_VALUE)
	lib.CheckFatalError(err)
	defer k.Close()

	proxyServer, _, err := k.GetStringValue("ProxyServer")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Current proxy settings: ", proxyServer)
}

func RemoveProxyConfig() {
	k, err := registry.OpenKey(registry.CURRENT_USER, `SOFTWARE\Microsoft\Windows\CurrentVersion\Internet Settings`, registry.QUERY_VALUE)
	lib.CheckFatalError(err)
	defer k.Close()

}
