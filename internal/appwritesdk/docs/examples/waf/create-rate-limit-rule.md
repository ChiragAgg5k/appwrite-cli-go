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

response, error := service.CreateRateLimitRule(
    "<RULE_ID>",
    "api",
    "<NAME>",
    1,
    1,
    waf.WithCreateRateLimitRuleResourceId("<RESOURCE_ID>"),
    waf.WithCreateRateLimitRuleDescription("<DESCRIPTION>"),
    waf.WithCreateRateLimitRuleKey("ip"),
    waf.WithCreateRateLimitRulePriority(-100000),
    waf.WithCreateRateLimitRuleEnabled(false),
    waf.WithCreateRateLimitRuleConditions(""),
)
```
