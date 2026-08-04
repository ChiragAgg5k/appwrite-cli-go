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

response, error := service.UpdateOAuth2Server(
    false,
    "https://example.com",
    project.WithUpdateOAuth2ServerScopes([]string{}),
    project.WithUpdateOAuth2ServerAuthorizationDetailsTypes([]string{}),
    project.WithUpdateOAuth2ServerAccessTokenDuration(60),
    project.WithUpdateOAuth2ServerRefreshTokenDuration(60),
    project.WithUpdateOAuth2ServerPublicAccessTokenDuration(60),
    project.WithUpdateOAuth2ServerPublicRefreshTokenDuration(60),
    project.WithUpdateOAuth2ServerInstallationAccessTokenDuration(60),
    project.WithUpdateOAuth2ServerConfidentialPkce(false),
    project.WithUpdateOAuth2ServerVerificationUrl("https://example.com"),
    project.WithUpdateOAuth2ServerUserCodeLength(6),
    project.WithUpdateOAuth2ServerUserCodeFormat("numeric"),
    project.WithUpdateOAuth2ServerDeviceCodeDuration(60),
    project.WithUpdateOAuth2ServerDefaultScopes([]string{}),
)
```
