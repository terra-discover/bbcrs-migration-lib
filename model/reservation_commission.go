package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

type ReservationCommission struct {
	basic.Base
	basic.DataOwner
	ReservationCommissionAPI
}

type ReservationCommissionAPI struct {
	ReservationID              *uuid.UUID `json:"reservation_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"`
	BusinessPartnerID          *uuid.UUID `json:"business_partner_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`
	CommissionTypeID           *uuid.UUID `json:"commission_type_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`
	ChargeTypeID               *uuid.UUID `json:"charge_type_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`
	Percent                    float64    `json:"percent,omitempty"`
	CurrencyID                 *uuid.UUID `json:"currency_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`
	CommissionableAmount       *float64   `json:"commissionable_amount,omitempty"`
	FlatCommission             *float64   `json:"flat_commission,omitempty"`
	PrepaidAmount              *float64   `json:"prepaid_amount,omitempty"`
	CommissionPayableAmount    *float64   `json:"commission_payable_amount,omitempty"`
	IsTaxInclusive             *bool      `json:"is_tax_inclusive,omitempty"`
	IsServiceFeeInclusive      *bool      `json:"is_service_fee_inclusive,omitempty"`
	IsSellableProductInclusive *bool      `json:"is_sellable_product_inclusive,omitempty"`
	CommissionReason           *string    `json:"commission_reason,omitempty" gorm:"type:text"`
	Comment                    *string    `json:"comment,omitempty" gorm:"type:text"`
	Description                *string    `json:"description,omitempty" gorm:"type:text"`
}
