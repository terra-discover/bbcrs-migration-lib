package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// IntegrationPartnerHotelAmenityTypeTranslation Integration Partner Hotel Amenity Type Translation
type IntegrationPartnerHotelAmenityTypeTranslation struct {
	basic.Base
	basic.DataOwner
	IntegrationPartnerHotelAmenityTypeTranslationAPI
	IntegrationPartnerHotelAmenityTypeID *uuid.UUID `json:"integration_partner_hotel_amenity_type_id,omitempty" gorm:"type:varchar(36);uniqueIndex:integration_partner_hotel_amenity_type_translation_unique;not null" swaggertype:"string" format:"uuid"` // Integration Partner Hotel Amenity Type id
}

// IntegrationPartnerHotelAmenityTypeTranslationAPI Integration Partner Hotel Amenity Type Translation API
type IntegrationPartnerHotelAmenityTypeTranslationAPI struct {
	LanguageCode         *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:integration_partner_hotel_amenity_type_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	HotelAmenityTypeName *string `json:"hotel_amenity_type_name,omitempty" gorm:"type:varchar(256)"`                                                                                 // Hotel Amenity Type Name
	Description          *string `json:"description,omitempty" gorm:"type:text"`                                                                                                     // Description
	Comment              *string `json:"comment,omitempty" gorm:"type:text"`                                                                                                         // Comment
}
