package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// AgentIntegrationPartner Agent Integration Partner
type AgentIntegrationPartner struct {
	basic.Base
	basic.DataOwner
	AgentIntegrationPartnerAPI
	Agent              *Agent              `json:"agent" gorm:"foreignKey:AgentID;references:ID"`
	IntegrationPartner *IntegrationPartner `json:"integration_partner" gorm:"foreignKey:IntegrationPartnerID;references:ID"` // Integration Partner
	// TODO: need review from master-service
	// Hotel              *Hotel              `json:"hotel" gorm:"foreignKey:ID"`
}

// AgentIntegrationPartnerAPI Agent Integration Partner Api
type AgentIntegrationPartnerAPI struct {
	AgentID              *uuid.UUID `json:"agent_id,omitempty" gorm:"type:varchar(36);not null;index:idx__agent_integration_partner_id,unique,where:deleted_at is null" swaggertype:"string" format:"uuid"`
	IntegrationPartnerID *uuid.UUID `json:"integration_partner_id,omitempty" gorm:"type:varchar(36);not null;index:idx__agent_integration_partner_id,unique,where:deleted_at is null" swaggertype:"string" format:"uuid"`
}
