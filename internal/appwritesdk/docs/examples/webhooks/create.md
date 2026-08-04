```go
package main

import (
    "fmt"
    "github.com/appwrite/sdk-for-go/client"
    "github.com/appwrite/sdk-for-go/webhooks"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
    client.WithProject("<YOUR_PROJECT_ID>")
)

service := webhooks.New(client)

response, error := service.Create(
    "<WEBHOOK_ID>",
    "",
    "<NAME>",
    []string{},
    webhooks.WithCreateEnabled(false),
    webhooks.WithCreateTls(false),
    webhooks.WithCreateAuthUsername("<AUTH_USERNAME>"),
    webhooks.WithCreateAuthPassword("password"),
    webhooks.WithCreateSecret("<SECRET>"),
)
```
