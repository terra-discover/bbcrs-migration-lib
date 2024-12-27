package model

import (
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// FlightCachingRequestRoute model | many to one relation
type FlightCachingRequestRoute struct {
	basic.Base
	basic.DataOwner
	SessionID            *uuid.UUID   `json:"session_id" gorm:"type:varchar(36);index:idx_flight_caching_request_route_session_id;not null"`
	FlightRequestID      *uuid.UUID   `json:"flight_request_id" gorm:"type:varchar(36)" format:"uuid"` // FlightRequestID
	DepartureAirportCode *string      `json:"departure_airport_code" gorm:"type:varchar(50)"`
	DepartureAirport     *string      `json:"departure_airport" gorm:"type:varchar(100)"`
	DepartureAirportID   *uuid.UUID   `json:"departure_airport_id" gorm:"type:varchar(36)"`
	DepartureCityID      *uuid.UUID   `json:"departure_city_id" gorm:"type:varchar(36)"`
	DepartureCityCode    *string      `json:"departure_city_code" gorm:"type:varchar(50)"`
	DepartureCity        *string      `json:"departure_city" gorm:"type:varchar(100)"`
	ArrivalAirportCode   *string      `json:"arrival_airport_code" gorm:"type:varchar(50)"`
	ArrivalAirport       *string      `json:"arrival_airport" gorm:"type:varchar(50)"`
	ArrivalAirportID     *uuid.UUID   `json:"arrival_airport_id" gorm:"type:varchar(36)"`
	ArrivalCityID        *uuid.UUID   `json:"arrival_city_id" gorm:"type:varchar(36)"`
	ArrivalCityCode      *string      `json:"arrival_city_code" gorm:"type:varchar(50)"`
	ArrivalCity          *string      `json:"arrival_city" gorm:"type:varchar(50)"`
	DepartureDate        *strfmt.Date `json:"departure_date" gorm:"type:date"`
	ReturnDate           *strfmt.Date `json:"return_date" gorm:"type:date"`
}
