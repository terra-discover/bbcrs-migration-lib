package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// RoomAmenityCategoryTranslation Amenity Category Translation
type RoomAmenityCategoryTranslation struct {
	basic.Base
	basic.DataOwner
	RoomAmenityCategoryTranslationAPI
	RoomAmenityCategoryID *uuid.UUID `json:"amenity_category_id,omitempty" gorm:"type:varchar(36);uniqueIndex:amenity_category_translation_unique;not null" swaggertype:"string" format:"uuid"` // Amenity Category id
}

// RoomAmenityCategoryTranslationAPI Amenity Category Translation API
type RoomAmenityCategoryTranslationAPI struct {
	LanguageCode            *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:amenity_category_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	RoomAmenityCategoryName *string `json:"room_amenity_category_name,omitempty" gorm:"type:varchar(256)" example:"Bedroom"`                                      // Room Amenity Category Name
	Description             *string `json:"description,omitempty" gorm:"type:text" example:"description"`                                                         // Description
}
