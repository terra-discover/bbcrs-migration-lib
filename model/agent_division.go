package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// AgentDivision Model
type AgentDivision struct {
	basic.Base
	basic.DataOwner
	AgentDivisionAPI
	Division *Division `json:"division" gorm:"foreignKey:DivisionID;references:ID"`
	Agent    *Agent    `json:"agent" gorm:"foreignKey:AgentID;references:ID"`
}

// AgentDivisionAPI API
type AgentDivisionAPI struct {
	AgentID    *uuid.UUID `json:"agent_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid"`
	DivisionID *uuid.UUID `json:"division_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid"`
	IsDefault  *bool      `json:"is_default,omitempty"`
}
