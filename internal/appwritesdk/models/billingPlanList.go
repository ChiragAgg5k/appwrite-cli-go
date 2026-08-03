package models

import (
    "encoding/json"
    "errors"
)

// BillingPlanList Model
type BillingPlanList struct {
    // Total number of plans that matched your query.
    Total int `json:"total"`
    // List of plans.
    Plans []BillingPlan `json:"plans"`

    // Used by Decode() method
    data []byte
}

func (model BillingPlanList) New(data []byte) *BillingPlanList {
    model.data = data
    return &model
}

func (model *BillingPlanList) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}