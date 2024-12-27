package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// IndustryTranslation Industry Translation
type IndustryTranslation struct {
	basic.Base
	basic.DataOwner
	IndustryTranslationAPI
	IndustryID *uuid.UUID `json:"industry_id,omitempty" gorm:"type:varchar(36);uniqueIndex:industry_translation_unique;not null" swaggertype:"string" format:"uuid"` // Industry id
}

// IndustryTranslationAPI Industry Translation API
type IndustryTranslationAPI struct {
	LanguageCode *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:industry_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	IndustryName *string `json:"industry_name,omitempty" gorm:"type:varchar(256)"`                                                             // Industry Name
}
