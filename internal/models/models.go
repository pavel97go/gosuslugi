package models

import "time"

type ApplicationType string
type ApplicationStatus string

const (
	TypePassport    ApplicationType = "passport"
	TypeCertificate ApplicationType = "certificate"

	StatusDraft     ApplicationStatus = "draft"
	StatusSubmitted ApplicationStatus = "submitted"
	StatusApproved  ApplicationStatus = "approved"
	StatusRejected  ApplicationStatus = "rejected"
)

type Application struct {
	ID           int64             `json:"id"`
	CitizenName  string            `json:"citizen_name"`
	DocumentType ApplicationType   `json:"document_type"`
	Data         map[string]any    `json:"data"`
	Status       ApplicationStatus `json:"status"`
	CreatedAt    time.Time         `json:"created_at"`
	UpdatedAt    time.Time         `json:"updated_at"`
}

type ApplicationFilter struct {
	Status       *ApplicationStatus
	DocumentType *ApplicationType
	Limit        int32
	Offset       int32
}
