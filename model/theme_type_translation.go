package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// ThemeTypeTranslation Theme Type Translation
type ThemeTypeTranslation struct {
	basic.Base
	basic.DataOwner
	ThemeTypeTranslationAPI
	ThemeTypeID *uuid.UUID `json:"theme_type_id,omitempty" gorm:"type:varchar(36);uniqueIndex:theme_type_translation_unique;not null" swaggertype:"string" format:"uuid"` // Theme Type id
}

// ThemeTypeTranslationAPI Theme Type Translation API
type ThemeTypeTranslationAPI struct {
	LanguageCode  *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:theme_type_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	ThemeTypeName *string `json:"theme_type_name,omitempty" gorm:"type:varchar(256)" example:"backend"`                                           // Theme Type Name
	Description   *string `json:"description,omitempty" gorm:"type:text" example:"description"`                                                   // Description
}
