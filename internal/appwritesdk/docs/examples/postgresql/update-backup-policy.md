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

response, error := service.UpdateBackupPolicy(
    "<DATABASE_ID>",
    "<POLICY_ID>",
    postgresql.WithUpdateBackupPolicyName("<NAME>"),
    postgresql.WithUpdateBackupPolicySchedule(""),
    postgresql.WithUpdateBackupPolicyRetention(1),
    postgresql.WithUpdateBackupPolicyEnabled(false),
)
```
