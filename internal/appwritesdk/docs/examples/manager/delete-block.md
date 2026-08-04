```go
package main

import (
    "fmt"
    "github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/client"
    "github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/manager"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
)

service := manager.New(client)

response, error := service.DeleteBlock(
    "<PROJECT_ID>",
    "projects",
    manager.WithDeleteBlockResourceId("<RESOURCE_ID>"),
)
```
