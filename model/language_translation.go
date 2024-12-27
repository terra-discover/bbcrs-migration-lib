package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// LanguageTranslation Language Translation
type LanguageTranslation struct {
	basic.Base
	basic.DataOwner
	LanguageTranslationAPI
	LanguageID *uuid.UUID `json:"language_id,omitempty" gorm:"type:varchar(36);uniqueIndex:language_translation_unique;not null" swaggertype:"string" format:"uuid"` // Language id
}

// LanguageTranslationAPI Language Translation API
type LanguageTranslationAPI struct {
	LanguageCode       *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:language_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	LanguageName       *string `json:"language_name,omitempty" gorm:"type:varchar(256)" example:"English"`                                           // International Language Name
	LanguageNativeName *string `json:"language_native_name,omitempty" gorm:"type:varchar(256)" example:"English"`                                    // Language Foreign Name
}
