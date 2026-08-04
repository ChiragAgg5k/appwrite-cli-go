```go
package main

import (
    "fmt"
    "github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/client"
    "github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/messaging"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
    client.WithProject("<YOUR_PROJECT_ID>")
)

service := messaging.New(client)

response, error := service.UpdateMailgunProvider(
    "<PROVIDER_ID>",
    messaging.WithUpdateMailgunProviderName("<NAME>"),
    messaging.WithUpdateMailgunProviderApiKey("<API_KEY>"),
    messaging.WithUpdateMailgunProviderDomain("<DOMAIN>"),
    messaging.WithUpdateMailgunProviderIsEuRegion(false),
    messaging.WithUpdateMailgunProviderEnabled(false),
    messaging.WithUpdateMailgunProviderFromName("<FROM_NAME>"),
    messaging.WithUpdateMailgunProviderFromEmail("email@example.com"),
    messaging.WithUpdateMailgunProviderReplyToName("<REPLY_TO_NAME>"),
    messaging.WithUpdateMailgunProviderReplyToEmail("<REPLY_TO_EMAIL>"),
)
```
