package models

import (
	"encoding/json"
	"errors"
)

// DomainSuggestion Model
type DomainSuggestion struct {
	// Domain suggestion.
	Domain string `json:"domain"`
	// Is the domain premium?
	Premium bool `json:"premium"`
	// Domain price.
	Price float64 `json:"price"`
	// Is the domain available?
	Available bool `json:"available"`

	// Used by Decode() method
	data []byte
}

func (model DomainSuggestion) New(data []byte) *DomainSuggestion {
	model.data = data
	return &model
}

func (model *DomainSuggestion) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
