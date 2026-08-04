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

response, error := service.UpdateOAuth2Oidc(
    project.WithUpdateOAuth2OidcClientId("<CLIENT_ID>"),
    project.WithUpdateOAuth2OidcClientSecret("<CLIENT_SECRET>"),
    project.WithUpdateOAuth2OidcWellKnownURL("https://example.com"),
    project.WithUpdateOAuth2OidcAuthorizationURL("https://example.com"),
    project.WithUpdateOAuth2OidcTokenURL("https://example.com"),
    project.WithUpdateOAuth2OidcUserInfoURL("https://example.com"),
    project.WithUpdateOAuth2OidcPrompt([]string{}),
    project.WithUpdateOAuth2OidcMaxAge(0),
    project.WithUpdateOAuth2OidcEnabled(false),
)
```
