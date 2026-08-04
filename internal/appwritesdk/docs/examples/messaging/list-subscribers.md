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

response, error := service.ListSubscribers(
    "<TOPIC_ID>",
    messaging.WithListSubscribersQueries([]string{}),
    messaging.WithListSubscribersSearch("<SEARCH>"),
    messaging.WithListSubscribersTotal(false),
)
```
