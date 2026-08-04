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

response, error := service.UpdateOAuth2Spotify(
    project.WithUpdateOAuth2SpotifyClientId("<CLIENT_ID>"),
    project.WithUpdateOAuth2SpotifyClientSecret("<CLIENT_SECRET>"),
    project.WithUpdateOAuth2SpotifyEnabled(false),
)
```
