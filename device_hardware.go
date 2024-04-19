// https://schema.ocsf.io/objects/geodhaiton
package gcs

import (
	"github.com/go-playground/validator/v10"
)

type DeviceHardware struct {
	Caption          string       `json:"caption" validate:"omitempty"`
	Name             string       `json:"name" validate:"omitempty"`
	BIOSDate         string       `json:"bios_date" validate:"omitempty"`
	BIOSManufacturer string       `json:"bios_manufacturer" validate:"omitempty"`
	BIOSVersion      string       `json:"bios_ver" validate:"omitempty"`
	CPUBits          int          `json:"cpu_bits" validate:"omitempty"`
	CPUCores         int          `json:"cpu_cores" validate:"omitempty"`
	CPUCount         int          `json:"cpu_count" validate:"omitempty"`
	Chassis          string       `json:"chassis" validate:"omitempty"`
	DesktopDisplay   Display      `json:"desktop_display" validate:"omitempty"`
	KeyboardInfo     KeyboardInfo `json:"keyboard_info" validate:"omitempty"`
	ProcessorSpeed   int          `json:"cpu_speed" validate:"omitempty"`
	ProcessorType    string       `json:"cpu_type" validate:"omitempty"`
	RAMSize          int          `json:"ram_size" validate:"omitempty"`
	SerialNumber     string       `json:"serial_number" validate:"omitempty"`
}

func ValidateDeviceHardware(dh *DeviceHardware) (*DeviceHardware, error) {
	err := validator.New().Struct(dh)
	if err != nil {
		ValidatorErrLog(err)
		return nil, err
	}

	return dh, nil
}
