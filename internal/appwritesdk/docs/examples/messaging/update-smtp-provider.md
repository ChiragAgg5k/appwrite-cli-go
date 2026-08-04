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

response, error := service.UpdateSMTPProvider(
    "<PROVIDER_ID>",
    messaging.WithUpdateSMTPProviderName("<NAME>"),
    messaging.WithUpdateSMTPProviderHost("<HOST>"),
    messaging.WithUpdateSMTPProviderPort(1),
    messaging.WithUpdateSMTPProviderUsername("<USERNAME>"),
    messaging.WithUpdateSMTPProviderPassword("password"),
    messaging.WithUpdateSMTPProviderEncryption("none"),
    messaging.WithUpdateSMTPProviderAutoTLS(false),
    messaging.WithUpdateSMTPProviderMailer("<MAILER>"),
    messaging.WithUpdateSMTPProviderFromName("<FROM_NAME>"),
    messaging.WithUpdateSMTPProviderFromEmail("email@example.com"),
    messaging.WithUpdateSMTPProviderReplyToName("<REPLY_TO_NAME>"),
    messaging.WithUpdateSMTPProviderReplyToEmail("<REPLY_TO_EMAIL>"),
    messaging.WithUpdateSMTPProviderEnabled(false),
)
```
