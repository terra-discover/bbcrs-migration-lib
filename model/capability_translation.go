package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// CapabilityTranslation Capability Translation
type CapabilityTranslation struct {
	basic.Base
	basic.DataOwner
	CapabilityTranslationAPI
	CapabilityID *uuid.UUID `json:"capability_id,omitempty" gorm:"type:varchar(36);uniqueIndex:capability_translation_unique;not null" swaggertype:"string" format:"uuid"` // Capability id
}

// CapabilityTranslationAPI Capability Translation API
type CapabilityTranslationAPI struct {
	LanguageCode   *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:capability_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	CapabilityName *string `json:"capability_name,omitempty" gorm:"type:varchar(256)"`                                                             // Capability Name
	Description    *string `json:"description,omitempty" gorm:"type:text"`                                                                         // Description
}
