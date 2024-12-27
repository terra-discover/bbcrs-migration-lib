package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// ThemeLocationTranslation Theme Location Translation
type ThemeLocationTranslation struct {
	basic.Base
	basic.DataOwner
	ThemeLocationTranslationAPI
	ThemeLocationID *uuid.UUID `json:"theme_location_id,omitempty" gorm:"type:varchar(36);uniqueIndex:theme_location_translation_unique;not null" swaggertype:"string" format:"uuid"` // Theme Location id
}

// ThemeLocationTranslationAPI Theme Location Translation API
type ThemeLocationTranslationAPI struct {
	LanguageCode      *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:theme_location_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	ThemeLocationName *string `json:"theme_location_name,omitempty" gorm:"type:varchar(256)"`                                                             // Theme Location Name
	Description       *string `json:"description,omitempty" gorm:"type:text"`                                                                             // Description
}
