package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// AgentViewership Model
type AgentViewership struct {
	basic.Base
	basic.DataOwner
	AgentViewershipAPI
	Agent *Agent `json:"agent" gorm:"foreignKey:AgentID;references:ID"`
}

// AgentViewershipAPI API
type AgentViewershipAPI struct {
	AgentID      *uuid.UUID `json:"agent_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid"`
	ViewershipID *uuid.UUID `json:"viewership_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid"`
	IsDefault    *bool      `json:"is_default,omitempty"`
}
