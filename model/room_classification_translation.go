package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// RoomClassificationTranslation Room Classification Translation
type RoomClassificationTranslation struct {
	basic.Base
	basic.DataOwner
	RoomClassificationTranslationAPI
	RoomClassificationID *uuid.UUID `json:"room_classification_id,omitempty" gorm:"type:varchar(36);uniqueIndex:room_classification_translation_unique;not null" swaggertype:"string" format:"uuid"` // Room Classification id
}

// RoomClassificationTranslationAPI Room Classification Translation API
type RoomClassificationTranslationAPI struct {
	LanguageCode           *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:room_classification_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	RoomClassificationName *string `json:"room_classification_name,omitempty" gorm:"type:varchar(256)" example:"Accessible rooms"`                                  // Room Classification Name
}
