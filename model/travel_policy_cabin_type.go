package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// TravelPolicyCabinType Travel Policy Job Title
type TravelPolicyCabinType struct {
	basic.Base
	basic.DataOwner
	TravelPolicyCabinTypeAPI
	TravelPolicy *TravelPolicy `json:"travel_policy,omitempty" gorm:"foreignKey:TravelPolicyID;references:ID"` // travel policy
	CabinType    *CabinType    `json:"cabin_type,omitempty" gorm:"foreignKey:CabinTypeID;reference:ID"`        // cabin type
}

// TravelPolicyCabinTypeAPI Travel Policy Job Title API
type TravelPolicyCabinTypeAPI struct {
	TravelPolicyID             *uuid.UUID `json:"travel_policy_id,omitempty" gorm:"type:varchar(36);not null" format:"uuid"` // travel policy id
	CabinTypeID                *uuid.UUID `json:"cabin_type_id,omitempty" gorm:"type:varchar(36);not null" format:"uuid"`    // cabin type id
	IsShortFlight              *bool      `json:"is_short_flight,omitempty" example:"true"`
	IsLongFlight               *bool      `json:"is_long_flight,omitempty" example:"true"`
	MaximumShortFlightDuration *float64   `json:"maximum_short_flight_duration,omitempty"`
}
