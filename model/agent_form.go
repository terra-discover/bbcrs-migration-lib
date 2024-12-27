package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// AgentForm Model
type AgentForm struct {
	basic.Base
	basic.DataOwner
	AgentFormAPI
	Form  *Form  `json:"form" gorm:"foreignKey:FormID;references:ID"`
	Agent *Agent `json:"agent" gorm:"foreignKey:AgentID;references:ID"`
}

// AgentFormAPI API
type AgentFormAPI struct {
	AgentID *uuid.UUID `json:"agent_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid"`
	FormID  *uuid.UUID `json:"form_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid"`
}
