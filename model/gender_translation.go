package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// GenderTranslation Gender Translation
type GenderTranslation struct {
	basic.Base
	basic.DataOwner
	GenderTranslationAPI
	GenderID *uuid.UUID `json:"gender_id,omitempty" gorm:"type:varchar(36);uniqueIndex:gender_translation_unique;not null" swaggertype:"string" format:"uuid"` // Gender id
}

// GenderTranslationAPI Gender Translation API
type GenderTranslationAPI struct {
	LanguageCode *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:gender_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	GenderName   *string `json:"gender_name,omitempty" gorm:"type:varchar(256)" example:"Laki-Laki"`                                         // Gender Name
}
