```go
package main

import (
    "fmt"
    "github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/client"
    "github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/teams"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
    client.WithProject("<YOUR_PROJECT_ID>")
)

service := teams.New(client)

response, error := service.UpdateInstallation(
    "<TEAM_ID>",
    "<INSTALLATION_ID>",
    teams.WithUpdateInstallationAuthorizationDetails("<AUTHORIZATION_DETAILS>"),
)
```
