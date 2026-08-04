```go
package main

import (
    "fmt"
    "github.com/appwrite/sdk-for-go/client"
    "github.com/appwrite/sdk-for-go/presences"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
    client.WithProject("<YOUR_PROJECT_ID>")
)

service := presences.New(client)

response, error := service.Upsert(
    "<PRESENCE_ID>",
    "<STATUS>",
    presences.WithUpsertPermissions([]string{"read("any")"}),
    presences.WithUpsertExpiresAt("2020-10-15T06:38:00.000+00:00"),
    presences.WithUpsertMetadata(map[string]interface{}{}),
)
```
