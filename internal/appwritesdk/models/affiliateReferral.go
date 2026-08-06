package models

import (
	"encoding/json"
	"errors"
)

// AffiliateReferral Model
type AffiliateReferral struct {
	// Referral ID.
	Id string `json:"$id"`
	// Referral creation date in ISO 8601 format.
	CreatedAt string `json:"$createdAt"`
	// Referral update date in ISO 8601 format.
	UpdatedAt string `json:"$updatedAt"`
	// Affiliate link ID used for attribution.
	LinkId string `json:"linkId"`
	// Privacy-safe truncated referred user ID.
	ReferredUserMaskedId string `json:"referredUserMaskedId"`
	// ISO 3166-1 alpha-2 country code of the referred user at signup, when
	// available.
	ReferredUserCountry string `json:"referredUserCountry"`
	// Referral status. Can be one of `pending`, `converted`, or `expired`.
	// `expired` is derived from `expiresAt` when still pending.
	Status string `json:"status"`
	// Attribution time in ISO 8601 format.
	AttributedAt string `json:"attributedAt"`
	// Attribution expiry time in ISO 8601 format.
	ExpiresAt string `json:"expiresAt"`
	// Conversion time in ISO 8601 format.
	ConvertedAt string `json:"convertedAt"`

	// Used by Decode() method
	data []byte
}

func (model AffiliateReferral) New(data []byte) *AffiliateReferral {
	model.data = data
	return &model
}

func (model *AffiliateReferral) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
