package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// ReligionTranslation Religion Translation
type ReligionTranslation struct {
	basic.Base
	basic.DataOwner
	ReligionTranslationAPI
	ReligionID *uuid.UUID `json:"religion_id,omitempty" gorm:"type:varchar(36);uniqueIndex:religion_translation_unique;not null" swaggertype:"string" format:"uuid"` // Religion id
}

// ReligionTranslationAPI Religion Translation API
type ReligionTranslationAPI struct {
	LanguageCode *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:religion_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	ReligionName *string `json:"religion_name,omitempty" gorm:"type:varchar(256)"`                                                             // Religion Name
}
