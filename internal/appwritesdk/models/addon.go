package models

import (
	"encoding/json"
	"errors"
)

// Addon Model
type Addon struct {
	// Addon ID.
	Id string `json:"$id"`
	// Addon creation time in ISO 8601 format.
	CreatedAt string `json:"$createdAt"`
	// Addon update date in ISO 8601 format.
	UpdatedAt string `json:"$updatedAt"`
	// Addon permissions. [Learn more about
	// permissions](https://appwrite.io/docs/permissions).
	Permissions []string `json:"$permissions"`
	// Addon key
	Key string `json:"key"`
	// Resource type (organization or project)
	ResourceType string `json:"resourceType"`
	// Resource ID
	ResourceId string `json:"resourceId"`
	// Payment status. Possible values: pending (awaiting payment confirmation
	// e.g. 3DS), active (payment confirmed and addon is running).
	Status string `json:"status"`
	// Current value for this billing cycle. For toggle addons: 1 (on) or 0 (off).
	// For numeric addons: the active quantity.
	CurrentValue int `json:"currentValue"`
	// Value to apply at the start of the next billing cycle. Null means no change
	// is scheduled. For toggle addons, 0 means the addon will be removed at the
	// next cycle.
	NextValue int `json:"nextValue"`

	// Used by Decode() method
	data []byte
}

func (model Addon) New(data []byte) *Addon {
	model.data = data
	return &model
}

func (model *Addon) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
