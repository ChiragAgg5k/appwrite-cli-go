package models

import (
    "encoding/json"
    "errors"
)

// UsageUsers Model
type UsageUsers struct {
    // Time range of the usage stats.
    Range string `json:"range"`
    // Total aggregated number of statistics of users.
    UsersTotal int `json:"usersTotal"`
    // Total aggregated number of active sessions.
    SessionsTotal int `json:"sessionsTotal"`
    // Aggregated number of users per period.
    Users []Metric `json:"users"`
    // Aggregated number of active sessions  per period.
    Sessions []Metric `json:"sessions"`

    // Used by Decode() method
    data []byte
}

func (model UsageUsers) New(data []byte) *UsageUsers {
    model.data = data
    return &model
}

func (model *UsageUsers) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}