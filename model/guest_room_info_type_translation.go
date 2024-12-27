package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// GuestRoomInfoTypeTranslation Guest Room Info Type Translation
type GuestRoomInfoTypeTranslation struct {
	basic.Base
	basic.DataOwner
	GuestRoomInfoTypeTranslationAPI
	GuestRoomInfoTypeID *uuid.UUID `json:"guest_room_info_type_id,omitempty" gorm:"type:varchar(36);uniqueIndex:guest_room_info_type_translation_unique;not null" swaggertype:"string" format:"uuid"` // Guest Room Info Type id
}

// GuestRoomInfoTypeTranslationAPI Guest Room Info Type Translation API
type GuestRoomInfoTypeTranslationAPI struct {
	LanguageCode          *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:guest_room_info_type_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	GuestRoomInfoTypeName *string `json:"guest_room_info_type_name,omitempty" gorm:"type:varchar(256)" example:"Accessible rooms"`                                  // Guest Room Info Type Name
	Description           *string `json:"description,omitempty" gorm:"type:text" example:"Accessible rooms"`                                                        // Description
}
