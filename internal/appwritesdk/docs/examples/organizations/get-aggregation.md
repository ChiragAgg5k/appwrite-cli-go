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

response, error := service.GetAggregation(
    "<ORGANIZATION_ID>",
    "<AGGREGATION_ID>",
    organizations.WithGetAggregationLimit(0),
    organizations.WithGetAggregationOffset(0),
)
```
