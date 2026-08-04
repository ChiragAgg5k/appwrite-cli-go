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

response, error := service.CreateIntegerColumn(
    "<DATABASE_ID>",
    "<TABLE_ID>",
    "",
    false,
    tablesdb.WithCreateIntegerColumnMin(0),
    tablesdb.WithCreateIntegerColumnMax(0),
    tablesdb.WithCreateIntegerColumnDefault(0),
    tablesdb.WithCreateIntegerColumnArray(false),
)
```
