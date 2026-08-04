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

response, error := service.UpdateVonageProvider(
    "<PROVIDER_ID>",
    messaging.WithUpdateVonageProviderName("<NAME>"),
    messaging.WithUpdateVonageProviderEnabled(false),
    messaging.WithUpdateVonageProviderApiKey("<API_KEY>"),
    messaging.WithUpdateVonageProviderApiSecret("<API_SECRET>"),
    messaging.WithUpdateVonageProviderFrom("<FROM>"),
)
```
