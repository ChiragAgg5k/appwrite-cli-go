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

response, error := service.UpdateBypassRule(
    "<RULE_ID>",
    waf.WithUpdateBypassRuleResourceType("api"),
    waf.WithUpdateBypassRuleResourceId("<RESOURCE_ID>"),
    waf.WithUpdateBypassRuleName("<NAME>"),
    waf.WithUpdateBypassRuleDescription("<DESCRIPTION>"),
    waf.WithUpdateBypassRulePriority(-100000),
    waf.WithUpdateBypassRuleEnabled(false),
    waf.WithUpdateBypassRuleConditions(""),
)
```
