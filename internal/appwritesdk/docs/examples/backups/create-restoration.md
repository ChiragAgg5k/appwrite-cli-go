```go
package main

import (
    "fmt"
    "github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/client"
    "github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/backups"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
    client.WithProject("<YOUR_PROJECT_ID>")
)

service := backups.New(client)

response, error := service.CreateRestoration(
    "<ARCHIVE_ID>",
    []string{},
    backups.WithCreateRestorationNewResourceId("<NEW_RESOURCE_ID>"),
    backups.WithCreateRestorationNewResourceName("<NEW_RESOURCE_NAME>"),
)
```
