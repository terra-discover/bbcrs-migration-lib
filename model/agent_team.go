package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// AgentTeam Model
type AgentTeam struct {
	basic.Base
	basic.DataOwner
	AgentTeamAPI
	Team                *Team                 `json:"team,omitempty" gorm:"foreignKey:TeamID;references:ID" swaggerignore:"true"`
	Agent               *Agent                `json:"agent,omitempty" gorm:"foreignKey:AgentID;references:ID"`
	AgentTeamConsultant []AgentTeamConsultant `json:"agent_team_consultant,omitempty" swaggerignore:"true"`
}

// AgentTeamAPI API
type AgentTeamAPI struct {
	AgentID   *uuid.UUID `json:"agent_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid"`
	TeamID    *uuid.UUID `json:"team_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid"`
	IsDefault *bool      `json:"is_default,omitempty"`
}

// AgentTeamResponse Response
type AgentTeamResponse struct {
	AgentCorporate *AgentCorporate `json:"-" gorm:"foreignKey:CorporateID" swaggerignore:"true"` // Agent Corporate
	AgentTeam
	NumberOfMembers      *int                  `json:"number_of_members,omitempty" gorm:"->;-:migration"`
	AgentTeamConsultants []AgentTeamConsultant `json:"agent_team_consultants,omitempty" gorm:"-"` // Agent Team Consultant Assign Team
}
