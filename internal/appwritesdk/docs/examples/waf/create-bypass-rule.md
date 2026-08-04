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

response, error := service.CreateBypassRule(
    "<RULE_ID>",
    "api",
    "<NAME>",
    waf.WithCreateBypassRuleResourceId("<RESOURCE_ID>"),
    waf.WithCreateBypassRuleDescription("<DESCRIPTION>"),
    waf.WithCreateBypassRulePriority(-100000),
    waf.WithCreateBypassRuleEnabled(false),
    waf.WithCreateBypassRuleConditions(""),
)
```
