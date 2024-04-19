package windows

import (
	"fmt"
	"os"
	"slices"

	"github.com/jercle/cloudini/lib"
	"golang.org/x/sys/windows/registry"
)

func SetProxySettings(config lib.ProxyConfig, ignoreOverrides bool) {
	k, err := registry.OpenKey(registry.CURRENT_USER, `SOFTWARE\Microsoft\Windows\CurrentVersion\Internet Settings`, registry.WRITE)
	lib.CheckFatalError(err)
	defer k.Close()

	err = k.SetDWordValue("ProxyEnable", 1)
	lib.CheckFatalError(err)
	err = k.SetStringValue("ProxyServer", config.Server+":"+config.Port)
	lib.CheckFatalError(err)

	if config.Overrides != "" || !ignoreOverrides {
		err = k.SetStringValue("ProxyOverride", config.Overrides)
		lib.CheckFatalError(err)
	}
}

func GetProxySettings() lib.ProxyConfig {
	k, err := registry.OpenKey(registry.CURRENT_USER, `SOFTWARE\Microsoft\Windows\CurrentVersion\Internet Settings`, registry.READ)
	lib.CheckFatalError(err)
	defer k.Close()

	valueNames, err := k.ReadValueNames(0)
	lib.CheckFatalError(err)
	if !slices.Contains(valueNames, "ProxyServer") || !slices.Contains(valueNames, "ProxyEnable") {
		fmt.Println("Proxy not enabled")
		os.Exit(0)
	}

	proxyServer, _, err := k.GetStringValue("ProxyServer")
	lib.CheckFatalError(err)
	proxyEnabledValue, _, err := k.GetIntegerValue("ProxyEnable")
	lib.CheckFatalError(err)
	var proxyOverrides string
	if slices.Contains(valueNames, "ProxyOverride") {
		proxyOverrides, _, err = k.GetStringValue("ProxyOverride")
		lib.CheckFatalError(err)
	}

	var proxyEnabled bool
	if proxyEnabledValue == 1 {
		proxyEnabled = true
	}

	if proxyServer == "" || !proxyEnabled {
		fmt.Println("Proxy not enabled")
		os.Exit(0)
	}

	var proxyConfig lib.ProxyConfig
	proxyConfig.Enabled = proxyEnabled
	proxyConfig.Server = proxyServer
	proxyConfig.Overrides = proxyOverrides

	return proxyConfig

	// fmt.Println("Proxy enabled?", proxyEnabled)
	// fmt.Println("Proxy Server:", proxyServer)
	// fmt.Println("Proxy Overrides:", proxyOverrides)
}

func RemoveProxyConfig() {
	k, err := registry.OpenKey(registry.CURRENT_USER, `SOFTWARE\Microsoft\Windows\CurrentVersion\Internet Settings`, registry.SET_VALUE)
	lib.CheckFatalError(err)
	defer k.Close()

	k.DeleteValue("ProxyServer")
	k.DeleteValue("ProxyEnable")
	k.DeleteValue("ProxyOverride")
}
