package models

import (
	"encoding/json"
	"errors"
)

// AffiliateLink Model
type AffiliateLink struct {
	// Link ID. This is the shareable referral code.
	Id string `json:"$id"`
	// Link creation date in ISO 8601 format.
	CreatedAt string `json:"$createdAt"`
	// Link update date in ISO 8601 format.
	UpdatedAt string `json:"$updatedAt"`
	// User ID of the link owner.
	UserId string `json:"userId"`
	// Optional link name.
	Name string `json:"name"`
	// Link status. Can be one of `active` or `disabled`.
	Status string `json:"status"`

	// Used by Decode() method
	data []byte
}

func (model AffiliateLink) New(data []byte) *AffiliateLink {
	model.data = data
	return &model
}

func (model *AffiliateLink) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
