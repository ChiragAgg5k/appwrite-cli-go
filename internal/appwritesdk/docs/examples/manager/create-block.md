```go
package main

import (
    "fmt"
    "github.com/appwrite/sdk-for-go/client"
    "github.com/appwrite/sdk-for-go/manager"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
)

service := manager.New(client)

response, error := service.CreateBlock(
    "<PROJECT_ID>",
    "projects",
    manager.WithCreateBlockResourceId("<RESOURCE_ID>"),
    manager.WithCreateBlockMode("full"),
    manager.WithCreateBlockReason("<REASON>"),
    manager.WithCreateBlockExpiredAt("2020-10-15T06:38:00.000+00:00"),
)
```
