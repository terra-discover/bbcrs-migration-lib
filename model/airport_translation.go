package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// AirportTranslation Airport Translation
type AirportTranslation struct {
	basic.Base
	basic.DataOwner
	AirportTranslationAPI
	AirportID *uuid.UUID `json:"airport_id,omitempty" gorm:"type:varchar(36);uniqueIndex:airport_translation_unique;not null" swaggertype:"string" format:"uuid"` // Airport id
}

// AirportTranslationAPI Airport Translation API
type AirportTranslationAPI struct {
	LanguageCode *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:airport_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	AirportName  *string `json:"airport_name,omitempty" gorm:"type:varchar(256)" example:"Bandara Soekarno Hatta"`                            // Airport Name
}
