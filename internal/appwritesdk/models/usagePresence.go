package models

import (
    "encoding/json"
    "errors"
)

// UsagePresence Model
type UsagePresence struct {
    // Time range of the usage stats.
    Range string `json:"range"`
    // Current total number of online users.
    UsersOnlineTotal int `json:"usersOnlineTotal"`
    // Aggregated number of online users per period.
    Presences []Metric `json:"presences"`

    // Used by Decode() method
    data []byte
}

func (model UsagePresence) New(data []byte) *UsagePresence {
    model.data = data
    return &model
}

func (model *UsagePresence) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}