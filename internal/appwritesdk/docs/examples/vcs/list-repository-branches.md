```go
package main

import (
    "fmt"
    "github.com/appwrite/sdk-for-go/client"
    "github.com/appwrite/sdk-for-go/vcs"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
    client.WithProject("<YOUR_PROJECT_ID>")
)

service := vcs.New(client)

response, error := service.ListRepositoryBranches(
    "<INSTALLATION_ID>",
    "<PROVIDER_REPOSITORY_ID>",
    vcs.WithListRepositoryBranchesSearch("<SEARCH>"),
    vcs.WithListRepositoryBranchesQueries([]string{}),
)
```
