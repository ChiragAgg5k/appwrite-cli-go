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

response, error := service.UpsertDocument(
    "<DATABASE_ID>",
    "<COLLECTION_ID>",
    "<DOCUMENT_ID>",
    databases.WithUpsertDocumentData(map[string]interface{}{
        "username": "walter.obrien",
        "email": "walter.obrien@example.com",
        "fullName": "Walter O'Brien",
        "age": 30,
        "isAdmin": false
    }),
    databases.WithUpsertDocumentPermissions([]string{"read("any")"}),
    databases.WithUpsertDocumentTransactionId("<TRANSACTION_ID>"),
)
```
