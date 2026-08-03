package models

import (
    "encoding/json"
    "errors"
)

// AggregationTeamList Model
type AggregationTeamList struct {
    // Total number of aggregations that matched your query.
    Total int `json:"total"`
    // List of aggregations.
    Aggregations []AggregationTeam `json:"aggregations"`

    // Used by Decode() method
    data []byte
}

func (model AggregationTeamList) New(data []byte) *AggregationTeamList {
    model.data = data
    return &model
}

func (model *AggregationTeamList) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}