```go
package main

import (
    "fmt"
    "github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/client"
    "github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/organizations"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
    client.WithProject("<YOUR_PROJECT_ID>")
)

service := organizations.New(client)

response, error := service.Create(
    "<ORGANIZATION_ID>",
    "<NAME>",
    "tier-0",
    organizations.WithCreatePaymentMethodId("<PAYMENT_METHOD_ID>"),
    organizations.WithCreateBillingAddressId("<BILLING_ADDRESS_ID>"),
    organizations.WithCreateInvites([]string{}),
    organizations.WithCreateCouponId("<COUPON_ID>"),
    organizations.WithCreateTaxId("<TAX_ID>"),
    organizations.WithCreateBudget(0),
    organizations.WithCreatePlatform("appwrite"),
)
```
