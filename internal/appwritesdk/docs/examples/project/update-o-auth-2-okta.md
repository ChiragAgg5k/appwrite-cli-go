```go
package main

import (
    "fmt"
    "github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/client"
    "github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/project"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
    client.WithProject("<YOUR_PROJECT_ID>")
)

service := project.New(client)

response, error := service.UpdateOAuth2Okta(
    project.WithUpdateOAuth2OktaClientId("<CLIENT_ID>"),
    project.WithUpdateOAuth2OktaClientSecret("<CLIENT_SECRET>"),
    project.WithUpdateOAuth2OktaDomain(""),
    project.WithUpdateOAuth2OktaAuthorizationServerId("<AUTHORIZATION_SERVER_ID>"),
    project.WithUpdateOAuth2OktaEnabled(false),
)
```
