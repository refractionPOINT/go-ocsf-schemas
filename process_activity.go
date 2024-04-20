// https://schema.ocsf.io/classes/process_activity?extensions=
package gcs

import (
	"github.com/go-playground/validator/v10"
)

// DestinationEndpoint *NetworkEndpoint `json:"dst_endpoint" validate:"required"`
// Enrichments *[]Enrichment `json:"enrichments" valiadte:"omitempty"`
// HTTPRequest *HTTPRequest `json:"http_request" validate:"omitempty"`
// Metadata *Metadata `json:"metadata" validate:"omitempty"`
// Observables *[]Observable `json:"observables" validate:"omitempty"`
// SourceEndpoint NetworkEndpoint `json:"network_endpoint" validate:"omitempty"`

type ProcessActivity struct {
	Activity             string                 `json:"activity_name" validate:"omitempty"`
	ActivityID           uint8                  `json:"activity_id" validate:"omitempty"`
	Actor                *Actor                 `json:"actor" validate:"omitempty"`
	ActualPermissions    int                    `json:"actual_permissions" validate:"omitempty"`
	Category             string                 `json:"category_name" validate:"omitempty"`
	CategoryID           *uint8                 `json:"category_uid" validate:"required"`
	Class                string                 `json:"class_name" validate:"omitempty"`
	ClassID              *uint                  `json:"class_uid" validate:"required"`
	Count                uint                   `json:"count" validate:"omitempty"`
	Device               *Device                `json:"device" validate:"omitempty"`
	Duration             uint                   `json:"duration" validate:"omitempty"`
	EndTime              string                 `json:"end_time" validate:"omitempty"`
	Enrichments          *[]Enrichment          `json:"enrichments" validate:"omitempty"`
	EventTime            string                 `json:"time" validate:"required"`
	ExitCode             int                    `json:"exit_code" validate:"omitempty"`
	InjectionType        *string                `json:"injection_type" validate:"omitempty"`
	InjectionTypeID      int                    `json:"injection_type_id" validate:"omitempty"`
	Message              string                 `json:"message" validate:"omitempty"`
	Metadata             Metadata               `json:"metadata" validate:"omitempty"`
	Module               *Module                `json:"module" validate:"omitempty"`
	Observables          *[]Observable          `json:"observables" validate:"omitempty"`
	Process              Process                `json:"process" validate:"omitempty"`
	RawData              string                 `json:"raw_data" validate:"omitempty"`
	RequestedPermissions int                    `json:"requested_permissions" validate:"omitempty"`
	Severity             string                 `json:"severity" validate:"omitempty"`
	SeverityID           uint8                  `json:"severity_id" validate:"omitempty"`
	StartTime            string                 `json:"start_time" validate:"omitempty"`
	Status               string                 `json:"status" validate:"omitempty"`
	StatusCode           string                 `json:"status_code" validate:"omitempty"`
	StatusDetails        string                 `json:"status_detail" validate:"omitempty"`
	StatusID             uint8                  `json:"status_id" validate:"omitempty"`
	TimezoneOffset       int                    `json:"timezone_offset" validate:"omitempty"`
	TypeID               *uint                  `json:"type_uid" validate:"required"`
	TypeName             string                 `json:"type_name" validate:"omitempty"`
	UnmappedData         map[string]interface{} `json:"unmapped" validate:"omitempty"`
}

func ValidateProcessActivity(activity *ProcessActivity) (*ProcessActivity, error) {
	err := validator.New().Struct(activity)
	if err != nil {
		ValidatorErrLog(err)
		return nil, err
	}

	return activity, nil
}
