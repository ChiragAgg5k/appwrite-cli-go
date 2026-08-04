```go
package main

import (
    "fmt"
    "github.com/appwrite/sdk-for-go/client"
    "github.com/appwrite/sdk-for-go/vectorsdb"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
    client.WithProject("<YOUR_PROJECT_ID>")
)

service := vectorsdb.New(client)

response, error := service.CreateCollection(
    "<DATABASE_ID>",
    "<COLLECTION_ID>",
    "<NAME>",
    1,
    vectorsdb.WithCreateCollectionPermissions([]string{"read("any")"}),
    vectorsdb.WithCreateCollectionDocumentSecurity(false),
    vectorsdb.WithCreateCollectionEnabled(false),
)
```
