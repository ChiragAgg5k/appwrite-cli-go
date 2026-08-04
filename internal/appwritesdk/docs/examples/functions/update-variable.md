```go
package main

import (
    "fmt"
    "github.com/appwrite/sdk-for-go/client"
    "github.com/appwrite/sdk-for-go/functions"
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
