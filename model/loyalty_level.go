package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// LoyaltyLevel Loyalty Level
type LoyaltyLevel struct {
	basic.Base
	basic.DataOwner
	LoyaltyLevelAPI
	LoyaltyLevelTranslation *LoyaltyLevelTranslation `json:"loyalty_level_translation,omitempty"` // Loyalty Level Translation
	BuiltInType             *BuiltInType             `json:"built_in_type,omitempty"`             // Built In Type
	LoyaltyProgram          *LoyaltyProgram          `json:"loyalty_program,omitempty"`           // Loyalty Program
	LoyaltyLevelAsset       *LoyaltyLevelAsset       `json:"loyalty_level_asset,omitempty"`       // Loyalty Level Asset
	LoyaltyLevelContent     *LoyaltyLevelContent     `json:"loyalty_level_content,omitempty"`     // Loyalty Level Content
	RewardType              *RewardType              `json:"reward_type,omitempty"`               // Reward Type
}

// LoyaltyLevelAPI Loyalty Level API
type LoyaltyLevelAPI struct {
	LoyaltyProgramID    *uuid.UUID              `json:"loyalty_program_id,omitempty" swaggertype:"string" format:"uuid"`                                       // Loyalty Program ID
	BuiltInTypeID       *uuid.UUID              `json:"built_in_type_id,omitempty" swaggertype:"string" format:"uuid"`                                         // Built In Type ID
	LoyaltyLevelCode    *int                    `json:"loyalty_level_code,omitempty" gorm:"type:smallint;index:,unique,where:deleted_at is null;not null"`     // Loyalty Level Code
	LoyaltyLevelName    *string                 `json:"loyalty_level_name,omitempty" gorm:"type:varchar(256);index:,unique,where:deleted_at is null;not null"` // Loyalty Level Name
	RewardTypeID        *uuid.UUID              `json:"reward_type_id,omitempty" swaggertype:"string" format:"uuid"`                                           // Reward Type ID
	MinRewardAmount     *float64                `json:"min_reward_amount,omitempty"`                                                                           // Min Reward Amount
	LoyaltyLevelAsset   *LoyaltyLevelAssetAPI   `json:"loyalty_level_asset,omitempty" gorm:"-"`                                                                // Loyalty Level Asset
	LoyaltyLevelContent *LoyaltyLevelContentAPI `json:"loyalty_level_content,omitempty" gorm:"-"`                                                              // Loyalty Level Content
}
