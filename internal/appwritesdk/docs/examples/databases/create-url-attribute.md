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

response, error := service.CreateUrlAttribute(
    "<DATABASE_ID>",
    "<COLLECTION_ID>",
    "",
    false,
    databases.WithCreateUrlAttributeDefault("https://example.com"),
    databases.WithCreateUrlAttributeArray(false),
)
```
