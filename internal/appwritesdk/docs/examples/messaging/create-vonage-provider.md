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

response, error := service.CreateVonageProvider(
    "<PROVIDER_ID>",
    "<NAME>",
    messaging.WithCreateVonageProviderFrom("+12065550100"),
    messaging.WithCreateVonageProviderApiKey("<API_KEY>"),
    messaging.WithCreateVonageProviderApiSecret("<API_SECRET>"),
    messaging.WithCreateVonageProviderEnabled(false),
)
```
