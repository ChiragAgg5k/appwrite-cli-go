```go
package main

import (
    "fmt"
    "github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/client"
    "github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/postgresql"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
    client.WithProject("<YOUR_PROJECT_ID>")
)

service := postgresql.New(client)

response, error := service.Create(
    "<DATABASE_ID>",
    "<NAME>",
    postgresql.WithCreateVersion("17"),
    postgresql.WithCreateSpecification("<SPECIFICATION>"),
    postgresql.WithCreateReplicas(0),
    postgresql.WithCreateSyncMode("async"),
    postgresql.WithCreateStandbyRegion("<STANDBY_REGION>"),
    postgresql.WithCreateNetworkIdleTimeoutSeconds(60),
    postgresql.WithCreateNetworkIPAllowlist([]string{}),
    postgresql.WithCreateIdleTimeoutMinutes(5),
    postgresql.WithCreatePitr(false),
    postgresql.WithCreatePitrRetentionDays(1),
    postgresql.WithCreateStorageAutoscaling(false),
    postgresql.WithCreateStorageAutoscalingThresholdPercent(50),
    postgresql.WithCreateStorageAutoscalingMaxGb(0),
)
```
