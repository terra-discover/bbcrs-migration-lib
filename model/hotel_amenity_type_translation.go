package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// HotelAmenityTypeTranslation Hotel Amenity Type Translation
type HotelAmenityTypeTranslation struct {
	basic.Base
	basic.DataOwner
	HotelAmenityTypeTranslationAPI
	HotelAmenityTypeID *uuid.UUID `json:"hotel_amenity_type_id,omitempty" gorm:"type:varchar(36);uniqueIndex:hotel_amenity_type_translation_unique;not null" swaggertype:"string" format:"uuid"` // Hotel Amenity Type id
}

// HotelAmenityTypeTranslationAPI Hotel Amenity Type Translation API
type HotelAmenityTypeTranslationAPI struct {
	LanguageCode         *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:hotel_amenity_type_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	HotelAmenityTypeName *string `json:"hotel_amenity_type_name,omitempty" gorm:"type:varchar(256)"`                                                             // Hotel Amenity Type Name
}
