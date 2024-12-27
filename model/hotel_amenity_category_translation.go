package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// HotelAmenityCategoryTranslation Hotel Amenity Category Translation
type HotelAmenityCategoryTranslation struct {
	basic.Base
	basic.DataOwner
	HotelAmenityCategoryTranslationAPI
	HotelAmenityCategoryID *uuid.UUID `json:"hotel_amenity_category_id,omitempty" gorm:"type:varchar(36);uniqueIndex:hotel_amenity_category_translation_unique;not null" swaggertype:"string" format:"uuid"` // Hotel Amenity Category id
}

// HotelAmenityCategoryTranslationAPI Hotel Amenity Category Translation API
type HotelAmenityCategoryTranslationAPI struct {
	LanguageCode             *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:hotel_amenity_category_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	HotelAmenityCategoryName *string `json:"hotel_amenity_category_name,omitempty" gorm:"type:varchar(256)"`                                                             // Hotel Amenity Category Name
	Description              *string `json:"description,omitempty" gorm:"type:text"`                                                                                     // Description
}
