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

response, error := service.AuthorizePost(
    oauth2.WithAuthorizePostClientId("<CLIENT_ID>"),
    oauth2.WithAuthorizePostRedirectUri("https://example.com"),
    oauth2.WithAuthorizePostResponseType(""),
    oauth2.WithAuthorizePostScope("<SCOPE>"),
    oauth2.WithAuthorizePostState("<STATE>"),
    oauth2.WithAuthorizePostNonce("<NONCE>"),
    oauth2.WithAuthorizePostCodeChallenge("<CODE_CHALLENGE>"),
    oauth2.WithAuthorizePostCodeChallengeMethod("s256"),
    oauth2.WithAuthorizePostPrompt("<PROMPT>"),
    oauth2.WithAuthorizePostMaxAge(0),
    oauth2.WithAuthorizePostAuthorizationDetails("<AUTHORIZATION_DETAILS>"),
    oauth2.WithAuthorizePostResource(""),
    oauth2.WithAuthorizePostAudience("<AUDIENCE>"),
    oauth2.WithAuthorizePostRequestUri("<REQUEST_URI>"),
)
```
