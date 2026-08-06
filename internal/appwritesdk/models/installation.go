package models

import (
	"encoding/json"
	"errors"
)

// Installation Model
type Installation struct {
	// Function ID.
	Id string `json:"$id"`
	// Function creation date in ISO 8601 format.
	CreatedAt string `json:"$createdAt"`
	// Function update date in ISO 8601 format.
	UpdatedAt string `json:"$updatedAt"`
	// VCS (Version Control System) provider name.
	Provider string `json:"provider"`
	// VCS (Version Control System) organization name.
	Organization string `json:"organization"`
	// VCS (Version Control System) installation ID.
	ProviderInstallationId string `json:"providerInstallationId"`

	// Used by Decode() method
	data []byte
}

func (model Installation) New(data []byte) *Installation {
	model.data = data
	return &model
}

func (model *Installation) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
