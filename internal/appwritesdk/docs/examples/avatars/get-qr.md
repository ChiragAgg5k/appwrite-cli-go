```go
package main

import (
    "fmt"
    "github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/client"
    "github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/avatars"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
    client.WithProject("<YOUR_PROJECT_ID>")
)

service := avatars.New(client)

response, error := service.GetQR(
    "<TEXT>",
    avatars.WithGetQRSize(1),
    avatars.WithGetQRMargin(0),
    avatars.WithGetQRDownload(false),
)
```
