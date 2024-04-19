// https://schema.ocsf.io/objects/enrichment
package gcs

import (
	"github.com/go-playground/validator/v10"
)

type Enrichment struct {
	Data     string `json:"data" validate:"required"`
	Name     string `json:"name" validate:"required"`
	Provider string `json:"provider"`
	Type     string `json:"type"`
	Value    string `json:"value" validate:"required"`
}

func ValidateEnrichment(en *Enrichment) (*Enrichment, error) {
	err := validator.New().Struct(en)
	if err != nil {
		ValidatorErrLog(err)
		return nil, err
	}

	return en, nil
}
