//go:build windows

package windows

import (
	"fmt"
	"log"
	"strings"

	"golang.org/x/sys/windows/registry"
)

type InstalledApp struct {
	DisplayName     string
	DisplayVersion  string
	Publisher       string
	InstallDate     string
	UninstallString string
	// Name            string
	InstallLocation string
	// LogFile         string
	// Installed       string
}

// var m = make(map[string]bool)
// var a = []string{}
// func add(s string) {
// 	if m[s] {
// 		return // Already in the map
// 	}
// 	a = append(a, s)
// 	m[s] = true
// }

func checkInstalledApp(searchString string) []InstalledApp {
	installedApps := getInstalledApps()
	found := false

	for _, app := range installedApps {
		// fmt.Println(app)
		if strings.Contains(strings.ToLower(app.DisplayName), strings.ToLower(searchString)) {
			found = true
		}
	}

	if !found {
		fmt.Println(searchString, "is not installed")
		return installedApps
		// os.Exit(1)
	} else {
		fmt.Println(searchString, "is installed")
		return installedApps
		// os.Exit(0)
	}
}

func getInstalledApps() []InstalledApp {
	regKeys := [2]string{
		`SOFTWARE\Wow6432Node\Microsoft\Windows\CurrentVersion\Uninstall`,
		`SOFTWARE\Microsoft\Windows\CurrentVersion\Uninstall`,
	}

	var installedApps []InstalledApp

	for _, regKey := range regKeys {
		key, err := registry.OpenKey(registry.LOCAL_MACHINE, regKey, registry.QUERY_VALUE|registry.ENUMERATE_SUB_KEYS)
		if err != nil {
			log.Fatal(err)
		}

		apps, err := key.ReadSubKeyNames(0)
		if err != nil {
			log.Fatal(err)
		}

		for _, app := range apps {
			program, err := registry.OpenKey(registry.LOCAL_MACHINE, regKey+`\`+app, registry.QUERY_VALUE|registry.ENUMERATE_SUB_KEYS)
			if err != nil {
				log.Fatal(err)
			}

			// valNames, _ := program.ReadValueNames(0)
			// for _, val := range valNames {
			// 	add(val)
			// }
			// valString, _ := json.MarshalIndent(valNames, "", "  ")

			var appStruct InstalledApp

			appStruct.DisplayName, _, _ = program.GetStringValue("DisplayName")
			appStruct.DisplayVersion, _, _ = program.GetStringValue("DisplayVersion")
			appStruct.InstallDate, _, _ = program.GetStringValue("InstallDate")
			appStruct.Publisher, _, _ = program.GetStringValue("Publisher")
			appStruct.UninstallString, _, _ = program.GetStringValue("UninstallString")
			appStruct.InstallLocation, _, _ = program.GetStringValue("InstallLocation")
			// appStruct.LogFile, _, _ = program.GetStringValue("LogFile")
			// appStruct.Name, _, _ = program.GetStringValue("Name")
			// appStruct.Installed, _, _ = program.GetStringValue("Installed")

			if appStruct.DisplayName != "" {
				installedApps = append(installedApps, appStruct)
			}
			defer program.Close()

		}
		defer key.Close()

	}

	// fields := []string{}
	// for k, _ := range m {
	// 	fields = append(fields, k)
	// }
	// sort.Slice(fields, func(i, j int) bool {
	// 	return strings.ToLower(fields[i]) < strings.ToLower(fields[j])
	// })
	// for _, name := range fields {
	// 	fmt.Println(name)
	// }

	return installedApps
}
