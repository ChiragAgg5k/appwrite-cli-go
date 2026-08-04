```go
package main

import (
    "fmt"
    "github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/client"
    "github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/functions"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
    client.WithProject("<YOUR_PROJECT_ID>")
)

service := functions.New(client)

response, error := service.Update(
    "<FUNCTION_ID>",
    "<NAME>",
    functions.WithUpdateRuntime("node-14.5"),
    functions.WithUpdateExecute([]string{"any"}),
    functions.WithUpdateEvents([]string{}),
    functions.WithUpdateSchedule(""),
    functions.WithUpdateTimeout(1),
    functions.WithUpdateEnabled(false),
    functions.WithUpdateLogging(false),
    functions.WithUpdateEntrypoint("<ENTRYPOINT>"),
    functions.WithUpdateCommands("<COMMANDS>"),
    functions.WithUpdateScopes([]string{}),
    functions.WithUpdateInstallationId("<INSTALLATION_ID>"),
    functions.WithUpdateProviderRepositoryId("<PROVIDER_REPOSITORY_ID>"),
    functions.WithUpdateProviderBranch("<PROVIDER_BRANCH>"),
    functions.WithUpdateProviderSilentMode(false),
    functions.WithUpdateProviderRootDirectory("<PROVIDER_ROOT_DIRECTORY>"),
    functions.WithUpdateProviderBranches([]string{}),
    functions.WithUpdateProviderPaths([]string{}),
    functions.WithUpdateBuildSpecification(""),
    functions.WithUpdateRuntimeSpecification(""),
    functions.WithUpdateDeploymentRetention(0),
)
```
