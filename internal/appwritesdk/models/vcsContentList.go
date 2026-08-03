package models

import (
    "encoding/json"
    "errors"
)

// VCSContentList Model
type VcsContentList struct {
    // Total number of contents that matched your query.
    Total int `json:"total"`
    // List of contents.
    Contents []VcsContent `json:"contents"`

    // Used by Decode() method
    data []byte
}

func (model VcsContentList) New(data []byte) *VcsContentList {
    model.data = data
    return &model
}

func (model *VcsContentList) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}