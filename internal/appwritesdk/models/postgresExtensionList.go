package models

import (
    "encoding/json"
    "errors"
)

// PostgresExtensionsList Model
type PostgresExtensionList struct {
    // Total number of extensions that matched your query.
    Total int `json:"total"`
    // List of extensions.
    Extensions []PostgresExtension `json:"extensions"`

    // Used by Decode() method
    data []byte
}

func (model PostgresExtensionList) New(data []byte) *PostgresExtensionList {
    model.data = data
    return &model
}

func (model *PostgresExtensionList) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}