```go
package main

import (
    "fmt"
    "github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/client"
    "github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/organizations"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
    client.WithProject("<YOUR_PROJECT_ID>")
)

service := organizations.New(client)

response, error := service.GetUsage(
    "<ORGANIZATION_ID>",
    organizations.WithGetUsageStartDate("2020-10-15T06:38:00.000+00:00"),
    organizations.WithGetUsageEndDate("2020-10-15T06:38:00.000+00:00"),
)
```
