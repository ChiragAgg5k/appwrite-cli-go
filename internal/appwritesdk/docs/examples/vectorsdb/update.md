```go
package main

import (
    "fmt"
    "github.com/appwrite/sdk-for-go/client"
    "github.com/appwrite/sdk-for-go/vectorsdb"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
    client.WithProject("<YOUR_PROJECT_ID>")
)

service := vectorsdb.New(client)

response, error := service.Update(
    "<DATABASE_ID>",
    "<NAME>",
    vectorsdb.WithUpdateEnabled(false),
    vectorsdb.WithUpdateSpecification("serverless"),
    vectorsdb.WithUpdateReplicas(0),
    vectorsdb.WithUpdateSyncMode("async"),
)
```
