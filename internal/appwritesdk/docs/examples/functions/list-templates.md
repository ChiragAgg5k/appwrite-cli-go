```go
package main

import (
    "fmt"
    "github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/client"
    "github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/functions"
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
