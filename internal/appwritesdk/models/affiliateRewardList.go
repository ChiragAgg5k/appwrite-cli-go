package models

import (
	"encoding/json"
	"errors"
)

// AffiliateRewardsList Model
type AffiliateRewardList struct {
	// Total number of rewards that matched your query.
	Total int `json:"total"`
	// List of rewards.
	Rewards []AffiliateReward `json:"rewards"`

	// Used by Decode() method
	data []byte
}

func (model AffiliateRewardList) New(data []byte) *AffiliateRewardList {
	model.data = data
	return &model
}

func (model *AffiliateRewardList) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
