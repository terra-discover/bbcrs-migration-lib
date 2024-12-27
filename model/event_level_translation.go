package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// EventLevelTranslation Event Level Translation
type EventLevelTranslation struct {
	basic.Base
	basic.DataOwner
	EventLevelTranslationAPI
	EventLevelID *uuid.UUID `json:"event_level_id,omitempty" gorm:"type:varchar(36);uniqueIndex:event_level_translation_unique;not null" swaggertype:"string" format:"uuid"` // Event Level id
}

// EventLevelTranslationAPI Event Level Translation API
type EventLevelTranslationAPI struct {
	LanguageCode   *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:event_level_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	EventLevelName *string `json:"event_level_name,omitempty" gorm:"type:varchar(256)" example:"Error"`                                             // Event Level Name
}
