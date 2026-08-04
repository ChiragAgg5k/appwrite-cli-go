```go
package main

import (
    "fmt"
    "github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/client"
    "github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/migrations"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
    client.WithProject("<YOUR_PROJECT_ID>")
)

service := migrations.New(client)

response, error := service.CreateJSONExport(
    "<DATABASE_ID>",
    "<COLLECTION_ID>",
    "<FILENAME>",
    migrations.WithCreateJSONExportColumns([]map[string]any{}),
    migrations.WithCreateJSONExportQueries([]string{}),
    migrations.WithCreateJSONExportNotify(false),
)
```
