package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// LoyaltyProgramProductType Loyalty Program Product Type
type LoyaltyProgramProductType struct {
	basic.Base
	basic.DataOwner
	LoyaltyProgram *LoyaltyProgram `json:"loyalty_program,omitempty"`
	ProductType    *ProductType    `json:"product_type,omitempty"`
	LoyaltyProgramProductTypeAPI
}

// LoyaltyProgramProductTypeAPI Loyalty Program Product Type Api
type LoyaltyProgramProductTypeAPI struct {
	LoyaltyProgramID *uuid.UUID `json:"loyalty_program_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"` // Loyalty Program Id
	ProductTypeID    *uuid.UUID `json:"product_type_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"`    // Product Type Id
}
