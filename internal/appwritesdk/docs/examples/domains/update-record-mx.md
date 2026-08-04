```go
package main

import (
    "fmt"
    "github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/client"
    "github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/domains"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
    client.WithProject("<YOUR_PROJECT_ID>")
)

service := domains.New(client)

response, error := service.UpdateRecordMX(
    "<DOMAIN_ID>",
    "<RECORD_ID>",
    "",
    "<VALUE>",
    1,
    0,
    domains.WithUpdateRecordMXComment("<COMMENT>"),
)
```
