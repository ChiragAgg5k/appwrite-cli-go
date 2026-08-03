package models

import (
    "encoding/json"
    "errors"
)

// FrameworkProviderRepositoriesList Model
type ProviderRepositoryFrameworkList struct {
    // Total number of frameworkProviderRepositories that matched your query.
    Total int `json:"total"`
    // List of frameworkProviderRepositories.
    FrameworkProviderRepositories []ProviderRepositoryFramework `json:"frameworkProviderRepositories"`
    // Provider repository list type.
    Type string `json:"type"`

    // Used by Decode() method
    data []byte
}

func (model ProviderRepositoryFrameworkList) New(data []byte) *ProviderRepositoryFrameworkList {
    model.data = data
    return &model
}

func (model *ProviderRepositoryFrameworkList) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}