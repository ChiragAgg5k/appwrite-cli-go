package models

import (
    "encoding/json"
    "errors"
)

// Stage Model
type Stage struct {
    // Stage ID.
    Id string `json:"id"`
    // SDK method key (namespace.name) for this stage.
    Sdk string `json:"sdk"`
    // Stage status.
    Status string `json:"status"`
    // When the stage was completed or skipped, in ISO 8601 format.
    At string `json:"at"`
    // Actor type when the stage was recorded.
    ActorType string `json:"actorType"`

    // Used by Decode() method
    data []byte
}

func (model Stage) New(data []byte) *Stage {
    model.data = data
    return &model
}

func (model *Stage) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}