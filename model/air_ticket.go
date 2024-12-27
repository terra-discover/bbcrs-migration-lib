package model

import (
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

type AirTicket struct {
	basic.Base
	basic.DataOwner
	AirTicketAPI
	AirItinerary           *AirItinerary             `json:"air_itinerary,omitempty" gorm:"foreignKey:AirItineraryID;references:ID"`
	AirTicketTraveler      *AirTicketTraveler        `json:"air_ticket_traveler,omitempty"`
	AirTicketFlightSegment *[]AirTicketFlightSegment `json:"air_ticket_flight_segment,omitempty"`
}

type AirTicketAPI struct {
	AirItineraryID           *uuid.UUID       `json:"air_itinerary_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"`
	TicketStatus             *int             `json:"ticket_status,omitempty" gorm:"type:smallint"`
	TicketDocumentNumber     *string          `json:"ticket_document_number,omitempty" gorm:"type:varchar(64);not null" swaggertype:"string"`
	BaseFareCurrencyID       *uuid.UUID       `json:"base_fare_currency_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string"`
	BaseFareAmount           *float64         `json:"base_fare_amount,omitempty"`
	EquivalentFareCurrencyID *uuid.UUID       `json:"equivalent_fare_currency_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string"`
	EquivalentFareAmount     *float64         `json:"equivalent_fare_amount,omitempty"`
	TotalTaxCurrencyID       *uuid.UUID       `json:"total_tax_currency_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string"`
	TotalTaxAmount           *float64         `json:"total_tax_amount,omitempty"`
	TotalFeeCurrencyID       *uuid.UUID       `json:"total_fee_currency_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string"`
	TotalFeeAmount           *float64         `json:"total_fee_amount,omitempty"`
	TotalFareCurrencyID      *uuid.UUID       `json:"total_fare_currency_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string"`
	TotalFareAmount          *float64         `json:"total_afare_amount,omitempty"`
	IssueTimestamp           *strfmt.DateTime `json:"issue_timestamp,omitempty" gorm:"type:timestamptz" format:"date-time"`
	IssueLocalTimestamp      *strfmt.DateTime `json:"issue_local_timestamp,omitempty" gorm:"type:timestamptz" format:"date-time"`
	IsHiddenFare             *bool            `json:"is_hidden_fare,omitempty"`
	NotValidBefore           *strfmt.DateTime `json:"not_valid_before,omitempty" gorm:"type:timestamptz" format:"date-time"`
	NotValidAfter            *strfmt.DateTime `json:"not_valid_after,omitempty" gorm:"type:timestamptz" format:"date-time"`
}
