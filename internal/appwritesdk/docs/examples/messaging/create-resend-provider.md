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

response, error := service.CreateResendProvider(
    "<PROVIDER_ID>",
    "<NAME>",
    messaging.WithCreateResendProviderApiKey("<API_KEY>"),
    messaging.WithCreateResendProviderFromName("<FROM_NAME>"),
    messaging.WithCreateResendProviderFromEmail("email@example.com"),
    messaging.WithCreateResendProviderReplyToName("<REPLY_TO_NAME>"),
    messaging.WithCreateResendProviderReplyToEmail("email@example.com"),
    messaging.WithCreateResendProviderEnabled(false),
)
```
