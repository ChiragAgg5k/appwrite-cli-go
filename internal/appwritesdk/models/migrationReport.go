package models

import (
	"encoding/json"
	"errors"
)

// MigrationReport Model
type MigrationReport struct {
	// Number of users to be migrated.
	User int `json:"user"`
	// Number of teams to be migrated.
	Team int `json:"team"`
	// Number of databases to be migrated.
	Database int `json:"database"`
	// Number of rows to be migrated.
	Row int `json:"row"`
	// Number of files to be migrated.
	File int `json:"file"`
	// Number of buckets to be migrated.
	Bucket int `json:"bucket"`
	// Number of functions to be migrated.
	Function int `json:"function"`
	// Number of platforms to be migrated.
	Platform int `json:"platform"`
	// Number of API keys to be migrated.
	ApiKey int `json:"api-key"`
	// Number of project variables to be migrated.
	ProjectVariable int `json:"project-variable"`
	// Number of webhooks to be migrated.
	Webhook int `json:"webhook"`
	// Number of auth-method configs to be migrated (always 0 or 1 — the
	// project-level flag bundle).
	AuthMethods int `json:"auth-methods"`
	// Number of protocol configs to be migrated (always 0 or 1 — the
	// project-level REST/GraphQL/WebSocket flags).
	ProjectProtocols int `json:"project-protocols"`
	// Number of label sets to be migrated (always 0 or 1 — the project-level
	// RBAC label array).
	ProjectLabels int `json:"project-labels"`
	// Number of service configs to be migrated (always 0 or 1 — the
	// project-level enable/disable flags for all 17 services).
	ProjectServices int `json:"project-services"`
	// Number of policy bundles to be migrated (always 0 or 1 — the
	// project-level security policies covering password rules, session behavior,
	// user limits, and membership privacy).
	Policies int `json:"policies"`
	// Number of SMTP configurations to be migrated (always 0 or 1 — the
	// project-level custom SMTP settings; password is not exposed by the source
	// API).
	Smtp int `json:"smtp"`
	// Number of custom-domain proxy rules to be migrated. Auto-generated
	// `.appwrite.network` rules are skipped — they are recreated by parent
	// Function/Site migration.
	Rule int `json:"rule"`
	// Number of custom email templates to be migrated (one per templateId ×
	// locale pair).
	ProjectEmailTemplate int `json:"project-email-template"`
	// Number of sites to be migrated.
	Site int `json:"site"`
	// Number of providers to be migrated.
	Provider int `json:"provider"`
	// Number of topics to be migrated.
	Topic int `json:"topic"`
	// Number of subscribers to be migrated.
	Subscriber int `json:"subscriber"`
	// Number of messages to be migrated.
	Message int `json:"message"`
	// Size of files to be migrated in mb.
	Size int `json:"size"`
	// Version of the Appwrite instance to be migrated.
	Version string `json:"version"`
	// Number of OAuth2 provider configurations to be migrated. Secrets
	// (clientSecret, p8File) are never migrated — destination admin must
	// re-enter them per provider.
	Oauth2Provider int `json:"oauth2-provider"`
	// Number of backup policies to be migrated.
	BackupPolicy int `json:"backup-policy"`

	// Used by Decode() method
	data []byte
}

func (model MigrationReport) New(data []byte) *MigrationReport {
	model.data = data
	return &model
}

func (model *MigrationReport) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
