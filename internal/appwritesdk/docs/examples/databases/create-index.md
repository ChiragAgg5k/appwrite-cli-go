```go
package main

import (
    "fmt"
    "github.com/appwrite/sdk-for-go/client"
    "github.com/appwrite/sdk-for-go/databases"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
    client.WithProject("<YOUR_PROJECT_ID>")
)

service := databases.New(client)

response, error := service.CreateIndex(
    "<DATABASE_ID>",
    "<COLLECTION_ID>",
    "",
    "key",
    []string{},
    databases.WithCreateIndexOrders([]string{}),
    databases.WithCreateIndexLengths([]int{}),
)
```
