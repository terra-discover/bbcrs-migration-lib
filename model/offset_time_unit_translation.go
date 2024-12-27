package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// OffsetTimeUnitTranslation Offset Time Unit Translation
type OffsetTimeUnitTranslation struct {
	basic.Base
	basic.DataOwner
	OffsetTimeUnitTranslationAPI
	OffsetTimeUnitID *uuid.UUID `json:"offset_time_unit_id,omitempty" gorm:"type:varchar(36);uniqueIndex:offset_time_unit_translation_unique;not null" swaggertype:"string" format:"uuid"` // Offset Time Unit id
}

// OffsetTimeUnitTranslationAPI Offset Time Unit Translation API
type OffsetTimeUnitTranslationAPI struct {
	LanguageCode       *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:offset_time_unit_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	OffsetTimeUnitName *string `json:"offset_time_unit_name,omitempty" gorm:"type:varchar(256)" example:"Year"`                                              // Offset Time Unit Name
}
