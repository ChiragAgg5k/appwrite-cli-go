```go
package main

import (
    "fmt"
    "github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/client"
    "github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/domains"
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
