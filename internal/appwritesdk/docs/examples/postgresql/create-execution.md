```go
package main

import (
    "fmt"
    "github.com/appwrite/sdk-for-go/client"
    "github.com/appwrite/sdk-for-go/postgresql"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
    client.WithProject("<YOUR_PROJECT_ID>")
)

service := postgresql.New(client)

response, error := service.CreateExecution(
    "<DATABASE_ID>",
    "<SQL>",
    postgresql.WithCreateExecutionBindings(map[string]interface{}{}),
    postgresql.WithCreateExecutionTimeoutSeconds(1),
)
```
