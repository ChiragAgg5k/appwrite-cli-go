package models

import (
	"encoding/json"
	"errors"
)

// PaymentMethod Model
type PaymentMethod struct {
	// Payment Method ID.
	Id string `json:"$id"`
	// Payment method creation time in ISO 8601 format.
	CreatedAt string `json:"$createdAt"`
	// Payment method update date in ISO 8601 format.
	UpdatedAt string `json:"$updatedAt"`
	// Payment method permissions. [Learn more about
	// permissions](/docs/permissions).
	Permissions []string `json:"$permissions"`
	// Payment method ID from the payment provider
	ProviderMethodId string `json:"providerMethodId"`
	// Client secret hash for payment setup
	ClientSecret string `json:"clientSecret"`
	// User ID from the payment provider.
	ProviderUserId string `json:"providerUserId"`
	// ID of the Team.
	UserId string `json:"userId"`
	// Expiry month of the payment method.
	ExpiryMonth int `json:"expiryMonth"`
	// Expiry year of the payment method.
	ExpiryYear int `json:"expiryYear"`
	// Last 4 digit of the payment method
	Last4 string `json:"last4"`
	// Payment method brand
	Brand string `json:"brand"`
	// Name of the owner
	Name string `json:"name"`
	// Mandate ID of the payment method
	MandateId string `json:"mandateId"`
	// Country of the payment method
	Country string `json:"country"`
	// State of the payment method
	State string `json:"state"`
	// Last payment error associated with the payment method.
	LastError string `json:"lastError"`
	// True when it's the default payment method.
	Default bool `json:"default"`
	// True when payment method has expired.
	Expired bool `json:"expired"`
	// True when payment method has failed to process multiple times.
	Failed bool `json:"failed"`

	// Used by Decode() method
	data []byte
}

func (model PaymentMethod) New(data []byte) *PaymentMethod {
	model.data = data
	return &model
}

func (model *PaymentMethod) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
