package models

import (
	"encoding/json"
	"errors"
)

// VCSNamespacesList Model
type VcsNamespaceList struct {
	// Total number of namespaces that matched your query.
	Total int `json:"total"`
	// List of namespaces.
	Namespaces []VcsNamespace `json:"namespaces"`

	// Used by Decode() method
	data []byte
}

func (model VcsNamespaceList) New(data []byte) *VcsNamespaceList {
	model.data = data
	return &model
}

func (model *VcsNamespaceList) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
