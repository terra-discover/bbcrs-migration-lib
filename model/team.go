package model

import (
	"github.com/terra-discover/bbcrs-migration-lib/lib"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Team struct Team
type Team struct {
	basic.Base
	basic.DataOwner
	TeamAPI
	TeamTranslation     *TeamTranslation     `json:"team_translation,omitempty"`
	AgentTeamConsultant *AgentTeamConsultant `json:"agent_team_consultant,omitempty" gorm:"foreignKey:LeaderID;references:EmployeeID"`
}

// TeamAPI Team API
type TeamAPI struct {
	TeamCode *string    `json:"team_code,omitempty" gorm:"type:varchar(36);not null" example:"Sales"`                                                // Team Code
	TeamName *string    `json:"team_name,omitempty" gorm:"type:varchar(256);not null;index:,unique,where:deleted_at is null" example:"Sales Team 1"` // Team Name
	LeaderID *uuid.UUID `json:"leader_id,omitempty" gorm:"type:varchar(36)"`                                                                         // Leader ID
}

// Seed Team
func (team *Team) Seed() *Team {
	seed := Team{
		TeamAPI: TeamAPI{
			TeamCode: lib.Strptr("Sales"),
			TeamName: lib.Strptr("Sales Team 1"),
			LeaderID: lib.UUIDPtr(uuid.New()),
		},
	}
	_ = lib.Merge(seed, &team)
	return team
}

// BeforeCreate Data
func (team *Team) BeforeCreate(tx *gorm.DB) error {
	err := team.Base.BeforeCreate(tx)
	if nil == team.TeamCode {
		id, _ := uuid.NewRandom()
		team.TeamCode = lib.Strptr(id.String())
	}
	return err
}
