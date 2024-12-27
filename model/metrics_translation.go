package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// MetricsTranslation Metrics Translation
type MetricsTranslation struct {
	basic.Base
	basic.DataOwner
	MetricsTranslationAPI
	MetricsID *uuid.UUID `json:"metrics_id,omitempty" gorm:"type:varchar(36);uniqueIndex:metrics_translation_unique;not null" swaggertype:"string" format:"uuid"` // Metrics id
}

// MetricsTranslationAPI Metrics Translation API
type MetricsTranslationAPI struct {
	LanguageCode *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:metrics_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	MetricsName  *string `json:"metrics_name,omitempty" gorm:"type:varchar(256)" example:"Average transaction value"`                         // Metrics Name
	Description  *string `json:"description,omitempty" gorm:"type:text"`                                                                      // Description
}
