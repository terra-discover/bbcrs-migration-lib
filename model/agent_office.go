package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// AgentOffice Model
type AgentOffice struct {
	basic.Base
	basic.DataOwner
	AgentOfficeAPI
	Office *Office `json:"office" gorm:"foreignKey:OfficeID;references:ID"`
	Agent  *Agent  `json:"agent" gorm:"foreignKey:AgentID;references:ID"`
}

// AgentOfficeAPI API
type AgentOfficeAPI struct {
	AgentID   *uuid.UUID `json:"agent_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid"`
	OfficeID  *uuid.UUID `json:"office_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid"`
	IsPrimary *bool      `json:"is_primary,omitempty"`
}
