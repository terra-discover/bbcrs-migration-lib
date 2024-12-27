package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// CommissionClaimOriginDestination Commission Claim Origin Destination
type CommissionClaimOriginDestination struct {
	basic.Base
	basic.DataOwner
	CommissionClaimOriginDestinationAPI
	CommissionClaim *CommissionClaim `json:"commission_claim,omitempty" gorm:"foreignKey:CommissionClaimID;references:ID" swaggerignore:"true"`
}

// CommissionClaimOriginDestinationAPI Commission Claim Origin Destination API
type CommissionClaimOriginDestinationAPI struct {
	CommissionClaimID            *uuid.UUID `json:"commision_claim_id,omitempty" gorm:"not null" swaggertype:"string" format:"uuid"`     // Commission Claim ID
	DepartureAirportLocationCode *string    `json:"departure_airport_location_code,omitempty" example:"code_123" gorm:"type:varchar(8)"` // Departure Airport Location Code
	DepartureCityCode            *string    `json:"departure_city_code,omitempty" example:"code 123" gorm:"type:varchar(8)"`             // Departure City Code
	ArrivalAirportLocationCode   *string    `json:"arrival_airport_location_code,omitempty" example:"code_123" gorm:"type:varchar(8)"`   // Arrival Airport Location Code
	ArrivalCityCode              *string    `json:"arrival_city_code,omitempty" example:"code 123" gorm:"type:varchar(8)"`               // Arrival City Code
}
