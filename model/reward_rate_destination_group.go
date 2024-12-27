package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// RewardRateDestinationGroup Reward Rate Destination Group
type RewardRateDestinationGroup struct {
	basic.Base
	basic.DataOwner
	RewardRateDestinationGroupAPI
}

// RewardRateDestinationGroupAPI Reward Rate Destination Group API
type RewardRateDestinationGroupAPI struct {
	RewardRateID       *uuid.UUID `json:"reward_rate_id,omitempty" gorm:"type:varchar(36);not null;index:,unique,where:deleted_at is null"`        //Reward Rate Id
	DestinationGroupID *uuid.UUID `json:"destination_group_id,omitempty" gorm:"type:varchar(256);not null;index:,unique,where:deleted_at is null"` // Destination Group Id
	IsIncluded         *bool      `json:"is_included,omitempty"`                                                                                   // Is Included
}
