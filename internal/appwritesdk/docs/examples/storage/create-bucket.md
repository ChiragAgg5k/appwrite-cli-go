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

response, error := service.CreateBucket(
    "<BUCKET_ID>",
    "<NAME>",
    storage.WithCreateBucketPermissions([]string{"read("any")"}),
    storage.WithCreateBucketFileSecurity(false),
    storage.WithCreateBucketEnabled(false),
    storage.WithCreateBucketMaximumFileSize(1),
    storage.WithCreateBucketAllowedFileExtensions([]string{}),
    storage.WithCreateBucketCompression("none"),
    storage.WithCreateBucketEncryption(false),
    storage.WithCreateBucketAntivirus(false),
    storage.WithCreateBucketTransformations(false),
)
```
