// https://schema.ocsf.io/objects/geolocaiton
package gcs

import (
	"github.com/go-playground/validator/v10"
)

type GeoLocation struct {
	City        string    `json:"city" validate:"required"`
	Continent   string    `json:"continent" validate:"required"`
	Coordinates []float64 `json:"coordinates" validate:"required"`
	Country     string    `json:"country" validate:"required"`
	Description string    `json:"desc"`
	ISP         string    `json:"isp"`
	OnPremises  bool      `json:"is_on_premises"`
	PostalCode  string    `json:"postal_code"`
	Provider    string    `json:"provider"`
	Region      string    `json:"region"`
}

func ValidateGeoLocation(loc *GeoLocation) (*GeoLocation, error) {
	err := validator.New().Struct(loc)
	if err != nil {
		ValidatorErrLog(err)
		return nil, err
	}

	return loc, nil
}
