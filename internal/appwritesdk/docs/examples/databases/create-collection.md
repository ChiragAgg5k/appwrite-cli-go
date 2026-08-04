```go
package main

import (
    "fmt"
    "github.com/appwrite/sdk-for-go/client"
    "github.com/appwrite/sdk-for-go/databases"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
    client.WithProject("<YOUR_PROJECT_ID>")
)

service := databases.New(client)

response, error := service.CreateCollection(
    "<DATABASE_ID>",
    "<COLLECTION_ID>",
    "<NAME>",
    databases.WithCreateCollectionPermissions([]string{"read("any")"}),
    databases.WithCreateCollectionDocumentSecurity(false),
    databases.WithCreateCollectionEnabled(false),
    databases.WithCreateCollectionAttributes([]interface{}{}),
    databases.WithCreateCollectionIndexes([]interface{}{}),
)
```
