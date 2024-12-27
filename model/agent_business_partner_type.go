package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// AgentBusinessPartnerType Model
type AgentBusinessPartnerType struct {
	basic.Base
	basic.DataOwner
	AgentBusinessPartnerTypeAPI
	Agent *Agent `json:"agent" gorm:"foreignKey:AgentID;references:ID"`
}

// AgentBusinessPartnerTypeAPI API
type AgentBusinessPartnerTypeAPI struct {
	AgentID               *uuid.UUID `json:"agent_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid"`
	BusinessPartnerTypeID *uuid.UUID `json:"business_partner_type_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid"`
}
