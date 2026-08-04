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

response, error := service.EstimationCreateOrganization(
    "tier-0",
    organizations.WithEstimationCreateOrganizationPaymentMethodId("<PAYMENT_METHOD_ID>"),
    organizations.WithEstimationCreateOrganizationInvites([]string{}),
    organizations.WithEstimationCreateOrganizationCouponId("<COUPON_ID>"),
    organizations.WithEstimationCreateOrganizationPlatform("appwrite"),
)
```
