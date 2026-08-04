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

response, error := service.ListNamespaces(
    "<INSTALLATION_ID>",
    vcs.WithListNamespacesSearch("<SEARCH>"),
    vcs.WithListNamespacesQueries([]string{}),
)
```
