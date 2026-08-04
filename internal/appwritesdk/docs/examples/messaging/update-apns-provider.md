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

response, error := service.UpdateAPNSProvider(
    "<PROVIDER_ID>",
    messaging.WithUpdateAPNSProviderName("<NAME>"),
    messaging.WithUpdateAPNSProviderEnabled(false),
    messaging.WithUpdateAPNSProviderAuthKey("<AUTH_KEY>"),
    messaging.WithUpdateAPNSProviderAuthKeyId("<AUTH_KEY_ID>"),
    messaging.WithUpdateAPNSProviderTeamId("<TEAM_ID>"),
    messaging.WithUpdateAPNSProviderBundleId("<BUNDLE_ID>"),
    messaging.WithUpdateAPNSProviderSandbox(false),
)
```
