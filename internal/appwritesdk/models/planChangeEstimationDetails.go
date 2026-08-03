package models

import (
    "encoding/json"
    "errors"
)

// PlanChangeEstimationDetails Model
type PlanChangeEstimationDetails struct {
    // Currency code
    Currency string `json:"currency"`
    // Gross amount after all discounts and credits
    GrossAmount float64 `json:"grossAmount"`
    // Credits applied from coupon
    Credits float64 `json:"credits"`
    // Organization's existing credits applied
    OrganizationCredits float64 `json:"organizationCredits"`
    // Discount amount from prorated invoices
    Discount float64 `json:"discount"`
    // Total amount before discounts and credits
    Amount float64 `json:"amount"`
    // Next invoice date
    NextInvoiceDate string `json:"nextInvoiceDate"`
    // Line items breakdown
    Items interface{} `json:"items"`
    // Applied discounts breakdown
    Discounts interface{} `json:"discounts"`

    // Used by Decode() method
    data []byte
}

func (model PlanChangeEstimationDetails) New(data []byte) *PlanChangeEstimationDetails {
    model.data = data
    return &model
}

func (model *PlanChangeEstimationDetails) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}