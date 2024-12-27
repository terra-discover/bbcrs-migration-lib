package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// LoyaltyBonus Operation Schedule
type LoyaltyBonus struct {
	basic.Base
	basic.DataOwner
	LoyaltyBonusAPI
}

// LoyaltyBonusAPI Operation Schedule API
type LoyaltyBonusAPI struct {
	LoyaltyProgramID *uuid.UUID `json:"loyalty_program_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"` // Loyalty Program Id
	BuiltInTypeID    *uuid.UUID `json:"built_in_type_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"`   // Built In Type Id
	LoyaltyLevelID   *uuid.UUID `json:"loyalty_level_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`            // Loyalty Level Id
	RewardTypeID     *uuid.UUID `json:"reward_type_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`              // Reward Type Id
	RewardAmount     *float32   `json:"description,omitempty" gorm:"type:decimal(19,4)" example:"1000.00"`                                // Reward Amount
}
