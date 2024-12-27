package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

type FlightCachingServiceFee struct {
	basic.Base
	basic.DataOwner
	SessionID            *uuid.UUID `json:"session_id" gorm:"type:varchar(36);index:idx_flight_caching_service_fee_session_id"`
	ServiceFeeCategoryID *uuid.UUID `json:"service_fee_category_id" example:"bfb52132-0d53-4ff2-81b7-ca574e69556c"`
	ServiceFeeRateID     *uuid.UUID `json:"service_fee_rate_id" example:"bfb52132-0d53-4ff2-81b7-ca574e69556c"`
	IsVisible            *bool      `json:"is_visible" example:"true"`
	IsHidden             *bool      `json:"is_hidden" example:"false"`
	IsIncluded           *bool      `json:"is_included" example:"true"`
	Percent              *float64   `json:"percent" example:"2.5"`
	Amount               *float64   `json:"amount" example:"50000"`
	FinalAmount          *float64   `json:"final_amount" example:"50000"`
	FinalAmountPerTicket *float64   `json:"final_amount_per_ticket" example:"12500"`
	TotalAmount          *float64   `json:"total_amount" example:"50000"`
	// ChargeType           *string    `json:"charge_type" example:"ticket" comment:"Available value : ticket/person/reservation"`
	TaxTypeName  *string    `json:"tax_type_name" gorm:"varchar(256)" example:"Emergency Service Assistance 24 Hours Surcharge"`
	CurrencyID   *uuid.UUID `json:"currency_id" gorm:"varchar(36)"`
	ChargeTypeID *uuid.UUID `json:"charge_type_id" gorm:"varchar(36)"`
	// CurrencyCode         *string    `json:"currency_code" gorm:"varchar(3)"`
	// CurrencyName         *string    `json:"currency_name" gorm:"varchar(64)"`
	// CurrencySymbol       *string    `json:"currency_symbol" gorm:"varchar(64)"`
	IsSplitInvoice *bool       `json:"is_split_invoice" example:"true"`
	ChargeType     *ChargeType `json:"charge_type" gorm:"-"`
	Currency       *Currency   `json:"currency" gorm:"-"`
}
