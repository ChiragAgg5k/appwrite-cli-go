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

response, error := service.UpdateOAuth2PaypalSandbox(
    project.WithUpdateOAuth2PaypalSandboxClientId("<CLIENT_ID>"),
    project.WithUpdateOAuth2PaypalSandboxSecretKey("<SECRET_KEY>"),
    project.WithUpdateOAuth2PaypalSandboxEnabled(false),
)
```
