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

response, error := service.CreateMailgunProvider(
    "<PROVIDER_ID>",
    "<NAME>",
    messaging.WithCreateMailgunProviderApiKey("<API_KEY>"),
    messaging.WithCreateMailgunProviderDomain("<DOMAIN>"),
    messaging.WithCreateMailgunProviderIsEuRegion(false),
    messaging.WithCreateMailgunProviderFromName("<FROM_NAME>"),
    messaging.WithCreateMailgunProviderFromEmail("email@example.com"),
    messaging.WithCreateMailgunProviderReplyToName("<REPLY_TO_NAME>"),
    messaging.WithCreateMailgunProviderReplyToEmail("email@example.com"),
    messaging.WithCreateMailgunProviderEnabled(false),
)
```
