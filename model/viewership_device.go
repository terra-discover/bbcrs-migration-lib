package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// ViewershipDevice Viewership Device
type ViewershipDevice struct {
	basic.Base
	basic.DataOwner
	ViewershipDeviceAPI
	Viewership *Viewership `json:"viewership" gorm:"foreignKey:ViewershipID;references:ID"`
	Device     *Device     `json:"device" gorm:"foreignKey:DeviceID;references:ID"`
}

// ViewershipDeviceAPI Viewership Device Api
type ViewershipDeviceAPI struct {
	ViewershipID      *uuid.UUID `json:"viewership_id,omitempty" gorm:"type:varchar(36);not null;index:,unique,where:deleted_at is null" swaggertype:"string" format:"uuid"`
	DeviceID          *uuid.UUID `json:"device_id,omitempty" gorm:"type:varchar(36);not null;index:,unique,where:deleted_at is null" swaggertype:"string" format:"uuid"`
	BusinessPartnerID *uuid.UUID `json:"business_partner_id" gorm:"type:varchar(36);not null;index:,unique,where:deleted_at is null" swaggertype:"string" format:"uuid"`
}
