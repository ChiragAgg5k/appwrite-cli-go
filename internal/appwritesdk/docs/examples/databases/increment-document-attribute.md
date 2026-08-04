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

response, error := service.IncrementDocumentAttribute(
    "<DATABASE_ID>",
    "<COLLECTION_ID>",
    "<DOCUMENT_ID>",
    "",
    databases.WithIncrementDocumentAttributeValue(0),
    databases.WithIncrementDocumentAttributeMax(0),
    databases.WithIncrementDocumentAttributeTransactionId("<TRANSACTION_ID>"),
)
```
