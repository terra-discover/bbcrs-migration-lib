package model

import (
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// FlightSegment Flight Segment
type FlightSegment struct {
	basic.Base
	basic.DataOwner
	FlightSegmentAPI
	FlightSegmentStopLocation []FlightSegmentStopLocation
	CabinType                 *CabinType `json:"cabin_type,omitempty"`
}

// FlightSegmentAPI Flight Segment API
type FlightSegmentAPI struct {
	AirOriginDestinationOptionID      *uuid.UUID       `json:"air_origin_destination_option_id,omitempty" gorm:"varchar(36);not null" swaggertype:"string" format:"uuid"` // The reference to the origin and destination.
	FlightSegmentStatus               *int             `json:"flight_segment_status,omitempty" gorm:"type:smallint"`                                                      // Indicates the status of the flight segment, e.g. confirmed, cancelled, etc.
	NumberInParty                     *int             `json:"number_in_party,omitempty" gorm:"type:smallint"`                                                            // Number of travelers associated with this segment.
	ETicketEligibility                *bool            `json:"e_ticket_eligibility,omitempty"`                                                                            // Specifies whether a flight segment is eligible for electronic ticketing.
	IsMealPlanIncluded                *bool            `json:"is_meal_plan_included,omitempty"`                                                                           // Indicates whether meal plan is included or not.
	MealPlanTypeID                    *uuid.UUID       `json:"meal_plan_type_id,omitempty" swaggertype:"string" format:"uuid"`                                            // The reference to the meal plan type, e.g. full board, half board, breakfast, etc.
	StopoverIndicator                 *bool            `json:"stopover_indicator,omitempty"`                                                                              // Indicates whether stopover is permitted or not.
	ConnectionType                    *string          `json:"connection_type,omitempty" gorm:"varchar(8)"`                                                               // The type of connection for this flight segment.
	Distance                          *string          `json:"distance,omitempty" gorm:"varchar(32)"`                                                                     // Distance acquired per flight segment, usually used for earning of frequent flyer miles.
	DistanceUnitOfMeasureID           *uuid.UUID       `json:"distance_unit_of_measure_id,omitempty" swaggertype:"string" format:"uuid"`                                  // The reference to the unit of measure, e.g. mile.
	FlightNumber                      *string          `json:"flight_number,omitempty" gorm:"varchar(16)"`                                                                // The flight number of the flight.
	ResBookDesignCode                 *string          `json:"res_book_design_code,omitempty" gorm:"varchar(2)"`                                                          // Specific Booking Class for this segment.
	ElapsedTime                       *int             `json:"elapsed_time,omitempty" gorm:"type:smallint"`                                                               // Flight duration
	CabinTypeID                       *uuid.UUID       `json:"cabin_type_id,omitempty" swaggertype:"string" format:"uuid"`                                                // The reference to the cabin type, e.g. economy, business, etc.
	CabinCode                         *string          `json:"cabin_code,omitempty" gorm:"varchar(2)"`                                                                    // Cabin Code for this segment.
	DepartureDatetime                 *strfmt.DateTime `json:"departure_datetime,omitempty" format:"date-time" swaggertype:"string" gorm:"type:timestamptz"`              // The date and time of the flight segment departure.
	ArrivalDatetime                   *strfmt.DateTime `json:"arrival_datetime,omitempty" format:"date-time" swaggertype:"string" gorm:"type:timestamptz"`                // The date and time of the flight segment arrival.
	StopQuantity                      *int             `json:"stop_quantity,omitempty" gorm:"type:smallint"`                                                              // The number of stops the flight makes.
	DepartureAirportLocationCode      *string          `json:"departure_airport_location_code,omitempty" gorm:"varchar(8)"`                                               // Location code used to identify a departure airport.
	DepartureAirportTerminal          *string          `json:"departure_airport_terminal,omitempty" gorm:"varchar(16)"`                                                   // Departure terminal
	DepartureAirportGate              *string          `json:"departure_airport_gate,omitempty" gorm:"varchar(16)"`                                                       // Departure gate.
	ArrivalAirportLocationCode        *string          `json:"arrival_airport_location_code,omitempty" gorm:"varchar(8)"`                                                 // Location code used to identify an arrival airport.
	ArrivalAirportTerminal            *string          `json:"arrival_airport_terminal,omitempty" gorm:"varchar(16)"`                                                     // Arrival terminal
	ArrivalAirportGate                *string          `json:"arrival_airport_gate,omitempty" gorm:"varchar(16)"`                                                         // Arrival gate
	OperatingAirlineFlightNumber      *string          `json:"operating_airline_flight_number,omitempty" gorm:"varchar(16)"`                                              // The flight number as assigned by the operating carrier.
	OperatingAirlineResBookDesignCode *string          `json:"operating_airline_res_book_design_code,omitempty" gorm:"varchar(2)"`                                        // The reservation booking designator of the operating carrier when different from the marketing carrier.
	OperatingAirlineCode              *string          `json:"operating_airline_code,omitempty" gorm:"varchar(16)"`                                                       // Identifies an airline by the code.
	OperatingAirlineCompanyName       *string          `json:"operating_airline_company_name,omitempty" gorm:"varchar(64)"`                                               // Used to provide the company name.
	MarketingAirlineFlightNumber      *string          `json:"marketing_airline_flight_number,omitempty" gorm:"varchar(16)"`                                              // The flight number as assigned by the marketing carrier.
	MarketingAirlineResBookDesignCode *string          `json:"marketing_airline_res_book_design_code,omitempty" gorm:"varchar(2)"`                                        // The reservation booking designator of the marketing carrier.
	MarketingAirlineCode              *string          `json:"marketing_airline_code,omitempty" gorm:"varchar(16)"`                                                       // Identifies an airline by the code.
	MarketingAirlineCompanyName       *string          `json:"marketing_airline_company_name,omitempty" gorm:"varchar(64)"`                                               // Used to provide the company name.
	MarriageGroup                     *string          `json:"marriage_group,omitempty" gorm:"varchar(1)"`                                                                // Indicator whether a set of air segments are considered married together | rule 'O': Not married, I': Married
	AirEquipmentType                  *string          `json:"air_equipment_type,omitempty" gorm:"varchar(3)"`                                                            // The type of equipment used for the flight (3 character IATA code).
	DepartureTimeZoneID               *uuid.UUID       `json:"departure_time_zone_id,omitempty" swaggertype:"string" format:"uuid"`                                       // The time zone of the flight segment departure.
	ArrivalTimeZoneID                 *uuid.UUID       `json:"arrival_time_zone_id,omitempty" swaggertype:"string" format:"uuid"`                                         // The time zone of the flight segment arrival.
	SeatsAvailable                    *int             `json:"seats_available,omitempty"`                                                                                 // Seats left
}
