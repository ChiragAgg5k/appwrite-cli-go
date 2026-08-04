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

response, error := service.UpdateChallengeRule(
    "<RULE_ID>",
    waf.WithUpdateChallengeRuleResourceType("api"),
    waf.WithUpdateChallengeRuleResourceId("<RESOURCE_ID>"),
    waf.WithUpdateChallengeRuleName("<NAME>"),
    waf.WithUpdateChallengeRuleDescription("<DESCRIPTION>"),
    waf.WithUpdateChallengeRuleChallengeType("compute"),
    waf.WithUpdateChallengeRulePriority(-100000),
    waf.WithUpdateChallengeRuleEnabled(false),
    waf.WithUpdateChallengeRuleConditions(""),
    waf.WithUpdateChallengeRuleDifficulty(1),
    waf.WithUpdateChallengeRuleTtl(900),
)
```
