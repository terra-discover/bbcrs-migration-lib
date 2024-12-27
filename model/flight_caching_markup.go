package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// FlightCachingMarkup will used as generic markup_fee data if exists
type FlightCachingMarkup struct {
	basic.Base
	basic.DataOwner
	SessionID        *uuid.UUID `json:"session_id" gorm:"type:varchar(36);index:idx_flight_caching_markup_session_id"`
	MarkupCategoryID *uuid.UUID `json:"markup_category_id" example:"bfb52132-0d53-4ff2-81b7-ca574e69556c"`
	MarkupRateID     *uuid.UUID `json:"markup_rate_id" example:"bfb52132-0d53-4ff2-81b7-ca574e69556c"`
	IsVisible        *bool      `json:"is_visible" example:"true"`
	Percent          *float64   `json:"percent" example:"2.5"`        // from master data
	Amount           *float64   `json:"amount" example:"50000"`       // from master data
	FinalAmount      *float64   `json:"final_amount" example:"50000"` // final amount from calculation per person
	TotalAmount      *float64   `json:"total_amount" example:"50000"` // final amount from calculation total per transaction
	ChargeType       *string    `json:"charge_type" example:"ticket" comment:"Available value : ticket/person/reservation"`
	CurrencyID       *uuid.UUID `json:"currency_id" gorm:"varchar(36)"`
	ChargeTypeID     *uuid.UUID `json:"charge_type_id" gorm:"varchar(36)"`
	IsTaxInclusive   *bool      `json:"is_tax_inclusive"`
}
