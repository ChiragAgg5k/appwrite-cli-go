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

response, error := service.CreateSendgridProvider(
    "<PROVIDER_ID>",
    "<NAME>",
    messaging.WithCreateSendgridProviderApiKey("<API_KEY>"),
    messaging.WithCreateSendgridProviderFromName("<FROM_NAME>"),
    messaging.WithCreateSendgridProviderFromEmail("email@example.com"),
    messaging.WithCreateSendgridProviderReplyToName("<REPLY_TO_NAME>"),
    messaging.WithCreateSendgridProviderReplyToEmail("email@example.com"),
    messaging.WithCreateSendgridProviderEnabled(false),
)
```
