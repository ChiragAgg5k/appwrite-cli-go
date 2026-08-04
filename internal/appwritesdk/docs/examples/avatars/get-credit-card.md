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

response, error := service.GetCreditCard(
    "amex",
    avatars.WithGetCreditCardWidth(0),
    avatars.WithGetCreditCardHeight(0),
    avatars.WithGetCreditCardQuality(-1),
)
```
