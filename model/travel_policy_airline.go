package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// TravelPolicyAirline Travel Policy Job Title
type TravelPolicyAirline struct {
	basic.Base
	basic.DataOwner
	TravelPolicyAirlineAPI
	TravelPolicy *TravelPolicy `json:"travel_policy,omitempty" gorm:"foreignKey:TravelPolicyID;references:ID"` // travel policy
	Airline      *Airline      `json:"airline,omitempty" gorm:"foreignKey:AirlineID;reference:ID"`             // airline
}

// TravelPolicyAirlineAPI Travel Policy Job Title API
type TravelPolicyAirlineAPI struct {
	TravelPolicyID *uuid.UUID `json:"travel_policy_id,omitempty" gorm:"type:varchar(36);not null" format:"uuid"` // travel policy id
	AirlineID      *uuid.UUID `json:"airline_id,omitempty" gorm:"type:varchar(36);not null" format:"uuid"`       // airline id
	IsPreferred    *bool      `json:"is_preferred,omitempty" example:"true"`
	IsRestricted   *bool      `json:"is_restricted,omitempty" example:"true"`
}
