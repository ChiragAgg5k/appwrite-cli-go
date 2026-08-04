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

response, error := service.ListOperations(
    "<DATABASE_ID>",
    mysql.WithListOperationsStatus("running"),
    mysql.WithListOperationsLimit(1),
    mysql.WithListOperationsOffset(0),
)
```
