package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// PreferenceTypeTranslation Preference Type Translation
type PreferenceTypeTranslation struct {
	basic.Base
	basic.DataOwner
	PreferenceTypeTranslationAPI
	PreferenceTypeID *uuid.UUID `json:"preference_type_id,omitempty" gorm:"type:varchar(36);uniqueIndex:preference_type_translation_unique;not null" swaggertype:"string" format:"uuid"` // Preference Type id
}

// PreferenceTypeTranslationAPI Preference Type Translation API
type PreferenceTypeTranslationAPI struct {
	LanguageCode       *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:preference_type_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	PreferenceTypeName *string `json:"preference_type_name,omitempty" gorm:"type:varchar(256)"`                                                             // Preference Type Name
	Description        *string `json:"description,omitempty" gorm:"type:text"`                                                                              // Description
}
