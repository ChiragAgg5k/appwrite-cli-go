package models

import (
    "encoding/json"
    "errors"
)

// Project Model
type UsageProject struct {
    // Total aggregated number of function executions.
    ExecutionsTotal int `json:"executionsTotal"`
    // Total aggregated  number of documents in legacy/tablesdb.
    DocumentsTotal int `json:"documentsTotal"`
    // Total aggregated  number of documents in documentsdb.
    DocumentsdbDocumentsTotal int `json:"documentsdbDocumentsTotal"`
    // Total aggregated  number of rows.
    RowsTotal int `json:"rowsTotal"`
    // Total aggregated number of databases.
    DatabasesTotal int `json:"databasesTotal"`
    // Total aggregated number of documentsdb.
    DocumentsdbTotal int `json:"documentsdbTotal"`
    // Total aggregated sum of databases storage size (in bytes).
    DatabasesStorageTotal int `json:"databasesStorageTotal"`
    // Total aggregated sum of documentsdb databases storage size (in bytes).
    DocumentsdbDatabasesStorageTotal int `json:"documentsdbDatabasesStorageTotal"`
    // Total aggregated number of users.
    UsersTotal int `json:"usersTotal"`
    // Total aggregated sum of files storage size (in bytes).
    FilesStorageTotal int `json:"filesStorageTotal"`
    // Total aggregated sum of functions storage size (in bytes).
    FunctionsStorageTotal int `json:"functionsStorageTotal"`
    // Total aggregated sum of builds storage size (in bytes).
    BuildsStorageTotal int `json:"buildsStorageTotal"`
    // Total aggregated sum of deployments storage size (in bytes).
    DeploymentsStorageTotal int `json:"deploymentsStorageTotal"`
    // Total aggregated number of buckets.
    BucketsTotal int `json:"bucketsTotal"`
    // Total aggregated number of function executions mbSeconds.
    ExecutionsMbSecondsTotal int `json:"executionsMbSecondsTotal"`
    // Total aggregated number of function builds mbSeconds.
    BuildsMbSecondsTotal int `json:"buildsMbSecondsTotal"`
    // Aggregated stats for total databases reads.
    DatabasesReadsTotal int `json:"databasesReadsTotal"`
    // Aggregated stats for total databases writes.
    DatabasesWritesTotal int `json:"databasesWritesTotal"`
    // Total number of documentsdb databases reads.
    DocumentsdbDatabasesReadsTotal int `json:"documentsdbDatabasesReadsTotal"`
    // Total number of documentsdb databases writes.
    DocumentsdbDatabasesWritesTotal int `json:"documentsdbDatabasesWritesTotal"`
    // Aggregated  number of requests per period.
    Requests []Metric `json:"requests"`
    // Aggregated number of consumed bandwidth per period.
    Network []Metric `json:"network"`
    // Aggregated number of users per period.
    Users []Metric `json:"users"`
    // Aggregated number of executions per period.
    Executions []Metric `json:"executions"`
    // Aggregated stats for total auth phone.
    AuthPhoneTotal int `json:"authPhoneTotal"`
    // Aggregated stats for total auth phone estimation.
    AuthPhoneEstimate int `json:"authPhoneEstimate"`
    // Aggregated breakdown in totals of phone auth by country.
    AuthPhoneCountryBreakdown []MetricBreakdown `json:"authPhoneCountryBreakdown"`
    // Aggregated stats for database reads.
    DatabasesReads []Metric `json:"databasesReads"`
    // Aggregated stats for database writes.
    DatabasesWrites []Metric `json:"databasesWrites"`
    // An array of aggregated number of documentsdb database reads.
    DocumentsdbDatabasesReads []Metric `json:"documentsdbDatabasesReads"`
    // An array of aggregated number of documentsdb database writes.
    DocumentsdbDatabasesWrites []Metric `json:"documentsdbDatabasesWrites"`
    // An array of aggregated sum of documentsdb databases storage size (in bytes)
    // per period.
    DocumentsdbDatabasesStorage []Metric `json:"documentsdbDatabasesStorage"`
    // An array of aggregated number of image transformations.
    ImageTransformations []Metric `json:"imageTransformations"`
    // Total aggregated number of image transformations.
    ImageTransformationsTotal int `json:"imageTransformationsTotal"`
    // Total aggregated number of VectorsDB databases.
    VectorsdbDatabasesTotal int `json:"vectorsdbDatabasesTotal"`
    // Total aggregated number of VectorsDB collections.
    VectorsdbCollectionsTotal int `json:"vectorsdbCollectionsTotal"`
    // Total aggregated number of VectorsDB documents.
    VectorsdbDocumentsTotal int `json:"vectorsdbDocumentsTotal"`
    // Total aggregated VectorsDB storage (bytes).
    VectorsdbDatabasesStorageTotal int `json:"vectorsdbDatabasesStorageTotal"`
    // Total aggregated number of VectorsDB reads.
    VectorsdbDatabasesReadsTotal int `json:"vectorsdbDatabasesReadsTotal"`
    // Total aggregated number of VectorsDB writes.
    VectorsdbDatabasesWritesTotal int `json:"vectorsdbDatabasesWritesTotal"`
    // Aggregated VectorsDB databases per period.
    VectorsdbDatabases []Metric `json:"vectorsdbDatabases"`
    // Aggregated VectorsDB collections per period.
    VectorsdbCollections []Metric `json:"vectorsdbCollections"`
    // Aggregated VectorsDB documents per period.
    VectorsdbDocuments []Metric `json:"vectorsdbDocuments"`
    // Aggregated VectorsDB storage per period.
    VectorsdbDatabasesStorage []Metric `json:"vectorsdbDatabasesStorage"`
    // Aggregated VectorsDB reads per period.
    VectorsdbDatabasesReads []Metric `json:"vectorsdbDatabasesReads"`
    // Aggregated VectorsDB writes per period.
    VectorsdbDatabasesWrites []Metric `json:"vectorsdbDatabasesWrites"`
    // Aggregated number of text embedding calls per period.
    EmbeddingsText Metric `json:"embeddingsText"`
    // Aggregated number of tokens processed by text embeddings per period.
    EmbeddingsTextTokens Metric `json:"embeddingsTextTokens"`
    // Aggregated duration spent generating text embeddings per period.
    EmbeddingsTextDuration Metric `json:"embeddingsTextDuration"`
    // Aggregated number of errors while generating text embeddings per period.
    EmbeddingsTextErrors Metric `json:"embeddingsTextErrors"`
    // Total aggregated number of text embedding calls.
    EmbeddingsTextTotal Metric `json:"embeddingsTextTotal"`
    // Total aggregated number of tokens processed by text.
    EmbeddingsTextTokensTotal Metric `json:"embeddingsTextTokensTotal"`
    // Total aggregated duration spent generating text embeddings.
    EmbeddingsTextDurationTotal Metric `json:"embeddingsTextDurationTotal"`
    // Total aggregated number of errors while generating text embeddings.
    EmbeddingsTextErrorsTotal Metric `json:"embeddingsTextErrorsTotal"`
    // Aggregated number of function executions per period.
    FunctionsExecutions []Metric `json:"functionsExecutions"`
    // Total aggregated number of function executions.
    FunctionsExecutionsTotal int `json:"functionsExecutionsTotal"`
    // Aggregated number of site executions per period.
    SitesExecutions []Metric `json:"sitesExecutions"`
    // Total aggregated number of site executions.
    SitesExecutionsTotal int `json:"sitesExecutionsTotal"`
    // Aggregated stats for total network bandwidth.
    NetworkTotal int `json:"networkTotal"`
    // Aggregated stats for total backups storage.
    BackupsStorageTotal int `json:"backupsStorageTotal"`
    // An array of aggregated number of screenshots generated.
    ScreenshotsGenerated []Metric `json:"screenshotsGenerated"`
    // Total aggregated number of screenshots generated.
    ScreenshotsGeneratedTotal int `json:"screenshotsGeneratedTotal"`
    // Current aggregated number of open Realtime connections.
    RealtimeConnectionsTotal int `json:"realtimeConnectionsTotal"`
    // Total number of Realtime messages sent to clients.
    RealtimeMessagesTotal int `json:"realtimeMessagesTotal"`
    // Total consumed Realtime bandwidth (in bytes).
    RealtimeBandwidthTotal int `json:"realtimeBandwidthTotal"`
    // Aggregated number of open Realtime connections per period.
    RealtimeConnections []Metric `json:"realtimeConnections"`
    // Aggregated number of Realtime messages sent to clients per period.
    RealtimeMessages []Metric `json:"realtimeMessages"`
    // Aggregated consumed Realtime bandwidth (in bytes) per period.
    RealtimeBandwidth []Metric `json:"realtimeBandwidth"`

    // Used by Decode() method
    data []byte
}

func (model UsageProject) New(data []byte) *UsageProject {
    model.data = data
    return &model
}

func (model *UsageProject) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}