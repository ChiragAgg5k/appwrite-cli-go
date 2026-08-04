```go
package main

import (
    "fmt"
    "github.com/appwrite/sdk-for-go/client"
    "github.com/appwrite/sdk-for-go/messaging"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
    client.WithProject("<YOUR_PROJECT_ID>")
)

service := messaging.New(client)

response, error := service.ListTopics(
    messaging.WithListTopicsQueries([]string{}),
    messaging.WithListTopicsSearch("<SEARCH>"),
    messaging.WithListTopicsTotal(false),
)
```
