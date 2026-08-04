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

response, error := service.UpdateSesProvider(
    "<PROVIDER_ID>",
    messaging.WithUpdateSesProviderName("<NAME>"),
    messaging.WithUpdateSesProviderEnabled(false),
    messaging.WithUpdateSesProviderAccessKey("<ACCESS_KEY>"),
    messaging.WithUpdateSesProviderSecretKey("<SECRET_KEY>"),
    messaging.WithUpdateSesProviderRegion("<REGION>"),
    messaging.WithUpdateSesProviderFromName("<FROM_NAME>"),
    messaging.WithUpdateSesProviderFromEmail("email@example.com"),
    messaging.WithUpdateSesProviderReplyToName("<REPLY_TO_NAME>"),
    messaging.WithUpdateSesProviderReplyToEmail("<REPLY_TO_EMAIL>"),
)
```
