package models

import (
	"encoding/json"
	"errors"
)

// DetectionFramework Model
type DetectionFramework struct {
	// Repository detection type.
	Type string `json:"type"`
	// Environment variables found in .env files
	Variables []DetectionVariable `json:"variables"`
	// Framework
	Framework string `json:"framework"`
	// Site Install Command
	InstallCommand string `json:"installCommand"`
	// Site Build Command
	BuildCommand string `json:"buildCommand"`
	// Site Output Directory
	OutputDirectory string `json:"outputDirectory"`

	// Used by Decode() method
	data []byte
}

func (model DetectionFramework) New(data []byte) *DetectionFramework {
	model.data = data
	return &model
}

func (model *DetectionFramework) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
