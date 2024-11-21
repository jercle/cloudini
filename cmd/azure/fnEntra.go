/*
Copyright © 2024 Evan Colwell ercolwell@gmail.com
*/
package azure

import (
	"encoding/json"

	"github.com/jercle/cloudini/lib"
)

type EntraRoleDefinition struct {
	Description     string `json:"description"`
	DisplayName     string `json:"displayName"`
	ID              string `json:"id"`
	IsBuiltIn       bool   `json:"isBuiltIn"`
	IsEnabled       bool   `json:"isEnabled"`
	RolePermissions []struct {
		AllowedResourceActions []string `json:"allowedResourceActions"`
		Condition              *string  `json:"condition"`
	} `json:"rolePermissions"`
}

type ListEntraRoleDefinitionsResponse struct {
	Odata_Context string                `json:"@odata.context"`
	Value         []EntraRoleDefinition `json:"value"`
}

func ListEntraRoleDefinitions(mat lib.AzureMultiAuthToken) ([]EntraRoleDefinition, error) {
	var (
		unmarshResponse ListEntraRoleDefinitionsResponse
		roleDefs        []EntraRoleDefinition
	)

	urlString := "https://graph.microsoft.com/v1.0/roleManagement/directory/roleDefinitions"

	response, err := HttpGet(urlString, mat)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(response, &unmarshResponse)
	if err != nil {
		return nil, err
	}

	jsonData, _ := json.Marshal(unmarshResponse.Value)

	err = json.Unmarshal(jsonData, &roleDefs)
	if err != nil {
		return nil, err
	}

	return roleDefs, nil
}
