package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// CapabilitySetting Capability Setting
type CapabilitySetting struct {
	basic.Base
	basic.DataOwner
	CapabilitySettingAPI
}

// CapabilitySettingAPI Capability Setting API
type CapabilitySettingAPI struct {
	CapabilityID *uuid.UUID `json:"capability_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"` // Capability ID
	Name         *string    `json:"name,omitempty" gorm:"type:varchar(256);not null"`                                            // Name
	Value        *string    `json:"value,omitempty" gorm:"type:text"`                                                            // Value
}
