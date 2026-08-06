package models

import (
	"encoding/json"
	"errors"
)

// MigrationsList Model
type MigrationList struct {
	// Total number of migrations that matched your query.
	Total int `json:"total"`
	// List of migrations.
	Migrations []Migration `json:"migrations"`

	// Used by Decode() method
	data []byte
}

func (model MigrationList) New(data []byte) *MigrationList {
	model.data = data
	return &model
}

func (model *MigrationList) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
