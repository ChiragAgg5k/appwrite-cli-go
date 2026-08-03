package models

import (
    "encoding/json"
    "errors"
)

// PlanChangeProjectCompliance Model
type PlanChangeProjectCompliance struct {
    // Project ID
    Id string `json:"$id"`
    // Project name
    Name string `json:"name"`
    // Whether the project complies with target plan limits
    IsCompliant bool `json:"isCompliant"`
    // Resource compliance details
    Resources []PlanChangeResourceCompliance `json:"resources"`
    // Failure reason when compliance could not be evaluated. Present only when
    // the project DB or Regions API was unreachable; in that case `isCompliant`
    // is false (fail closed) and `resources` is empty.
    Error string `json:"error"`

    // Used by Decode() method
    data []byte
}

func (model PlanChangeProjectCompliance) New(data []byte) *PlanChangeProjectCompliance {
    model.data = data
    return &model
}

func (model *PlanChangeProjectCompliance) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}