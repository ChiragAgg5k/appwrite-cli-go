```go
package main

import (
    "fmt"
    "github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/client"
    "github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/account"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
    client.WithProject("<YOUR_PROJECT_ID>")
)

service := account.New(client)

response, error := service.CreateOAuth2Session(
    "amazon",
    account.WithCreateOAuth2SessionSuccess("https://example.com"),
    account.WithCreateOAuth2SessionFailure("https://example.com"),
    account.WithCreateOAuth2SessionScopes([]string{}),
)
```
