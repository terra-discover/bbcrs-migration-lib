package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// PreferenceOptionTranslation Preference Option Translation
type PreferenceOptionTranslation struct {
	basic.Base
	basic.DataOwner
	PreferenceOptionTranslationAPI
	PreferenceOptionID *uuid.UUID `json:"preference_option_id,omitempty" gorm:"type:varchar(36);uniqueIndex:preference_option_translation_unique;not null" swaggertype:"string" format:"uuid"` // Preference Option id
}

// PreferenceOptionTranslationAPI Preference Option Translation API
type PreferenceOptionTranslationAPI struct {
	LanguageCode         *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:preference_option_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	PreferenceOptionName *string `json:"preference_option_name,omitempty" gorm:"type:varchar(256)"`                                                             // Preference Option Name
	Description          *string `json:"description,omitempty" gorm:"type:text"`                                                                                // Description
}
