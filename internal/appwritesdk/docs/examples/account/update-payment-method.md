```go
package main

import (
    "fmt"
    "github.com/appwrite/sdk-for-go/client"
    "github.com/appwrite/sdk-for-go/account"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
    client.WithProject("<YOUR_PROJECT_ID>")
)

service := account.New(client)

response, error := service.UpdatePaymentMethod(
    "<PAYMENT_METHOD_ID>",
    1,
    2026,
    account.WithUpdatePaymentMethodState("<STATE>"),
)
```
