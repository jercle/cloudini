package lib

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func GetCldConfig(options CldConfigOptions) CldConfig {
	var (
		config     CldConfig
		jsonConfig *os.File
		byteValue  []byte
		err        error
	)

	configFilePath := InitConfig(CldConfigOptions{})

	if _, err = os.Stat(configFilePath); err != nil {
		jsonConfig, err = os.Create(configFilePath)
		CheckFatalError(err)
		defer jsonConfig.Close()

		byteValue, err = json.Marshal(config)
		CheckFatalError(err)

		os.WriteFile(configFilePath, byteValue, os.ModePerm)

	} else {
		jsonConfig, err = os.Open(configFilePath)
		CheckFatalError(err)
		defer jsonConfig.Close()
		byteValue, _ = io.ReadAll(jsonConfig)
		json.Unmarshal(byteValue, &config)
		CheckFatalError(err)
	}

	// fmt.Println(config)

	return config
}

func InitConfig(options CldConfigOptions) string {
	usrHomeDir, err := os.UserHomeDir()
	var (
		configPath     = usrHomeDir + "/.config/cld"
		configFile     = configPath + "/cldConf.json"
		configFilePath string
	)
	CheckFatalError(err)

	if options.ConfigFilePath != "" {
		configFilePath = options.ConfigFilePath
	} else {
		configFilePath = configFile
	}

	if _, err := os.Stat(configFilePath); err != nil {
		os.MkdirAll(configFilePath, os.ModePerm)
	}

	// fmt.Println(configFilePath)

	return configFilePath
}

func (config *CldConfig) SaveToFile() {
	// fmt.Println(config)

	jsonData, _ := json.MarshalIndent(config, "", "  ")

	// fmt.Println(string(jsonData))
	configPath := InitConfig(CldConfigOptions{})

	os.WriteFile(configPath, jsonData, os.ModePerm)
	// fmt.Println("test")
	// configPath := InitConfig(CldConfigOptions{})

	// byteData, err := json.Marshal(config)
	// lib.CheckFatalError(err)
	// if _, err := os.Stat(tCacheFile); err != nil {
	// 	os.Create(tCacheFile)
	// }
	// encodedData := b64.StdEncoding.EncodeToString(byteData)
	// os.WriteFile(tCacheFile, []byte(encodedData), os.ModePerm)
	// fmt.Println(encodedData)
}

func (config *CldConfig) AddAzureTenant(tenantId string, tenantName string) {

	for _, t := range config.Azure.Tenants {
		fmt.Println(t.TenantID)
		if t.TenantID == tenantId {
			CheckFatalError(fmt.Errorf("Identical tenant config already exists"))
		}
	}

	if tenantName != "" {
		var newTenant CldConfigTenantAuth
		newTenant.TenantName = tenantName
		newTenant.TenantID = tenantId
		config.Azure.Tenants = append(config.Azure.Tenants, newTenant)
		jsonData, _ := json.MarshalIndent(config, "", "  ")
		fmt.Println(string(jsonData))
		// config.SaveToFile()
	}
}

func (config *CldConfig) UpdateAzureTenantCreds(tenantName string, updateWriter bool, clientId string, clientSecret string) {
	var tenant CldConfigTenantAuth

	if updateWriter {
		tenant.Writer.ClientID = clientId
		tenant.Writer.ClientSecret = clientSecret
	} else {
		tenant.Reader.ClientID = clientId
		tenant.Reader.ClientSecret = clientSecret
	}

}
