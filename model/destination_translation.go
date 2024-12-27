package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// DestinationTranslation Destination Translation
type DestinationTranslation struct {
	basic.Base
	basic.DataOwner
	DestinationTranslationAPI
	DestinationID *uuid.UUID `json:"destination_id,omitempty" gorm:"type:varchar(36);uniqueIndex:destination_translation_unique;not null" swaggertype:"string" format:"uuid"` // Destination id
}

// DestinationTranslationAPI Destination Translation API
type DestinationTranslationAPI struct {
	LanguageCode    *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:destination_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	DestinationName *string `json:"destination_name,omitempty" gorm:"type:varchar(256)"`                                                             // Destination Name
	Description     *string `json:"description,omitempty" gorm:"type:text" example:"deskripsi"`
}
