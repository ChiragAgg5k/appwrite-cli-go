package models

import (
    "encoding/json"
    "errors"
)

// UsageDataPoint Model
type UsageDataPoint struct {
    // Bucket start timestamp (ISO 8601). When `interval` is omitted this is the
    // request end time, marking the aggregate as-of moment.
    Time string `json:"time"`
    // Aggregated value for the bucket. Counters are whole numbers; gauge rates
    // (e.g. queries per second) may be fractional.
    Value float64 `json:"value"`
    // API endpoint path when broken down by `path`.
    Path string `json:"path"`
    // HTTP method when broken down by `method`.
    Method string `json:"method"`
    // HTTP status code when broken down by `status`.
    Status string `json:"status"`
    // API service segment when broken down by `service`.
    Service string `json:"service"`
    // Country code when broken down by `country`.
    Country string `json:"country"`
    // Continent code when broken down by `continentCode`.
    ContinentCode string `json:"continentCode"`
    // City name when broken down by `city`.
    City string `json:"city"`
    // Region/state chain when broken down by `subdivisions`.
    Subdivisions string `json:"subdivisions"`
    // Internet service provider when broken down by `isp`.
    Isp string `json:"isp"`
    // Autonomous System Number (ASN) when broken down by
    // `autonomousSystemNumber`.
    AutonomousSystemNumber string `json:"autonomousSystemNumber"`
    // Organization owning the ASN when broken down by
    // `autonomousSystemOrganization`.
    AutonomousSystemOrganization string `json:"autonomousSystemOrganization"`
    // Connection type (e.g. cable, cellular, corporate) when broken down by
    // `connectionType`.
    ConnectionType string `json:"connectionType"`
    // User type (e.g. residential, business, hosting) when broken down by
    // `connectionUsageType`.
    ConnectionUsageType string `json:"connectionUsageType"`
    // Registered organization of the IP when broken down by
    // `connectionOrganization`.
    ConnectionOrganization string `json:"connectionOrganization"`
    // Appwrite region when broken down by `region`.
    Region string `json:"region"`
    // Caller origin hostname when broken down by `hostname`.
    Hostname string `json:"hostname"`
    // Caller IP address when broken down by `ip`.
    Ip string `json:"ip"`
    // Operating system name when broken down by `osName`.
    OsName string `json:"osName"`
    // Client type when broken down by `clientType`.
    ClientType string `json:"clientType"`
    // Client name when broken down by `clientName`.
    ClientName string `json:"clientName"`
    // SDK name when broken down by `sdk`.
    Sdk string `json:"sdk"`
    // SDK version when broken down by `sdkVersion`.
    SdkVersion string `json:"sdkVersion"`
    // Device classification when broken down by `deviceName`.
    DeviceName string `json:"deviceName"`
    // Owning team ID when broken down by `teamId`.
    TeamId string `json:"teamId"`
    // External resource ID when broken down by `resourceId`.
    ResourceId string `json:"resourceId"`
    // Resource type when broken down by `resourceType`.
    ResourceType string `json:"resourceType"`
    // Replica ordinal when broken down by `ordinal`. 0 is the primary; 1+ are
    // replicas.
    Ordinal string `json:"ordinal"`

    // Used by Decode() method
    data []byte
}

func (model UsageDataPoint) New(data []byte) *UsageDataPoint {
    model.data = data
    return &model
}

func (model *UsageDataPoint) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}