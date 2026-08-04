```go
package main

import (
    "fmt"
    "github.com/appwrite/sdk-for-go/client"
    "github.com/appwrite/sdk-for-go/backups"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
    client.WithProject("<YOUR_PROJECT_ID>")
)

service := backups.New(client)

response, error := service.UpdatePolicy(
    "<POLICY_ID>",
    backups.WithUpdatePolicyName("<NAME>"),
    backups.WithUpdatePolicyRetention(1),
    backups.WithUpdatePolicySchedule(""),
    backups.WithUpdatePolicyEnabled(false),
)
```
