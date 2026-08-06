package models

import (
	"encoding/json"
	"errors"
)

// PaymentMethodsList Model
type PaymentMethodList struct {
	// Total number of paymentMethods that matched your query.
	Total int `json:"total"`
	// List of paymentMethods.
	PaymentMethods []PaymentMethod `json:"paymentMethods"`

	// Used by Decode() method
	data []byte
}

func (model PaymentMethodList) New(data []byte) *PaymentMethodList {
	model.data = data
	return &model
}

func (model *PaymentMethodList) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
