package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// IntegrationPartnerZone Integration Partner Zone
type IntegrationPartnerZone struct {
	basic.Base
	basic.DataOwner
	IntegrationPartnerZoneAPI
	IntegrationPartnerZoneTranslation *IntegrationPartnerZoneTranslation `json:"integration_partner_zone_translation,omitempty"` // Integration Partner Zone Translation
	Zone                              *Zone                              `json:"zone,omitempty"`                                 // Zone
}

// IntegrationPartnerZoneAPI Integration Partner Zone API
type IntegrationPartnerZoneAPI struct {
	IntegrationPartnerID *uuid.UUID `json:"integration_partner_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36);index:idx_integration_partner_zone__zone_id,unique,where:deleted_at is null;not null"` // Integration Partner ID
	ZoneID               *uuid.UUID `json:"zone_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36);index:idx_integration_partner_zone__zone_id,unique,where:deleted_at is null;not null"`                // Zone ID
	DestinationCode      *string    `json:"destination_code,omitempty" gorm:"type:varchar(36)"`                                                                                                                               // Destination Code
	ZoneCode             *string    `json:"zone_code,omitempty" gorm:"type:varchar(36)"`                                                                                                                                      // Zone Code
	ZoneName             *string    `json:"zone_name,omitempty" gorm:"type:varchar(256)"`                                                                                                                                     // Zone Name
	CityCode             *string    `json:"city_code,omitempty" gorm:"type:varchar(36)"`                                                                                                                                      // City Code
	Description          *string    `json:"description,omitempty" gorm:"type:text"`                                                                                                                                           // Description
}
