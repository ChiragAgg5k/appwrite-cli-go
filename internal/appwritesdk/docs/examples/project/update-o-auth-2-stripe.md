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

response, error := service.UpdateOAuth2Stripe(
    project.WithUpdateOAuth2StripeClientId("<CLIENT_ID>"),
    project.WithUpdateOAuth2StripeApiSecretKey("<API_SECRET_KEY>"),
    project.WithUpdateOAuth2StripeEnabled(false),
)
```
