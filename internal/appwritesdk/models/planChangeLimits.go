package models

import (
	"encoding/json"
	"errors"
)

// PlanChangeLimits Model
type PlanChangeLimits struct {
	// Total number of projects in the organization
	TotalProjects int `json:"totalProjects"`
	// Number of projects exceeding target plan limits
	NonCompliantProjects int `json:"nonCompliantProjects"`
	// Whether the plan change is allowed
	CanChangePlan bool `json:"canChangePlan"`
	// Project compliance details
	Projects []PlanChangeProjectCompliance `json:"projects"`
	// Active addon keys that the target plan does not support. When non-empty,
	// `canChangePlan` is false.
	UnsupportedAddons []string `json:"unsupportedAddons"`

	// Used by Decode() method
	data []byte
}

func (model PlanChangeLimits) New(data []byte) *PlanChangeLimits {
	model.data = data
	return &model
}

func (model *PlanChangeLimits) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
