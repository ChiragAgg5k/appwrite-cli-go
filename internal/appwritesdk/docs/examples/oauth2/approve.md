```go
package main

import (
    "fmt"
    "github.com/appwrite/sdk-for-go/client"
    "github.com/appwrite/sdk-for-go/oauth2"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
    client.WithProject("<YOUR_PROJECT_ID>")
)

service := oauth2.New(client)

response, error := service.Approve(
    "<GRANT_ID>",
    oauth2.WithApproveAuthorizationDetails("<AUTHORIZATION_DETAILS>"),
    oauth2.WithApproveScope("<SCOPE>"),
)
```
