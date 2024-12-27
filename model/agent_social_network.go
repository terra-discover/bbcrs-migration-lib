package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// AgentSocialNetwork Agent Social Network
type AgentSocialNetwork struct {
	basic.Base
	basic.DataOwner
	AgentSocialNetworkAPI
	Agent         *Agent         `json:"agent" gorm:"foreignKey:AgentID;references:ID"`
	SocialNetwork *SocialNetwork `json:"social_network" gorm:"foreignKey:SocialNetworkID;references:ID"`
}

// AgentSocialNetworkAPI Agent Social Network Api
type AgentSocialNetworkAPI struct {
	AgentID         *uuid.UUID `json:"agent_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid"`
	SocialNetworkID *uuid.UUID `json:"social_network_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid"`
}
