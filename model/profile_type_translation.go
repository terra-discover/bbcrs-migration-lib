package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// ProfileTypeTranslation Profile Type Translation
type ProfileTypeTranslation struct {
	basic.Base
	basic.DataOwner
	ProfileTypeTranslationAPI
	ProfileTypeID *uuid.UUID `json:"profile_type_id,omitempty" gorm:"type:varchar(36);uniqueIndex:profile_type_translation_unique;not null" swaggertype:"string" format:"uuid"` // Profile Type id
}

// ProfileTypeTranslationAPI Profile Type Translation API
type ProfileTypeTranslationAPI struct {
	LanguageCode    *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:profile_type_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	ProfileTypeName *string `json:"profile_type_name,omitempty" example:"Contact" gorm:"type:varchar(256)"`                                           // Profile Type Name
}
