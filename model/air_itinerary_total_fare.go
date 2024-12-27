package model

import (
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// AirItineraryTotalFare Air Itinerary Total Fare
type AirItineraryTotalFare struct {
	basic.Base
	basic.DataOwner
	AirItineraryTotalFareAPI
}

// AirItineraryTotalFareAPI Air Itinerary Total Fare Api
type AirItineraryTotalFareAPI struct {
	AirItineraryID                   *uuid.UUID       `json:"air_itinerary_id" gorm:"type:varchar(36);not null"`
	PricingSource                    *string          `json:"pricing_source" gorm:"type:varchar(100)"` //Get Pricing Source from Air Passenger Fare Info[].Pricing Source
	ValidatingAirlineCode            *string          `json:"validationg_airline_code" gorm:"type:varchar(50)"`
	QuoteID                          *string          `json:"quote_id,omitempty" gorm:"type:varchar(128)"`
	ETicketEligibility               *string          `json:"e_ticket_eligibility,omitempty" gorm:"type:varchar(16)"`
	BaseFareCurrencyID               *uuid.UUID       `json:"base_fare_currency_id" gorm:"type:varchar(36)"`
	BaseFareAmount                   *float64         `json:"base_fare_amount"`
	EquivalentFareCurrencyID         *uuid.UUID       `json:"equivalent_fare_currency_id" gorm:"type:varchar(36)"`
	EquivalentFareAmount             *float64         `json:"equivalent_fare_amount"`
	TotalTaxCurrencyID               *uuid.UUID       `json:"total_tax_currency_id" gorm:"type:varchar(36)"`
	TotalTaxAmount                   *float64         `json:"total_tax_amount"`
	TotalFareCurrencyID              *uuid.UUID       `json:"total_fare_currency_id" gorm:"type:varchar(36)"`
	TotalFareAmount                  *float64         `json:"total_fare_amount"`
	CurrencyID                       *uuid.UUID       `json:"currency_id" gorm:"type:varchar(36)"`
	TotalAmountBeforeTax             *float64         `json:"total_amount_before_tax"`
	TotalAmountAfterTax              *float64         `json:"total_amount_after_tax"`
	CommissionPercent                *float64         `json:"commission_percent,omitempty" gorm:"type:decimal(19,4)"`
	CommissionCurrencyID             *uuid.UUID       `json:"commission_currency_id,omitempty" gorm:"type:varchar(36)"`
	CommissionAmount                 *float64         `json:"commision_amount,omitempty" gorm:"type:decimal(19,4)"`
	BasePenaltyCurrencyID            *uuid.UUID       `json:"base_penalty_currency_id,omitempty" gorm:"type:varchar(36)"`
	BasePenaltyAmount                *float64         `json:"base_penalty_amount,omitempty" gorm:"type:decimal(19,4)"`
	EquivalentPenaltyCurrencyID      *uuid.UUID       `json:"equivalent_penalty_currency_id" gorm:"type:varchar(36)"`
	EquivalentPenaltyAmount          *float64         `json:"equivalent_penalty_amount" gorm:"type:decimal(19,4)"`
	TotalPenaltyAmountBeforeTax      *float64         `json:"total_penalty_amount_before_tax" gorm:"type:decimal(19,4)"`
	TotalPenaltyAmountAfterTax       *float64         `json:"total_penalty_amount_after_tax" gorm:"type:decimal(19,4)"`
	CurrencyConversionFromCurrencyID *uuid.UUID       `json:"currency_conversion_from_currency_id,omitempty" gorm:"type:varchar(36)"`
	CurrencyConversionToCurrencyID   *uuid.UUID       `json:"currency_conversion_to_currency_id,omitempty" gorm:"type:varchar(36)"`
	CurrencyConversionMultipleRate   *uuid.UUID       `json:"currency_conversion_multiple_rate,omitempty" gorm:"type:decimal(19,4)"`
	CancellationDeadline             *strfmt.DateTime `json:"cancellation_deadline,omitempty" gorm:"type:timestamptz"`
	CancellationAt                   *strfmt.DateTime `json:"cancellation_at,omitempty" gorm:"type:timestamptz"`
	CancellationReason               *string          `json:"cancellation_reason,omitempty" gorm:"type:varchar(4000)"`
}
