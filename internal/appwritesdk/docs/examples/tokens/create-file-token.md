```go
package main

import (
    "fmt"
    "github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/client"
    "github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/tokens"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
    client.WithProject("<YOUR_PROJECT_ID>")
)

service := tokens.New(client)

response, error := service.CreateFileToken(
    "<BUCKET_ID>",
    "<FILE_ID>",
    tokens.WithCreateFileTokenExpire("2020-10-15T06:38:00.000+00:00"),
)
```
