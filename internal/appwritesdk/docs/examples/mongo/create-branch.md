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

response, error := service.CreateBranch(
    "<DATABASE_ID>",
    mongo.WithCreateBranchBranchId("<BRANCH_ID>"),
    mongo.WithCreateBranchTtl(300),
)
```
