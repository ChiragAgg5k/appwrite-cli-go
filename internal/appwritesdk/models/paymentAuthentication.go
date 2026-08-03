package models

import (
    "encoding/json"
    "errors"
)

// PaymentAuthentication Model
type PaymentAuthentication struct {
    // Message for the end user to show on Console.
    Message string `json:"message"`
    // Stripe client secret to use for validation.
    ClientSecret string `json:"clientSecret"`
    // Organization ID for which the payment authentication is needed.
    OrganizationId string `json:"organizationId"`
    // Invoice ID against which the payment needs to be validated.
    InvoiceId string `json:"invoiceId"`
    // Addon ID to use when calling the addon validate endpoint. Empty when
    // authentication is not for an addon.
    AddonId string `json:"addonId"`
    // Project ID for project-level addon payments. Empty for organization-level
    // addons.
    ProjectId string `json:"projectId"`

    // Used by Decode() method
    data []byte
}

func (model PaymentAuthentication) New(data []byte) *PaymentAuthentication {
    model.data = data
    return &model
}

func (model *PaymentAuthentication) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}