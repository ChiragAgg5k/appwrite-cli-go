```go
package main

import (
    "fmt"
    "github.com/appwrite/sdk-for-go/client"
    "github.com/appwrite/sdk-for-go/migrations"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
    client.WithProject("<YOUR_PROJECT_ID>")
)

service := migrations.New(client)

response, error := service.List(
    migrations.WithListQueries([]string{}),
    migrations.WithListSearch("<SEARCH>"),
    migrations.WithListTotal(false),
)
```
