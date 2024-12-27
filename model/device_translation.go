package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// DeviceTranslation Device Translation
type DeviceTranslation struct {
	basic.Base
	basic.DataOwner
	DeviceTranslationAPI
	DeviceID *uuid.UUID `json:"device_id,omitempty" gorm:"type:varchar(36);uniqueIndex:device_translation_unique;not null" swaggertype:"string" format:"uuid"` // Device id
}

// DeviceTranslationAPI Device Translation API
type DeviceTranslationAPI struct {
	LanguageCode *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:device_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	DeviceName   *string `json:"device_name,omitempty" example:"Desktop" gorm:"type:varchar(256)"`                                           // Device Name
}
