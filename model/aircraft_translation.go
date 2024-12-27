package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// AircraftTranslation Aircraft Translation
type AircraftTranslation struct {
	basic.Base
	basic.DataOwner
	AircraftTranslationAPI
	AircraftID *uuid.UUID `json:"aircraft_id,omitempty" gorm:"type:varchar(36);uniqueIndex:aircraft_translation_unique;not null" swaggertype:"string" format:"uuid"` // Aircraft id
}

// AircraftTranslationAPI Aircraft Translation API
type AircraftTranslationAPI struct {
	LanguageCode *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:aircraft_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	AircraftName *string `json:"aircraft_name,omitempty" gorm:"type:varchar(64)"`                                                              // Aircraft Name
	Model        *string `json:"model,omitempty" gorm:"type:varchar(64)"`                                                                      // Model
}
