package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// AgentTeamConsultant Model
type AgentTeamConsultant struct {
	basic.Base
	basic.DataOwner
	AgentTeamConsultantAPI
	Employee        *Employee        `json:"employee,omitempty" gorm:"foreignKey:EmployeeID;references:ID"`
	AgentTeam       *AgentTeam       `json:"agent_team,omitempty" gorm:"foreignKey:AgentTeamID;references:ID"`
	AgentConsultant *AgentConsultant `json:"agent_consultant,omitempty" gorm:"foreignKey:EmployeeID;references:EmployeeID"`
}

// AgentTeamConsultantAPI API
type AgentTeamConsultantAPI struct {
	AgentTeamID *uuid.UUID `json:"agent_team_id,omitempty" gorm:"type:varchar(36);index:idx_agent_team_id__employee_id,unique,where:deleted_at is null;not null;" format:"uuid"`
	EmployeeID  *uuid.UUID `json:"employee_id,omitempty" gorm:"type:varchar(36);index:idx_agent_team_id__employee_id,unique,where:deleted_at is null;not null;" format:"uuid"`
}

// AgentTeamConsultantRequest fake model schema
type AgentTeamConsultantRequest struct {
	TeamName    *string                           `json:"team_name" example:"Team BayuBuana" validate:"required"`
	TeamMembers *[]AgentTeamConsultantTeamMembers `json:"team_members" validate:"required"`
}

// AgentTeamConsultantTeamMembers fake model schema
type AgentTeamConsultantTeamMembers struct {
	EmployeeID       *uuid.UUID `json:"employee_id" format:"uuid"`
	GivenName        *string    `json:"given_name,omitempty" swaggerignore:"true"`
	MiddleName       *string    `json:"middle_name,omitempty" swaggerignore:"true"`
	Surname          *string    `json:"surname,omitempty" swaggerignore:"true"`
	IsLeader         *bool      `json:"is_leader" example:"true"`
	IsCanIssueTicket *bool      `json:"is_can_issue_ticket" example:"true"`
}

// AgentTeamConsultantData fake data model record
type AgentTeamConsultantData struct {
	Team
	TeamLeaderGivenName   *string               `json:"team_leader_given_name,omitempty"`
	TeamLeaderMiddleName  *string               `json:"team_leader_middle_name,omitempty"`
	TeamLeaderSurname     *string               `json:"team_leader_surname,omitempty"`
	NumberOfMembers       *int                  `json:"number_of_members,omitempty"`
	NameAndEmailOfMembers *string               `json:"name_and_email_of_members,omitempty"`
	TeamMembers           []AgentTeamConsultant `json:"team_members,omitempty" gorm:"foreignKey:ID"`
	TeamLeader            *AgentTeamConsultant  `json:"team_leader,omitempty" gorm:"foreignKey:ID"`
}
