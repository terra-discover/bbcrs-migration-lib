package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// ThemeTranslation Theme Translation
type ThemeTranslation struct {
	basic.Base
	basic.DataOwner
	ThemeTranslationAPI
	ThemeID *uuid.UUID `json:"theme_id,omitempty" gorm:"type:varchar(36);uniqueIndex:theme_translation_unique;not null" swaggertype:"string" format:"uuid"` // Theme id
}

// ThemeTranslationAPI Theme Translation API
type ThemeTranslationAPI struct {
	LanguageCode *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:theme_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	ThemeName    *string `json:"theme_name,omitempty" gorm:"type:varchar(256)"`                                                             // Theme Name
	Description  *string `json:"description,omitempty" gorm:"type:text" example:"description"`                                              // Description
}
