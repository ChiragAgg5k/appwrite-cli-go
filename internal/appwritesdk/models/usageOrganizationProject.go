package models

import (
    "encoding/json"
    "errors"
)

// OrganizationProject Model
type UsageOrganizationProject struct {
    // projectId
    ProjectId string `json:"projectId"`
    // Aggregated stats for number of requests.
    Bandwidth []Metric `json:"bandwidth"`
    // Aggregated stats for consumed bandwidth.
    Users []Metric `json:"users"`
    // Aggregated stats for function executions.
    Executions int `json:"executions"`
    // Aggregated stats for database reads.
    DatabasesReads []Metric `json:"databasesReads"`
    // Aggregated stats for database writes.
    DatabasesWrites []Metric `json:"databasesWrites"`
    // Aggregated stats for function executions in mb seconds.
    ExecutionsMBSeconds int `json:"executionsMBSeconds"`
    // Aggregated stats for function builds in mb seconds.
    BuildsMBSeconds int `json:"buildsMBSeconds"`
    // Aggregated stats for number of documents.
    Storage int `json:"storage"`
    // Aggregated stats for phone authentication.
    AuthPhoneTotal int `json:"authPhoneTotal"`
    // Aggregated stats for phone authentication estimated cost.
    AuthPhoneEstimate float64 `json:"authPhoneEstimate"`
    // Aggregated stats for total databases reads.
    DatabasesReadsTotal int `json:"databasesReadsTotal"`
    // Aggregated stats for total databases writes.
    DatabasesWritesTotal int `json:"databasesWritesTotal"`
    // Aggregated stats for file transformations.
    ImageTransformations []Metric `json:"imageTransformations"`
    // Aggregated stats for total file transformations.
    ImageTransformationsTotal int `json:"imageTransformationsTotal"`
    // Aggregated stats for file transformations.
    ScreenshotsGenerated []Metric `json:"screenshotsGenerated"`
    // Aggregated stats for total file transformations.
    ScreenshotsGeneratedTotal int `json:"screenshotsGeneratedTotal"`
    // Aggregated stats for realtime connections.
    RealtimeConnections int `json:"realtimeConnections"`
    // Aggregated stats for realtime messages.
    RealtimeMessages int `json:"realtimeMessages"`
    // Aggregated stats for realtime bandwidth.
    RealtimeBandwidth int `json:"realtimeBandwidth"`

    // Used by Decode() method
    data []byte
}

func (model UsageOrganizationProject) New(data []byte) *UsageOrganizationProject {
    model.data = data
    return &model
}

func (model *UsageOrganizationProject) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}