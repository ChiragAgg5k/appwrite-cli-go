```go
package main

import (
    "fmt"
    "github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/client"
    "github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/apps"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
    client.WithProject("<YOUR_PROJECT_ID>")
)

service := apps.New(client)

response, error := service.Update(
    "<APP_ID>",
    "<NAME>",
    apps.WithUpdateDescription("<DESCRIPTION>"),
    apps.WithUpdateClientUri("https://example.com"),
    apps.WithUpdateLogoUri("https://example.com"),
    apps.WithUpdatePrivacyPolicyUrl("https://example.com"),
    apps.WithUpdateTermsUrl("https://example.com"),
    apps.WithUpdateContacts([]string{}),
    apps.WithUpdateTagline("<TAGLINE>"),
    apps.WithUpdateTags([]string{}),
    apps.WithUpdateImages([]string{}),
    apps.WithUpdateSupportUrl("https://example.com"),
    apps.WithUpdateDataDeletionUrl("https://example.com"),
    apps.WithUpdateEnabled(false),
    apps.WithUpdateRedirectUris([]string{}),
    apps.WithUpdatePostLogoutRedirectUris([]string{}),
    apps.WithUpdateType("public"),
    apps.WithUpdateDeviceFlow(false),
    apps.WithUpdateInstallationScopes([]string{}),
    apps.WithUpdateInstallationRedirectUrl("https://example.com"),
)
```
