package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// RoomLocationTypeTranslation Room Location Type Translation
type RoomLocationTypeTranslation struct {
	basic.Base
	basic.DataOwner
	RoomLocationTypeTranslationAPI
	RoomLocationTypeID *uuid.UUID `json:"room_location_type_id,omitempty" gorm:"type:varchar(36);uniqueIndex:room_location_type_translation_unique;not null" swaggertype:"string" format:"uuid"` // Room Location Type id
}

// RoomLocationTypeTranslationAPI Room Location Type Translation API
type RoomLocationTypeTranslationAPI struct {
	LanguageCode         *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:room_location_type_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	RoomLocationTypeName *string `json:"room_location_type_name,omitempty" gorm:"type:varchar(256)" example:"Away from the elevator"`                            // Room Location Type Name
}
