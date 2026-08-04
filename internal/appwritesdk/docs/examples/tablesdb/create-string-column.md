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

response, error := service.CreateStringColumn(
    "<DATABASE_ID>",
    "<TABLE_ID>",
    "",
    1,
    false,
    tablesdb.WithCreateStringColumnDefault("<DEFAULT>"),
    tablesdb.WithCreateStringColumnArray(false),
    tablesdb.WithCreateStringColumnEncrypt(false),
)
```
