package main

import (
	"fmt"

	"github.com/jercle/azg/lib"
	"golang.org/x/sys/windows/registry"
)

func main() {
	// GetProxySettings()
	// RemoveProxyConfig()
	// GetProxySettings()

	confLocation := lib.InitConfig(&lib.CldConfigOptions{})
	fmt.Println(confLocation)

	// k, err := registry.OpenKey(registry.CURRENT_USER, `SOFTWARE\Microsoft\Windows\CurrentVersion\Internet Settings`, registry.READ)
	// lib.CheckFatalError(err)

	// defer k.Close()

	// subKeyNames, err := k.ReadSubKeyNames(0)
	// lib.CheckFatalError(err)

	// valueNames, err := k.ReadValueNames(0)
	// lib.CheckFatalError(err)

	// fmt.Println("subKeyNames")
	// subKeyBytes, _ := json.MarshalIndent(subKeyNames, "", "  ")
	// fmt.Println(string(subKeyBytes))
	// fmt.Println("valueNames")
	// valueBytes, _ := json.MarshalIndent(valueNames, "", "  ")
	// fmt.Println(string(valueBytes))
	// fmt.Println(valueNames)

	// proxyServer, _, err := k.GetStringValue("ProxyServer")

}
func SetProxySettings(config string) {
	k, err := registry.OpenKey(registry.CURRENT_USER, `SOFTWARE\Microsoft\Windows\CurrentVersion\Internet Settings`, registry.WRITE)
	lib.CheckFatalError(err)
	defer k.Close()

}

func GetProxySettings() {
	k, err := registry.OpenKey(registry.CURRENT_USER, `SOFTWARE\Microsoft\Windows\CurrentVersion\Internet Settings`, registry.QUERY_VALUE)
	lib.CheckFatalError(err)
	defer k.Close()

	proxyServer, _, err := k.GetStringValue("ProxyServer")
	lib.CheckFatalError(err)
	proxyEnable, _, err := k.GetIntegerValue("ProxyEnable")
	lib.CheckFatalError(err)
	proxyOverride, _, err := k.GetStringValue("ProxyOverride")
	lib.CheckFatalError(err)

	fmt.Println("Current proxy server settings:", proxyServer)
	fmt.Println("Current proxy enabl settings:", proxyEnable)
	fmt.Println("Current proxy enabl settings:", proxyOverride)
}

func RemoveProxyConfig() {
	k, err := registry.OpenKey(registry.CURRENT_USER, `SOFTWARE\Microsoft\Windows\CurrentVersion\Internet Settings`, registry.SET_VALUE)
	lib.CheckFatalError(err)
	defer k.Close()

	// k.DeleteValue("ProxyServer")

	k.SetStringValue("ProxyServer", "")

	// k.

}

// func GetRegistryValues
// *.apetp.gov.au;172.*;<local>
// *.apetp.gov.au;172.*;localhost;127.0.0.;10.;192.168.*;*.apetp.gov.au;*.azure.net;*.azure.com;*.windows.net;*.visualstudio.com;*.microsoft.com
