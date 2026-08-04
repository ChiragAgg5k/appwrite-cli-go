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

response, error := service.CreateSchedule(
    "<PROJECT_ID>",
    "function",
    "<RESOURCE_ID>",
    "",
    projects.WithCreateScheduleActive(false),
    projects.WithCreateScheduleData(map[string]interface{}{}),
)
```
