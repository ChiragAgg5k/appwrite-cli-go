package models

import (
	"encoding/json"
	"errors"
)

// EstimationPlanChange Model
type EstimationPlanChange struct {
	// Current billing plan ID
	CurrentBillingPlanId string `json:"currentBillingPlanId"`
	// Target billing plan ID
	TargetBillingPlanId string `json:"targetBillingPlanId"`
	// Direction of plan change: upgrade, downgrade, or same
	Direction string `json:"direction"`
	// Cost estimation details
	Estimation PlanChangeEstimationDetails `json:"estimation"`
	// Plan limits and compliance information
	Limits PlanChangeLimits `json:"limits"`

	// Used by Decode() method
	data []byte
}

func (model EstimationPlanChange) New(data []byte) *EstimationPlanChange {
	model.data = data
	return &model
}

func (model *EstimationPlanChange) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
