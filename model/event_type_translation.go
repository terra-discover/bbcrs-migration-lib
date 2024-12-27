package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// EventTypeTranslation Event Type Translation
type EventTypeTranslation struct {
	basic.Base
	basic.DataOwner
	EventTypeTranslationAPI
	EventTypeID *uuid.UUID `json:"event_type_id,omitempty" gorm:"type:varchar(36);uniqueIndex:event_type_translation_unique;not null" swaggertype:"string" format:"uuid"` // Event Type id
}

// EventTypeTranslationAPI Event Type Translation API
type EventTypeTranslationAPI struct {
	LanguageCode  *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:event_type_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	EventTypeName *string `json:"event_type_name,omitempty" gorm:"type:varchar(256)" example:"Error"`                                             // Event Type Name
}
