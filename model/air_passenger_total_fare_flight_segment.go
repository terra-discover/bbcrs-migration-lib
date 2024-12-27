package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

type AirPassengerTotalFareFlightSegment struct {
	basic.Base
	basic.DataOwner
	AirPassengerTotalFareFlightSegmentAPI
}

type AirPassengerTotalFareFlightSegmentAPI struct {
	AirPassengerTotalFareID *uuid.UUID `json:"air_passenger_total_fare_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"`
	FlightSegmentID         *uuid.UUID `json:"flight_segment_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"`
}
