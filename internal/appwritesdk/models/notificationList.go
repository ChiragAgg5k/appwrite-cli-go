package models

import (
    "encoding/json"
    "errors"
)

// NotificationsList Model
type NotificationList struct {
    // Total number of notifications that matched your query.
    Total int `json:"total"`
    // List of notifications.
    Notifications []Notification `json:"notifications"`

    // Used by Decode() method
    data []byte
}

func (model NotificationList) New(data []byte) *NotificationList {
    model.data = data
    return &model
}

func (model *NotificationList) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}