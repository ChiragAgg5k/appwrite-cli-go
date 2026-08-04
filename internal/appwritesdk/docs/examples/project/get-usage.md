```go
package main

import (
    "fmt"
    "github.com/appwrite/sdk-for-go/client"
    "github.com/appwrite/sdk-for-go/project"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
    client.WithProject("<YOUR_PROJECT_ID>")
)

service := project.New(client)

response, error := service.GetUsage(
    "2020-10-15T06:38:00.000+00:00",
    "2020-10-15T06:38:00.000+00:00",
    project.WithGetUsagePeriod("1h"),
)
```
