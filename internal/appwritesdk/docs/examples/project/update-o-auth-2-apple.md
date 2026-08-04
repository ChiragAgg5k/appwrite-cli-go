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

response, error := service.UpdateOAuth2Apple(
    project.WithUpdateOAuth2AppleServiceId("<SERVICE_ID>"),
    project.WithUpdateOAuth2AppleKeyId("<KEY_ID>"),
    project.WithUpdateOAuth2AppleTeamId("<TEAM_ID>"),
    project.WithUpdateOAuth2AppleP8File("<P8_FILE>"),
    project.WithUpdateOAuth2AppleEnabled(false),
)
```
