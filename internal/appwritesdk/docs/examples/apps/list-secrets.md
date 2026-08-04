```go
package main

import (
    "fmt"
    "github.com/appwrite/sdk-for-go/client"
    "github.com/appwrite/sdk-for-go/apps"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
    client.WithProject("<YOUR_PROJECT_ID>")
)

service := apps.New(client)

response, error := service.ListSecrets(
    "<APP_ID>",
    apps.WithListSecretsQueries([]string{}),
    apps.WithListSecretsTotal(false),
)
```
