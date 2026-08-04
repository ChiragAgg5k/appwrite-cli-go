```go
package main

import (
    "fmt"
    "github.com/appwrite/sdk-for-go/client"
    "github.com/appwrite/sdk-for-go/manager"
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
