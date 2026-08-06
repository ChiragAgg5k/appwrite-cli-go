package models

import (
	"encoding/json"
	"errors"
)

// TemplateSite Model
type TemplateSite struct {
	// Site Template ID.
	Key string `json:"key"`
	// Site Template Name.
	Name string `json:"name"`
	// Short description of template
	Tagline string `json:"tagline"`
	// URL hosting a template demo.
	DemoUrl string `json:"demoUrl"`
	// File URL with preview screenshot in dark theme preference.
	ScreenshotDark string `json:"screenshotDark"`
	// File URL with preview screenshot in light theme preference.
	ScreenshotLight string `json:"screenshotLight"`
	// Site use cases.
	UseCases []string `json:"useCases"`
	// List of frameworks that can be used with this template.
	Frameworks []TemplateFramework `json:"frameworks"`
	// VCS (Version Control System) Provider.
	VcsProvider string `json:"vcsProvider"`
	// VCS (Version Control System) Repository ID
	ProviderRepositoryId string `json:"providerRepositoryId"`
	// VCS (Version Control System) Owner.
	ProviderOwner string `json:"providerOwner"`
	// VCS (Version Control System) branch version (tag).
	ProviderVersion string `json:"providerVersion"`
	// Site variables.
	Variables []TemplateVariable `json:"variables"`

	// Used by Decode() method
	data []byte
}

func (model TemplateSite) New(data []byte) *TemplateSite {
	model.data = data
	return &model
}

func (model *TemplateSite) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
