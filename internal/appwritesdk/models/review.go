package models

import (
	"encoding/json"
	"errors"
)

// Review Model
type Review struct {
	// Name of user
	Name string `json:"name"`
	// Reviewer image
	Image string `json:"image"`
	// Reviewer description
	Description string `json:"description"`
	// Review
	Review string `json:"review"`

	// Used by Decode() method
	data []byte
}

func (model Review) New(data []byte) *Review {
	model.data = data
	return &model
}

func (model *Review) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
