// https://schema.ocsf.io/objects/logger
package gcs

import (
	"github.com/go-playground/validator/v10"
)

type Logger struct {
	Device           string  `json:"device"`
	LogLevel         string  `json:"log_level,omitempty"`
	LogName          string  `json:"log_name"`
	LogProvider      string  `json:"log_provider"`
	LogVersion       string  `json:"log_version,omitempty"`
	LoggedTime       string  `json:"logged_time,omitempty"`
	Name             string  `json:"name"`
	Product          Product `json:"product"`
	TransmissionTime string  `json:"transmit_time,omitempty"`
	UniqueID         string  `json:"uid"`
	Version          string  `json:"version,omitempty"`
}

func ValidateLogger(log *Logger) (*Logger, error) {
	err := validator.New().Struct(log)
	if err != nil {
		ValidatorErrLog(err)
		return nil, err
	}

	return log, nil
}
