package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// TaxRate model
type TaxRate struct {
	basic.Base
	basic.DataOwner
	TaxRateAPI
	TaxRateProductType      *TaxRateProductType      `json:"tax_rate_product_type,omitempty"`
	TaxRateDestinationGroup *TaxRateDestinationGroup `json:"tax_rate_destination_group,omitempty"`
	TaxCategory             *TaxCategory             `json:"tax_category,omitempty"`
	FeeTaxType              *FeeTaxType              `json:"fee_tax_type,omitempty"`
	ChargeType              *ChargeType              `json:"charge_type,omitempty"`
	Currency                *Currency                `json:"currency,omitempty"`
}

// TaxRateAPI model
type TaxRateAPI struct {
	TaxCategoryID              *uuid.UUID `json:"tax_category_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid" swaggerignore:"true"`
	FeeTaxTypeID               *uuid.UUID `json:"fee_tax_type_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid"`
	Amount                     *float64   `json:"amount,omitempty"`
	CurrencyID                 *uuid.UUID `json:"currency_id,omitempty" gorm:"type:varchar(36);" format:"uuid"`
	Percent                    *float64   `json:"percent,omitempty"`
	ChargeTypeID               *uuid.UUID `json:"charge_type_id,omitempty" gorm:"column:charge_type_id;type:varchar(36)" format:"uuid"`
	IsSummary                  *bool      `json:"is_summary,omitempty"`
	ParentTaxID                *uuid.UUID `json:"parent_tax_id,omitempty" gorm:"type:varchar(36);" format:"uuid"`
	Cascade                    *bool      `json:"cascade,omitempty"`
	ApplicationOrder           *int       `json:"application_order,omitempty"`
	AssignAllProductTypes      *bool      `json:"assign_all_product_types,omitempty"`
	AssignAllProductCategories *bool      `json:"assign_all_product_categories,omitempty"`
	AssignAllDestinationGroup  *bool      `json:"assign_all_destination_group,omitempty"`
	Description                *string    `json:"description,omitempty" gorm:"type:text"`
	ChargeTypeString           *string    `json:"charge_type_name,omitempty" gorm:"-"`
}
