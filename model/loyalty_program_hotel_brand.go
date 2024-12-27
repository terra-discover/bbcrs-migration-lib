package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// LoyaltyProgramHotelBrand Loyalty Program Hotel Brand
type LoyaltyProgramHotelBrand struct {
	basic.Base
	basic.DataOwner
	LoyaltyProgramHotelBrandAPI
}

// LoyaltyProgramHotelBrandAPI Loyalty Program Hotel Brand
type LoyaltyProgramHotelBrandAPI struct {
	LoyaltyProgramID *uuid.UUID `json:"loyalty_program_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"` // Loyalty Program Id
	HotelBrandID     *uuid.UUID `json:"hotel_brand_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"`     // Hotel Brand Id
}
