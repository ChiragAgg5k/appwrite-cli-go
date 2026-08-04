```go
package main

import (
    "fmt"
    "github.com/appwrite/sdk-for-go/client"
    "github.com/appwrite/sdk-for-go/waf"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
    client.WithProject("<YOUR_PROJECT_ID>")
)

service := waf.New(client)

response, error := service.UpdateRateLimitRule(
    "<RULE_ID>",
    waf.WithUpdateRateLimitRuleResourceType("api"),
    waf.WithUpdateRateLimitRuleResourceId("<RESOURCE_ID>"),
    waf.WithUpdateRateLimitRuleName("<NAME>"),
    waf.WithUpdateRateLimitRuleDescription("<DESCRIPTION>"),
    waf.WithUpdateRateLimitRuleLimit(1),
    waf.WithUpdateRateLimitRuleInterval(1),
    waf.WithUpdateRateLimitRuleKey("ip"),
    waf.WithUpdateRateLimitRulePriority(-100000),
    waf.WithUpdateRateLimitRuleEnabled(false),
    waf.WithUpdateRateLimitRuleConditions(""),
)
```
