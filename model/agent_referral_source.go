package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// AgentReferralSource Agent Referral Source
type AgentReferralSource struct {
	basic.Base
	basic.DataOwner
	AgentReferralSourceAPI
	Agent          *Agent          `json:"agent" gorm:"foreignKey:AgentID;references:ID"`
	ReferralSource *ReferralSource `json:"referral_source" gorm:"foreignKey:ReferralSourceID;references:ID"`
}

// AgentReferralSourceAPI Agent Referral Source Api
type AgentReferralSourceAPI struct {
	AgentID          *uuid.UUID `json:"agent_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid"`
	ReferralSourceID *uuid.UUID `json:"referral_source_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid"`
}
