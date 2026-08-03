package models

import (
    "encoding/json"
    "errors"
)

// BillingInvoicesList Model
type InvoiceList struct {
    // Total number of invoices that matched your query.
    Total int `json:"total"`
    // List of invoices.
    Invoices []Invoice `json:"invoices"`

    // Used by Decode() method
    data []byte
}

func (model InvoiceList) New(data []byte) *InvoiceList {
    model.data = data
    return &model
}

func (model *InvoiceList) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}