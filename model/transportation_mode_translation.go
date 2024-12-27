package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// TransportationModeTranslation Transportation Mode Translation
type TransportationModeTranslation struct {
	basic.Base
	basic.DataOwner
	TransportationModeTranslationAPI
	TransportationModeID *uuid.UUID `json:"transportation_mode_id,omitempty" gorm:"type:varchar(36);uniqueIndex:transportation_mode_translation_unique;not null" swaggertype:"string" format:"uuid"` // Transportation Mode id
}

// TransportationModeTranslationAPI Transportation Mode Translation API
type TransportationModeTranslationAPI struct {
	LanguageCode           *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:transportation_mode_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	TransportationModeName *string `json:"transportation_mode_name,omitempty" gorm:"type:varchar(256)" example:"Bicycle"`                                           // Transportation Mode Name
}
