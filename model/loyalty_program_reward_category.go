package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// LoyaltyProgramRewardCategory Loyalty Program Reward Category
type LoyaltyProgramRewardCategory struct {
	basic.Base
	basic.DataOwner
	LoyaltyProgramRewardCategoryAPI
}

// LoyaltyProgramRewardCategoryAPI Loyalty Program Reward Category API
type LoyaltyProgramRewardCategoryAPI struct {
	LoyaltyProgramID  *uuid.UUID `json:"loyalty_program_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"` // Loyalty Program Id
	RewardCategoryID  *uuid.UUID `json:"reward_category_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"` // Reward Category Id
	LoyaltyLevelID    *uuid.UUID `json:"loyalty_level_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"`   // Loyalty Level Id
	IsSelfBookingTool *bool      `json:"is_self_booking_tool,omitempty"`                                                                   // Is Self Booking Tool
}
