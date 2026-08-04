```go
package main

import (
    "fmt"
    "github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/client"
    "github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/postgresql"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
    client.WithProject("<YOUR_PROJECT_ID>")
)

service := postgresql.New(client)

response, error := service.UpdatePooler(
    "<DATABASE_ID>",
    postgresql.WithUpdatePoolerMode("transaction"),
    postgresql.WithUpdatePoolerMaxConnections(10),
    postgresql.WithUpdatePoolerDefaultPoolSize(1),
    postgresql.WithUpdatePoolerReadWriteSplitting(false),
    postgresql.WithUpdatePoolerPoolerCpuRequest("<POOLER_CPU_REQUEST>"),
    postgresql.WithUpdatePoolerPoolerCpuLimit("<POOLER_CPU_LIMIT>"),
    postgresql.WithUpdatePoolerPoolerMemoryRequest("<POOLER_MEMORY_REQUEST>"),
    postgresql.WithUpdatePoolerPoolerMemoryLimit("<POOLER_MEMORY_LIMIT>"),
)
```
