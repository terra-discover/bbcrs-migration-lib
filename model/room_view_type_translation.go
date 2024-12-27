package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// RoomViewTypeTranslation Room View Type Translation
type RoomViewTypeTranslation struct {
	basic.Base
	basic.DataOwner
	RoomViewTypeTranslationAPI
	RoomViewTypeID *uuid.UUID `json:"room_view_type_id,omitempty" gorm:"type:varchar(36);uniqueIndex:room_view_type_translation_unique;not null" swaggertype:"string" format:"uuid"` // Room View Type id
}

// RoomViewTypeTranslationAPI Room View Type Translation API
type RoomViewTypeTranslationAPI struct {
	LanguageCode     *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:room_view_type_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	RoomViewTypeName *string `json:"room_view_type_name,omitempty" gorm:"type:varchar(256)" example:"Airport view"`                                      // Room View Type Name
}
