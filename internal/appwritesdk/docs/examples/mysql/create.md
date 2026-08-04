```go
package main

import (
    "fmt"
    "github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/client"
    "github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/mysql"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
    client.WithProject("<YOUR_PROJECT_ID>")
)

service := mysql.New(client)

response, error := service.Create(
    "<DATABASE_ID>",
    "<NAME>",
    mysql.WithCreateVersion("17"),
    mysql.WithCreateSpecification("<SPECIFICATION>"),
    mysql.WithCreateReplicas(0),
    mysql.WithCreateSyncMode("async"),
    mysql.WithCreateStandbyRegion("<STANDBY_REGION>"),
    mysql.WithCreateNetworkIdleTimeoutSeconds(60),
    mysql.WithCreateNetworkIPAllowlist([]string{}),
    mysql.WithCreateIdleTimeoutMinutes(5),
    mysql.WithCreatePitr(false),
    mysql.WithCreatePitrRetentionDays(1),
    mysql.WithCreateStorageAutoscaling(false),
    mysql.WithCreateStorageAutoscalingThresholdPercent(50),
    mysql.WithCreateStorageAutoscalingMaxGb(0),
)
```
