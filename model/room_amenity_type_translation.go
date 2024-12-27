package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// RoomAmenityTypeTranslation Room Amenity Type Translation
type RoomAmenityTypeTranslation struct {
	basic.Base
	basic.DataOwner
	RoomAmenityTypeTranslationAPI
	RoomAmenityTypeID *uuid.UUID `json:"room_amenity_type_id,omitempty" gorm:"type:varchar(36);uniqueIndex:room_amenity_type_translation_unique;not null" swaggertype:"string" format:"uuid"` // Room Amenity Type id
}

// RoomAmenityTypeTranslationAPI Room Amenity Type Translation API
type RoomAmenityTypeTranslationAPI struct {
	LanguageCode        *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:room_amenity_type_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	RoomAmenityTypeName *string `json:"room_amenity_type_name,omitempty" gorm:"type:varchar(256)" example:"Adjoining rooms"`                                   // Room Amenity Type Name
}
