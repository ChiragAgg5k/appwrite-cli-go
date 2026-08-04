```go
package main

import (
    "fmt"
    "github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/client"
    "github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/tablesdb"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
    client.WithProject("<YOUR_PROJECT_ID>")
)

service := tablesdb.New(client)

response, error := service.CreateRow(
    "<DATABASE_ID>",
    "<TABLE_ID>",
    "<ROW_ID>",
    map[string]interface{}{
        "username": "walter.obrien",
        "email": "walter.obrien@example.com",
        "fullName": "Walter O'Brien",
        "age": 30,
        "isAdmin": false
    },
    tablesdb.WithCreateRowPermissions([]string{"read("any")"}),
    tablesdb.WithCreateRowTransactionId("<TRANSACTION_ID>"),
)
```
