package models

import (
	"encoding/json"
	"errors"
)

// Branch Model
type Branch struct {
	// Branch Name.
	Name string `json:"name"`

	// Used by Decode() method
	data []byte
}

func (model Branch) New(data []byte) *Branch {
	model.data = data
	return &model
}

func (model *Branch) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
