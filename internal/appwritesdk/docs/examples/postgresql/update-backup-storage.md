```go
package main

import (
    "fmt"
    "github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/client"
    "github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/postgresql"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
    client.WithProject("<YOUR_PROJECT_ID>")
)

service := postgresql.New(client)

response, error := service.UpdateBackupStorage(
    "<DATABASE_ID>",
    "s3",
    "<BUCKET>",
    "<ACCESS_KEY>",
    "<SECRET_KEY>",
    postgresql.WithUpdateBackupStorageRegion("<REGION>"),
    postgresql.WithUpdateBackupStoragePrefix("<PREFIX>"),
    postgresql.WithUpdateBackupStorageEndpoint("<ENDPOINT>"),
)
```
