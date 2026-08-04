```go
package main

import (
    "fmt"
    "github.com/appwrite/sdk-for-go/client"
    "github.com/appwrite/sdk-for-go/console"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
    client.WithProject("<YOUR_PROJECT_ID>")
)

service := console.New(client)

response, error := service.SuggestColumns(
    "<DATABASE_ID>",
    "<TABLE_ID>",
    console.WithSuggestColumnsContext("<CONTEXT>"),
    console.WithSuggestColumnsMin(1),
    console.WithSuggestColumnsMax(1),
)
```
