package models

import (
    "encoding/json"
    "errors"
)

// AddonsList Model
type AddonList struct {
    // Total number of addons that matched your query.
    Total int `json:"total"`
    // List of addons.
    Addons []Addon `json:"addons"`

    // Used by Decode() method
    data []byte
}

func (model AddonList) New(data []byte) *AddonList {
    model.data = data
    return &model
}

func (model *AddonList) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}