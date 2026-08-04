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

response, error := service.UpdateEmailTemplate(
    "verification",
    project.WithUpdateEmailTemplateLocale("af"),
    project.WithUpdateEmailTemplateSubject("<SUBJECT>"),
    project.WithUpdateEmailTemplateMessage("<MESSAGE>"),
    project.WithUpdateEmailTemplateSenderName("<SENDER_NAME>"),
    project.WithUpdateEmailTemplateSenderEmail("email@example.com"),
    project.WithUpdateEmailTemplateReplyToEmail("email@example.com"),
    project.WithUpdateEmailTemplateReplyToName("<REPLY_TO_NAME>"),
)
```
