package models

import (
	"encoding/json"
	"errors"
)

// OrganizationsList Model
type OrganizationList struct {
	// Total number of teams that matched your query.
	Total int `json:"total"`
	// List of teams.
	Teams []Organization `json:"teams"`

	// Used by Decode() method
	data []byte
}

func (model OrganizationList) New(data []byte) *OrganizationList {
	model.data = data
	return &model
}

func (model *OrganizationList) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
