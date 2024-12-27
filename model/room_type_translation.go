package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// RoomTypeTranslation Task Type Translation
type RoomTypeTranslation struct {
	basic.Base
	basic.DataOwner
	RoomTypeTranslationAPI
	RoomTypeID *uuid.UUID `json:"room_type_id,omitempty" gorm:"type:varchar(36);uniqueIndex:room_type_translation_unique;not null" swaggertype:"string" format:"uuid"` // Markup Rate id
}

// RoomTypeTranslationAPI Task Type Translation API
type RoomTypeTranslationAPI struct {
	LanguageCode *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:room_type_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	RoomTypeCode *string `json:"room_type_code,omitempty" gorm:"type:varchar(36);not null;"`
	RoomTypeName *string `json:"room_type_name,omitempty" gorm:"type:varchar(256);"`
	Description  *string `json:"description,omitempty" gorm:"type:text"`
}
