package models

import (
    "encoding/json"
    "errors"
)

// UpdatePlan Model
type EstimationUpdatePlan struct {
    // Total amount
    Amount float64 `json:"amount"`
    // Gross payable amount
    GrossAmount float64 `json:"grossAmount"`
    // Discount amount
    Discount float64 `json:"discount"`
    // Credits amount
    Credits float64 `json:"credits"`
    // Estimation items
    Items []EstimationItem `json:"items"`
    // Estimation discount items
    Discounts []EstimationItem `json:"discounts"`
    // Trial days
    TrialDays int `json:"trialDays"`
    // Trial end date
    TrialEndDate string `json:"trialEndDate"`
    // Organization's existing credits
    OrganizationCredits float64 `json:"organizationCredits"`

    // Used by Decode() method
    data []byte
}

func (model EstimationUpdatePlan) New(data []byte) *EstimationUpdatePlan {
    model.data = data
    return &model
}

func (model *EstimationUpdatePlan) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}