package models

import (
    "encoding/json"
    "errors"
)

// Address Model
type BillingAddress struct {
    // Region ID
    Id string `json:"$id"`
    // User ID
    UserId string `json:"userId"`
    // Street address
    StreetAddress string `json:"streetAddress"`
    // Address line 2
    AddressLine2 string `json:"addressLine2"`
    // Address country
    Country string `json:"country"`
    // city
    City string `json:"city"`
    // state
    State string `json:"state"`
    // postal code
    PostalCode string `json:"postalCode"`

    // Used by Decode() method
    data []byte
}

func (model BillingAddress) New(data []byte) *BillingAddress {
    model.data = data
    return &model
}

func (model *BillingAddress) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}