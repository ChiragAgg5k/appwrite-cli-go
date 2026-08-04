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

response, error := service.UpdateOAuth2Yandex(
    project.WithUpdateOAuth2YandexClientId("<CLIENT_ID>"),
    project.WithUpdateOAuth2YandexClientSecret("<CLIENT_SECRET>"),
    project.WithUpdateOAuth2YandexEnabled(false),
)
```
