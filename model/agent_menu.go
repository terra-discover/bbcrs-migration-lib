package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// AgentMenu AgentMenu
type AgentMenu struct {
	basic.Base
	basic.DataOwner
	AgentMenuAPI
	Agent *Agent `json:"agent,omitempty"`
	Menu  *Menu  `json:"menu,omitempty"`
}

// AgentMenuAPI AgentMenu API
type AgentMenuAPI struct {
	AgentID *uuid.UUID `json:"agent_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36);not null;"`
	MenuID  *uuid.UUID `json:"menu_id,omitempty" format:"uuid" gorm:"type:varchar(36);not null;"`
}
