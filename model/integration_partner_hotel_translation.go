package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// IntegrationPartnerHotelTranslation Integration Partner Hotel Translation
type IntegrationPartnerHotelTranslation struct {
	basic.Base
	basic.DataOwner
	IntegrationPartnerHotelTranslationAPI
	IntegrationPartnerHotelID *uuid.UUID `json:"integration_partner_hotel_id,omitempty" gorm:"type:varchar(36);uniqueIndex:integration_partner_hotel_translation_unique;not null" swaggertype:"string" format:"uuid"` // Integration Partner Hotel id
}

// IntegrationPartnerHotelTranslationAPI Integration Partner Hotel Translation API
type IntegrationPartnerHotelTranslationAPI struct {
	LanguageCode *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:integration_partner_hotel_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	HotelDetails *string `json:"hotel_details,omitempty" gorm:"type:text"`                                                                                      // Hotel Details
	Comment      *string `json:"comment,omitempty" gorm:"type:text"`                                                                                            // Comment
}
