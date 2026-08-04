```go
package main

import (
    "fmt"
    "github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/client"
    "github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/sites"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
    client.WithProject("<YOUR_PROJECT_ID>")
)

service := sites.New(client)

response, error := service.ListTemplates(
    sites.WithListTemplatesFrameworks([]string{}),
    sites.WithListTemplatesUseCases([]string{}),
    sites.WithListTemplatesLimit(1),
    sites.WithListTemplatesOffset(0),
)
```
