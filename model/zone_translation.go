package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// ZoneTranslation Zone Translation
type ZoneTranslation struct {
	basic.Base
	basic.DataOwner
	ZoneTranslationAPI
	ZoneID *uuid.UUID `json:"zone_id,omitempty" gorm:"type:varchar(36);uniqueIndex:zone_translation_unique;not null" swaggertype:"string" format:"uuid"` // Zone id
}

// ZoneTranslationAPI Zone Translation API
type ZoneTranslationAPI struct {
	LanguageCode *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:zone_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	ZoneName     *string `json:"zone_name,omitempty" gorm:"type:varchar(256)" example:"Jakarta"`                                           // Zone Name
	Description  *string `json:"description,omitempty" gorm:"type:text" example:"Jakarta adalah Ibukota Negara Indonesia"`                 // Description
}
