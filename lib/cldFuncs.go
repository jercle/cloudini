package lib

import (
	"encoding/json"
	"os"
)

var ConfigPath = InitConfig(nil)
var TokenCacheFile = ConfigPath + "/tCache"

func GetCldConfig(options *CldConfigOptions) CldConfigRoot {
	var (
		config CldConfigRoot
	)

	configFilePath := InitConfig(nil)

	byteValue, err := os.ReadFile(configFilePath)
	CheckFatalError(err)

	// decodedBytes, err := b64.StdEncoding.DecodeString(string(byteValue))
	// CheckFatalError(err)

	// fmt.Println(string(byteValue))
	// os.Exit(0)

	// err = json.Unmarshal(decodedBytes, &config)
	err = json.Unmarshal(byteValue, &config)
	CheckFatalError(err)

	return config
}

func InitConfig(options *CldConfigOptions) string {
	usrHomeDir, err := os.UserHomeDir()
	var (
		configPath     = usrHomeDir + "/.config/cld"
		configFile     = configPath + "/cldConf.json"
		configFilePath string
	)
	CheckFatalError(err)

	if options != nil && options.ConfigFilePath != "" {
		configFilePath = options.ConfigFilePath
	} else {
		configFilePath = configFile
	}

	if _, err := os.Stat(configPath); err != nil {
		os.MkdirAll(configPath, os.ModePerm)
	}

	fileStat, err := os.Stat(configFilePath)

	if err != nil || fileStat.Size() == 0 {
		config := CldConfigRoot{}
		config.Azure.MultiTenantAuth.Tenants = make(CldConfigTenants)
		jsonBytes, _ := json.MarshalIndent(config, "", "  ")

		// encodedData := b64.StdEncoding.EncodeToString(jsonBytes)
		os.WriteFile(configFilePath, jsonBytes, os.ModePerm)
	}

	return configFilePath
}
