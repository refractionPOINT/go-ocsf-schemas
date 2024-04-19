// https://schema.ocsf.io/objects/metadata
package gcs

import (
	"github.com/go-playground/validator/v10"
)

type Metadata struct {
	CorrelationUID   string            `json:"correlation_uid,omitempty"`
	EventCode        string            `json:"event_code,omitempty"`
	EventUID         string            `json:"uid,omitempty"`
	Labels           []string          `json:"labels,omitempty"`
	LogLevel         string            `json:"log_level,omitempty"`
	LogName          string            `json:"log_name,omitempty"`
	LogProvider      string            `json:"log_provider,omitempty"`
	LogVersion       string            `json:"log_version,omitempty"`
	LoggedTime       string            `json:"logged_time,omitempty"`
	Loggers          []Logger          `json:"loggers,omitempty"`
	ModifiedTime     string            `json:"modified_time,omitempty"`
	OriginalTime     string            `json:"original_time,omitempty"`
	ProcessedTime    string            `json:"processed_time,omitempty"`
	Product          Product           `json:"product,omitempty"`
	Profiles         []string          `json:"profiles,omitempty"`
	SchemaExtension  SchemaExtension   `json:"extension,omitempty"`
	SchemaExtensions []SchemaExtension `json:"extensions,omitempty"`
	SequenceNumber   int               `json:"sequence,omitempty"`
	TenantUID        string            `json:"tenant_uid,omitempty"`
	Version          string            `json:"version,omitempty"`
}

func ValidateMetadata(met *Metadata) (*Metadata, error) {
	err := validator.New().Struct(met)
	if err != nil {
		ValidatorErrLog(err)
		return nil, err
	}

	return met, nil
}
