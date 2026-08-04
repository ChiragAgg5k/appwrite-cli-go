```go
package main

import (
    "fmt"
    "github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/client"
    "github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/domains"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
    client.WithProject("<YOUR_PROJECT_ID>")
)

service := domains.New(client)

response, error := service.CreatePurchase(
    "",
    "<ORGANIZATION_ID>",
    "<FIRST_NAME>",
    "<LAST_NAME>",
    "email@example.com",
    "+12065550100",
    "<BILLING_ADDRESS_ID>",
    "<PAYMENT_METHOD_ID>",
    domains.WithCreatePurchaseAddressLine3("<ADDRESS_LINE3>"),
    domains.WithCreatePurchaseCompanyName("<COMPANY_NAME>"),
    domains.WithCreatePurchasePeriodYears(1),
    domains.WithCreatePurchaseAutoRenewal(false),
)
```
