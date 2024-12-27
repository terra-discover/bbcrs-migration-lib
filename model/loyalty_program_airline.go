package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// LoyaltyProgramAirline Loyalty Program Airline
type LoyaltyProgramAirline struct {
	basic.Base
	basic.DataOwner
	LoyaltyProgramAirlineAPI
}

// LoyaltyProgramAirlineAPI Loyalty Program Airline API
type LoyaltyProgramAirlineAPI struct {
	LoyaltyProgramID *uuid.UUID `json:"loyalty_program_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"` // Loyalty Program Id
	AirlineID        *uuid.UUID `json:"airline_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`                  // Airline Id
}
