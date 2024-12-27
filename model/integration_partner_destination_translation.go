package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// IntegrationPartnerDestinationTranslation Integration Partner Destination Translation
type IntegrationPartnerDestinationTranslation struct {
	basic.Base
	basic.DataOwner
	IntegrationPartnerDestinationTranslationAPI
	IntegrationPartnerDestinationID *uuid.UUID `json:"integration_partner_destination_id,omitempty" gorm:"type:varchar(36);uniqueIndex:integration_partner_destination_translation_unique;not null" swaggertype:"string" format:"uuid"` // Integration Partner Destination id
}

// IntegrationPartnerDestinationTranslationAPI Integration Partner Destination Translation API
type IntegrationPartnerDestinationTranslationAPI struct {
	LanguageCode    *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:integration_partner_destination_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	DestinationName *string `json:"destination_name,omitempty" gorm:"type:varchar(256)"`                                                                                 // Destination Name
	Description     *string `json:"description,omitempty" gorm:"type:text"`                                                                                              // Description
}
