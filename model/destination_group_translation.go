package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// DestinationGroupTranslation Destination Group Translation
type DestinationGroupTranslation struct {
	basic.Base
	basic.DataOwner
	DestinationGroupTranslationAPI
	DestinationGroupID *uuid.UUID `json:"destination_group_id,omitempty" gorm:"type:varchar(36);uniqueIndex:destination_group_translation_unique;not null" swaggertype:"string" format:"uuid"` // Destination Group id
}

// DestinationGroupTranslationAPI Destination Group Translation API
type DestinationGroupTranslationAPI struct {
	LanguageCode         *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:destination_group_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	DestinationGroupName *string `json:"destination_group_name,omitempty" gorm:"type:varchar(256)" example:"pulau lombok"`                                      // Destination Group Name
}
