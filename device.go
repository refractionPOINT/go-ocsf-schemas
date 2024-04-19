// https://schema.ocsf.io/objects/device
package gcs

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type Device struct {
	AlternateID          string             `json:"uid_alt,omitempty"`
	AutoscaleUID         string             `json:"autoscale_uid,omitempty"`
	CompliantDevice      bool               `json:"is_compliant,omitempty"`
	CreatedTime          time.Time          `json:"created_time,omitempty"`
	Description          string             `json:"desc,omitempty"`
	Domain               string             `json:"domain,omitempty"`
	FirstSeen            time.Time          `json:"first_seen_time,omitempty"`
	GeoLocation          GeoLocation        `json:"location,omitempty"`
	Groups               []string           `json:"groups,omitempty"`
	HardwareInfo         DeviceHardware     `json:"hw_info,omitempty"`
	Hostname             string             `json:"hostname,omitempty"`
	Hypervisor           string             `json:"hypervisor,omitempty"`
	IMEI                 string             `json:"imei,omitempty"`
	IPAddress            string             `json:"ip,omitempty"`
	Image                Image              `json:"image,omitempty"`
	InstanceID           string             `json:"instance_uid,omitempty"`
	LastSeen             time.Time          `json:"last_seen_time,omitempty"`
	MACAddress           string             `json:"mac,omitempty"`
	ManagedDevice        bool               `json:"is_managed,omitempty"`
	ModifiedTime         time.Time          `json:"modified_time,omitempty"`
	Name                 string             `json:"name,omitempty"`
	NetworkInterfaceID   string             `json:"interface_uid,omitempty"`
	NetworkInterfaceName string             `json:"interface_name,omitempty"`
	NetworkInterfaces    []NetworkInterface `json:"network_interfaces,omitempty"`
	NetworkZone          string             `json:"zone,omitempty"`
	OS                   string             `json:"os,omitempty"`
	Organization         string             `json:"org,omitempty"`
	PersonalDevice       bool               `json:"is_personal,omitempty"`
	Region               string             `json:"region,omitempty"`
	RiskLevel            string             `json:"risk_level,omitempty"`
	RiskLevelID          int                `json:"risk_level_id,omitempty"`
	RiskScore            int                `json:"risk_score,omitempty"`
	Subnet               string             `json:"subnet,omitempty"`
	SubnetUID            string             `json:"subnet_uid,omitempty"`
	TrustedDevice        bool               `json:"is_trusted,omitempty"`
	Type                 string             `json:"type,omitempty"`
	TypeID               int                `json:"type_id,omitempty"`
	UniqueID             string             `json:"uid,omitempty"`
	VLAN                 string             `json:"vlan_uid,omitempty"`
	VPCUID               string             `json:"vpc_uid,omitempty"`
}

func ValidateDevice(device *Device) (*Device, error) {
	err := validator.New().Struct(device)
	if err != nil {
		ValidatorErrLog(err)
		return nil, err
	}

	return device, nil
}
