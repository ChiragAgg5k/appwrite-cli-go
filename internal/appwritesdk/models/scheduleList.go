package models

import (
	"encoding/json"
	"errors"
)

// SchedulesList Model
type ScheduleList struct {
	// Total number of schedules that matched your query.
	Total int `json:"total"`
	// List of schedules.
	Schedules []Schedule `json:"schedules"`

	// Used by Decode() method
	data []byte
}

func (model ScheduleList) New(data []byte) *ScheduleList {
	model.data = data
	return &model
}

func (model *ScheduleList) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
