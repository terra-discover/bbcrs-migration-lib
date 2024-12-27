package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// Reservation Product
type ReservationProduct struct {
	basic.Base
	basic.DataOwner
	ReservationProductAPI
	Product         *Product         `json:"product,omitempty" gorm:"foreignKey:ProductID;references:ID"`                 // Product
	Reservation     *Reservation     `json:"proposal,omitempty" gorm:"foreignKey:ReservationID;references:ID"`            // Reservation
	ProductPurchase *ProductPurchase `json:"product_purchase,omitempty" gorm:"foreignKey:ProductID;references:ProductID"` // ProductID
}

// ReservationProductAPI Reservation Product API
type ReservationProductAPI struct {
	ReservationID      *uuid.UUID `json:"proposal_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"`
	ProductID          *uuid.UUID `json:"product_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`
	ProductCode        *string    `json:"product_code,omitempty" gorm:"type:varchar(36)"`
	ProductName        *string    `json:"product_name,omitempty" gorm:"type:varchar(256)"`
	ProductUnit        *string    `json:"product_unit,omitempty" gorm:"type:varchar(32)"`
	ChargeTypeID       *uuid.UUID `json:"charge_type_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`
	CurrencyID         *uuid.UUID `json:"currency_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`
	IsMandatory        *bool      `json:"is_mandatory,omitempty"`
	IsTaxInclusive     *bool      `json:"is_tax_inclusive,omitempty"`
	TaxPercent         *float64   `json:"tax_percent,omitempty"`
	TaxAmount          *float64   `json:"tax_amount,omitempty"`
	AssignAllRatePlans *bool      `json:"assign_all_rate_plans,omitempty"`
	IsAccrual          *bool      `json:"is_accrual,omitempty"`
	IsRedeemable       *bool      `json:"is_redeemable,omitempty"`
	Description        *string    `json:"description,omitempty" gorm:"type:text"`
}
