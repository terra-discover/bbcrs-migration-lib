package model

import (
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

type SabreCache struct {
	basic.Base
	basic.DataOwner
	SessionID          *uuid.UUID             `json:"session_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`
	OldSession         *uuid.UUID             `json:"old_session,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`
	IsModifyBooking    *bool                  `json:"is_modify_booking,omitempty"`
	IsModifyFlight     *bool                  `json:"is_modify_flight,omitempty"`
	DepartureDate      *strfmt.DateTime       `json:"departure_date,omitempty" format:"date-time" gorm:"type:timestamptz" default:"2022-06-05T10:10:00" swaggertype:"string"`
	ConversationID     *string                `json:"conversation_id,omitempty" gorm:"type:varchar(60)" swaggertype:"string" format:"uuid"`
	RefToMessageId     *string                `json:"ref_to_message_id,omitempty" gorm:"type:varchar(400)"`
	Security           *string                `json:"security,omitempty" gorm:"type:varchar(400)"`
	FlightID           *uuid.UUID             `json:"flight_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`
	FareID             *uuid.UUID             `json:"fare_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`
	ResBookCode        *string                `json:"res_book_code,omitempty" gorm:"type:varchar(1)"`
	AirItinerary       *string                `json:"air_itinerary,omitempty" gorm:"type:varchar(60)"`
	ValidatingCarrier  *string                `json:"validating_carrier,omitempty" gorm:"type:varchar(60)"`
	NumberInParty      *int                   `json:"number_in_party,omitempty"`
	CabinCode          *string                `json:"cabin_code,omitempty" gorm:"type:varchar(10)"`
	PriceQuote         *int                   `json:"price_quote,omitempty"`
	EmployeeName       *string                `json:"employee_name,omitempty" gorm:"type:varchar(60)"`
	FlightStatus       *string                `json:"flight_status,omitempty" gorm:"type:varchar(2)"`
	Expired            *strfmt.DateTime       `json:"expired,omitempty"`
	SeatAvailability   *bool                  `json:"seat_availability,omitempty"`
	AddonsAvailability *bool                  `json:"addons_availability,omitempty"`
	Ticket             *bool                  `json:"ticket,omitempty"`
	Currency           *string                `json:"currency"`
	Price              *float64               `json:"price"`
	PriceQuotes        *string                `json:"price_quotes,omitempty"`
	TicketedDate       *strfmt.DateTime       `json:"ticketed_date,omitempty" swaggertype:"string"`
	Flight             SabreCacheFlight       `json:"flight" gorm:"foreignKey:FlightID;references:ID"`
	Fare               SabreCacheFares        `json:"fare" gorm:"foreignKey:FareID;references:ID"`
	Traverllers        []SabreCacheTravellers `json:"travellers" gorm:"foreignKey:SessionID;references:SessionID"`
}

type SabreCalendarFares struct {
	basic.Base
	basic.DataOwner
	SessionID     *uuid.UUID       `json:"session_id,omitempty" gorm:"type:varchar(36)"`
	FlightID      *uuid.UUID       `json:"flight_id,omitempty" gorm:"type:varchar(36)"`
	Unique        *string          `json:"unique,omitempty"`
	DepartureDate *strfmt.DateTime `json:"departure_date,omitempty"`
	ArrivalDate   *strfmt.DateTime `json:"arrival_date,omitempty"`
	Price         *float64         `json:"price,omitempty"`
}

type SabreCachePriceQuote struct {
	basic.Base
	basic.DataOwner
	SessionID     *uuid.UUID `json:"session_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`
	Number        *int       `json:"number,omitempty"`
	NameID        *int       `json:"name_id,omitempty"`
	PassengerType *string    `json:"passenger_type,omitempty"`
	Quantity      *float64   `json:"quantity,omitempty"`
	Price         *float64   `json:"price,omitempty"`
}
type SabreAncillaryNumber struct {
	basic.Base
	basic.DataOwner
	SessionID *uuid.UUID `json:"session_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`
	Number    *int       `json:"number,omitempty"`
}

type SabreCacheTicket struct {
	basic.Base
	basic.DataOwner
	SessionID          *uuid.UUID       `json:"session_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`
	Index              *int             `json:"index,omitempty"`
	TicketNumber       *string          `json:"ticket_number,omitempty" gorm:"type:varchar(60)"`
	TicketType         *string          `json:"ticket_type,omitempty" gorm:"type:varchar(60)"`
	IndexTraveller     *int             `json:"index_traveller,omitempty"`
	FirstName          *string          `json:"first_name,omitempty" gorm:"type:varchar(60)"`
	LastName           *string          `json:"last_name,omitempty" gorm:"type:varchar(60)"`
	Currency           *string          `json:"currency,omitempty" gorm:"type:varchar(60)"`
	Amount             *float64         `json:"amount,omitempty"`
	NotValidBefore     *strfmt.Date     `json:"not_valid_before,omitempty"`
	NotValidAfter      *strfmt.Date     `json:"not_valid_after,omitempty"`
	LocalIssueDateTime *strfmt.DateTime `json:"local_issue_date_time,omitempty"`
}

type SabreIgnoreSegmentHX struct {
	basic.Base
	basic.DataOwner
	Locator *string `json:"locator,omitempty" gorm:"unique;not null"`
}
