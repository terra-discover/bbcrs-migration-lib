package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// AgentModule AgentModule
type AgentModule struct {
	basic.Base
	basic.DataOwner
	AgentModuleAPI
	Agent  *Agent  `json:"agent,omitempty"`
	Module *Module `json:"module,omitempty"`
}

// AgentModuleAPI AgentModule API
type AgentModuleAPI struct {
	AgentID  *uuid.UUID `json:"agent_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36);not null;"`
	ModuleID *uuid.UUID `json:"module_id,omitempty" format:"uuid" gorm:"type:varchar(36);not null;"`
}
