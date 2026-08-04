```go
package main

import (
    "fmt"
    "github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/client"
    "github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/migrations"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
    client.WithProject("<YOUR_PROJECT_ID>")
)

service := migrations.New(client)

response, error := service.GetSupabaseReport(
    []string{},
    "https://example.com",
    "<API_KEY>",
    "<DATABASE_HOST>",
    "<USERNAME>",
    "password",
    migrations.WithGetSupabaseReportPort(0),
)
```
