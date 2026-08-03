package models

import (
    "encoding/json"
    "errors"
)

// ConsoleKeyScope Model
type ConsoleKeyScope struct {
    // Scope ID.
    Id string `json:"$id"`
    // Scope description.
    Description string `json:"description"`
    // Scope category.
    Category string `json:"category"`
    // Scope is deprecated.
    Deprecated bool `json:"deprecated"`

    // Used by Decode() method
    data []byte
}

func (model ConsoleKeyScope) New(data []byte) *ConsoleKeyScope {
    model.data = data
    return &model
}

func (model *ConsoleKeyScope) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}