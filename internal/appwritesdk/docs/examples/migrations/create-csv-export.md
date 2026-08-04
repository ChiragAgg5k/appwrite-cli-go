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

response, error := service.CreateCSVExport(
    "<DATABASE_ID>",
    "<COLLECTION_ID>",
    "<FILENAME>",
    migrations.WithCreateCSVExportColumns([]map[string]any{}),
    migrations.WithCreateCSVExportQueries([]string{}),
    migrations.WithCreateCSVExportDelimiter("<DELIMITER>"),
    migrations.WithCreateCSVExportEnclosure("<ENCLOSURE>"),
    migrations.WithCreateCSVExportEscape("<ESCAPE>"),
    migrations.WithCreateCSVExportHeader(false),
    migrations.WithCreateCSVExportNotify(false),
)
```
