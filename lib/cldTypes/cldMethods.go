package cldTypes

import (
	"encoding/json"
	"fmt"
)

func (config *CldConfig) AddAzureTenant(tenantId string, tenantName string) {

	for _, t := range config.Azure.MultiTenantAuth.Tenants {
		fmt.Println(t.TenantID)
		if t.TenantID == tenantId {
			CheckFatalError(fmt.Errorf("Identical tenant config already exists"))
		}
	}

	if tenantName != "" {
		var newTenant CldConfigTenantAuth
		newTenant.TenantName = tenantName
		newTenant.TenantID = tenantId
		config.Azure.MultiTenantAuth.Tenants[newTenant.TenantName] = newTenant
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
