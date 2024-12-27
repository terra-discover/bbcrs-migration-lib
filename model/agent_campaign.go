package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// AgentCampaign Model
type AgentCampaign struct {
	basic.Base
	basic.DataOwner
	AgentCampaignAPI
	Campaign *Campaign `json:"campaign" gorm:"foreignKey:CampaignID;references:ID"`
	Agent    *Agent    `json:"agent" gorm:"foreignKey:AgentID;references:ID"`
}

// AgentCampaignAPI API
type AgentCampaignAPI struct {
	AgentID    *uuid.UUID `json:"agent_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid"`
	CampaignID *uuid.UUID `json:"campaign_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid"`
}
