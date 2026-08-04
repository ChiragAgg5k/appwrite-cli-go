```go
package main

import (
    "fmt"
    "github.com/appwrite/sdk-for-go/client"
    "github.com/appwrite/sdk-for-go/projects"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
    client.WithProject("<YOUR_PROJECT_ID>")
)

service := projects.New(client)

response, error := service.UpdateStage(
    "<PROJECT_ID>",
    "<STAGE_ID>",
    projects.WithUpdateStageSkip(false),
)
```
