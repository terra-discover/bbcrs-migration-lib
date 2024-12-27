package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// AgentAdPlacement AgentAdPlacement
type AgentAdPlacement struct {
	basic.Base
	basic.DataOwner
	AgentAdPlacementAPI
	Agent       *Agent       `json:"agent,omitempty" gorm:"foreignKey:AgentID;references:ID"`
	AdPlacement *AdPlacement `json:"ad_placement,omitempty" gorm:"foreignKey:AdPlacementID;references:ID"`
}

// AgentAdPlacementAPI AgentAdPlacement API
type AgentAdPlacementAPI struct {
	AgentID       *uuid.UUID `json:"agent_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36);not null;"`
	AdPlacementID *uuid.UUID `json:"ad_placement_id,omitempty" format:"uuid" gorm:"type:varchar(36);not null;"`
}
