```go
package main

import (
    "fmt"
    "github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/client"
    "github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/databases"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
    client.WithProject("<YOUR_PROJECT_ID>")
)

service := databases.New(client)

response, error := service.UpdatePolygonAttribute(
    "<DATABASE_ID>",
    "<COLLECTION_ID>",
    "",
    false,
    databases.WithUpdatePolygonAttributeDefault([][]interface{}{[[1, 2], [3, 4], [5, 6], [1, 2]]}),
    databases.WithUpdatePolygonAttributeNewKey(""),
)
```
