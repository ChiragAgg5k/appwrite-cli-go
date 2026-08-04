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

response, error := service.DeleteBlock(
    "<PROJECT_ID>",
    "projects",
    manager.WithDeleteBlockResourceId("<RESOURCE_ID>"),
)
```
