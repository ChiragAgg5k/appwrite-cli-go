package models

import (
	"encoding/json"
	"errors"
)

// DowngradeFeedback Model
type DowngradeFeedback struct {
	// Feedback ID.
	Id string `json:"$id"`
	// Feedback creation date in ISO 8601 format.
	CreatedAt string `json:"$createdAt"`
	// Feedback update date in ISO 8601 format.
	UpdatedAt string `json:"$updatedAt"`
	// Feedback reason
	Title string `json:"title"`
	// Feedback message
	Message string `json:"message"`
	// Plan ID downgrading from
	FromPlanId string `json:"fromPlanId"`
	// Plan ID downgrading to
	ToPlanId string `json:"toPlanId"`
	// Organization ID
	TeamId string `json:"teamId"`
	// User ID who submitted feedback
	UserId string `json:"userId"`
	// Console version
	Version string `json:"version"`

	// Used by Decode() method
	data []byte
}

func (model DowngradeFeedback) New(data []byte) *DowngradeFeedback {
	model.data = data
	return &model
}

func (model *DowngradeFeedback) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
