package models

import (
    "encoding/json"
    "errors"
)

// Notification Model
type Notification struct {
    // Notification ID.
    Id string `json:"$id"`
    // Notification creation date in ISO 8601 format.
    CreatedAt string `json:"$createdAt"`
    // Notification update date in ISO 8601 format.
    UpdatedAt string `json:"$updatedAt"`
    // Stable message ID used for dedup.
    MessageId string `json:"messageId"`
    // Notification type: info, warning, error.
    Type string `json:"type"`
    // Channel: email, sms, push, console, webhook.
    Channel string `json:"channel"`
    // Resource type this notification is addressed to.
    ResourceType string `json:"resourceType"`
    // Resource ID this notification is addressed to.
    ResourceId string `json:"resourceId"`
    // Parent resource type for the notification.
    ParentResourceType string `json:"parentResourceType"`
    // Parent resource ID for the notification.
    ParentResourceId string `json:"parentResourceId"`
    // Project the notification pertains to.
    ProjectId string `json:"projectId"`
    // Notification title.
    Title string `json:"title"`
    // Notification body.
    Body string `json:"body"`
    // Whether the notification has been read.
    Read bool `json:"read"`
    // First time the notification was viewed from a notification logo.
    FirstSeen string `json:"firstSeen"`
    // Most recent time the notification was viewed from a notification logo.
    LastSeen string `json:"lastSeen"`

    // Used by Decode() method
    data []byte
}

func (model Notification) New(data []byte) *Notification {
    model.data = data
    return &model
}

func (model *Notification) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}