package models

import (
	"encoding/json"
	"errors"
)

// Region Model
type ConsoleRegion struct {
	// Region ID
	Id string `json:"$id"`
	// Region name
	Name string `json:"name"`
	// Does the organization have access to this region. Null when access has not
	// been resolved.
	Available bool `json:"available"`
	// Does the backend support this region.
	Disabled bool `json:"disabled"`
	// Is this the region default.
	Default bool `json:"default"`
	// Region flag code.
	Flag string `json:"flag"`

	// Used by Decode() method
	data []byte
}

func (model ConsoleRegion) New(data []byte) *ConsoleRegion {
	model.data = data
	return &model
}

func (model *ConsoleRegion) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
