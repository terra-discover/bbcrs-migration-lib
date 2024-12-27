package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

type FlightCachingFareDetailBaggageAllowance struct {
	basic.Base
	basic.DataOwner
	SessionID *uuid.UUID `json:"session_id,omitempty" gorm:"type:varchar(36);index:idx_flight_caching_fare_detail_baggage_allowance_session_id;not null"`
	FlightCachingFareDetailBaggageAllowanceAPI
}

type FlightCachingFareDetailBaggageAllowanceAPI struct {
	FlightCachingRouteID      *uuid.UUID `json:"flight_caching_route_id,omitempty" gorm:"type:varchar(36)" format:"uuid" swaggertype:"string"`
	FlightCachingFareDetailID *uuid.UUID `json:"flight_caching_fare_detail_id,omitempty" gorm:"type:varchar(36)" format:"uuid" swaggertype:"string"`
	Quantity                  *float64   `json:"quantity,omitempty"`
	UnitOfMeasureName         *string    `json:"unit_of_measure_name,omitempty" gorm:"type:varchar(36)"`
}
