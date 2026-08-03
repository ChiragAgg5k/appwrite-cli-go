package models

import (
    "encoding/json"
    "errors"
)

// DomainPurchase Model
type DomainPurchase struct {
    // Purchase/invoice ID.
    Id string `json:"$id"`
    // Purchase creation time in ISO 8601 format.
    CreatedAt string `json:"$createdAt"`
    // Purchase update date in ISO 8601 format.
    UpdatedAt string `json:"$updatedAt"`
    // Domain document ID.
    DomainId string `json:"domainId"`
    // Domain name.
    Domain string `json:"domain"`
    // Team ID that owns the domain.
    OrganizationId string `json:"organizationId"`
    // Domain purchase status.
    Status string `json:"status"`
    // Stripe client secret for 3DS; empty when not applicable.
    ClientSecret string `json:"clientSecret"`
    // Purchase amount.
    Amount float64 `json:"amount"`
    // Currency code.
    Currency string `json:"currency"`

    // Used by Decode() method
    data []byte
}

func (model DomainPurchase) New(data []byte) *DomainPurchase {
    model.data = data
    return &model
}

func (model *DomainPurchase) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}