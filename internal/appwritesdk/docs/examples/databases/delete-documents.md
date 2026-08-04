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

response, error := service.DeleteDocuments(
    "<DATABASE_ID>",
    "<COLLECTION_ID>",
    databases.WithDeleteDocumentsQueries([]string{}),
    databases.WithDeleteDocumentsTransactionId("<TRANSACTION_ID>"),
)
```
