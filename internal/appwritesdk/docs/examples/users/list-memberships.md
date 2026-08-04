```go
package main

import (
    "fmt"
    "github.com/appwrite/sdk-for-go/client"
    "github.com/appwrite/sdk-for-go/users"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
    client.WithProject("<YOUR_PROJECT_ID>")
)

service := users.New(client)

response, error := service.ListMemberships(
    "<USER_ID>",
    users.WithListMembershipsQueries([]string{}),
    users.WithListMembershipsSearch("<SEARCH>"),
    users.WithListMembershipsTotal(false),
)
```
