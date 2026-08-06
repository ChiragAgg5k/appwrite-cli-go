package models

import (
	"encoding/json"
	"errors"
)

// DetectionVariable Model
type DetectionVariable struct {
	// Name of environment variable
	Name string `json:"name"`
	// Value of environment variable
	Value string `json:"value"`

	// Used by Decode() method
	data []byte
}

func (model DetectionVariable) New(data []byte) *DetectionVariable {
	model.data = data
	return &model
}

func (model *DetectionVariable) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
