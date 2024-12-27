package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// AdditionalInfoOptionTranslation Additional Info Option Translation
type AdditionalInfoOptionTranslation struct {
	basic.Base
	basic.DataOwner
	AdditionalInfoOptionTranslationAPI
	AdditionalInfoOptionID *uuid.UUID `json:"additional_info_option_id,omitempty" gorm:"type:varchar(36);uniqueIndex:additional_info_option_translation_unique;not null" swaggertype:"string" format:"uuid"` // Additional Info Option id
}

// AdditionalInfoOptionTranslationAPI Additional Info Option Translation API
type AdditionalInfoOptionTranslationAPI struct {
	LanguageCode             *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:additional_info_option_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	AdditionalInfoOptionName *string `json:"additional_info_option_name,omitempty" gorm:"type:varchar(256)"`                                                             // Additional Info Option Name
	Description              *string `json:"description,omitempty" gorm:"type:text"`                                                                                     // Description
}
