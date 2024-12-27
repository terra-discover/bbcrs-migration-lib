package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// AgentTravelPurpose Model
type AgentTravelPurpose struct {
	basic.Base
	basic.DataOwner
	AgentTravelPurposeAPI
	TravelPurpose *TravelPurpose `json:"travel_purpose" gorm:"foreignKey:TravelPurposeID;references:ID"`
	Agent         *Agent         `json:"agent" gorm:"foreignKey:AgentID;references:ID"`
}

// AgentTravelPurposeAPI API
type AgentTravelPurposeAPI struct {
	AgentID         *uuid.UUID `json:"agent_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid"`
	TravelPurposeID *uuid.UUID `json:"travel_purpose_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid"`
	IsDefault       *bool      `json:"is_default,omitempty"`
}
