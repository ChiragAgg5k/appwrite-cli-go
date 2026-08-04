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

response, error := service.UpdateBillingAddress(
    "<BILLING_ADDRESS_ID>",
    "<COUNTRY>",
    "<CITY>",
    "<STREET_ADDRESS>",
    account.WithUpdateBillingAddressAddressLine2("<ADDRESS_LINE2>"),
    account.WithUpdateBillingAddressState("<STATE>"),
    account.WithUpdateBillingAddressPostalCode("<POSTAL_CODE>"),
)
```
