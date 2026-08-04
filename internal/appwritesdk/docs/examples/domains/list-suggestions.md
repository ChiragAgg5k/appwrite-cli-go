```go
package main

import (
    "fmt"
    "github.com/appwrite/sdk-for-go/client"
    "github.com/appwrite/sdk-for-go/domains"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
    client.WithProject("<YOUR_PROJECT_ID>")
)

service := domains.New(client)

response, error := service.ListSuggestions(
    "<QUERY>",
    domains.WithListSuggestionsTlds([]string{}),
    domains.WithListSuggestionsLimit(0),
    domains.WithListSuggestionsFilterType("premium"),
    domains.WithListSuggestionsPriceMax(0),
    domains.WithListSuggestionsPriceMin(0),
)
```
