package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// FlightCachingCabins model
type FlightCachingCabins struct {
	basic.Base
	basic.DataOwner
	SessionID          *uuid.UUID                           `json:"session_id,omitempty" gorm:"type:varchar(36);index:idx_flight_caching_cabins_session_id;not null"`
	FlightCachingID    *uuid.UUID                           `json:"flight_caching_id,omitempty" gorm:"type:varchar(36);index:idx_flight_caching_cabins_flight_caching_id;not null" format:"uuid"` // FlightCachingID
	CabinTypeID        *uuid.UUID                           `json:"cabin_type_id"`                                                                                                                // database cabin type id
	CabinTypeName      *string                              `json:"cabin_type_name"`                                                                                                              // database cabin type name
	Name               *string                              `json:"name,omitempty" gorm:"type:varchar(255)"`
	FlightCachingFares *[]FlightCachingFares                `json:"fares,omitempty" gorm:"foreignKey:FlightCachingCabinsID;references:ID"`
	Violations         []FlightCachingTravelPolicyViolation `json:"violations"`
}
