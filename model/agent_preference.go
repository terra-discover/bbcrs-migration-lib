package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// AgentPreference Agent Preference
type AgentPreference struct {
	basic.Base
	basic.DataOwner
	AgentPreferenceAPI
	Agent      *Agent      `json:"agent" gorm:"foreignKey:AgentID;references:ID"`
	Preference *Preference `json:"preference" gorm:"foreignKey:PreferenceID;references:ID"`
}

// AgentPreferenceAPI Agent Preference Api
type AgentPreferenceAPI struct {
	AgentID      *uuid.UUID `json:"agent_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid"`
	PreferenceID *uuid.UUID `json:"preference_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid"`
}
