```go
package main

import (
    "fmt"
    "github.com/appwrite/sdk-for-go/client"
    "github.com/appwrite/sdk-for-go/organizations"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
    client.WithProject("<YOUR_PROJECT_ID>")
)

service := organizations.New(client)

response, error := service.UpdatePlan(
    "<ORGANIZATION_ID>",
    "tier-0",
    organizations.WithUpdatePlanPaymentMethodId("<PAYMENT_METHOD_ID>"),
    organizations.WithUpdatePlanBillingAddressId("<BILLING_ADDRESS_ID>"),
    organizations.WithUpdatePlanInvites([]string{}),
    organizations.WithUpdatePlanCouponId("<COUPON_ID>"),
    organizations.WithUpdatePlanTaxId("<TAX_ID>"),
    organizations.WithUpdatePlanBudget(0),
)
```
