package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

type FlightCachingFareDetail struct {
	basic.Base
	basic.DataOwner
	SessionID *uuid.UUID `json:"session_id,omitempty" gorm:"type:varchar(36);index:idx_flight_caching_fare_detail_session_id_traveller_id_flight_caching_fare_id;not null"`
	FlightCachingFareDetailAPI
	FlightCachingFares *FlightCachingFares                      `json:"flight_caching_fares,omitempty" gorm:"foreignKey:FlightCachingFareID"`
	FlightTraveller    *FlightCachingTraveller                  `json:"flight_traveller,omitempty" gorm:"foreignKey:TravellerID"`
	FareTax            *[]FlightCachingFareDetailTax            `json:"fare_tax,omitempty" gorm:"foreignKey:FlightCachingFareDetailID;references:ID"`
	FareTaxSummary     *[]FlightCachingFareDetailTaxSummary     `json:"fare_tax_summary,omitempty" gorm:"foreignKey:FlightCachingFareDetailID;references:ID"`
	BaggageAllowance   *FlightCachingFareDetailBaggageAllowance `json:"baggage_allowance,omitempty" gorm:"foreignKey:FlightCachingFareDetailID;references:ID"`
}

type FlightCachingFareDetailAPI struct {
	TravellerID         *uuid.UUID `json:"traveller_id,omitempty" validate:"required" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"`
	FlightCachingFareID *uuid.UUID `json:"flight_caching_fare_id,omitempty" validate:"required" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"`
	TrxID               *string    `json:"trx_id,omitempty"`
	Price               *float64   `json:"price,omitempty"`
	PriceBeforeTax      *float64   `json:"price_before_tax,omitempty"`
	PriceAfterMarkup    *float64   `json:"price_after_markup,omitempty"`
	TaxAmount           *float64   `json:"tax_amount,omitempty"`
	PrevPrice           *float64   `json:"prev_price,omitempty"`            // Prevous price amount include tax
	PrevPriceBeforeTax  *float64   `json:"prev_price_before_tax,omitempty"` // Prevous price amount exclude tax
	PrevTaxAmount       *float64   `json:"prev_tax_amount,omitempty"`       // Prevous amount
	PenaltyPrice        *float64   `json:"penalty_price,omitempty"`         // Represents the penalty price fee, if applicable
	PriceDifference     *float64   `json:"price_difference,omitempty"`      // Difference between new and previous prices (exclude penalty price). Positive: Upgrade, Negative: Downgrade
	PriceAmountDue      *float64   `json:"price_amount_due,omitempty"`      // Represents the amount due to pay for the current pricing
	CurrencyCode        *string    `json:"currency_code,omitempty" gorm:"type:varchar(36)"`
	CurrencyName        *string    `json:"currency_name,omitempty" gorm:"type:varchar(255)"`
	Description         *string    `json:"description,omitempty"`
	FareMessage         *string    `json:"fare_message,omitempty" gorm:"type:text"`
	IsSetToTraveller    *bool      `json:"is_set_to_traveller,omitempty"` // this value is used to inform whether this price has been set to the passenger or not
	PassengerType       *string    `json:"passenger_type,omitempty"`      // this value is used to inform the price refers to the type of passenger
	Additional          *string    `json:"additional,omitempty" gorm:"type:text"`
	IsPaid              *bool      `json:"is_paid,omitempty"` // IsPaid
}
