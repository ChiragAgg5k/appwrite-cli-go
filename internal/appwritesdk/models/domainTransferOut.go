package models

import (
	"encoding/json"
	"errors"
)

// DomainTransferOut Model
type DomainTransferOut struct {
	// Domain transfer authorization code.
	AuthCode string `json:"authCode"`

	// Used by Decode() method
	data []byte
}

func (model DomainTransferOut) New(data []byte) *DomainTransferOut {
	model.data = data
	return &model
}

func (model *DomainTransferOut) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
