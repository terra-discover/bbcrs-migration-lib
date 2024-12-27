package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// IntegrationPartnerHotelGroupTranslation Integration Partner Hotel Group Translation
type IntegrationPartnerHotelGroupTranslation struct {
	basic.Base
	basic.DataOwner
	IntegrationPartnerHotelGroupTranslationAPI
	IntegrationPartnerHotelGroupID *uuid.UUID `json:"integration_partner_hotel_group_id,omitempty" gorm:"type:varchar(36);uniqueIndex:integration_partner_hotel_group_translation_unique;not null" swaggertype:"string" format:"uuid"` // Integration Partner Hotel Group id
}

// IntegrationPartnerHotelGroupTranslationAPI Integration Partner Hotel Group Translation API
type IntegrationPartnerHotelGroupTranslationAPI struct {
	LanguageCode   *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:integration_partner_hotel_group_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	HotelGroupName *string `json:"hotel_group_name,omitempty" gorm:"type:varchar(512)"`                                                                                 // Hotel Group Name
	Description    *string `json:"description,omitempty" gorm:"type:text"`                                                                                              // Description
	Comment        *string `json:"comment,omitempty" gorm:"type:text"`                                                                                                  // Comment
}
