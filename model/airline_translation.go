package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// AirlineTranslation Airline Translation
type AirlineTranslation struct {
	basic.Base
	basic.DataOwner
	AirlineTranslationAPI
	AirlineID *uuid.UUID `json:"airline_id,omitempty" gorm:"type:varchar(36);uniqueIndex:airline_translation_unique;not null" swaggertype:"string" format:"uuid"` // Airline id
}

// AirlineTranslationAPI Airline Translation API
type AirlineTranslationAPI struct {
	LanguageCode *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:airline_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	AirlineName  *string `json:"airline_name,omitempty" gorm:"type:varchar(64)" example:"Garuda Indonesia"`                                   // Airline Name
}
