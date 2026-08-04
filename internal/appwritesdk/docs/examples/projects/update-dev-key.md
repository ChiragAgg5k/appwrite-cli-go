```go
package main

import (
    "fmt"
    "github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/client"
    "github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/projects"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
    client.WithProject("<YOUR_PROJECT_ID>")
)

service := projects.New(client)

response, error := service.UpdateDevKey(
    "<PROJECT_ID>",
    "<KEY_ID>",
    "<NAME>",
    "2020-10-15T06:38:00.000+00:00",
)
```
