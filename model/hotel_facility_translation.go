package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// HotelFacilityTranslation Hotel Facility Translation
type HotelFacilityTranslation struct {
	basic.Base
	basic.DataOwner
	HotelFacilityTranslationAPI
	HotelFacilityID *uuid.UUID `json:"hotel_facility_id,omitempty" gorm:"type:varchar(36);uniqueIndex:hotel_facility_translation_unique;not null" swaggertype:"string" format:"uuid"` // Hotel Facility id
}

// HotelFacilityTranslationAPI Hotel Facility Translation API
type HotelFacilityTranslationAPI struct {
	LanguageCode      *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:hotel_facility_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	HotelFacilityName *string `json:"hotel_facility_name,omitempty" gorm:"type:varchar(256)"`                                                             // Hotel Facility Name
	Description       *string `json:"description,omitempty" gorm:"type:text"`                                                                             // Description
}
