```go
package main

import (
    "fmt"
    "github.com/appwrite/sdk-for-go/client"
    "github.com/appwrite/sdk-for-go/users"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
    client.WithProject("<YOUR_PROJECT_ID>")
)

service := users.New(client)

response, error := service.CreateTarget(
    "<USER_ID>",
    "<TARGET_ID>",
    "email",
    "<IDENTIFIER>",
    users.WithCreateTargetProviderId("<PROVIDER_ID>"),
    users.WithCreateTargetName("<NAME>"),
)
```
