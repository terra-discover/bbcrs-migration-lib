package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// PreferenceTranslation Preference Translation
type PreferenceTranslation struct {
	basic.Base
	basic.DataOwner
	PreferenceTranslationAPI
	PreferenceID *uuid.UUID `json:"preference_id,omitempty" gorm:"type:varchar(36);uniqueIndex:preference_translation_unique;not null" swaggertype:"string" format:"uuid"` // Preference id
}

// PreferenceTranslationAPI Preference Translation API
type PreferenceTranslationAPI struct {
	LanguageCode   *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:preference_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	PreferenceName *string `json:"preference_name,omitempty" gorm:"type:varchar(256)" example:"queen bed"`                                         // Preference Name
	Description    *string `json:"description,omitempty" gorm:"type:text" example:"deskripsi"`                                                     // Description
}
