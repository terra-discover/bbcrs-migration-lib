package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// DestinationSystemTypeTranslation Destination System Type Translation
type DestinationSystemTypeTranslation struct {
	basic.Base
	basic.DataOwner
	DestinationSystemTypeTranslationAPI
	DestinationSystemTypeID *uuid.UUID `json:"destination_system_type_id,omitempty" gorm:"type:varchar(36);uniqueIndex:destination_system_type_translation_unique;not null" swaggertype:"string" format:"uuid"` // Destination System Type id
}

// DestinationSystemTypeTranslationAPI Destination System Type Translation API
type DestinationSystemTypeTranslationAPI struct {
	LanguageCode              *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:destination_system_type_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	DestinationSystemTypeName *string `json:"destination_system_type_name,omitempty" gorm:"type:varchar(256)" example:"Sample"`                                            // Destination System Type Name
}
