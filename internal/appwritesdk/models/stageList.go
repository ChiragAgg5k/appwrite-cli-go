package models

import (
    "encoding/json"
    "errors"
)

// StagesList Model
type StageList struct {
    // List of stages.
    Stages []Stage `json:"stages"`

    // Used by Decode() method
    data []byte
}

func (model StageList) New(data []byte) *StageList {
    model.data = data
    return &model
}

func (model *StageList) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}