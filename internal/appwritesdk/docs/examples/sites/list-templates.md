```go
package main

import (
    "fmt"
    "github.com/appwrite/sdk-for-go/client"
    "github.com/appwrite/sdk-for-go/sites"
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
