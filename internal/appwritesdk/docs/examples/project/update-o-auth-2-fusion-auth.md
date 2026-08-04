```go
package main

import (
    "fmt"
    "github.com/appwrite/sdk-for-go/client"
    "github.com/appwrite/sdk-for-go/project"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
    client.WithProject("<YOUR_PROJECT_ID>")
)

service := project.New(client)

response, error := service.UpdateOAuth2FusionAuth(
    project.WithUpdateOAuth2FusionAuthClientId("<CLIENT_ID>"),
    project.WithUpdateOAuth2FusionAuthClientSecret("<CLIENT_SECRET>"),
    project.WithUpdateOAuth2FusionAuthEndpoint("<ENDPOINT>"),
    project.WithUpdateOAuth2FusionAuthEnabled(false),
)
```
