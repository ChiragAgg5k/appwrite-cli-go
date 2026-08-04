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

response, error := service.CreateChallengeRule(
    "<RULE_ID>",
    "api",
    "<NAME>",
    waf.WithCreateChallengeRuleResourceId("<RESOURCE_ID>"),
    waf.WithCreateChallengeRuleDescription("<DESCRIPTION>"),
    waf.WithCreateChallengeRuleChallengeType("compute"),
    waf.WithCreateChallengeRulePriority(-100000),
    waf.WithCreateChallengeRuleEnabled(false),
    waf.WithCreateChallengeRuleConditions(""),
    waf.WithCreateChallengeRuleDifficulty(1),
    waf.WithCreateChallengeRuleTtl(900),
)
```
