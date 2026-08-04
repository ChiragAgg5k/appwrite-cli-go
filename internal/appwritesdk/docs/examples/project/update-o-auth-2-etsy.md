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

response, error := service.UpdateOAuth2Etsy(
    project.WithUpdateOAuth2EtsyKeyString("<KEY_STRING>"),
    project.WithUpdateOAuth2EtsySharedSecret("<SHARED_SECRET>"),
    project.WithUpdateOAuth2EtsyEnabled(false),
)
```
