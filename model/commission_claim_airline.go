package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// CommissionClaimAirline Commission Claim Airlane
type CommissionClaimAirline struct {
	basic.Base
	basic.DataOwner
	CommissionClaimAirlineAPI
	CommissionClaim *CommissionClaim `json:"commission_claim,omitempty" gorm:"foreignKey:CommissionClaimID;references:ID" swaggerignore:"true"`
	Airline         *Airline         `json:"airline" gorm:"foreignKey:AirlineID;references:ID"`
}

// CommissionClaimAirlineAPI Commission Claim Airlane API
type CommissionClaimAirlineAPI struct {
	CommissionClaimID *uuid.UUID `json:"commision_claim_id,omitempty" gorm:"not null" swaggertype:"string" format:"uuid"` // Commission Claim ID
	AirlineID         *uuid.UUID `json:"airline_id,omitempty" gorm:"not null" swaggertype:"string" format:"uuid"`         // Airlane ID
}
