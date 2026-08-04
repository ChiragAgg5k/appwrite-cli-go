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

response, error := service.SuggestIndexes(
    "<DATABASE_ID>",
    "<TABLE_ID>",
    console.WithSuggestIndexesMin(1),
    console.WithSuggestIndexesMax(1),
)
```
