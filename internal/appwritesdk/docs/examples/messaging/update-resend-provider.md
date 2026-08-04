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

response, error := service.UpdateResendProvider(
    "<PROVIDER_ID>",
    messaging.WithUpdateResendProviderName("<NAME>"),
    messaging.WithUpdateResendProviderEnabled(false),
    messaging.WithUpdateResendProviderApiKey("<API_KEY>"),
    messaging.WithUpdateResendProviderFromName("<FROM_NAME>"),
    messaging.WithUpdateResendProviderFromEmail("email@example.com"),
    messaging.WithUpdateResendProviderReplyToName("<REPLY_TO_NAME>"),
    messaging.WithUpdateResendProviderReplyToEmail("<REPLY_TO_EMAIL>"),
)
```
