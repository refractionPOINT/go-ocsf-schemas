// https://schema.ocsf.io/objects/schema_extension
package gcs

import (
	"github.com/go-playground/validator/v10"
)

type SchemaExtension struct {
	Name    string `json:"name" validate:"required"`
	UID     string `json:"uid" validate:"required"`
	Version string `json:"version" validate:"required"`
}

func ValidateSchemaExtension(se *SchemaExtension) (*SchemaExtension, error) {
	err := validator.New().Struct(se)
	if err != nil {
		ValidatorErrLog(err)
		return nil, err
	}

	return se, nil
}
