package lib

import (
	"encoding/json"
	"fmt"
	"os"
)

func (config *CldConfigRoot) Save(options *CldConfigOptions) {
	configPath := InitConfig(options)

	jsonBytes, _ := json.MarshalIndent(config, "", "  ")

	os.WriteFile(configPath, jsonBytes, os.ModePerm)
	// fmt.Println(string(jsonBytes))
	// configPath := InitConfig(CldConfigOptions{})

	// byteData, err := json.Marshal(config)
	// lib.CheckFatalError(err)
	// if _, err := os.Stat(tCacheFile); err != nil {
	// 	os.Create(tCacheFile)
	// }
	// fmt.Println(jsonBytes)
	// encodedData := b64.StdEncoding.EncodeToString(jsonBytes)
	// os.WriteFile(tCacheFile, []byte(encodedData), os.ModePerm)
	// fmt.Println(encodedData)
	// jsonBytes, _ := json.MarshalIndent(config, "", "  ")
	// os.WriteFile(configPath, jsonBytes, os.ModePerm)
}

func (config *CldConfigRoot) Show() {
	jsonBytes, _ := json.MarshalIndent(config, "", "  ")
	fmt.Println(string(jsonBytes))
}

func (tenants *CldConfigTenants) AddOrUpdateTenant(tenant CldConfigTenantAuth) {
	if *tenants == nil {
		*tenants = make(CldConfigTenants)
	}
	(*tenants)[tenant.TenantName] = tenant

}

func (tenants *CldConfigTenants) RemoveTenant(tenantName string) {
	delete(*tenants, tenantName)
}
