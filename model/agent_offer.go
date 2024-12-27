package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// AgentOffer Agent Offer
type AgentOffer struct {
	basic.Base
	basic.DataOwner
	AgentOfferAPI
	Agent *Agent `json:"agent" gorm:"foreignKey:AgentID;references:ID"`
}

// AgentOfferAPI Agent Offer Api
type AgentOfferAPI struct {
	AgentID *uuid.UUID `json:"agent_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid"`
	OfferID *uuid.UUID `json:"offer_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid"`
}
