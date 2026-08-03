package models

import (
    "encoding/json"
    "errors"
)

// Schedule Model
type Schedule struct {
    // Schedule ID.
    Id string `json:"$id"`
    // Schedule creation date in ISO 8601 format.
    CreatedAt string `json:"$createdAt"`
    // Schedule update date in ISO 8601 format.
    UpdatedAt string `json:"$updatedAt"`
    // The resource type associated with this schedule.
    ResourceType string `json:"resourceType"`
    // The resource ID associated with this schedule.
    ResourceId string `json:"resourceId"`
    // Change-tracking timestamp used by the scheduler to detect resource changes
    // in ISO 8601 format.
    ResourceUpdatedAt string `json:"resourceUpdatedAt"`
    // The project ID associated with this schedule.
    ProjectId string `json:"projectId"`
    // The CRON schedule expression.
    Schedule string `json:"schedule"`
    // Schedule data used to store resource-specific context needed for execution.
    Data interface{} `json:"data"`
    // Whether the schedule is active.
    Active bool `json:"active"`
    // The region where the schedule is deployed.
    Region string `json:"region"`

    // Used by Decode() method
    data []byte
}

func (model Schedule) New(data []byte) *Schedule {
    model.data = data
    return &model
}

func (model *Schedule) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}