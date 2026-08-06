package models

import (
	"encoding/json"
	"errors"
)

// AddonPrice Model
type AddonPrice struct {
	// Addon key.
	AddonKey string `json:"addonKey"`
	// Addon display name.
	Name string `json:"name"`
	// Full monthly price of the addon.
	MonthlyPrice float64 `json:"monthlyPrice"`
	// Calculated prorated amount for the current billing cycle.
	ProratedAmount float64 `json:"proratedAmount"`
	// Days remaining in the current billing cycle.
	RemainingDays int `json:"remainingDays"`
	// Total days in the billing cycle.
	TotalCycleDays int `json:"totalCycleDays"`
	// Currency code.
	Currency string `json:"currency"`
	// When the current billing cycle ends.
	BillingCycleEnd string `json:"billingCycleEnd"`

	// Used by Decode() method
	data []byte
}

func (model AddonPrice) New(data []byte) *AddonPrice {
	model.data = data
	return &model
}

func (model *AddonPrice) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
