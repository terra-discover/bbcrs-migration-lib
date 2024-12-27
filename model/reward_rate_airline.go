package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// RewardRateAirline Reward Rate Airline
type RewardRateAirline struct {
	basic.Base
	basic.DataOwner
	RewardRateAirlineAPI
}

// RewardRateAirlineAPI Religion API
type RewardRateAirlineAPI struct {
	RewardRateID *uuid.UUID `json:"reward_rate_id,omitempty" gorm:"type:varchar(36);not null;index:,unique,where:deleted_at is null"` //Reward Rate Id
	AirlineID    *uuid.UUID `json:"airline_id,omitempty" gorm:"type:varchar(256);not null;index:,unique,where:deleted_at is null"`    // Airline Id
	IsIncluded   *bool      `json:"is_included,omitempty"`                                                                            // Is Included
}
