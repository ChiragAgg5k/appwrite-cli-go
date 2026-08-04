```go
package main

import (
    "fmt"
    "github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/client"
    "github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/functions"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
    client.WithProject("<YOUR_PROJECT_ID>")
)

service := functions.New(client)

response, error := service.UpdateVariable(
    "<FUNCTION_ID>",
    "<VARIABLE_ID>",
    functions.WithUpdateVariableKey("<KEY>"),
    functions.WithUpdateVariableValue("<VALUE>"),
    functions.WithUpdateVariableSecret(false),
)
```
