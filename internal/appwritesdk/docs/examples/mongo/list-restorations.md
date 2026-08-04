```go
package main

import (
    "fmt"
    "github.com/appwrite/sdk-for-go/client"
    "github.com/appwrite/sdk-for-go/mongo"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
    client.WithProject("<YOUR_PROJECT_ID>")
)

service := mongo.New(client)

response, error := service.ListRestorations(
    "<DATABASE_ID>",
    mongo.WithListRestorationsStatus("pending"),
    mongo.WithListRestorationsType("backup"),
    mongo.WithListRestorationsLimit(1),
    mongo.WithListRestorationsOffset(0),
)
```
