// https://schema.ocsf.io/objects/image
package gcs

import (
	"github.com/go-playground/validator/v10"
)

type Image struct {
	Tag    string   `json:"tag,omitempty"`
	Labels []string `json:"labels,omitempty"`
	Name   string   `json:"name,omitempty"`
	Path   string   `json:"path,omitempty"`
	UID    string   `json:"uid,omitempty"`
}

func ValidateImage(image *Image) (*Image, error) {
	err := validator.New().Struct(image)
	if err != nil {
		ValidatorErrLog(err)
		return nil, err
	}

	return image, nil
}
