package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// AgentCorporateTeam Model
type AgentCorporateTeam struct {
	basic.Base
	basic.DataOwner
	AgentCorporateTeamAPI
	AgentCorporate *AgentCorporate `json:"agent_corporate" gorm:"foreignKey:AgentCorporateID;references:ID"`
	Team           *Team           `json:"team" gorm:"foreignKey:TeamID;references:ID"`
}

// AgentCorporateTeamAPI API
type AgentCorporateTeamAPI struct {
	AgentCorporateID *uuid.UUID `json:"agent_corporate_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid"`
	TeamID           *uuid.UUID `json:"team_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid"`
	IsPrimary        *bool      `json:"is_primary,omitempty" gorm:"type:bool"`
}
