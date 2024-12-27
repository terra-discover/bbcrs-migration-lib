package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// FlightCachingTax will used as generic tax fee data if exists
type FlightCachingTax struct {
	basic.Base
	basic.DataOwner
	SessionID                 *uuid.UUID `json:"session_id" gorm:"type:varchar(36);index:idx_flight_caching_tax_session_id"`
	InvoiceSeparateServiceFee *bool      `json:"invoice_separate_service_fee"` // check if "Corporate Setting.Invoice Separate Service Fee" is true/false
	FeeTaxTypeCode            *string    `json:"fee_tax_type_code"`
	TaxRateID                 *uuid.UUID `json:"tax_rate_id" example:"bfb52132-0d53-4ff2-81b7-ca574e69556c"`
	TaxCategoryID             *uuid.UUID `json:"tax_category_id" example:"bfb52132-0d53-4ff2-81b7-ca574e69556c"`
	ChargeTypeID              *uuid.UUID `json:"charge_type_id" example:"bfb52132-0d53-4ff2-81b7-ca574e69556c"`
	IsVisible                 *bool      `json:"is_visible" example:"true"`
	Percent                   *float64   `json:"percent" example:"2.5"`
	Amount                    *float64   `json:"amount" example:"50000"`
	FinalAmount               *float64   `json:"final_amount" example:"50000"`
	TotalAmount               *float64   `json:"total_amount" example:"50000"`
	ChargeType                *string    `json:"charge_type" example:"ticket" comment:"Available value : ticket/person/reservation"`
	CurrencyID                *uuid.UUID `json:"currency_id" gorm:"varchar(36)"`
}
