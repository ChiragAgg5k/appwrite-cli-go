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

response, error := service.CreatePlanEstimation(
    "<ORGANIZATION_ID>",
    "tier-0",
    organizations.WithCreatePlanEstimationInvites([]string{}),
    organizations.WithCreatePlanEstimationCouponId("<COUPON_ID>"),
)
```
