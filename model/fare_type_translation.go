package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// FareTypeTranslation Fare Type Translation
type FareTypeTranslation struct {
	basic.Base
	basic.DataOwner
	FareTypeTranslationAPI
	FareTypeID *uuid.UUID `json:"fare_type_id,omitempty" gorm:"type:varchar(36);uniqueIndex:fare_type_translation_unique;not null" swaggertype:"string" format:"uuid"` // Fare Type id
}

// FareTypeTranslationAPI Fare Type Translation API
type FareTypeTranslationAPI struct {
	LanguageCode *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:fare_type_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	FareTypeName *string `json:"fare_type_name,omitempty" gorm:"type:varchar(256)" example:"Lite"`                                              // Fare Type Name
}
