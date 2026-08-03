package models

import (
    "encoding/json"
    "errors"
)

// DevKeysList Model
type DevKeyList struct {
    // Total number of devKeys that matched your query.
    Total int `json:"total"`
    // List of devKeys.
    DevKeys []DevKey `json:"devKeys"`

    // Used by Decode() method
    data []byte
}

func (model DevKeyList) New(data []byte) *DevKeyList {
    model.data = data
    return &model
}

func (model *DevKeyList) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}