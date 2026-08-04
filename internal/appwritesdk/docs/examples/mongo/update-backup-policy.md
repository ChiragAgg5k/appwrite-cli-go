```go
package main

import (
    "fmt"
    "github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/client"
    "github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/mongo"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
    client.WithProject("<YOUR_PROJECT_ID>")
)

service := mongo.New(client)

response, error := service.UpdateBackupPolicy(
    "<DATABASE_ID>",
    "<POLICY_ID>",
    mongo.WithUpdateBackupPolicyName("<NAME>"),
    mongo.WithUpdateBackupPolicySchedule(""),
    mongo.WithUpdateBackupPolicyRetention(1),
    mongo.WithUpdateBackupPolicyEnabled(false),
)
```
