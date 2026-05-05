package azure

import "time"

type RoleAssignmentListResponse struct {
	Value []RoleAssignment `json:"value,omitempty,omitzero" bson:"value,omitempty,omitzero"`
}

//
//

type RoleAssignment struct {
	ID         string `json:"id,omitempty,omitzero" bson:"_id,omitempty,omitzero"`
	Name       string `json:"name,omitempty,omitzero" bson:"name,omitempty,omitzero"`
	Properties struct {
		Condition                          *string   `json:"condition,omitempty,omitzero" bson:"condition,omitempty,omitzero"`
		ConditionVersion                   *string   `json:"conditionVersion,omitempty,omitzero" bson:"conditionVersion,omitempty,omitzero"`
		CreatedBy                          string    `json:"createdBy,omitempty,omitzero" bson:"createdBy,omitempty,omitzero"`
		CreatedOn                          time.Time `json:"createdOn,omitempty,omitzero" bson:"createdOn,omitempty,omitzero"`
		DelegatedManagedIdentityResourceID any       `json:"delegatedManagedIdentityResourceId,omitempty,omitzero" bson:"delegatedManagedIdentityResourceId,omitempty,omitzero"`
		Description                        *string   `json:"description,omitempty,omitzero" bson:"description,omitempty,omitzero"`
		PrincipalID                        string    `json:"principalId,omitempty,omitzero" bson:"principalId,omitempty,omitzero"`
		PrincipalType                      string    `json:"principalType,omitempty,omitzero" bson:"principalType,omitempty,omitzero"`
		RoleDefinitionID                   string    `json:"roleDefinitionId,omitempty,omitzero" bson:"roleDefinitionId,omitempty,omitzero"`
		Scope                              string    `json:"scope,omitempty,omitzero" bson:"scope,omitempty,omitzero"`
		UpdatedBy                          string    `json:"updatedBy,omitempty,omitzero" bson:"updatedBy,omitempty,omitzero"`
		UpdatedOn                          time.Time `json:"updatedOn,omitempty,omitzero" bson:"updatedOn,omitempty,omitzero"`
	} `json:"properties,omitempty,omitzero" bson:"properties,omitempty,omitzero"`
	Type      string `json:"type,omitempty,omitzero" bson:"type,omitempty,omitzero"`
	ScopeType string `json:"scopeType,omitempty" bson:"scopeType,omitempty"`
}
