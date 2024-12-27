package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// RateTimeUnitTranslation Rate Time Unit Translation
type RateTimeUnitTranslation struct {
	basic.Base
	basic.DataOwner
	RateTimeUnitTranslationAPI
	RateTimeUnitID *uuid.UUID `json:"rate_time_unit_id,omitempty" gorm:"type:varchar(36);uniqueIndex:rate_time_unit_translation_unique;not null" swaggertype:"string" format:"uuid"` // Rate Time Unit id
}

// RateTimeUnitTranslationAPI Rate Time Unit Translation API
type RateTimeUnitTranslationAPI struct {
	LanguageCode     *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:rate_time_unit_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	RateTimeUnitName *string `json:"rate_time_unit_name,omitempty" gorm:"type:varchar(256)" example:"Year"`                                              // Rate Time Unit Name
}
