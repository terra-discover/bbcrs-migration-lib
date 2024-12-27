package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// RewardRateAirlineCategory Reward Rate Airline Category
type RewardRateAirlineCategory struct {
	basic.Base
	basic.DataOwner
	RewardRateAirlineCategoryAPI
}

// RewardRateAirlineCategoryAPI Reward Rate Airline Category API
type RewardRateAirlineCategoryAPI struct {
	RewardRateID      *uuid.UUID `json:"reward_rate_id,omitempty" gorm:"type:varchar(36);not null;index:,unique,where:deleted_at is null"`       //Reward Rate Id
	AirlineCategoryID *uuid.UUID `json:"airline_category_id,omitempty" gorm:"type:varchar(256);not null;index:,unique,where:deleted_at is null"` // Airline Category Id
	IsIncluded        *bool      `json:"is_included,omitempty"`                                                                                  // Is Included
}
