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

response, error := service.UpdateOAuth2Keycloak(
    project.WithUpdateOAuth2KeycloakClientId("<CLIENT_ID>"),
    project.WithUpdateOAuth2KeycloakClientSecret("<CLIENT_SECRET>"),
    project.WithUpdateOAuth2KeycloakEndpoint("<ENDPOINT>"),
    project.WithUpdateOAuth2KeycloakRealmName("<REALM_NAME>"),
    project.WithUpdateOAuth2KeycloakEnabled(false),
)
```
