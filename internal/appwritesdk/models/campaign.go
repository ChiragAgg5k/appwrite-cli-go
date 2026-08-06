package models

import (
	"encoding/json"
	"errors"
)

// Campaign Model
type Campaign struct {
	// Campaign ID
	Id string `json:"$id"`
	// Campaign template
	Template string `json:"template"`
	// Campaign title
	Title string `json:"title"`
	// Campaign description
	Description string `json:"description"`
	// Billing plan campaign is associated with
	Plan string `json:"plan"`
	// Campaign CTA
	Cta string `json:"cta"`
	// Campaign info when claimed
	Claimed string `json:"claimed"`
	// Campaign infor when unclaimed
	Unclaimed string `json:"unclaimed"`
	// Campaign images
	Image interface{} `json:"image"`
	// Campaign reviews
	Reviews []Review `json:"reviews"`
	// Campaign valid only for new orgs.
	OnlyNewOrgs bool `json:"onlyNewOrgs"`
	// Is footer
	Footer bool `json:"footer"`

	// Used by Decode() method
	data []byte
}

func (model Campaign) New(data []byte) *Campaign {
	model.data = data
	return &model
}

func (model *Campaign) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
