// https://schema.ocsf.io/objects/network_interface
package gcs

import (
	"github.com/go-playground/validator/v10"
)

type NetworkInterface struct {
	Hostname           string `json:"hostname" validate:"required"`
	IPAddress          string `json:"ip" validate:"required"`
	MACAddress         string `json:"mac" validate:"required"`
	Name               string `json:"name" validate:"required"`
	Namespace          string `json:"namespace"`
	SubnetPrefixLength int    `json:"subnet_prefix"`
	Type               string `json:"type"`
	TypeID             int    `json:"type_id" validate:"required"`
	UniqueID           string `json:"uid"`
}

func ValidateNetworkInterface(ni *NetworkInterface) (*NetworkInterface, error) {
	err := validator.New().Struct(ni)
	if err != nil {
		ValidatorErrLog(err)
		return nil, err
	}

	return ni, nil
}
