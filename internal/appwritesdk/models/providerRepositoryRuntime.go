package models

import (
	"encoding/json"
	"errors"
)

// ProviderRepositoryRuntime Model
type ProviderRepositoryRuntime struct {
	// VCS (Version Control System) repository ID.
	Id string `json:"id"`
	// VCS (Version Control System) repository name.
	Name string `json:"name"`
	// VCS (Version Control System) organization name
	Organization string `json:"organization"`
	// VCS (Version Control System) provider name.
	Provider string `json:"provider"`
	// Is VCS (Version Control System) repository private?
	Private bool `json:"private"`
	// VCS (Version Control System) repository's default branch name.
	DefaultBranch string `json:"defaultBranch"`
	// VCS (Version Control System) installation ID.
	ProviderInstallationId string `json:"providerInstallationId"`
	// Is VCS (Version Control System) repository authorized for the installation?
	Authorized bool `json:"authorized"`
	// Last commit date in ISO 8601 format.
	PushedAt string `json:"pushedAt"`
	// Environment variables found in .env files
	Variables []string `json:"variables"`
	// Auto-detected runtime. Empty if type is not "runtime".
	Runtime string `json:"runtime"`

	// Used by Decode() method
	data []byte
}

func (model ProviderRepositoryRuntime) New(data []byte) *ProviderRepositoryRuntime {
	model.data = data
	return &model
}

func (model *ProviderRepositoryRuntime) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
