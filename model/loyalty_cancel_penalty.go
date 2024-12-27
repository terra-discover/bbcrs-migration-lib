package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// LoyaltyCancelPenalty Loyalty Cancel Penalty
type LoyaltyCancelPenalty struct {
	basic.Base
	basic.DataOwner
	LoyaltyCancelPenaltyAPI
}

// LoyaltyCancelPenaltyAPI Loyalty Cancel Penalty API
type LoyaltyCancelPenaltyAPI struct {
	LoyaltyProgramID *uuid.UUID `json:"loyalty_program_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"` // Loyalty Program Id
	LoyaltyLevelID   *uuid.UUID `json:"loyalty_level_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`            // Loyalty Level Id
	ProductTypeID    *uuid.UUID `json:"product_type_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`             // Product type Id
	Percent          *float32   `json:"percent,omitempty" gorm:"type:decimal(19,4)" example:"1000.00"`                                    // Percent
}
