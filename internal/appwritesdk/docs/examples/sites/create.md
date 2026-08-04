```go
package main

import (
    "fmt"
    "github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/client"
    "github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/sites"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
    client.WithProject("<YOUR_PROJECT_ID>")
)

service := sites.New(client)

response, error := service.Create(
    "<SITE_ID>",
    "<NAME>",
    "analog",
    "node-14.5",
    sites.WithCreateEnabled(false),
    sites.WithCreateLogging(false),
    sites.WithCreateTimeout(1),
    sites.WithCreateInstallCommand("<INSTALL_COMMAND>"),
    sites.WithCreateBuildCommand("<BUILD_COMMAND>"),
    sites.WithCreateStartCommand("<START_COMMAND>"),
    sites.WithCreateOutputDirectory("<OUTPUT_DIRECTORY>"),
    sites.WithCreateAdapter("static"),
    sites.WithCreateInstallationId("<INSTALLATION_ID>"),
    sites.WithCreateFallbackFile("<FALLBACK_FILE>"),
    sites.WithCreateProviderRepositoryId("<PROVIDER_REPOSITORY_ID>"),
    sites.WithCreateProviderBranch("<PROVIDER_BRANCH>"),
    sites.WithCreateProviderSilentMode(false),
    sites.WithCreateProviderRootDirectory("<PROVIDER_ROOT_DIRECTORY>"),
    sites.WithCreateProviderBranches([]string{}),
    sites.WithCreateProviderPaths([]string{}),
    sites.WithCreateBuildSpecification(""),
    sites.WithCreateRuntimeSpecification(""),
    sites.WithCreateDeploymentRetention(0),
)
```
