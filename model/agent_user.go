package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// AgentUser Agent User
type AgentUser struct {
	basic.Base
	basic.DataOwner
	AgentUserAPI
	Agent       *Agent       `json:"agent,omitempty"`
	UserAccount *UserAccount `json:"user_account,omitempty"`
}

// AgentUserAPI Agent User Api
type AgentUserAPI struct {
	AgentID                     *uuid.UUID `json:"agent_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36);not null;"`
	UserAccountID               *uuid.UUID `json:"user_account_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36);not null;"`
	AssignAllDestinationSystems *bool      `json:"assign_all_destination_systems,omitempty"`
	AssignAllHotels             *bool      `json:"assign_all_hotels,omitempty"`
	AssignAllVendors            *bool      `json:"assign_all_vendors,omitempty"`
	AssignAllCorporates         *bool      `json:"assign_all_corporates,omitempty"`
	AssignAllSubagents          *bool      `json:"assign_all_subagents,omitempty"`
	AssignAllMembers            *bool      `json:"assign_all_members,omitempty"`
	IsPrimary                   *bool      `json:"is_primary,omitempty"`
}
