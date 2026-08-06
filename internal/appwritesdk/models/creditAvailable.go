package models

import (
	"encoding/json"
	"errors"
)

// CreditAvailable Model
type CreditAvailable struct {
	// Total available credits for the organization.
	Available int `json:"available"`

	// Used by Decode() method
	data []byte
}

func (model CreditAvailable) New(data []byte) *CreditAvailable {
	model.data = data
	return &model
}

func (model *CreditAvailable) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
