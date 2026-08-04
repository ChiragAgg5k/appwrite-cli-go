```go
package main

import (
    "fmt"
    "github.com/appwrite/sdk-for-go/client"
    "github.com/appwrite/sdk-for-go/functions"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
    client.WithProject("<YOUR_PROJECT_ID>")
)

service := functions.New(client)

response, error := service.ListTemplates(
    functions.WithListTemplatesRuntimes([]string{}),
    functions.WithListTemplatesUseCases([]string{}),
    functions.WithListTemplatesLimit(1),
    functions.WithListTemplatesOffset(0),
    functions.WithListTemplatesTotal(false),
)
```
