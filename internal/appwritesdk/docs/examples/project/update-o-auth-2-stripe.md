```go
package main

import (
    "fmt"
    "github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/client"
    "github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/project"
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
