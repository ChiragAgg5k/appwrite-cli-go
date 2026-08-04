```go
package main

import (
    "fmt"
    "github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/client"
    "github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/waf"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
    client.WithProject("<YOUR_PROJECT_ID>")
)

service := waf.New(client)

response, error := service.CreateDenyRule(
    "<RULE_ID>",
    "api",
    "<NAME>",
    waf.WithCreateDenyRuleResourceId("<RESOURCE_ID>"),
    waf.WithCreateDenyRuleDescription("<DESCRIPTION>"),
    waf.WithCreateDenyRulePriority(-100000),
    waf.WithCreateDenyRuleEnabled(false),
    waf.WithCreateDenyRuleConditions(""),
)
```
