```go
package main

import (
    "fmt"
    "github.com/appwrite/sdk-for-go/client"
    "github.com/appwrite/sdk-for-go/oauth2"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
    client.WithProject("<YOUR_PROJECT_ID>")
)

service := oauth2.New(client)

response, error := service.ListProjects(
    oauth2.WithListProjectsLimit(1),
    oauth2.WithListProjectsOffset(0),
    oauth2.WithListProjectsSearch("<SEARCH>"),
)
```
