```go
package main

import (
    "fmt"
    "github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/client"
    "github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/project"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
    client.WithProject("<YOUR_PROJECT_ID>")
)

service := project.New(client)

response, error := service.UpdateSMTP(
    project.WithUpdateSMTPHost(""),
    project.WithUpdateSMTPPort(0),
    project.WithUpdateSMTPUsername("<USERNAME>"),
    project.WithUpdateSMTPPassword("password"),
    project.WithUpdateSMTPSenderEmail("email@example.com"),
    project.WithUpdateSMTPSenderName("<SENDER_NAME>"),
    project.WithUpdateSMTPReplyToEmail("email@example.com"),
    project.WithUpdateSMTPReplyToName("<REPLY_TO_NAME>"),
    project.WithUpdateSMTPSecure("tls"),
    project.WithUpdateSMTPEnabled(false),
)
```
