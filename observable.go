// https://schema.ocsf.io/objects/observable
package gcs

import (
	"github.com/go-playground/validator/v10"
)

type Observable struct {
	Name       string      `json:"name" validate:"required"`
	Reputation *Reputation `json:"reputation"`
	Type       string      `json:"type"`
	TypeID     int         `json:"type_id" validate:"required"`
	Value      string      `json:"value"`
}

func ValidateObservable(ob *Observable) (*Observable, error) {
	err := validator.New().Struct(ob)
	if err != nil {
		ValidatorErrLog(err)
		return nil, err
	}

	return ob, nil
}
