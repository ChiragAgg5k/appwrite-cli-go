```go
package main

import (
    "fmt"
    "github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/client"
    "github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/usage"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
    client.WithProject("<YOUR_PROJECT_ID>")
)

service := usage.New(client)

response, error := service.ListEvents(
    []string{},
    usage.WithListEventsQueries([]string{}),
    usage.WithListEventsInterval("1m"),
    usage.WithListEventsDimensions([]string{}),
    usage.WithListEventsStartAt("2020-10-15T06:38:00.000+00:00"),
    usage.WithListEventsEndAt("2020-10-15T06:38:00.000+00:00"),
    usage.WithListEventsOrderBy("time"),
    usage.WithListEventsOrderDir("asc"),
    usage.WithListEventsLimit(1),
    usage.WithListEventsOffset(0),
)
```
