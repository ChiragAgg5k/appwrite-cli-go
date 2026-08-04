```go
package main

import (
    "fmt"
    "github.com/appwrite/sdk-for-go/client"
    "github.com/appwrite/sdk-for-go/migrations"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
    client.WithProject("<YOUR_PROJECT_ID>")
)

service := migrations.New(client)

response, error := service.CreateJSONImport(
    "<BUCKET_ID>",
    "<FILE_ID>",
    "<DATABASE_ID>",
    "<COLLECTION_ID>",
    migrations.WithCreateJSONImportInternalFile(false),
    migrations.WithCreateJSONImportOnDuplicate("fail"),
)
```
