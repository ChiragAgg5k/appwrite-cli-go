package models

import (
	"encoding/json"
	"errors"
)

// DatabaseMigration Model
type DatabaseMigration struct {
	// Database migration ID.
	Id string `json:"$id"`
	// Migration creation time in ISO 8601 format.
	CreatedAt string `json:"$createdAt"`
	// Migration update time in ISO 8601 format.
	UpdatedAt string `json:"$updatedAt"`
	// Project ID that owns the migrating database.
	ProjectId string `json:"projectId"`
	// Logical database ID being migrated.
	DatabaseId string `json:"databaseId"`
	// Dedicated compute specification provisioned for the migration target.
	Specification string `json:"specification"`
	// Migration phase. Possible values: pending, provisioned, capturing,
	// backfilling, catching_up, verifying, ready_to_cutover, cutover, soaking,
	// done, failed, rolled_back.
	Phase string `json:"phase"`
	// Number of times a migration step has failed and been recorded.
	Attempt int `json:"attempt"`
	// Reason the most recent migration step failed, empty while none has.
	LastError string `json:"lastError"`
	// Number of documents still pending replication to the target.
	LagDocuments int `json:"lagDocuments"`
	// Time the migrated data was verified against the source in ISO 8601 format.
	VerifiedAt string `json:"verifiedAt"`
	// Time routing was flipped to the target in ISO 8601 format.
	CutoverAt string `json:"cutoverAt"`
	// Time the post-cutover soak window ends in ISO 8601 format.
	SoakUntil string `json:"soakUntil"`
	// Whether the migration cuts over automatically once ready.
	AutoCutover bool `json:"autoCutover"`
	// Whether the migration is paused.
	Paused bool `json:"paused"`

	// Used by Decode() method
	data []byte
}

func (model DatabaseMigration) New(data []byte) *DatabaseMigration {
	model.data = data
	return &model
}

func (model *DatabaseMigration) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
