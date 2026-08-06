package models

import (
	"encoding/json"
	"errors"
)

// UsageEventList Model
type UsageEventList struct {
	// Time interval size (1h or 1d). Empty when the request omits `interval` —
	// points then carry the request end time as their as-of marker.
	Interval string `json:"interval"`
	// One entry per requested metric, each carrying its own points[] time series
	// (sums per bucket / dimension over time).
	Metrics []UsageMetric `json:"metrics"`

	// Used by Decode() method
	data []byte
}

func (model UsageEventList) New(data []byte) *UsageEventList {
	model.data = data
	return &model
}

func (model *UsageEventList) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
