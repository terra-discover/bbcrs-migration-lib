package model

import (
	"github.com/terra-discover/bbcrs-migration-lib/lib"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// AirTravelerFlightSegment Air Traveler Flight Segment
type AirTravelerFlightSegment struct {
	basic.Base
	basic.DataOwner
	AirTravelerFlightSegmentAPI
}

// AirTravelerFlightSegmentAPI Air Traveler Flight Segment API
type AirTravelerFlightSegmentAPI struct {
	AirTravelerID   *uuid.UUID `json:"air_traveler_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"`
	FlightSegmentID *uuid.UUID `json:"flight_segment_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"`
	SeatColumn      *string    `json:"seat_column,omitempty" gorm:"type:varchar(1)"`
	SeatRow         *string    `json:"seat_row,omitempty" gorm:"type:varchar(3)"`
	SeatNumber      *string    `json:"seat_number,omitempty" gorm:"type:varchar(4)"`
}

// GetAirTravelerFlightSegment by query filter
func (data *AirTravelerFlightSegment) GetAirTravelerFlightSegment(tx *gorm.DB, queryFilters string) {
	qf, wf, _, _ := lib.CustomFilters(queryFilters, "", "")
	tx.Where(qf, wf...).First(&data)
}
