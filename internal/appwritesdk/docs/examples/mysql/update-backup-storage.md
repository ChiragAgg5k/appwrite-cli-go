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

response, error := service.UpdateBackupStorage(
    "<DATABASE_ID>",
    "s3",
    "<BUCKET>",
    "<ACCESS_KEY>",
    "<SECRET_KEY>",
    mysql.WithUpdateBackupStorageRegion("<REGION>"),
    mysql.WithUpdateBackupStoragePrefix("<PREFIX>"),
    mysql.WithUpdateBackupStorageEndpoint("<ENDPOINT>"),
)
```
