package model

import basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"

// Device Device
type Device struct {
	basic.Base
	basic.DataOwner
	DeviceAPI
	DeviceTranslation *DeviceTranslation `json:"device_translation,omitempty"` // Device Translation
}

// DeviceAPI Device API
type DeviceAPI struct {
	DeviceCode *string `json:"device_code,omitempty" example:"desktop" gorm:"type:varchar(36);index:,unique,where:deleted_at is null;not null"` // Device Code
	DeviceName *string `json:"device_name,omitempty" example:"Desktop" gorm:"type:varchar(256)"`                                                // Device Name
}
