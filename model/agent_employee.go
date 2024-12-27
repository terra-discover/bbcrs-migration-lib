package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// AgentEmployee Model
type AgentEmployee struct {
	basic.Base
	basic.DataOwner
	AgentEmployeeAPI
	Employee *Employee `json:"employee" gorm:"foreignKey:EmployeeID;references:ID"`
	Agent    *Agent    `json:"agent" gorm:"foreignKey:AgentID;references:ID"`
}

// AgentEmployeeAPI API
type AgentEmployeeAPI struct {
	AgentID    *uuid.UUID `json:"agent_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid"`
	EmployeeID *uuid.UUID `json:"employee_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid"`
}
