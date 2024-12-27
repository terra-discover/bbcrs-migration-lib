package model

import (
	"github.com/terra-discover/bbcrs-migration-lib/lib"

	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
	"gorm.io/gorm"
)

// AirOriginDestinationOption Air Origin Destination Option
type AirOriginDestinationOption struct {
	basic.Base
	basic.DataOwner
	AirOriginDestinationOptionAPI
	FlightSegment         []FlightSegment         `json:"flight_segment,omitempty"`
	DepartureAirport      Airport                 `json:"departure_airport,omitempty" gorm:"foreignKey:DepartureAirportLocationCode;references:AirportCode"`
	ArrivalAirport        Airport                 `json:"arrival_airport,omitempty" gorm:"foreignKey:ArrivalAirportLocationCode;references:AirportCode"`
	FlightViolationReason []FlightViolationReason `json:"flight_violation_reason,omitempty"`
}

// AirOriginDestinationOptionAPI Air Origin Destination Option API
type AirOriginDestinationOptionAPI struct {
	AirItineraryID               *uuid.UUID       `json:"air_itinerary_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"`
	OriginDestinationStatus      *int             `json:"origin_destination_status,omitempty" gorm:"type:smallint"`
	DepartureDatetime            *strfmt.DateTime `json:"departure_datetime,omitempty" format:"date-time" swaggertype:"string" gorm:"type:timestamptz"`
	ArrivalDatetime              *strfmt.DateTime `json:"arrival_datetime,omitempty" format:"date-time" swaggertype:"string" gorm:"type:timestamptz"`
	ElapsedTime                  *int             `json:"elapsed_time,omitempty" gorm:"type:smallint"` // value stored in minute, not second
	DepartureAirportLocationCode *string          `json:"departure_airport_location_code,omitempty" gorm:"type:varchar(8)"`
	ArrivalAirportLocationCode   *string          `json:"arrival_airport_location_code,omitempty" gorm:"type:varchar(8)"`
}

// GetAirOriginDestinationOption by query filter
func (data *AirOriginDestinationOption) GetAirOriginDestinationOption(tx *gorm.DB, queryFilters string) {
	qf, wf, _, _ := lib.CustomFilters(queryFilters, "", "")
	tx.Where(qf, wf...).Take(&data)
}

// GetAirOriginDestinationOptions by query filter
func (data *AirOriginDestinationOption) GetAirOriginDestinationOptions(tx *gorm.DB, queryFilters string) []AirOriginDestinationOption {
	res := []AirOriginDestinationOption{}
	qf, wf, _, _ := lib.CustomFilters(queryFilters, "", "")
	tx.Where(qf, wf...).Find(&res)
	return res
}
