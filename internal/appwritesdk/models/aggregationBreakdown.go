package models

import (
	"encoding/json"
	"errors"
)

// Breakdown Model
type AggregationBreakdown struct {
	// Aggregation ID.
	Id string `json:"$id"`
	// Project name
	Name string `json:"name"`
	// Project region
	Region string `json:"region"`
	// Aggregated amount
	Amount int `json:"amount"`
	//
	Resources []UsageResources `json:"resources"`

	// Used by Decode() method
	data []byte
}

func (model AggregationBreakdown) New(data []byte) *AggregationBreakdown {
	model.data = data
	return &model
}

func (model *AggregationBreakdown) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
