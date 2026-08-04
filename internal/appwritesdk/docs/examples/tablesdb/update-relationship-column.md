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

response, error := service.UpdateRelationshipColumn(
    "<DATABASE_ID>",
    "<TABLE_ID>",
    "",
    tablesdb.WithUpdateRelationshipColumnOnDelete("cascade"),
    tablesdb.WithUpdateRelationshipColumnNewKey(""),
)
```
