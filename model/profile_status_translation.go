package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// ProfileStatusTranslation Profile Status Translation
type ProfileStatusTranslation struct {
	basic.Base
	basic.DataOwner
	ProfileStatusTranslationAPI
	ProfileStatusID *uuid.UUID `json:"profile_status_id,omitempty" gorm:"type:varchar(36);uniqueIndex:profile_status_translation_unique;not null" swaggertype:"string" format:"uuid"` // Profile Status id
}

// ProfileStatusTranslationAPI Profile Status Translation API
type ProfileStatusTranslationAPI struct {
	LanguageCode      *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:profile_status_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	ProfileStatusName *string `json:"profile_status_name,omitempty" gorm:"type:varchar(256)" example:"active"`                                            // Profile Status Name
}
