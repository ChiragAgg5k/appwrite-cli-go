package models

import (
	"encoding/json"
	"errors"
)

// UsageMetric Model
type UsageMetric struct {
	// Metric key this series describes.
	Metric string `json:"metric"`
	// Data points for this metric, ordered by time ascending. With `interval`,
	// each entry is one bucket; without, each entry is one row of the dimensional
	// or aggregate breakdown.
	Points []UsageDataPoint `json:"points"`

	// Used by Decode() method
	data []byte
}

func (model UsageMetric) New(data []byte) *UsageMetric {
	model.data = data
	return &model
}

func (model *UsageMetric) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
