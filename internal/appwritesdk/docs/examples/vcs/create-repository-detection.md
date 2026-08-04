```go
package main

import (
    "fmt"
    "github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/client"
    "github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/vcs"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
    client.WithProject("<YOUR_PROJECT_ID>")
)

service := vcs.New(client)

response, error := service.CreateRepositoryDetection(
    "<INSTALLATION_ID>",
    "<PROVIDER_REPOSITORY_ID>",
    "runtime",
    vcs.WithCreateRepositoryDetectionProviderRootDirectory("<PROVIDER_ROOT_DIRECTORY>"),
)
```
