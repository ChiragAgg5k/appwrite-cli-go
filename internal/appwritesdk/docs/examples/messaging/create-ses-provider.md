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

response, error := service.CreateSesProvider(
    "<PROVIDER_ID>",
    "<NAME>",
    messaging.WithCreateSesProviderAccessKey("<ACCESS_KEY>"),
    messaging.WithCreateSesProviderSecretKey("<SECRET_KEY>"),
    messaging.WithCreateSesProviderRegion("<REGION>"),
    messaging.WithCreateSesProviderFromName("<FROM_NAME>"),
    messaging.WithCreateSesProviderFromEmail("email@example.com"),
    messaging.WithCreateSesProviderReplyToName("<REPLY_TO_NAME>"),
    messaging.WithCreateSesProviderReplyToEmail("email@example.com"),
    messaging.WithCreateSesProviderEnabled(false),
)
```
