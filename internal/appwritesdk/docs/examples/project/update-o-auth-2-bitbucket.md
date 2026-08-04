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

response, error := service.UpdateOAuth2Bitbucket(
    project.WithUpdateOAuth2BitbucketKey("<KEY>"),
    project.WithUpdateOAuth2BitbucketSecret("<SECRET>"),
    project.WithUpdateOAuth2BitbucketEnabled(false),
)
```
