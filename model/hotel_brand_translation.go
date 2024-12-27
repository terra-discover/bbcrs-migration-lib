package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// HotelBrandTranslation Hotel Brand Translation
type HotelBrandTranslation struct {
	basic.Base
	basic.DataOwner
	HotelBrandTranslationAPI
	HotelBrandID *uuid.UUID `json:"hotel_brand_id,omitempty" gorm:"type:varchar(36);uniqueIndex:hotel_brand_translation_unique;not null" swaggertype:"string" format:"uuid"` // Hotel Brand id
}

// HotelBrandTranslationAPI Hotel Brand Translation API
type HotelBrandTranslationAPI struct {
	LanguageCode   *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:hotel_brand_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	HotelBrandName *string `json:"hotel_brand_name,omitempty" gorm:"type:varchar(256)" example:"Hilton Hotel and Resorts"`                          // Hotel Brand Name
}
