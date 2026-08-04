```go
package main

import (
    "fmt"
    "github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/client"
    "github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/users"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
    client.WithProject("<YOUR_PROJECT_ID>")
)

service := users.New(client)

response, error := service.Create(
    "<USER_ID>",
    users.WithCreateEmail("email@example.com"),
    users.WithCreatePhone("+12065550100"),
    users.WithCreatePassword("password"),
    users.WithCreateName("<NAME>"),
)
```
