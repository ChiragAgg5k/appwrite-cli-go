```go
package main

import (
    "fmt"
    "github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/client"
    "github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/mongo"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
    client.WithProject("<YOUR_PROJECT_ID>")
)

service := mongo.New(client)

response, error := service.CreateRestoration(
    "<DATABASE_ID>",
    mongo.WithCreateRestorationType("backup"),
    mongo.WithCreateRestorationBackupId("<BACKUP_ID>"),
    mongo.WithCreateRestorationTargetTime("2020-10-15T06:38:00.000+00:00"),
)
```
