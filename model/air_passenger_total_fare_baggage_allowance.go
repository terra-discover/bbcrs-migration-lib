package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

type AirPassengerTotalFareBaggageAllowance struct {
	basic.Base
	basic.DataOwner
	AirPassengerTotalFareBaggageAllowanceAPI
}

type AirPassengerTotalFareBaggageAllowanceAPI struct {
	AirPassengerTotalFareID *uuid.UUID `json:"air_passenger_total_fare_id,omitempty" gorm:"type:varchar(36);not null;" swaggertype:"string" format:"uuid"`
	Quantity                *float64   `json:"quantity,omitempty"`
	UnitOfMeasureID         *string    `json:"unit_of_measure_id,omitempty" gorm:"type:varchar(36)"`
	FlightSegmentID         *uuid.UUID `json:"flight_segment_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"`
}
