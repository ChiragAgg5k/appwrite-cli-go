```go
package main

import (
    "fmt"
    "github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/client"
    "github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/mysql"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
    client.WithProject("<YOUR_PROJECT_ID>")
)

service := mysql.New(client)

response, error := service.ListRestorations(
    "<DATABASE_ID>",
    mysql.WithListRestorationsStatus("pending"),
    mysql.WithListRestorationsType("backup"),
    mysql.WithListRestorationsLimit(1),
    mysql.WithListRestorationsOffset(0),
)
```
