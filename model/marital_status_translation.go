package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// MaritalStatusTranslation Marital Status Translation
type MaritalStatusTranslation struct {
	basic.Base
	basic.DataOwner
	MaritalStatusTranslationAPI
	MaritalStatusID *uuid.UUID `json:"marital_status_id,omitempty" gorm:"type:varchar(36);uniqueIndex:marital_status_translation_unique;not null" swaggertype:"string" format:"uuid"` // Marital Status id
}

// MaritalStatusTranslationAPI Marital Status Translation API
type MaritalStatusTranslationAPI struct {
	LanguageCode      *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:marital_status_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	MaritalStatusName *string `json:"marital_status_name,omitempty" gorm:"type:varchar(256)" example:"married"`                                           // Marital Status Name
}
