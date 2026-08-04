```go
package main

import (
    "fmt"
    "github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/client"
    "github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/manager"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
)

service := manager.New(client)

response, error := service.DeleteCache(
    manager.WithDeleteCacheRegion("fra"),
    manager.WithDeleteCacheCache("cache"),
    manager.WithDeleteCacheAll(false),
    manager.WithDeleteCacheDatabase("console"),
    manager.WithDeleteCacheProjectId("<PROJECT_ID>"),
    manager.WithDeleteCacheCollectionId("<COLLECTION_ID>"),
    manager.WithDeleteCacheDocumentId("<DOCUMENT_ID>"),
)
```
