package models

import (
	"encoding/json"
	"errors"
)

// BlockDelete Model
type BlockDelete struct {
	// Number of blocks deleted
	Deleted int `json:"deleted"`
	// List of deleted blocks
	Blocks []Block `json:"blocks"`

	// Used by Decode() method
	data []byte
}

func (model BlockDelete) New(data []byte) *BlockDelete {
	model.data = data
	return &model
}

func (model *BlockDelete) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
