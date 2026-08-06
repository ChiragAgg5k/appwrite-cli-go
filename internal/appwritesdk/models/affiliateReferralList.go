package models

import (
	"encoding/json"
	"errors"
)

// AffiliateReferralsList Model
type AffiliateReferralList struct {
	// Total number of referrals that matched your query.
	Total int `json:"total"`
	// List of referrals.
	Referrals []AffiliateReferral `json:"referrals"`

	// Used by Decode() method
	data []byte
}

func (model AffiliateReferralList) New(data []byte) *AffiliateReferralList {
	model.data = data
	return &model
}

func (model *AffiliateReferralList) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
