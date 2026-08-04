```go
package main

import (
    "fmt"
    "github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/client"
    "github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/advisor"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
    client.WithProject("<YOUR_PROJECT_ID>")
)

service := advisor.New(client)

response, error := service.ListInsights(
    "<REPORT_ID>",
    advisor.WithListInsightsQueries([]string{}),
    advisor.WithListInsightsTotal(false),
)
```
