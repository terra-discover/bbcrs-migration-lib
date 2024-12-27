package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

type ReservationTaxRate struct {
	basic.Base
	basic.DataOwner
	ReservationTaxRateAPI
}

type ReservationTaxRateAPI struct {
	ReservationID    *uuid.UUID `json:"reservation_id,omitempty" gorm:"type:varchar(36);not null;" swaggertype:"string" format:"uuid"`
	TaxCategoryID    *uuid.UUID `json:"tax_category_id,omitempty" gorm:"type:varchar(36);not null;" swaggertype:"string" format:"uuid"`
	FeeTaxTypeID     *uuid.UUID `json:"fee_tax_type_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`
	FeeTaxTypeCode   *string    `json:"fee_tax_type_code,omitempty" gorm:"type:varchar(36)" swaggertype:"string"`
	FeeTaxTypeName   *string    `json:"fee_tax_type_name,omitempty" gorm:"type:varchar(256)" swaggertype:"string"`
	IsSummary        *bool      `json:"is_summary,omitempty"`
	Percent          *float64   `json:"percent,omitempty"`
	CurrencyID       *string    `json:"currency_id,omitempty" gorm:"type:varchar(36)"`
	Amount           *float64   `json:"amount,omitempty"`
	ChargeTypeID     *string    `json:"charge_type_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string"`
	ParentTaxRateID  *string    `json:"parent_tax_rate_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string"`
	Cascade          *bool      `json:"cascade,omitempty"`
	ApplicationOrder *int       `json:"application_order,omitempty" gorm:"type:int"`
	TotalAmount      *float64   `json:"total_amount,omitempty"`
	Description      *string    `json:"description,omitempty" gorm:"type:text"`
}
