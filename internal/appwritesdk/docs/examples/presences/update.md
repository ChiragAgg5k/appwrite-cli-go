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

response, error := service.Update(
    "<PRESENCE_ID>",
    presences.WithUpdateStatus("<STATUS>"),
    presences.WithUpdateExpiresAt("2020-10-15T06:38:00.000+00:00"),
    presences.WithUpdateMetadata(map[string]interface{}{}),
    presences.WithUpdatePermissions([]string{"read("any")"}),
    presences.WithUpdatePurge(false),
)
```
