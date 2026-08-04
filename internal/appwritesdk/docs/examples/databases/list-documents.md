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

response, error := service.ListDocuments(
    "<DATABASE_ID>",
    "<COLLECTION_ID>",
    databases.WithListDocumentsQueries([]string{}),
    databases.WithListDocumentsTransactionId("<TRANSACTION_ID>"),
    databases.WithListDocumentsTotal(false),
    databases.WithListDocumentsTtl(0),
)
```
