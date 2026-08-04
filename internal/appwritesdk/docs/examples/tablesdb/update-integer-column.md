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

response, error := service.UpdateIntegerColumn(
    "<DATABASE_ID>",
    "<TABLE_ID>",
    "",
    false,
    0,
    tablesdb.WithUpdateIntegerColumnMin(0),
    tablesdb.WithUpdateIntegerColumnMax(0),
    tablesdb.WithUpdateIntegerColumnNewKey(""),
)
```
