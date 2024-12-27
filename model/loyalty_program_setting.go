package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// LoyaltyProgramSetting Loyalty Program Setting
type LoyaltyProgramSetting struct {
	basic.Base
	basic.DataOwner
	LoyaltyProgramSettingAPI
}

// LoyaltyProgramSettingAPI Loyalty Program Setting Api
type LoyaltyProgramSettingAPI struct {
	LoyaltyProgramID   *uuid.UUID `json:"loyalty_program_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"` // Loyalty Program Id
	LoyaltyLevelAge    *int       `json:"loyalty_level_age,omitempty" gorm:"type:smallint"`                                                 // Loyalty Level Age
	LoyaltyRewardAge   *int       `json:"loyalty_reward_age,omitempty" gorm:"type:smallint"`                                                // Loyalty Reward Age
	LoyaltyPurchaseAge *int       `json:"loyalty_purchase_age,omitempty" gorm:"type:smallint"`                                              // Loyalty Purchase Age
}
