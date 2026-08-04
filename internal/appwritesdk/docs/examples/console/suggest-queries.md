```go
package main

import (
    "fmt"
    "github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/client"
    "github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/console"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
    client.WithProject("<YOUR_PROJECT_ID>")
)

service := console.New(client)

response, error := service.SuggestQueries(
    "activities",
    "<INPUT>",
    console.WithSuggestQueriesDatabaseId("<DATABASE_ID>"),
    console.WithSuggestQueriesTableId("<TABLE_ID>"),
)
```
