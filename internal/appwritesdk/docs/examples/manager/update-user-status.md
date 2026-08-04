```go
package main

import (
    "fmt"
    "github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/client"
    "github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/manager"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
)

service := manager.New(client)

response, error := service.UpdateUserStatus(
    false,
    manager.WithUpdateUserStatusUserId("<USER_ID>"),
    manager.WithUpdateUserStatusEmail("<EMAIL>"),
    manager.WithUpdateUserStatusReason("<REASON>"),
)
```
