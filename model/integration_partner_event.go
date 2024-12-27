package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// IntegrationPartnerEvent Integration Partner Event
type IntegrationPartnerEvent struct {
	basic.Base
	basic.DataOwner
	IntegrationPartnerEventAPI
	Event              *Event              `json:"event" gorm:"foreignKey:EventID;references:ID"`
	IntegrationPartner *IntegrationPartner `json:"integration_partner" gorm:"foreignKey:IntegrationPartnerID;references:ID"`
}

// IntegrationPartnerEventAPI Integration PartnerEvent API
type IntegrationPartnerEventAPI struct {
	IntegrationPartnerID *uuid.UUID `json:"integration_partner_id,omitempty" gorm:"type:varchar(36);not null;" swaggerignore:"true"`
	EventID              *uuid.UUID `json:"event_id,omitempty" gorm:"type:varchar(36);not null;" swaggertype:"string" format:"uuid"`
	EventCode            *string    `json:"event_code,omitempty" gorm:"type:varchar(256);" example:"102"`
	EventName            *string    `json:"event_name,omitempty" gorm:"type:varchar(256);" example:"Error parsing XML"`
	ShortDescription     *string    `json:"short_description,omitempty" gorm:"type:text"`
	Description          *string    `json:"description,omitempty" gorm:"type:text"`
}
