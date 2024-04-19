// https://schema.ocsf.io/objects/module
package gcs

import (
	"github.com/go-playground/validator/v10"
)

type Module struct {
	BaseAddress  string `json:"base_address" validate:"required"`
	File         File   `json:"file" validate:"required"`
	FunctionName string `json:"function_name"`
	LoadType     string `json:"load_type"`
	LoadTypeID   int    `json:"load_type_id" validate:"required"`
	StartAddress string `json:"start_address" validate:"required"`
	Type         string `json:"type" validate:"required"`
}

func ValidateModule(mod *Module) (*Module, error) {
	err := validator.New().Struct(mod)
	if err != nil {
		ValidatorErrLog(err)
		return nil, err
	}

	return mod, nil
}
