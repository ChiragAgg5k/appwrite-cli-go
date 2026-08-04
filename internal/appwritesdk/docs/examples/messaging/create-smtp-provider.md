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

response, error := service.CreateSMTPProvider(
    "<PROVIDER_ID>",
    "<NAME>",
    "<HOST>",
    messaging.WithCreateSMTPProviderPort(1),
    messaging.WithCreateSMTPProviderUsername("<USERNAME>"),
    messaging.WithCreateSMTPProviderPassword("password"),
    messaging.WithCreateSMTPProviderEncryption("none"),
    messaging.WithCreateSMTPProviderAutoTLS(false),
    messaging.WithCreateSMTPProviderMailer("<MAILER>"),
    messaging.WithCreateSMTPProviderFromName("<FROM_NAME>"),
    messaging.WithCreateSMTPProviderFromEmail("email@example.com"),
    messaging.WithCreateSMTPProviderReplyToName("<REPLY_TO_NAME>"),
    messaging.WithCreateSMTPProviderReplyToEmail("email@example.com"),
    messaging.WithCreateSMTPProviderEnabled(false),
)
```
