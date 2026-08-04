```go
package main

import (
    "fmt"
    "github.com/appwrite/sdk-for-go/client"
    "github.com/appwrite/sdk-for-go/documentsdb"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
    client.WithProject("<YOUR_PROJECT_ID>")
)

service := documentsdb.New(client)

response, error := service.Create(
    "<DATABASE_ID>",
    "<NAME>",
    documentsdb.WithCreateEnabled(false),
    documentsdb.WithCreateSpecification("serverless"),
    documentsdb.WithCreateReplicas(0),
    documentsdb.WithCreateSyncMode("async"),
)
```
