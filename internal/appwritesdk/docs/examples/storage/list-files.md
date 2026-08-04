```go
package main

import (
    "fmt"
    "github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/client"
    "github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/storage"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
    client.WithProject("<YOUR_PROJECT_ID>")
)

service := storage.New(client)

response, error := service.ListFiles(
    "<BUCKET_ID>",
    storage.WithListFilesQueries([]string{}),
    storage.WithListFilesSearch("<SEARCH>"),
    storage.WithListFilesTotal(false),
)
```
