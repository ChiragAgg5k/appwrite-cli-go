package models

import (
    "encoding/json"
    "errors"
)

// BranchesList Model
type BranchList struct {
    // Total number of branches that matched your query.
    Total int `json:"total"`
    // List of branches.
    Branches []Branch `json:"branches"`

    // Used by Decode() method
    data []byte
}

func (model BranchList) New(data []byte) *BranchList {
    model.data = data
    return &model
}

func (model *BranchList) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}