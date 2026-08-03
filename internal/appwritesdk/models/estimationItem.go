package models

import (
    "encoding/json"
    "errors"
)

// Item Model
type EstimationItem struct {
    // Label
    Label string `json:"label"`
    // Gross payable amount
    Value float64 `json:"value"`

    // Used by Decode() method
    data []byte
}

func (model EstimationItem) New(data []byte) *EstimationItem {
    model.data = data
    return &model
}

func (model *EstimationItem) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}