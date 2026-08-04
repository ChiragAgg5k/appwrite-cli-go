```go
package main

import (
    "fmt"
    "github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/client"
    "github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/backups"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
    client.WithProject("<YOUR_PROJECT_ID>")
)

service := backups.New(client)

response, error := service.CreatePolicy(
    "<POLICY_ID>",
    []string{},
    1,
    "",
    backups.WithCreatePolicyName("<NAME>"),
    backups.WithCreatePolicyResourceId("<RESOURCE_ID>"),
    backups.WithCreatePolicyEnabled(false),
)
```
