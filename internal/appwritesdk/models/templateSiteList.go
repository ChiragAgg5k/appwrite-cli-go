package models

import (
	"encoding/json"
	"errors"
)

// SiteTemplatesList Model
type TemplateSiteList struct {
	// Total number of templates that matched your query.
	Total int `json:"total"`
	// List of templates.
	Templates []TemplateSite `json:"templates"`

	// Used by Decode() method
	data []byte
}

func (model TemplateSiteList) New(data []byte) *TemplateSiteList {
	model.data = data
	return &model
}

func (model *TemplateSiteList) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
