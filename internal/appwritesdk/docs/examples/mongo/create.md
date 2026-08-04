```go
package main

import (
    "fmt"
    "github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/client"
    "github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/mongo"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
    client.WithProject("<YOUR_PROJECT_ID>")
)

service := mongo.New(client)

response, error := service.Create(
    "<DATABASE_ID>",
    "<NAME>",
    mongo.WithCreateVersion("17"),
    mongo.WithCreateSpecification("<SPECIFICATION>"),
    mongo.WithCreateReplicas(0),
    mongo.WithCreateSyncMode("async"),
    mongo.WithCreateStandbyRegion("<STANDBY_REGION>"),
    mongo.WithCreateNetworkIdleTimeoutSeconds(60),
    mongo.WithCreateNetworkIPAllowlist([]string{}),
    mongo.WithCreateIdleTimeoutMinutes(5),
    mongo.WithCreatePitr(false),
    mongo.WithCreatePitrRetentionDays(1),
    mongo.WithCreateStorageAutoscaling(false),
    mongo.WithCreateStorageAutoscalingThresholdPercent(50),
    mongo.WithCreateStorageAutoscalingMaxGb(0),
)
```
