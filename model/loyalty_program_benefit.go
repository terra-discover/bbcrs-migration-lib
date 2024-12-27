package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// LoyaltyProgramBenefit Loyalty Program Benefit
type LoyaltyProgramBenefit struct {
	basic.Base
	basic.DataOwner
	LoyaltyProgramBenefitAPI
}

// LoyaltyProgramBenefitAPI Loyalty Program Benefit API
type LoyaltyProgramBenefitAPI struct {
	LoyaltyProgramID *uuid.UUID `json:"loyalty_program_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"` // Loyalty Program Id
	BenefitID        *uuid.UUID `json:"benefit_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"`         // Benefit Id
	LoyaltyLevelID   *uuid.UUID `json:"loyalty_level_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"`   // Loyalty Level Id
}
