package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// FlightTypeTranslation Flight Type Translation
type FlightTypeTranslation struct {
	basic.Base
	basic.DataOwner
	FlightTypeTranslationAPI
	FlightTypeID *uuid.UUID `json:"flight_type_id,omitempty" gorm:"type:varchar(36);uniqueIndex:flight_type_translation_unique;not null" swaggertype:"string" format:"uuid"` // Flight Type id
}

// FlightTypeTranslationAPI Flight Type Translation API
type FlightTypeTranslationAPI struct {
	LanguageCode   *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:flight_type_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	FlightTypeName *string `json:"flight_type_name,omitempty" gorm:"type:varchar(256)" example:"Forces"`                                            // Flight Type Name
}
