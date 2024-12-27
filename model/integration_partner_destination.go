package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// IntegrationPartnerDestination Integration Partner Destination
type IntegrationPartnerDestination struct {
	basic.Base
	basic.DataOwner
	IntegrationPartnerDestinationAPI
	IntegrationPartnerDestinationTranslation *IntegrationPartnerDestinationTranslation `json:"integration_partner_destination_translation,omitempty"` // Integration Partner Destination Translation
	Destination                              *Destination                              `json:"destination,omitempty"`                                 // Destination
}

// IntegrationPartnerDestinationAPI Integration Partner Destination API
type IntegrationPartnerDestinationAPI struct {
	IntegrationPartnerID *uuid.UUID `json:"integration_partner_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36);index:idx_idx_integration_partner_destination__destination_id,unique,where:deleted_at is null;not null"` // Integration Partner ID
	DestinationID        *uuid.UUID `json:"destination_id,omitempty" swaggertype:"string" format:"uuid" gorm:"index:idx_idx_integration_partner_destination__destination_id,unique,where:deleted_at is null;not null"`                          // Destination ID
	DestinationCode      *string    `json:"destination_code,omitempty" gorm:"type:varchar(36)"`                                                                                                                                                 // Destination Code
	DestinationName      *string    `json:"destination_name,omitempty" gorm:"type:varchar(256)"`                                                                                                                                                // Destination Name
	CountryCode          *string    `json:"country_code,omitempty" gorm:"type:varchar(36)"`                                                                                                                                                     // Country Code
	CityCode             *string    `json:"city_code,omitempty" gorm:"type:varchar(36)"`                                                                                                                                                        // City Code
	Description          *string    `json:"description,omitempty" gorm:"type:text"`                                                                                                                                                             // Description
}
