package model

import (
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

type SabreFlightSession struct {
	OldSession *uuid.UUID `json:"old_session,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`
	SessionID  *uuid.UUID `json:"session_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`
	FlightID   *uuid.UUID `json:"flight_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`
}

// SabreCacheFlight caching search flight
type SabreCacheFlight struct {
	basic.Base
	basic.DataOwner
	SessionID         *uuid.UUID             `json:"session_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`
	Code              *string                `json:"code,omitempty" gorm:"type:varchar(60)"`              // FareBasis.Code
	DepartureDate     *strfmt.DateTime       `json:"departure_date,omitempty" gorm:"type:timestamp"`      // DepartureDateTime
	Airline           *string                `json:"airline,omitempty" gorm:"type:varchar(60)"`           // MarketingCarrier
	DepartureAirport  *string                `json:"departure_airport,omitempty" gorm:"type:varchar(60)"` // OriginLocation
	ArrivalAirport    *string                `json:"arrival_airport,omitempty" gorm:"type:varchar(60)"`   // DestinationLocation
	SequenceNumber    *int64                 `json:"sequence_number,omitempty" `
	DirectionInd      *string                `json:"direction_ind,omitempty" gorm:"type:varchar(60)"`
	ValidatingCarrier *string                `json:"validating_carrier,omitempty" gorm:"type:varchar(10)"`
	PriceQuote        *int                   `json:"price_quote,omitempty"`
	Airlines          []SabreCacheAirlines   `json:"airlines" gorm:"foreignKey:FlightID;references:ID"`
	Travellers        []SabreCacheTravellers `json:"travellers" gorm:"foreignKey:SessionID;references:SessionID"`
	Seats             []SabreCacheSeats      `json:"seats" gorm:"foreignKey:FlightID;references:ID"`
	Rules             []SabreCacheAirRules   `json:"rules" gorm:"foreignKey:FlightID;references:ID"`
	// Tickets

}
type SabreCacheAirlines struct {
	basic.Base
	basic.DataOwner
	SessionID        *uuid.UUID         `json:"session_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`
	FlightID         *uuid.UUID         `json:"flight_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`
	Index            *int               `json:"direction_ind,omitempty" gorm:"type:int"`
	ElapsedTime      *int64             `json:"elapsed_time,omitempty"`
	Key              *string            `json:"key,omitempty" gorm:"type:varchar(400)"`
	Unique           *string            `json:"unique,omitempty" gorm:"type:varchar(400)"`
	DepartureCountry *string            `json:"departure_country,omitempty" gorm:"type:varchar(60)"`
	ArrivalCountry   *string            `json:"arrival_country,omitempty" gorm:"type:varchar(60)"`
	Routes           []SabreCacheRoutes `json:"-" gorm:"foreignKey:AirlineID;references:ID"`
}
type SabreCacheRoutes struct {
	basic.Base
	basic.DataOwner
	SessionID        *uuid.UUID       `json:"session_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`
	AirlineID        *uuid.UUID       `json:"division_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`
	FlightID         *uuid.UUID       `json:"flight_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`
	Index            *int             `json:"index,omitempty" gorm:"type:int"`         // ex = 1
	IndexAirline     *int             `json:"index_airline,omitempty" gorm:"type:int"` // ex = 3
	DepartureAirport *string          `json:"departure_airport,omitempty" gorm:"type:varchar(60)"`
	ArrivalAirport   *string          `json:"arrival_airport,omitempty" gorm:"type:varchar(60)"`
	DepartureCity    *string          `json:"departure_city,omitempty" gorm:"type:varchar(60)"`
	ArrivalCity      *string          `json:"arrival_city,omitempty" gorm:"type:varchar(60)"`
	DepTerminalID    *string          `json:"dep_terminal_id,omitempty" gorm:"type:varchar(60)"`
	ArrTerminalID    *string          `json:"arr_terminal_id,omitempty" gorm:"type:varchar(60)"`
	DepDateTime      *strfmt.DateTime `json:"dep_date_time,omitempty" format:"date-time" swaggertype:"string"`
	ArrDateTime      *strfmt.DateTime `json:"arr_date_time,omitempty" format:"date-time" swaggertype:"string"`
	StopQuantity     *int             `json:"stop_quantity,omitempty" gorm:"type:int"`
	FlightNumber     *string          `json:"flight_number,omitempty" gorm:"type:varchar(60)"`
	ResBookDesigCode *string          `json:"res_book_desig_code,omitempty" gorm:"type:varchar(60)"`
	ElapsedTime      *int             `json:"profile_status,omitempty" gorm:"type:int"`
	OperatingAirline *string          `json:"operating_airline,omitempty" gorm:"type:varchar(60)"`
	MarketingAirline *string          `json:"marketing_airline,omitempty" gorm:"type:varchar(60)"`
	Milage           *int             `json:"milage,omitempty" gorm:"type:int"`
	Equipment        *string          `json:"equipment,omitempty" gorm:"type:varchar(60)"`
	ETicket          *string          `json:"e_ticket,omitempty" gorm:"type:varchar(60)"`
	DepTimeZone      *float64         `json:"dep_time_zone,omitempty"`
	ArrTimeZone      *float64         `json:"arr_time_zone,omitempty"`
	Cabin            *string          `json:"cabin,omitempty" gorm:"type:varchar(60)"`
	Meal             *string          `json:"meal,omitempty" gorm:"type:varchar(60)"`
	Baggage          *float64         `json:"baggage,omitempty"`
	BaggageInfo      *string          `json:"baggage_info,omitempty"`
	MarriageGrp      *string          `json:"marriage_grp,omitempty" gorm:"type:varchar(60)"`
	Aircraft         *string          `json:"aircraft,omitempty" gorm:"type:varchar(60)"`
}

type SabreCacheFares struct {
	basic.Base
	basic.DataOwner
	SessionID                   *uuid.UUID `json:"session_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`
	FlightID                    *uuid.UUID `json:"flight_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`
	FareID                      *uuid.UUID `json:"fare_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`
	TrxID                       *string    `json:"trx_id,omitempty" gorm:"type:varchar(36)"`
	PricingSource               *string    `json:"direction_ind,omitempty" gorm:"type:varchar(60)"`
	PricingSubSource            *string    `json:"pricing_sub_source,omitempty" gorm:"type:varchar(60)"`
	FareReturned                *string    `json:"FareReturned,omitempty" gorm:"type:varchar(60)"`
	LastTicketDate              *string    `json:"LastTicketDate,omitempty" gorm:"type:varchar(60)"`
	LastTicketTime              *string    `json:"last_ticket_time,omitempty" gorm:"type:varchar(60)"`
	Price                       *float64   `json:"price,omitempty" gorm:"type:decimal(22,2)"`
	Currency                    *string    `json:"currency,omitempty" gorm:"type:varchar(60)"`
	IsprivateFare               *bool      `json:"is_private_fare,omitempty"`
	ADT                         *int       `json:"adt,omitempty" gorm:"type:int"`
	CNN                         *int       `json:"cnn,omitempty" gorm:"type:int"`
	INF                         *int       `json:"inf,omitempty" gorm:"type:int"`
	NonRefundableIndicator      *bool      `json:"non_refundable_indicator,omitempty"`
	AvailabilityBreak           *string    `json:"availability_break,omitempty" gorm:"type:varchar(60)"`
	BookingCode                 *string    `json:"booking_code,omitempty" gorm:"type:varchar(60)"`
	DepartureAirportCode        *string    `json:"departure_airport_code,omitempty" gorm:"type:varchar(60)"`
	ArrivalAirportCode          *string    `json:"arrival_airport_code,omitempty" gorm:"type:varchar(60)"`
	FareComponentBeginAirport   *string    `json:"fare_component_begin_airport,omitempty" gorm:"type:varchar(60)"`
	FareComponentEndAirport     *string    `json:"fare_component_end_airport,omitempty" gorm:"type:varchar(60)"`
	FareComponentDirectionality *string    `json:"fare_component_directionality,omitempty" gorm:"type:varchar(60)"`
	FareComponentVendorCode     *string    `json:"fare_component_vendor_code,omitempty" gorm:"type:varchar(60)"`
	FareComponentFareTypeBitmap *string    `json:"fare_component_fare_type_bitmap,omitempty" gorm:"type:varchar(60)"`
	FareComponentFareType       *string    `json:"fare_component_fare_type,omitempty" gorm:"type:varchar(60)"`
	FareComponentFareTariff     *string    `json:"fare_component_fare_tariff,omitempty" gorm:"type:varchar(60)"`
	FareComponentFareRule       *string    `json:"fare_component_fare_rule,omitempty" gorm:"type:varchar(60)"`
	GovCarrier                  *string    `json:"gov_carrier,omitempty" gorm:"type:varchar(60)"`
	BasisCode                   *string    `json:"basis_code,omitempty" gorm:"type:varchar(60)"`
	BaggageSegmentID            *int       `json:"profile_status,omitempty" gorm:"type:int"`
	BaggagePieces               *int64     `json:"baggage_pieces,omitempty" example:"0"`
	BaggageWeight               *float64   `json:"baggage_weight,omitempty" example:"0"`
	ExchangeBefore              *bool      `json:"exchange_before,omitempty"`
	ExchangeAfter               *bool      `json:"exchange_after,omitempty"`
	ExchangeBeforeAmount        *int       `json:"exchange_before_amount,omitempty" example:"0"`
	ExchangeAfterAmount         *int       `json:"exchange_after_amount,omitempty" example:"0"`
	RefundBefore                *bool      `json:"refund_before,omitempty"`
	RefundAfter                 *bool      `json:"Refund_after,omitempty"`
	RefundbeforeFee             *int       `json:"Refund_before_ee,omitempty" example:"0"`
	RefundAfterFee              *int       `json:"refund_after_fee,omitempty" example:"0"`
	ItineraryRef                *string    `json:"itinerary_ref,omitempty" gorm:"type:varchar(60)"`
	ValidatingCarrier           *string    `json:"validating_carrier,omitempty" gorm:"type:varchar(60)"`
	Taxes                       *string    `json:"taxes,omitempty" gorm:"type:text"`
	TaxesInfant                 *string    `json:"taxes_infant,omitempty" gorm:"type:text"`
	TaxSummaries                *string    `json:"tax_summaries,omitempty" gorm:"type:text"`
	TaxSummariesInfant          *string    `json:"tax_summaries_infant,omitempty" gorm:"type:text"`
	AvailableCabins             *string    `json:"available_cabins,omitempty" gorm:"type:text"`
}

type SabreCacheCabin struct {
	basic.Base
	basic.DataOwner
	SessionID     *uuid.UUID `json:"session_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`
	FlightID      *uuid.UUID `json:"flight_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`
	FareID        *uuid.UUID `json:"fare_id,omitempty" gorm:"type:varchar(36); index:idx_sabre_cache_cabin_fare_id" swaggertype:"string" format:"uuid"`
	FareReference *string    `json:"fare_reference,omitempty" gorm:"type:varchar(60)"`
	FareName      *string    `json:"fare_name,omitempty" gorm:"type:varchar(60)"`
	Number        *int       `json:"number,omitempty" gorm:"type:int"`
	Cabin         *string    `json:"cabin,omitempty" gorm:"type:varchar(60)"`
	Meal          *string    `json:"meal,omitempty" gorm:"type:varchar(60)"`
	Baggage       *float64   `json:"baggage,omitempty"`
	BaggageInfo   *string    `json:"baggage_info,omitempty"`
	SegmentID     *int       `json:"segment_id,omitempty"`
}
type SabreCacheFareBasis struct {
	basic.Base
	basic.DataOwner
	SessionID                   *uuid.UUID `json:"session_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`
	FlightID                    *uuid.UUID `json:"flight_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`
	FareID                      *uuid.UUID `json:"FareID,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`
	Index                       *int       `json:"profile_status,omitempty" gorm:"type:int"`
	PrivateFareType             *string    `json:"GovCarrier,omitempty" gorm:"type:varchar(60)"`
	FareComponentReferenceID    *string    `json:"fare_component_reference_id,omitempty" gorm:"type:varchar(60)"`
	AccountCode                 *string    `json:"account_code,omitempty" gorm:"type:varchar(60)"`
	Mileage                     *string    `json:"mileage,omitempty" gorm:"type:varchar(60)"`
	BookingCode                 *string    `json:"booking_code,omitempty" gorm:"type:varchar(60)"`
	AvailabilityBreak           *string    `json:"availability_break,omitempty" gorm:"type:varchar(60)"`
	DepartureAirportCode        *string    `json:"departure_airport_code,omitempty" gorm:"type:varchar(60)"`
	ArrivalAirportCode          *string    `json:"ArrivalAirportCode,omitempty" gorm:"type:varchar(60)"`
	FareComponentBeginAirport   *string    `json:"fare_component_begin_airport,omitempty" gorm:"type:varchar(60)"`
	FareComponentEndAirport     *string    `json:"fare_component_end_airport,omitempty" gorm:"type:varchar(60)"`
	FareComponentDirectionality *string    `json:"fare_component_directionality,omitempty" gorm:"type:varchar(60)"`
	FareComponentFareTypeBitmap *string    `json:"fare_component_fare_type_bitmap,omitempty" gorm:"type:varchar(60)"`
	FareComponentFareType       *string    `json:"fare_component_fare_type,omitempty" gorm:"type:varchar(60)"`
	FareComponentFareTariff     *string    `json:"fare_component_fare_tariff,omitempty" gorm:"type:varchar(60)"`
	FareComponentFareRule       *string    `json:"fare_component_fare_rule,omitempty" gorm:"type:varchar(60)"`
	GovCarrier                  *string    `json:"gov_carrier,omitempty" gorm:"type:varchar(60)"`
	Code                        *string    `json:"code,omitempty" gorm:"type:varchar(60)"`
}

type SabreCacheSeats struct {
	basic.Base
	basic.DataOwner
	SessionID     *uuid.UUID `json:"session_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`
	FlightID      *uuid.UUID `json:"flight_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`
	PersonID      *uuid.UUID `json:"person_id,omitempty" gorm:"type:varchar(36)" validate:"required" format:"uuid"`
	FareID        *uuid.UUID `json:"fare_id,omitempty" gorm:"type:varchar(36);not null" validate:"required" format:"uuid"`
	RouteID       *uuid.UUID `json:"route_id,omitempty" gorm:"type:varchar(36);not null" validate:"required" format:"uuid"`
	SegmentNumber *string    `json:"segment_number,omitempty" gorm:"type:varchar(10)" example:"1" validate:"required"`
	Number        *string    `json:"number,omitempty" gorm:"type:varchar(100)" example:"16A" validate:"required"`
	Column        *string    `json:"column,omitempty" gorm:"type:varchar(100)" validate:"required" example:"A"`
	Row           *string    `json:"row,omitempty" gorm:"type:varchar(100)" validate:"required" example:"16"`
	NameNumber    *string    `json:"name_number,omitempty" gorm:"type:varchar(10)" validate:"required" example:"1.1"`
	NoInfantInd   *bool      `json:"no_infant_ind,omitempty" gorm:"type:boolean;default:false" example:"false"`
	Price         *float64   `json:"price,omitempty"`
	ChangeOfGauge *bool      `json:"change_of_gauge,omitempty" gorm:"type:boolean;default:false" example:"false"`
}
