package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// OperatingSystemTranslation Operating System Translation
type OperatingSystemTranslation struct {
	basic.Base
	basic.DataOwner
	OperatingSystemTranslationAPI
	OperatingSystemID *uuid.UUID `json:"operating_system_id,omitempty" gorm:"type:varchar(36);uniqueIndex:operating_system_translation_unique;not null" swaggertype:"string" format:"uuid"` // Operating System id
}

// OperatingSystemTranslationAPI Operating System Translation API
type OperatingSystemTranslationAPI struct {
	LanguageCode        *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:operating_system_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	OperatingSystemName *string `json:"operating_system_name,omitempty" gorm:"type:varchar(256)" example:"Microsoft Windows"`                                 // Operating System Name
}
