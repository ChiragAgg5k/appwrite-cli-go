```go
package main

import (
    "fmt"
    "github.com/appwrite/sdk-for-go/client"
    "github.com/appwrite/sdk-for-go/oauth2"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
    client.WithProject("<YOUR_PROJECT_ID>")
)

service := oauth2.New(client)

response, error := service.LogoutPost(
    oauth2.WithLogoutPostIdTokenHint("<ID_TOKEN_HINT>"),
    oauth2.WithLogoutPostLogoutHint("<LOGOUT_HINT>"),
    oauth2.WithLogoutPostClientId("<CLIENT_ID>"),
    oauth2.WithLogoutPostPostLogoutRedirectUri("https://example.com"),
    oauth2.WithLogoutPostState("<STATE>"),
    oauth2.WithLogoutPostUiLocales("<UI_LOCALES>"),
)
```
