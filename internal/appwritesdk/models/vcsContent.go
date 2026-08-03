package models

import (
    "encoding/json"
    "errors"
)

// VcsContents Model
type VcsContent struct {
    // Content size in bytes. Only files have size, and for directories, 0 is
    // returned.
    Size int `json:"size"`
    // If a content is a directory. Directories can be used to check nested
    // contents.
    IsDirectory bool `json:"isDirectory"`
    // Name of directory or file.
    Name string `json:"name"`

    // Used by Decode() method
    data []byte
}

func (model VcsContent) New(data []byte) *VcsContent {
    model.data = data
    return &model
}

func (model *VcsContent) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}