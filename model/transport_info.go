package model

import (
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

type TransportInfo struct {
	basic.Base
	basic.DataOwner
	TransportInfoAPI
}

type TransportInfoAPI struct {
	TransportationModeID   *uuid.UUID       `json:"transportation_mode_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"` // The reference to the transportation mode, e.g. plane, train, etc.
	AirlineCode            *string          `json:"airline_code,omitempty" gorm:"type:varchar(64)"`                                              // Identifies an airline by the code.
	FlightNumber           *string          `json:"flight_number,omitempty" gorm:"type:varchar(16)"`                                             // The flight number of the flight.
	AirportLocationCode    *string          `json:"airport_location_code,omitempty" gorm:"type:varchar(256)"`                                    // Location code used to identify an airport.
	TransportationDatetime *strfmt.DateTime `json:"transportation_datetime,omitempty" gorm:"type:timestamptz" format:"date-time"`                // The date and time of the transportation.
	AirEquipmentType       *string          `json:"air_equipment_type,omitempty" gorm:"type:varchar(64)"`                                        // The type of equipment used for the flight (3 character IATA code).
	Description            *string          `json:"description,omitempty" gorm:"type:text"`                                                      // The description of the attraction.
}
