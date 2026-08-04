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

response, error := service.UpdateOAuth2Twitch(
    project.WithUpdateOAuth2TwitchClientId("<CLIENT_ID>"),
    project.WithUpdateOAuth2TwitchClientSecret("<CLIENT_SECRET>"),
    project.WithUpdateOAuth2TwitchEnabled(false),
)
```
