```go
package main

import (
    "fmt"
    "github.com/appwrite/sdk-for-go/client"
    "github.com/appwrite/sdk-for-go/teams"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
    client.WithProject("<YOUR_PROJECT_ID>")
)

service := teams.New(client)

response, error := service.ListMemberships(
    "<TEAM_ID>",
    teams.WithListMembershipsQueries([]string{}),
    teams.WithListMembershipsSearch("<SEARCH>"),
    teams.WithListMembershipsTotal(false),
)
```
