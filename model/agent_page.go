package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// AgentPage Model
type AgentPage struct {
	basic.Base
	basic.DataOwner
	AgentPageAPI
	// Page *Pages `json:"page" gorm:"foreignKey:PageID;references:ID"`
	Agent *Agent `json:"agent" gorm:"foreignKey:AgentID;references:ID"`
}

// AgentPageAPI API
type AgentPageAPI struct {
	AgentID *uuid.UUID `json:"agent_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid"`
	PageID  *uuid.UUID `json:"page_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid"`
}
