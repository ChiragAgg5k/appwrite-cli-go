package models

import (
    "encoding/json"
    "errors"
)

// TemplateFramework Model
type TemplateFramework struct {
    // Parent framework key.
    Key string `json:"key"`
    // Framework Name.
    Name string `json:"name"`
    // The install command used to install the dependencies.
    InstallCommand string `json:"installCommand"`
    // The build command used to build the deployment.
    BuildCommand string `json:"buildCommand"`
    // The output directory to store the build output.
    OutputDirectory string `json:"outputDirectory"`
    // Path to site in VCS (Version Control System) repository
    ProviderRootDirectory string `json:"providerRootDirectory"`
    // Runtime used during build step of template.
    BuildRuntime string `json:"buildRuntime"`
    // Site framework runtime
    Adapter string `json:"adapter"`
    // Fallback file for SPA. Only relevant for static serve runtime.
    FallbackFile string `json:"fallbackFile"`

    // Used by Decode() method
    data []byte
}

func (model TemplateFramework) New(data []byte) *TemplateFramework {
    model.data = data
    return &model
}

func (model *TemplateFramework) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}