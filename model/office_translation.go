package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// OfficeTranslation Office Translation
type OfficeTranslation struct {
	basic.Base
	basic.DataOwner
	OfficeTranslationAPI
	OfficeID *uuid.UUID `json:"office_id,omitempty" gorm:"type:varchar(36);uniqueIndex:office_translation_unique;not null" swaggertype:"string" format:"uuid"` // Office id
}

// OfficeTranslationAPI Office Translation API
type OfficeTranslationAPI struct {
	LanguageCode *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:office_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	OfficeName   *string `json:"office_name,omitempty" gorm:"type:varchar(256)"`                                                             // Office Name
	Description  *string `json:"description,omitempty" gorm:"type:text"`                                                                     // Description
}
