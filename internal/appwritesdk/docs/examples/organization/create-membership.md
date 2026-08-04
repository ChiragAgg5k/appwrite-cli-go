```go
package main

import (
    "fmt"
    "github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/client"
    "github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/organization"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
    client.WithProject("<YOUR_PROJECT_ID>")
)

service := organization.New(client)

response, error := service.CreateMembership(
    []string{},
    organization.WithCreateMembershipEmail("email@example.com"),
    organization.WithCreateMembershipUserId("<USER_ID>"),
    organization.WithCreateMembershipPhone("+12065550100"),
    organization.WithCreateMembershipUrl("https://example.com"),
    organization.WithCreateMembershipName("<NAME>"),
)
```
