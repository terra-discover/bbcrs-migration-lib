package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// OffsetDropTimeTranslation Offset Drop Time Translation
type OffsetDropTimeTranslation struct {
	basic.Base
	basic.DataOwner
	OffsetDropTimeTranslationAPI
	OffsetDropTimeID *uuid.UUID `json:"offset_drop_time_id,omitempty" gorm:"type:varchar(36);uniqueIndex:offset_drop_time_translation_unique;not null" swaggertype:"string" format:"uuid"` // Offset Drop Time id
}

// OffsetDropTimeTranslationAPI Offset Drop Time Translation API
type OffsetDropTimeTranslationAPI struct {
	LanguageCode       *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:offset_drop_time_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	OffsetDropTimeName *string `json:"offset_drop_time_name,omitempty" gorm:"type:varchar(256)" example:"Before Arrival"`                                    // Offset Drop Time Name
}
