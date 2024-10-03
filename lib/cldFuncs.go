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

	configFile := InitConfig(options)

	// fmt.Println("Getting config from", configFile)
	byteValue, err := os.ReadFile(configFile)

	CheckFatalError(err)

	err = json.Unmarshal(byteValue, &config)
	CheckFatalError(err)

	return config
}

func InitConfig(options *CldConfigOptions) string {
	usrHomeDir, err := os.UserHomeDir()
	var (
		configPath     = usrHomeDir + "/.config/cld"
		configFilePath = configPath + "/cldConf.json"
		configFile     string
	)
	CheckFatalError(err)

	CLD_CONFIG_PATH := os.Getenv("CLD_CONFIG_PATH")

	if options != nil && options.ConfigFile != "" {
		configFile = options.ConfigFile
	} else if CLD_CONFIG_PATH != "" {
		configFile = CLD_CONFIG_PATH
	} else {
		configFile = configFilePath
	}

	if _, err := os.Stat(configPath); err != nil {
		os.MkdirAll(configPath, os.ModePerm)
	}

	fileStat, err := os.Stat(configFile)

	if err != nil || fileStat.Size() == 0 {
		config := CldConfigRoot{}
		config.Azure.MultiTenantAuth.Tenants = make(CldConfigTenants)
		jsonBytes, _ := json.MarshalIndent(config, "", "  ")

		// encodedData := b64.StdEncoding.EncodeToString(jsonBytes)
		os.WriteFile(configFile, jsonBytes, os.ModePerm)
	}

	return configFile
}
