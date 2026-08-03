package models

import (
    "encoding/json"
    "errors"
)

// PlanChangeResourceCompliance Model
type PlanChangeResourceCompliance struct {
    // Resource type
    Type string `json:"type"`
    // Current usage count
    CurrentUsage int `json:"currentUsage"`
    // Allowed limit in target plan
    Limit int `json:"limit"`
    // Compliance status
    Status string `json:"status"`
    // Number of resources exceeding the limit
    Excess int `json:"excess"`
    // Suggestion for resolving the compliance issue
    ResolutionHint string `json:"resolutionHint"`

    // Used by Decode() method
    data []byte
}

func (model PlanChangeResourceCompliance) New(data []byte) *PlanChangeResourceCompliance {
    model.data = data
    return &model
}

func (model *PlanChangeResourceCompliance) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}