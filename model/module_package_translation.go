package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// ModulePackageTranslation Module Package Translation
type ModulePackageTranslation struct {
	basic.Base
	basic.DataOwner
	ModulePackageTranslationAPI
	ModulePackageID *uuid.UUID `json:"module_package_id,omitempty" gorm:"type:varchar(36);uniqueIndex:module_package_translation_unique;not null" swaggertype:"string" format:"uuid"` // Module Package id
}

// ModulePackageTranslationAPI Module Package Translation API
type ModulePackageTranslationAPI struct {
	LanguageCode      *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:module_package_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	ModulePackageName *string `json:"module_package_name,omitempty" gorm:"type:varchar(256)" example:"Help Center"`                                       // Module package name
	Description       *string `json:"description,omitempty" gorm:"type:text" example:"Help Center"`                                                       // Description
}
