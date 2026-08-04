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

response, error := service.CreateRedirectRule(
    "<RULE_ID>",
    "api",
    "<NAME>",
    "<LOCATION>",
    300,
    waf.WithCreateRedirectRuleResourceId("<RESOURCE_ID>"),
    waf.WithCreateRedirectRuleDescription("<DESCRIPTION>"),
    waf.WithCreateRedirectRulePriority(-100000),
    waf.WithCreateRedirectRuleEnabled(false),
    waf.WithCreateRedirectRuleConditions(""),
)
```
