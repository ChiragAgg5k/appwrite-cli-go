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

response, error := service.UpdateOAuth2Microsoft(
    project.WithUpdateOAuth2MicrosoftApplicationId("<APPLICATION_ID>"),
    project.WithUpdateOAuth2MicrosoftApplicationSecret("<APPLICATION_SECRET>"),
    project.WithUpdateOAuth2MicrosoftTenant("<TENANT>"),
    project.WithUpdateOAuth2MicrosoftEnabled(false),
)
```
