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

response, error := service.UpdateOAuth2Dailymotion(
    project.WithUpdateOAuth2DailymotionApiKey("<API_KEY>"),
    project.WithUpdateOAuth2DailymotionApiSecret("<API_SECRET>"),
    project.WithUpdateOAuth2DailymotionEnabled(false),
)
```
