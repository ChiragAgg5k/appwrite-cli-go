```go
package main

import (
    "fmt"
    "github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/client"
    "github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/messaging"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
    client.WithProject("<YOUR_PROJECT_ID>")
)

service := messaging.New(client)

response, error := service.CreateTelesignProvider(
    "<PROVIDER_ID>",
    "<NAME>",
    messaging.WithCreateTelesignProviderFrom("+12065550100"),
    messaging.WithCreateTelesignProviderCustomerId("<CUSTOMER_ID>"),
    messaging.WithCreateTelesignProviderApiKey("<API_KEY>"),
    messaging.WithCreateTelesignProviderEnabled(false),
)
```
