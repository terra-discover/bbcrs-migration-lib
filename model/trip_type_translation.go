package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// TripTypeTranslation Cabin Type Translation
type TripTypeTranslation struct {
	basic.Base
	basic.DataOwner
	TripTypeTranslationAPI
	TripTypeID *uuid.UUID `json:"trip_type_id,omitempty" gorm:"type:varchar(36);uniqueIndex:trip_type_translation_unique;not null" swaggertype:"string" format:"uuid"` // Cabin Type id
}

// TripTypeTranslationAPI Cabin Type Translation API
type TripTypeTranslationAPI struct {
	LanguageCode *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:trip_type_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	TripTypeName *string `json:"trip_type_name,omitempty" gorm:"type:varchar(256)"`                                                             // Cabin Type Name
}
