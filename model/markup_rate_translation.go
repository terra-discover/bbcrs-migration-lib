package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// MarkupRateTranslation Task Type Translation
type MarkupRateTranslation struct {
	basic.Base
	basic.DataOwner
	MarkupRateTranslationAPI
	MarkupRateID *uuid.UUID `json:"markup_rate_id,omitempty" gorm:"type:varchar(36);uniqueIndex:markup_rate_translation_unique;not null" swaggertype:"string" format:"uuid"` // Markup Rate id
}

// MarkupRateTranslationAPI Task Type Translation API
type MarkupRateTranslationAPI struct {
	LanguageCode *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:markup_rate_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	Description  *string `json:"description,omitempty" gorm:"type:text"`
}
