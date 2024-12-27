package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// LoyaltyConversionRate Loyalty Conversion Rate
type LoyaltyConversionRate struct {
	basic.Base
	basic.DataOwner
	LoyaltyConversionRateAPI
}

// LoyaltyConversionRateAPI Loyalty Conversion Rate API
type LoyaltyConversionRateAPI struct {
	LoyaltyProgramID *uuid.UUID `json:"loyalty_program_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"` // Loyalty Program Id
	CurrencyID       *uuid.UUID `json:"currency_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`                 // Currency Id
	RewardTypeID     *uuid.UUID `json:"reward_type_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`              // Reward type Id
	Amount           *float32   `json:"amount,omitempty" gorm:"type:decimal(19,4)" example:"1000.00"`                                     // Amount
	RewardAmount     *float32   `json:"reward_amount,omitempty" gorm:"type:decimal(19,4)" example:"1000.00"`                              // Reward Amount
}
