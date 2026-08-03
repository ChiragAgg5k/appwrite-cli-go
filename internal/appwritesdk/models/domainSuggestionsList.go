package models

import (
    "encoding/json"
    "errors"
)

// DomainSuggestionsList Model
type DomainSuggestionsList struct {
    // Total number of suggestions that matched your query.
    Total int `json:"total"`
    // List of suggestions.
    Suggestions []DomainSuggestion `json:"suggestions"`

    // Used by Decode() method
    data []byte
}

func (model DomainSuggestionsList) New(data []byte) *DomainSuggestionsList {
    model.data = data
    return &model
}

func (model *DomainSuggestionsList) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}