```go
package main

import (
    "fmt"
    "github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/client"
    "github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/console"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
    client.WithProject("<YOUR_PROJECT_ID>")
)

service := console.New(client)

response, error := service.CreateSource(
    console.WithCreateSourceRef("<REF>"),
    console.WithCreateSourceReferrer("https://example.com"),
    console.WithCreateSourceUtmSource("<UTM_SOURCE>"),
    console.WithCreateSourceUtmCampaign("<UTM_CAMPAIGN>"),
    console.WithCreateSourceUtmMedium("<UTM_MEDIUM>"),
)
```
