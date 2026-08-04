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

response, error := service.CreatePAR(
    "<CLIENT_ID>",
    "https://example.com",
    "code",
    oauth2.WithCreatePARScope("<SCOPE>"),
    oauth2.WithCreatePARState("<STATE>"),
    oauth2.WithCreatePARNonce("<NONCE>"),
    oauth2.WithCreatePARCodeChallenge("<CODE_CHALLENGE>"),
    oauth2.WithCreatePARCodeChallengeMethod("s256"),
    oauth2.WithCreatePARPrompt("<PROMPT>"),
    oauth2.WithCreatePARMaxAge(0),
    oauth2.WithCreatePARAuthorizationDetails("<AUTHORIZATION_DETAILS>"),
    oauth2.WithCreatePARResource(""),
    oauth2.WithCreatePARAudience("<AUDIENCE>"),
)
```
