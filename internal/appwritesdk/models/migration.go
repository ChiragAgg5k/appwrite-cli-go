package models

import (
    "encoding/json"
    "errors"
)

// Migration Model
type Migration struct {
    // Migration ID.
    Id string `json:"$id"`
    // Migration creation date in ISO 8601 format.
    CreatedAt string `json:"$createdAt"`
    // Variable creation date in ISO 8601 format.
    UpdatedAt string `json:"$updatedAt"`
    // Migration status ( pending, processing, failed, completed )
    Status string `json:"status"`
    // Migration stage ( init, processing, source-check, destination-check,
    // migrating, finished )
    Stage string `json:"stage"`
    // A string containing the type of source of the migration.
    Source string `json:"source"`
    // A string containing the type of destination of the migration.
    Destination string `json:"destination"`
    // Resources to migrate.
    Resources []string `json:"resources"`
    // ID of the resource being migrated.
    ResourceId string `json:"resourceId"`
    // Internal ID of the resource being migrated.
    ResourceInternalId string `json:"resourceInternalId"`
    // Type of the resource being migrated.
    ResourceType string `json:"resourceType"`
    // ID of the parent resource that contains the migrated resource.
    ParentResourceId string `json:"parentResourceId"`
    // Internal ID of the parent resource that contains the migrated resource.
    ParentResourceInternalId string `json:"parentResourceInternalId"`
    // Type of the parent resource that contains the migrated resource.
    ParentResourceType string `json:"parentResourceType"`
    // ID of the destination resource created or overwritten by the migration.
    DestinationResourceId string `json:"destinationResourceId"`
    // Internal ID of the destination resource created or overwritten by the
    // migration.
    DestinationResourceInternalId string `json:"destinationResourceInternalId"`
    // Type of the destination resource created or overwritten by the migration.
    DestinationResourceType string `json:"destinationResourceType"`
    // A group of counters that represent the total progress of the migration.
    StatusCounters interface{} `json:"statusCounters"`
    // An array of objects containing the report data of the resources that were
    // migrated.
    ResourceData interface{} `json:"resourceData"`
    // All errors that occurred during the migration process.
    Errors []string `json:"errors"`
    // Migration options used during the migration process.
    Options interface{} `json:"options"`

    // Used by Decode() method
    data []byte
}

func (model Migration) New(data []byte) *Migration {
    model.data = data
    return &model
}

func (model *Migration) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}