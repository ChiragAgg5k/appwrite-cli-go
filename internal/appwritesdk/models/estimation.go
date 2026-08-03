package models

import (
    "encoding/json"
    "errors"
)

// Estimation Model
type Estimation struct {
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

    // Used by Decode() method
    data []byte
}

func (model Estimation) New(data []byte) *Estimation {
    model.data = data
    return &model
}

func (model *Estimation) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}