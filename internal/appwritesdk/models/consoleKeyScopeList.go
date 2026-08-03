package models

import (
    "encoding/json"
    "errors"
)

// ConsoleKeyScopesList Model
type ConsoleKeyScopeList struct {
    // Total number of key scopes exposed by the server.
    Total int `json:"total"`
    // List of key scopes, each with its ID and description.
    Scopes []ConsoleKeyScope `json:"scopes"`

    // Used by Decode() method
    data []byte
}

func (model ConsoleKeyScopeList) New(data []byte) *ConsoleKeyScopeList {
    model.data = data
    return &model
}

func (model *ConsoleKeyScopeList) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}