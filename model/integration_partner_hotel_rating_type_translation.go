package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// IntegrationPartnerHotelRatingTypeTranslation Integration Partner Hotel Rating Type Translation
type IntegrationPartnerHotelRatingTypeTranslation struct {
	basic.Base
	basic.DataOwner
	IntegrationPartnerHotelRatingTypeTranslationAPI
	IntegrationPartnerHotelRatingTypeID *uuid.UUID `json:"integration_partner_hotel_rating_type_id,omitempty" gorm:"type:varchar(36);uniqueIndex:integration_partner_hotel_rating_type_translation_unique;not null" swaggertype:"string" format:"uuid"` // Integration Partner Hotel Rating Type id
}

// IntegrationPartnerHotelRatingTypeTranslationAPI Integration Partner Hotel Rating Type Translation API
type IntegrationPartnerHotelRatingTypeTranslationAPI struct {
	LanguageCode        *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:integration_partner_hotel_rating_type_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	HotelRatingTypeName *string `json:"hotel_rating_type_name,omitempty" gorm:"type:varchar(256)"`                                                                                 // Hotel Rating Type Name
	Description         *string `json:"description,omitempty" gorm:"type:text"`                                                                                                    // Description
	Comment             *string `json:"comment,omitempty" gorm:"type:text"`                                                                                                        // Comment
}
