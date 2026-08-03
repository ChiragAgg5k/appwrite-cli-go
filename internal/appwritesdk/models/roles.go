package models

import (
    "encoding/json"
    "errors"
)

// Roles Model
type Roles struct {
    // Array of scopes accessible to current user.
    Scopes []string `json:"scopes"`
    // Array of roles assigned to current user.
    Roles []string `json:"roles"`

    // Used by Decode() method
    data []byte
}

func (model Roles) New(data []byte) *Roles {
    model.data = data
    return &model
}

func (model *Roles) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}