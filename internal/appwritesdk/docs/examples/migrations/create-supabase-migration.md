```go
package main

import (
    "fmt"
    "github.com/appwrite/sdk-for-go/client"
    "github.com/appwrite/sdk-for-go/migrations"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
    client.WithProject("<YOUR_PROJECT_ID>")
)

service := migrations.New(client)

response, error := service.CreateSupabaseMigration(
    []string{},
    "https://example.com",
    "<API_KEY>",
    "<DATABASE_HOST>",
    "<USERNAME>",
    "password",
    migrations.WithCreateSupabaseMigrationPort(0),
)
```
