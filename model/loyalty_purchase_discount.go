package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// LoyaltyPurchaseDiscount Loyalty Purchase Discount
type LoyaltyPurchaseDiscount struct {
	basic.Base
	basic.DataOwner
	LoyaltyPurchaseDiscountAPI
}

// LoyaltyPurchaseDiscountAPI Loyalty Purchase Discount API
type LoyaltyPurchaseDiscountAPI struct {
	LoyaltyProgramID *uuid.UUID `json:"loyalty_program_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"` // Loyalty Program Id
	RewardTypeID     *uuid.UUID `json:"reward_type_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"`     // Reward Type Id
	MinRewardAmount  *float64   `json:"min_reward_amount,omitempty" gorm:"type:decimal(19,4)"`                                            // Min Reward Amount
	DiscountPercent  *float64   `json:"discount_percent,omitempty" gorm:"type:decimal(19,4)"`                                             // Discount Percent
}
