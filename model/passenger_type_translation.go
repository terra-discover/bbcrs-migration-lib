package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// PassengerTypeTranslation Passenger Type Translation
type PassengerTypeTranslation struct {
	basic.Base
	basic.DataOwner
	PassengerTypeTranslationAPI
	PassengerTypeID *uuid.UUID `json:"passenger_type_id,omitempty" gorm:"type:varchar(36);uniqueIndex:passenger_type_translation_unique;not null" swaggertype:"string" format:"uuid"` // Passenger Type id
}

// PassengerTypeTranslationAPI Passenger Type Translation API
type PassengerTypeTranslationAPI struct {
	LanguageCode      *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:passenger_type_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	PassengerTypeName *string `json:"passenger_type_name,omitempty" gorm:"type:varchar(256)" example:"ADULT"`                                             // Passenger Type Name
}
