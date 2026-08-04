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

response, error := service.CreateDeployment(
    "<FUNCTION_ID>",
    file.NewInputFile("/path/to/file.png", "file.png"),
    false,
    functions.WithCreateDeploymentEntrypoint("<ENTRYPOINT>"),
    functions.WithCreateDeploymentCommands("<COMMANDS>"),
)
```
