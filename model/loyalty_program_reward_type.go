package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// LoyaltyProgramRewardType Loyalty Program RewardType
type LoyaltyProgramRewardType struct {
	basic.Base
	basic.DataOwner
	LoyaltyProgramRewardTypeAPI
}

// LoyaltyProgramRewardTypeAPI Loyalty Program RewardType API
type LoyaltyProgramRewardTypeAPI struct {
	LoyaltyProgramID *uuid.UUID `json:"loyalty_program_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"` // Loyalty Program Id
	RewardTypeID     *uuid.UUID `json:"reward_type_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"`     // Reward Type Id

}
