// https://schema.ocsf.io/objects/keyboard_info
package gcs

import (
	"github.com/go-playground/validator/v10"
)

type KeyboardInfo struct {
	FunctionKeys    int    `json:"function_keys,omitempty"`
	IME             string `json:"ime,omitempty"`
	KeyboardLayout  string `json:"keyboard_layout,omitempty"`
	KeyboardSubtype int    `json:"keyboard_subtype,omitempty"`
	KeyboardType    string `json:"keyboard_type,omitempty"`
}

func ValidateKeyboardInfo(ki *KeyboardInfo) (*KeyboardInfo, error) {
	err := validator.New().Struct(ki)
	if err != nil {
		ValidatorErrLog(err)
		return nil, err
	}

	return ki, nil
}
