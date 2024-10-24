package azure

import (
	"time"
)

// Represents a single environment variable with its value
type Variable struct {
    Value string `json:"value"`
}

// Represents the set of variables as a map
type Variables map[string]Variable

// Represents user information
type User struct {
    DisplayName string `json:"displayName"`
    ID          string `json:"id"`
    UniqueName  string `json:"uniqueName"`
}

// Struct representing Azure-specific variable group
type AzureVariableGroup struct {
    Variables                   Variables   `json:"variables"`
    ID                          int         `json:"id"`
    Type                        string      `json:"type"`
    Name                        string      `json:"name"`
    Description                 string      `json:"description"`
    CreatedBy                   User        `json:"createdBy"`
    CreatedOn                   time.Time   `json:"createdOn"`
    ModifiedBy                  User        `json:"modifiedBy"`
    ModifiedOn                  time.Time   `json:"modifiedOn"`
    IsShared                    bool        `json:"isShared"`
    VariableGroupProjectReferences interface{} `json:"variableGroupProjectReferences"`
}

// Struct representing variable groups list
type AzureVariableGroupList []AzureVariableGroup