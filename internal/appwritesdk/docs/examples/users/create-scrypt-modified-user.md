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

response, error := service.CreateScryptModifiedUser(
    "<USER_ID>",
    "email@example.com",
    "password",
    "<PASSWORD_SALT>",
    "<PASSWORD_SALT_SEPARATOR>",
    "<PASSWORD_SIGNER_KEY>",
    users.WithCreateScryptModifiedUserName("<NAME>"),
)
```
