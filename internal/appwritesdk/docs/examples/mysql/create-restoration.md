```go
package main

import (
    "fmt"
    "github.com/appwrite/sdk-for-go/client"
    "github.com/appwrite/sdk-for-go/mysql"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
    client.WithProject("<YOUR_PROJECT_ID>")
)

service := mysql.New(client)

response, error := service.CreateRestoration(
    "<DATABASE_ID>",
    mysql.WithCreateRestorationType("backup"),
    mysql.WithCreateRestorationBackupId("<BACKUP_ID>"),
    mysql.WithCreateRestorationTargetTime("2020-10-15T06:38:00.000+00:00"),
)
```
