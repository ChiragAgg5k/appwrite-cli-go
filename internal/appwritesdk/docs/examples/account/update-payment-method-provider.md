```go
package main

import (
    "fmt"
    "github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/client"
    "github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/account"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
    client.WithProject("<YOUR_PROJECT_ID>")
)

service := account.New(client)

response, error := service.UpdatePaymentMethodProvider(
    "<PAYMENT_METHOD_ID>",
    "<PROVIDER_METHOD_ID>",
    "<NAME>",
    account.WithUpdatePaymentMethodProviderState("<STATE>"),
)
```
