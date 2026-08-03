package models

import (
    "encoding/json"
    "errors"
)

// MetricBreakdown Model
type MetricBreakdown struct {
    // Resource ID.
    ResourceId string `json:"resourceId"`
    // Resource name.
    Name string `json:"name"`
    // The value of this metric at the timestamp.
    Value int `json:"value"`
    // The estimated value of this metric at the end of the period.
    Estimate float64 `json:"estimate"`

    // Used by Decode() method
    data []byte
}

func (model MetricBreakdown) New(data []byte) *MetricBreakdown {
    model.data = data
    return &model
}

func (model *MetricBreakdown) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}