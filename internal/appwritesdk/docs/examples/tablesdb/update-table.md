```go
package main

import (
    "fmt"
    "github.com/appwrite/sdk-for-go/client"
    "github.com/appwrite/sdk-for-go/tablesdb"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
    client.WithProject("<YOUR_PROJECT_ID>")
)

service := tablesdb.New(client)

response, error := service.UpdateTable(
    "<DATABASE_ID>",
    "<TABLE_ID>",
    tablesdb.WithUpdateTableName("<NAME>"),
    tablesdb.WithUpdateTablePermissions([]string{"read("any")"}),
    tablesdb.WithUpdateTableRowSecurity(false),
    tablesdb.WithUpdateTableEnabled(false),
    tablesdb.WithUpdateTablePurge(false),
)
```
