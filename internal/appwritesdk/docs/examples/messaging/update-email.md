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

response, error := service.UpdateEmail(
    "<MESSAGE_ID>",
    messaging.WithUpdateEmailTopics([]string{}),
    messaging.WithUpdateEmailUsers([]string{}),
    messaging.WithUpdateEmailTargets([]string{}),
    messaging.WithUpdateEmailSubject("<SUBJECT>"),
    messaging.WithUpdateEmailContent("<CONTENT>"),
    messaging.WithUpdateEmailDraft(false),
    messaging.WithUpdateEmailHtml(false),
    messaging.WithUpdateEmailCc([]string{}),
    messaging.WithUpdateEmailBcc([]string{}),
    messaging.WithUpdateEmailScheduledAt("2020-10-15T06:38:00.000+00:00"),
    messaging.WithUpdateEmailAttachments([]string{}),
)
```
