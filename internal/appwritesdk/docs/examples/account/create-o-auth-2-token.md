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

response, error := service.CreateOAuth2Token(
    "amazon",
    account.WithCreateOAuth2TokenSuccess("https://example.com"),
    account.WithCreateOAuth2TokenFailure("https://example.com"),
    account.WithCreateOAuth2TokenScopes([]string{}),
)
```
