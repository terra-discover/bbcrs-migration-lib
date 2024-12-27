package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// AgentDistributionPartner Model
type AgentDistributionPartner struct {
	basic.Base
	basic.DataOwner
	AgentDistributionPartnerAPI
	Agent *Agent `json:"agent" gorm:"foreignKey:AgentID;references:ID"`
}

// AgentDistributionPartnerAPI API
type AgentDistributionPartnerAPI struct {
	AgentID               *uuid.UUID `json:"agent_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid"`
	DistributionPartnerID *uuid.UUID `json:"distribution_partner_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid"`
}
