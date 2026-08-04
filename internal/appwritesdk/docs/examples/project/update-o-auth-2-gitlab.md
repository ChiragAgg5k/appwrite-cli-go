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

response, error := service.UpdateOAuth2Gitlab(
    project.WithUpdateOAuth2GitlabApplicationId("<APPLICATION_ID>"),
    project.WithUpdateOAuth2GitlabSecret("<SECRET>"),
    project.WithUpdateOAuth2GitlabEndpoint("https://example.com"),
    project.WithUpdateOAuth2GitlabEnabled(false),
)
```
