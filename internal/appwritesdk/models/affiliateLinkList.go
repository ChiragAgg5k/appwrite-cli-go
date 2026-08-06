package models

import (
	"encoding/json"
	"errors"
)

// AffiliateLinksList Model
type AffiliateLinkList struct {
	// Total number of links that matched your query.
	Total int `json:"total"`
	// List of links.
	Links []AffiliateLink `json:"links"`

	// Used by Decode() method
	data []byte
}

func (model AffiliateLinkList) New(data []byte) *AffiliateLinkList {
	model.data = data
	return &model
}

func (model *AffiliateLinkList) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
