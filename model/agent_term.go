package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// AgentTerm Model
type AgentTerm struct {
	basic.Base
	basic.DataOwner
	AgentTermAPI
	Term  *Term  `json:"term" gorm:"foreignKey:TermID;references:ID"`
	Agent *Agent `json:"agent" gorm:"foreignKey:AgentID;references:ID"`
}

// AgentTermAPI API
type AgentTermAPI struct {
	AgentID *uuid.UUID `json:"agent_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid"`
	TermID  *uuid.UUID `json:"term_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid"`
}
