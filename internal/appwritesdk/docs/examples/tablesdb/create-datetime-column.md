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

response, error := service.CreateDatetimeColumn(
    "<DATABASE_ID>",
    "<TABLE_ID>",
    "",
    false,
    tablesdb.WithCreateDatetimeColumnDefault("2020-10-15T06:38:00.000+00:00"),
    tablesdb.WithCreateDatetimeColumnArray(false),
)
```
