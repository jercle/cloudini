package lib

import (
	"encoding/json"
	"io"
	"os"
)

var usrHomeDir, err = os.UserHomeDir()
var configPath = usrHomeDir + "/.config/cld"
var configFile = configPath + "/cldConf.json"

func GetCldConfig(options CldConfigOptions) CldConfig {
	var (
		config     CldConfig
		configPath string
		jsonConfig *os.File
		byteValue  []byte
	)

	if options.ConfigFilePath != "" {
		configPath = options.ConfigFilePath
	} else {
		configPath = configFile
	}

	if _, err := os.Stat(configPath); err != nil {
		os.MkdirAll(configPath, os.ModePerm)
	}
	if _, err = os.Stat(configFile); err != nil {
		jsonConfig, err = os.Create(configFile)
		CheckFatalError(err)
		defer jsonConfig.Close()

		byteValue, err = json.Marshal(config)
		CheckFatalError(err)

		os.WriteFile(configPath, byteValue, os.ModePerm)

	} else {
		jsonConfig, err = os.Open(configPath)
		CheckFatalError(err)
		defer jsonConfig.Close()
		byteValue, _ = io.ReadAll(jsonConfig)
		json.Unmarshal(byteValue, &config)
		CheckFatalError(err)
	}

	// fmt.Println(config)

	return config
}
