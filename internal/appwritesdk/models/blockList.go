package models

import (
    "encoding/json"
    "errors"
)

// BlocksList Model
type BlockList struct {
    // Total number of blocks that matched your query.
    Total int `json:"total"`
    // List of blocks.
    Blocks []Block `json:"blocks"`

    // Used by Decode() method
    data []byte
}

func (model BlockList) New(data []byte) *BlockList {
    model.data = data
    return &model
}

func (model *BlockList) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}