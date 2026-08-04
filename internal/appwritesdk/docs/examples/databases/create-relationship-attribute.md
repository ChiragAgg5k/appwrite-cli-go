```go
package main

import (
    "fmt"
    "github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/client"
    "github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/databases"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
    client.WithProject("<YOUR_PROJECT_ID>")
)

service := databases.New(client)

response, error := service.CreateRelationshipAttribute(
    "<DATABASE_ID>",
    "<COLLECTION_ID>",
    "<RELATED_COLLECTION_ID>",
    "oneToOne",
    databases.WithCreateRelationshipAttributeTwoWay(false),
    databases.WithCreateRelationshipAttributeKey(""),
    databases.WithCreateRelationshipAttributeTwoWayKey(""),
    databases.WithCreateRelationshipAttributeOnDelete("cascade"),
)
```
