package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// TravelPolicyAirlineCategory struct
type TravelPolicyAirlineCategory struct {
	basic.Base
	basic.DataOwner
	TravelPolicyAirlineCategoryAPI
	TravelPolicy    *TravelPolicy    `json:"travel_policy,omitempty" gorm:"foreignKey:TravelPolicyID;references:ID"` // travel policy
	AirlineCategory *AirlineCategory `json:"airline_category,omitempty"`                                             // airline category
}

// TravelPolicyAirlineCategoryAPI struct
type TravelPolicyAirlineCategoryAPI struct {
	TravelPolicyID    *uuid.UUID `json:"travel_policy_id,omitempty" gorm:"type:varchar(36);not null" format:"uuid"`    // travel policy id
	AirlineCategoryID *uuid.UUID `json:"airline_category_id,omitempty" gorm:"type:varchar(36);not null" format:"uuid"` // airline id
	IsPreferred       *bool      `json:"is_preferred,omitempty" gorm:"type:bool" example:"true"`
	IsRestricted      *bool      `json:"is_restricted,omitempty" example:"true"`
}
