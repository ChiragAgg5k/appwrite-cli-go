```go
package main

import (
    "fmt"
    "github.com/appwrite/sdk-for-go/client"
    "github.com/appwrite/sdk-for-go/domains"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
    client.WithProject("<YOUR_PROJECT_ID>")
)

service := domains.New(client)

response, error := service.UpdateRecordAAAA(
    "<DOMAIN_ID>",
    "<RECORD_ID>",
    "",
    "",
    1,
    domains.WithUpdateRecordAAAAComment("<COMMENT>"),
)
```
