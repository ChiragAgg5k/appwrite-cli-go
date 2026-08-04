```go
package main

import (
    "fmt"
    "github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/client"
    "github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/tablesdb"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
    client.WithProject("<YOUR_PROJECT_ID>")
)

service := tablesdb.New(client)

response, error := service.Update(
    "<DATABASE_ID>",
    tablesdb.WithUpdateName("<NAME>"),
    tablesdb.WithUpdateEnabled(false),
    tablesdb.WithUpdateSpecification("serverless"),
    tablesdb.WithUpdateReplicas(0),
    tablesdb.WithUpdateSyncMode("async"),
)
```
