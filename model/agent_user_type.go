package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// AgentUserType Agent User Type
type AgentUserType struct {
	basic.Base
	basic.DataOwner
	AgentUserTypeAPI
	Agent *Agent `json:"agent,omitempty"`
}

// AgentUserTypeAPI Agent User Type Api
type AgentUserTypeAPI struct {
	AgentID                     *uuid.UUID `json:"agent_id,omitempty" gorm:"type:varchar(36); not null" swaggertype:"string" format:"uuid"`
	UserTypeID                  *uuid.UUID `json:"user_type_id,omitempty" gorm:"type:varchar(36); not null" swaggertype:"string" format:"uuid"`
	AssignAllDestinationSystems *bool      `json:"assign_all_destination_systems,omitempty" gorm:"type:bool"`
	AssignAllHotels             *bool      `json:"assign_all_hotels,omitempty" gorm:"type:bool"`
	AssignAllVendors            *bool      `json:"assign_all_vendors,omitempty" gorm:"type:bool"`
	AssignAllCorporates         *bool      `json:"assign_all_corporates,omitempty" gorm:"type:bool"`
	AssignAllSubagents          *bool      `json:"assign_all_subagents,omitempty" gorm:"type:bool"`
	AssignAllMembers            *bool      `json:"assign_all_members,omitempty" gorm:"type:bool"`
	IsDefault                   *bool      `json:"is_default,omitempty" gorm:"type:bool"`
}
