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

response, error := service.CreateVarcharAttribute(
    "<DATABASE_ID>",
    "<COLLECTION_ID>",
    "",
    1,
    false,
    databases.WithCreateVarcharAttributeDefault("<DEFAULT>"),
    databases.WithCreateVarcharAttributeArray(false),
    databases.WithCreateVarcharAttributeEncrypt(false),
)
```
