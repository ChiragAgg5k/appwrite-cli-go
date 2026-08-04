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

response, error := service.UpdateRedirectRule(
    "<RULE_ID>",
    waf.WithUpdateRedirectRuleResourceType("api"),
    waf.WithUpdateRedirectRuleResourceId("<RESOURCE_ID>"),
    waf.WithUpdateRedirectRuleName("<NAME>"),
    waf.WithUpdateRedirectRuleDescription("<DESCRIPTION>"),
    waf.WithUpdateRedirectRuleLocation("<LOCATION>"),
    waf.WithUpdateRedirectRuleStatusCode(300),
    waf.WithUpdateRedirectRulePriority(-100000),
    waf.WithUpdateRedirectRuleEnabled(false),
    waf.WithUpdateRedirectRuleConditions(""),
)
```
