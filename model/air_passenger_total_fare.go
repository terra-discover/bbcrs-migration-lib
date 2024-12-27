package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// AirPassengerTotalFare Air Passenger Total Fare
type AirPassengerTotalFare struct {
	basic.Base
	basic.DataOwner
	AirPassengerTotalFareAPI
	AirPassengerTotalFareTax              []AirPassengerTotalFareTax             `json:"air_passenger_total_fare_tax,omitempty"`
	AirPassengerTotalFareTaxSummary       []AirPassengerTotalFareTaxSummary      `json:"air_passenger_total_fare_tax_summary,omitempty"`
	AirPassengerTotalFareBasis            *AirPassengerTotalFareBasis            `json:"air_passenger_total_fare_basis,omitempty"`
	AirPassengerTotalFareBaggageAllowance *AirPassengerTotalFareBaggageAllowance `json:"air_passenger_total_fare_baggage_allowance,omitempty"`
	AirPassengerTotalFareMessage          []AirPassengerTotalFareMessage         `json:"air_passenger_total_fare_message,omitempty"`
	AirPassengerTotalFareFlightSegment    *AirPassengerTotalFareFlightSegment    `json:"air_passenger_total_fare_flight_segment,omitempty"`
}

// AirPassengerTotalFareAPI Air Passenger Total Fare API
type AirPassengerTotalFareAPI struct {
	AirItineraryID           *uuid.UUID `json:"air_itinerary_id" gorm:"type:varchar(36);not null"`
	PassengerIndexNumber     *int       `json:"passenger_index_number,omitempty"`
	PassengerTypeID          *uuid.UUID `json:"passenger_type_id,omitempty" gorm:"type:varchar(36)"` //Get Passenger Type Id from Air Passenger Fare Info[].Passenger Type Id
	PricingSource            *string    `json:"pricing_source,omitempty" gorm:"type:varchar(100)"`   //Get Pricing Source from Air Passenger Fare Info[].Pricing Source
	NonRefundable            *bool      `json:"non_refundable,omitempty"`
	BaseFareCurrencyID       *uuid.UUID `json:"base_fare_currency_id,omitempty" gorm:"type:varchar(36)"`
	BaseFareAmount           *float64   `json:"base_fare_amount,omitempty"`
	EquivalentFareCurrencyID *uuid.UUID `json:"equivalent_fare_currency_id,omitempty" gorm:"type:varchar(36)"`
	EquivalentFareAmount     *float64   `json:"equivalent_fare_amount,omitempty"`
	TotalTaxCurrencyID       *uuid.UUID `json:"total_tax_currency_id,omitempty" gorm:"type:varchar(36)"`
	TotalTaxAmount           *float64   `json:"total_tax_amount,omitempty"`
	TotalFareCurrencyID      *uuid.UUID `json:"total_fare_currency_id,omitempty" gorm:"type:varchar(36)"`
	TotalFareAmount          *float64   `json:"total_fare_amount,omitempty"`
	CurrencyID               *uuid.UUID `json:"currency_id,omitempty" gorm:"type:varchar(36)"`
	TotalAmountBeforeTax     *float64   `json:"total_amount_before_tax,omitempty"`
	TotalAmountAfterTax      *float64   `json:"total_amount_after_tax,omitempty"`
	CommisionPercent         *float64   `json:"commision_percent,omitempty"`
	CommisionAmount          *float64   `json:"commision_amount,omitempty"`
}
