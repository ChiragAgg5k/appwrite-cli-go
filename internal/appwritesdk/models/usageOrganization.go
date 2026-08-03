package models

import (
    "encoding/json"
    "errors"
)

// Organization Model
type UsageOrganization struct {
    // Aggregated stats for number of requests.
    Bandwidth []Metric `json:"bandwidth"`
    // Aggregated stats for consumed bandwidth.
    Users []Metric `json:"users"`
    // Aggregated stats for function executions.
    Executions []Metric `json:"executions"`
    // Aggregated stats for database reads.
    DatabasesReads []Metric `json:"databasesReads"`
    // Aggregated stats for database writes.
    DatabasesWrites []Metric `json:"databasesWrites"`
    // Aggregated stats for file transformations.
    ImageTransformations []Metric `json:"imageTransformations"`
    // Aggregated stats for total file transformations.
    ImageTransformationsTotal int `json:"imageTransformationsTotal"`
    // Aggregated stats for file transformations.
    ScreenshotsGenerated []Metric `json:"screenshotsGenerated"`
    // Aggregated stats for total file transformations.
    ScreenshotsGeneratedTotal int `json:"screenshotsGeneratedTotal"`
    // Aggregated stats for total users.
    UsersTotal int `json:"usersTotal"`
    // Aggregated stats for total executions.
    ExecutionsTotal int `json:"executionsTotal"`
    // Aggregated stats for function executions in mb seconds.
    ExecutionsMBSecondsTotal int `json:"executionsMBSecondsTotal"`
    // Aggregated stats for function builds in mb seconds.
    BuildsMBSecondsTotal int `json:"buildsMBSecondsTotal"`
    // Aggregated stats for total file storage.
    FilesStorageTotal int `json:"filesStorageTotal"`
    // Aggregated stats for total builds storage.
    BuildsStorageTotal int `json:"buildsStorageTotal"`
    // Aggregated stats for total deployments storage.
    DeploymentsStorageTotal int `json:"deploymentsStorageTotal"`
    // Aggregated stats for total databases storage.
    DatabasesStorageTotal int `json:"databasesStorageTotal"`
    // Aggregated stats for total databases  reads.
    DatabasesReadsTotal int `json:"databasesReadsTotal"`
    // Aggregated stats for total databases  writes.
    DatabasesWritesTotal int `json:"databasesWritesTotal"`
    // Aggregated stats for total backups storage.
    BackupsStorageTotal int `json:"backupsStorageTotal"`
    // Aggregated stats for total storage.
    StorageTotal int `json:"storageTotal"`
    // Aggregated stats for total auth phone.
    AuthPhoneTotal int `json:"authPhoneTotal"`
    // Aggregated stats for total auth phone estimation.
    AuthPhoneEstimate int `json:"authPhoneEstimate"`
    // Aggregated stats for each projects.
    Projects []UsageOrganizationProject `json:"projects"`
    // Aggregated stats for realtime connections.
    RealtimeConnections []Metric `json:"realtimeConnections"`
    // Aggregated stats for total realtime connections.
    RealtimeConnectionsTotal int `json:"realtimeConnectionsTotal"`
    // Aggregated stats for realtime messages.
    RealtimeMessages []Metric `json:"realtimeMessages"`
    // Aggregated stats for total realtime messages.
    RealtimeMessagesTotal int `json:"realtimeMessagesTotal"`
    // Aggregated stats for realtime bandwidth.
    RealtimeBandwidth []Metric `json:"realtimeBandwidth"`
    // Aggregated stats for total realtime bandwidth.
    RealtimeBandwidthTotal int `json:"realtimeBandwidthTotal"`

    // Used by Decode() method
    data []byte
}

func (model UsageOrganization) New(data []byte) *UsageOrganization {
    model.data = data
    return &model
}

func (model *UsageOrganization) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}