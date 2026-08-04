```go
package main

import (
    "fmt"
    "github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/client"
    "github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/oauth2"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
    client.WithProject("<YOUR_PROJECT_ID>")
)

service := oauth2.New(client)

response, error := service.CreateDeviceAuthorization(
    oauth2.WithCreateDeviceAuthorizationClientId("<CLIENT_ID>"),
    oauth2.WithCreateDeviceAuthorizationScope("<SCOPE>"),
    oauth2.WithCreateDeviceAuthorizationAuthorizationDetails("<AUTHORIZATION_DETAILS>"),
    oauth2.WithCreateDeviceAuthorizationResource(""),
    oauth2.WithCreateDeviceAuthorizationAudience("<AUDIENCE>"),
)
```
