// https://schema.ocsf.io/objects/geodhaiton
package gcs

import (
	"github.com/go-playground/validator/v10"
)

type Display struct {
}

func ValidateDisplay(display *Display) (*Display, error) {
	err := validator.New().Struct(display)
	if err != nil {
		ValidatorErrLog(err)
		return nil, err
	}

	return display, nil
}
