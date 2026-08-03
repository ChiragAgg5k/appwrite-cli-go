package models

import (
    "encoding/json"
    "errors"
)

// DetectionRuntime Model
type DetectionRuntime struct {
    // Repository detection type.
    Type string `json:"type"`
    // Environment variables found in .env files
    Variables []DetectionVariable `json:"variables"`
    // Runtime
    Runtime string `json:"runtime"`
    // Function Entrypoint
    Entrypoint string `json:"entrypoint"`
    // Function install and build commands
    Commands string `json:"commands"`

    // Used by Decode() method
    data []byte
}

func (model DetectionRuntime) New(data []byte) *DetectionRuntime {
    model.data = data
    return &model
}

func (model *DetectionRuntime) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}