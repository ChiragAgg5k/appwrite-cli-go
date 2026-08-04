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

response, error := service.CreateSHAUser(
    "<USER_ID>",
    "email@example.com",
    "password",
    users.WithCreateSHAUserPasswordVersion("sha1"),
    users.WithCreateSHAUserName("<NAME>"),
)
```
