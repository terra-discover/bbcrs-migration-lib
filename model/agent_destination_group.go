package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// AgentDestinationGroup Model
type AgentDestinationGroup struct {
	basic.Base
	basic.DataOwner
	AgentDestinationGroupAPI
	DestinationGroup                 *DestinationGroup                  `json:"destination_group" gorm:"foreignKey:DestinationGroupID;references:ID"`
	Agent                            *Agent                             `json:"agent" gorm:"foreignKey:AgentID;references:ID"`
	AgentDestinationGroupDestination []AgentDestinationGroupDestination `json:"agent_destination_group_destination,omitempty"`
}

// AgentDestinationGroupAPI API
type AgentDestinationGroupAPI struct {
	AgentID            *uuid.UUID `json:"agent_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid"`
	DestinationGroupID *uuid.UUID `json:"destination_group_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid"`
}
