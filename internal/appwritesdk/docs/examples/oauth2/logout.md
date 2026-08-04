```go
package main

import (
    "fmt"
    "github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/client"
    "github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/oauth2"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
    client.WithProject("<YOUR_PROJECT_ID>")
)

service := oauth2.New(client)

response, error := service.Logout(
    oauth2.WithLogoutIdTokenHint("<ID_TOKEN_HINT>"),
    oauth2.WithLogoutLogoutHint("<LOGOUT_HINT>"),
    oauth2.WithLogoutClientId("<CLIENT_ID>"),
    oauth2.WithLogoutPostLogoutRedirectUri("https://example.com"),
    oauth2.WithLogoutState("<STATE>"),
    oauth2.WithLogoutUiLocales("<UI_LOCALES>"),
)
```
