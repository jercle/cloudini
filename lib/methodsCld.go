package lib

import (
	"encoding/json"
	"fmt"
	"os"
)

func (config *CldConfigRoot) Save(options *CldConfigOptions) {
	configFile, _, _ := InitConfig(options)
	jsonBytes, _ := json.MarshalIndent(config, "", "  ")
	os.WriteFile(configFile, jsonBytes, os.ModePerm)
}

func (config *CldConfigRoot) Show() {
	jsonBytes, _ := json.MarshalIndent(config, "", "  ")
	fmt.Println(string(jsonBytes))
}

func (config *CldConfigRoot) Path(options *CldConfigOptions) string {
	_, configPath, _ := InitConfig(options)
	return configPath
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
