package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// FlightSegmentStopLocation Flight Segment Stop Location
type FlightSegmentStopLocation struct {
	basic.Base
	basic.DataOwner
	FlightSegmentStopLocationAPI
}

// FlightSegmentStopLocationAPI Flight Segment Stop Location API
type FlightSegmentStopLocationAPI struct {
	FlightSegmentID *uuid.UUID `json:"flight_segment_id,omitempty" gorm:"varchar(36);not null" swaggertype:"string" format:"uuid"` // he reference to the flight segment.
	LocationCode    *string    `json:"location_code,omitempty" gorm:"varchar(8)"`                                                  // A code used to identify a location.
}
