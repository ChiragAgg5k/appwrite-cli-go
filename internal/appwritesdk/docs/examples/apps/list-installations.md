```go
package main

import (
    "fmt"
    "github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/client"
    "github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/apps"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
    client.WithProject("<YOUR_PROJECT_ID>")
)

service := apps.New(client)

response, error := service.ListInstallations(
    "<APP_ID>",
    apps.WithListInstallationsQueries([]string{}),
    apps.WithListInstallationsTotal(false),
)
```
