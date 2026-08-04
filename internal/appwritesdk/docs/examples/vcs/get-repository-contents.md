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

response, error := service.GetRepositoryContents(
    "<INSTALLATION_ID>",
    "<PROVIDER_REPOSITORY_ID>",
    vcs.WithGetRepositoryContentsProviderRootDirectory("<PROVIDER_ROOT_DIRECTORY>"),
    vcs.WithGetRepositoryContentsProviderReference("<PROVIDER_REFERENCE>"),
)
```
