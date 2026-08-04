```go
package main

import (
    "fmt"
    "github.com/appwrite/sdk-for-go/client"
    "github.com/appwrite/sdk-for-go/postgresql"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
    client.WithProject("<YOUR_PROJECT_ID>")
)

service := postgresql.New(client)

response, error := service.CreateBackupPolicy(
    "<DATABASE_ID>",
    "<POLICY_ID>",
    "<NAME>",
    "",
    1,
    postgresql.WithCreateBackupPolicyType("full"),
    postgresql.WithCreateBackupPolicyEnabled(false),
)
```
