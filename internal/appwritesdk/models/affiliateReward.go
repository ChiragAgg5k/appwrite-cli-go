package models

import (
	"encoding/json"
	"errors"
)

// AffiliateReward Model
type AffiliateReward struct {
	// Reward ID.
	Id string `json:"$id"`
	// Reward creation date in ISO 8601 format.
	CreatedAt string `json:"$createdAt"`
	// Reward update date in ISO 8601 format.
	UpdatedAt string `json:"$updatedAt"`
	// User ID of the reward owner.
	UserId string `json:"userId"`
	// Affiliate link ID that earned this reward.
	LinkId string `json:"linkId"`
	// Referral ID that earned this reward.
	ReferralId string `json:"referralId"`
	// Reward amount in USD.
	Amount float64 `json:"amount"`
	// Reward status. Can be one of `pending` or `claimed`.
	Status string `json:"status"`
	// Organization ID the reward was claimed on.
	TeamId string `json:"teamId"`
	// Credit document ID created when the reward was claimed.
	CreditId string `json:"creditId"`

	// Used by Decode() method
	data []byte
}

func (model AffiliateReward) New(data []byte) *AffiliateReward {
	model.data = data
	return &model
}

func (model *AffiliateReward) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
