```go
package main

import (
    "fmt"
    "github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/client"
    "github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/mysql"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
    client.WithProject("<YOUR_PROJECT_ID>")
)

service := mysql.New(client)

response, error := service.UpdatePooler(
    "<DATABASE_ID>",
    mysql.WithUpdatePoolerMode("transaction"),
    mysql.WithUpdatePoolerMaxConnections(10),
    mysql.WithUpdatePoolerDefaultPoolSize(1),
    mysql.WithUpdatePoolerReadWriteSplitting(false),
    mysql.WithUpdatePoolerPoolerCpuRequest("<POOLER_CPU_REQUEST>"),
    mysql.WithUpdatePoolerPoolerCpuLimit("<POOLER_CPU_LIMIT>"),
    mysql.WithUpdatePoolerPoolerMemoryRequest("<POOLER_MEMORY_REQUEST>"),
    mysql.WithUpdatePoolerPoolerMemoryLimit("<POOLER_MEMORY_LIMIT>"),
)
```
