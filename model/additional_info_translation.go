package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// AdditionalInfoTranslation Additional Info Translation
type AdditionalInfoTranslation struct {
	basic.Base
	basic.DataOwner
	AdditionalInfoTranslationAPI
	AdditionalInfoID *uuid.UUID `json:"additional_info_id,omitempty" gorm:"type:varchar(36);uniqueIndex:additional_info_translation_unique;not null" swaggertype:"string" format:"uuid"` // Additional Info id
}

// AdditionalInfoTranslationAPI Additional Info Translation API
type AdditionalInfoTranslationAPI struct {
	LanguageCode       *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:additional_info_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	AdditionalInfoName *string `json:"additional_info_name,omitempty" gorm:"type:varchar(256)"`                                                             // Additional Info Name
	Description        *string `json:"description,omitempty" gorm:"type:text"`                                                                              // Description
}
