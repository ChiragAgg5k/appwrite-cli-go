```go
package main

import (
    "fmt"
    "github.com/appwrite/sdk-for-go/client"
    "github.com/appwrite/sdk-for-go/tokens"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
    client.WithProject("<YOUR_PROJECT_ID>")
)

service := tokens.New(client)

response, error := service.List(
    "<BUCKET_ID>",
    "<FILE_ID>",
    tokens.WithListQueries([]string{}),
    tokens.WithListTotal(false),
)
```
