package models

import (
    "encoding/json"
    "errors"
)

// BillingAddressList Model
type BillingAddressList struct {
    // Total number of billingAddresses that matched your query.
    Total int `json:"total"`
    // List of billingAddresses.
    BillingAddresses []BillingAddress `json:"billingAddresses"`

    // Used by Decode() method
    data []byte
}

func (model BillingAddressList) New(data []byte) *BillingAddressList {
    model.data = data
    return &model
}

func (model *BillingAddressList) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}