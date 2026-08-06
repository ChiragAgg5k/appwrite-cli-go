package models

import (
	"encoding/json"
	"errors"
)

// RuntimeProviderRepositoriesList Model
type ProviderRepositoryRuntimeList struct {
	// Total number of runtimeProviderRepositories that matched your query.
	Total int `json:"total"`
	// List of runtimeProviderRepositories.
	RuntimeProviderRepositories []ProviderRepositoryRuntime `json:"runtimeProviderRepositories"`
	// Provider repository list type.
	Type string `json:"type"`

	// Used by Decode() method
	data []byte
}

func (model ProviderRepositoryRuntimeList) New(data []byte) *ProviderRepositoryRuntimeList {
	model.data = data
	return &model
}

func (model *ProviderRepositoryRuntimeList) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
