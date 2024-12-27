package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// ModuleTranslation Module Translation
type ModuleTranslation struct {
	basic.Base
	basic.DataOwner
	ModuleTranslationAPI
	ModuleID *uuid.UUID `json:"module_id,omitempty" gorm:"type:varchar(36);uniqueIndex:module_translation_unique;not null" swaggertype:"string" format:"uuid"`
}

// ModuleTranslationAPI Module Translation API
type ModuleTranslationAPI struct {
	LanguageCode *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:module_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	ModuleName   *string `json:"module_name,omitempty" gorm:"type:varchar(256);"`
	Description  *string `json:"description,omitempty" gorm:"type:text"`
}
