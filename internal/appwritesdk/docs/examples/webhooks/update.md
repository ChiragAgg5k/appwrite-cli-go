```go
package main

import (
    "fmt"
    "github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/client"
    "github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/webhooks"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
    client.WithProject("<YOUR_PROJECT_ID>")
)

service := webhooks.New(client)

response, error := service.Update(
    "<WEBHOOK_ID>",
    "<NAME>",
    "",
    []string{},
    webhooks.WithUpdateEnabled(false),
    webhooks.WithUpdateTls(false),
    webhooks.WithUpdateAuthUsername("<AUTH_USERNAME>"),
    webhooks.WithUpdateAuthPassword("password"),
)
```
