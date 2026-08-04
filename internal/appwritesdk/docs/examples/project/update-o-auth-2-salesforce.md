```go
package main

import (
    "fmt"
    "github.com/appwrite/sdk-for-go/client"
    "github.com/appwrite/sdk-for-go/project"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
    client.WithProject("<YOUR_PROJECT_ID>")
)

service := project.New(client)

response, error := service.UpdateOAuth2Salesforce(
    project.WithUpdateOAuth2SalesforceCustomerKey("<CUSTOMER_KEY>"),
    project.WithUpdateOAuth2SalesforceCustomerSecret("<CUSTOMER_SECRET>"),
    project.WithUpdateOAuth2SalesforceEnabled(false),
)
```
