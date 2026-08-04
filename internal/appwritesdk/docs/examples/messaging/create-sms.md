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

response, error := service.CreateSMS(
    "<MESSAGE_ID>",
    "<CONTENT>",
    messaging.WithCreateSMSTopics([]string{}),
    messaging.WithCreateSMSUsers([]string{}),
    messaging.WithCreateSMSTargets([]string{}),
    messaging.WithCreateSMSDraft(false),
    messaging.WithCreateSMSScheduledAt("2020-10-15T06:38:00.000+00:00"),
)
```
