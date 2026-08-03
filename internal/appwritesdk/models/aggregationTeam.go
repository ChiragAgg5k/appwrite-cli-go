package models

import (
    "encoding/json"
    "errors"
)

// Team Model
type AggregationTeam struct {
    // Aggregation ID.
    Id string `json:"$id"`
    // Aggregation creation time in ISO 8601 format.
    CreatedAt string `json:"$createdAt"`
    // Aggregation update date in ISO 8601 format.
    UpdatedAt string `json:"$updatedAt"`
    // Aggregation permissions. [Learn more about permissions](/docs/permissions).
    Permissions []string `json:"$permissions"`
    // Beginning date of the invoice
    From string `json:"from"`
    // End date of the invoice
    To string `json:"to"`
    // Total storage usage
    UsageStorage int `json:"usageStorage"`
    // Total storage usage with builds storage
    UsageTotalStorage int `json:"usageTotalStorage"`
    // Total files storage usage
    UsageFilesStorage int `json:"usageFilesStorage"`
    // Total deployments storage usage
    UsageDeploymentsStorage int `json:"usageDeploymentsStorage"`
    // Total builds storage usage
    UsageBuildsStorage int `json:"usageBuildsStorage"`
    // Total databases storage usage
    UsageDatabasesStorage int `json:"usageDatabasesStorage"`
    // Total active users for the billing period
    UsageUsers int `json:"usageUsers"`
    // Total number of executions for the billing period
    UsageExecutions int `json:"usageExecutions"`
    // Total bandwidth usage for the billing period
    UsageBandwidth int `json:"usageBandwidth"`
    // Peak concurrent realtime connections for the billing period
    UsageRealtime int `json:"usageRealtime"`
    // Total realtime messages sent for the billing period
    UsageRealtimeMessages int `json:"usageRealtimeMessages"`
    // Total realtime bandwidth usage for the billing period
    UsageRealtimeBandwidth int `json:"usageRealtimeBandwidth"`
    // Additional members
    AdditionalMembers int `json:"additionalMembers"`
    // Additional members cost
    AdditionalMemberAmount int `json:"additionalMemberAmount"`
    // Additional storage usage cost
    AdditionalStorageAmount int `json:"additionalStorageAmount"`
    // Additional users usage cost.
    AdditionalUsersAmount int `json:"additionalUsersAmount"`
    // Additional executions usage cost
    AdditionalExecutionsAmount int `json:"additionalExecutionsAmount"`
    // Additional bandwidth usage cost
    AdditionalBandwidthAmount int `json:"additionalBandwidthAmount"`
    // Additional realtime usage cost
    AdditionalRealtimeAmount int `json:"additionalRealtimeAmount"`
    // Billing plan
    Plan string `json:"plan"`
    // Aggregated amount
    Amount int `json:"amount"`
    // Aggregation project breakdown
    Breakdown []AggregationBreakdown `json:"breakdown"`
    // Usage resources
    Resources []UsageResources `json:"resources"`

    // Used by Decode() method
    data []byte
}

func (model AggregationTeam) New(data []byte) *AggregationTeam {
    model.data = data
    return &model
}

func (model *AggregationTeam) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}