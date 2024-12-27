package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// MarkupRate Model
type MarkupRate struct {
	basic.Base
	basic.DataOwner
	MarkupRateAPI
	MarkupCategory             *MarkupCategory              `json:"markup_category" gorm:"foreignKey:MarkupCategoryID;references:ID"`
	FeeTaxType                 *FeeTaxType                  `json:"fee_tax_type" gorm:"foreignKey:FeeTaxTypeID;references:ID"`
	Currency                   *Currency                    `json:"currency" gorm:"foreignKey:CurrencyID;references:ID"`
	ChargeType                 *ChargeType                  `json:"charge_type,omitempty"`
	ParentMarkupRate           *MarkupRate                  `json:"parent_markup_rate,omitempty" gorm:"foreignKey:ParentMarkupRateID" swaggerignore:"true"`
	MarkupRateTranslation      *MarkupRateTranslation       `json:"markup_rate_translation,omitempty" swaggerignore:"true"`
	MarkupRateProductType      *MarkupRateProductType       `json:"markup_rate_product_type,omitempty" gorm:"foreignKey:MarkupRateID;references:ID" swaggerignore:"true"`
	MarkupRateDestinationGroup []MarkupRateDestinationGroup `json:"markup_rate_destination_group,omitempty" gorm:"foreignKey:MarkupRateID;references:ID" swaggerignore:"true"`
	MarkupRateAirline          []MarkupRateAirline          `json:"markup_rate_airline,omitempty" gorm:"foreignKey:MarkupRateID;references:ID" swaggerignore:"true"`
	MarkupRateAirlineCategory  []MarkupRateAirlineCategory  `json:"markup_rate_airline_category,omitempty" gorm:"foreignKey:MarkupRateID;references:ID" swaggerignore:"true"`
	MarkupRateSupplierType     []MarkupRateSupplierType     `json:"markup_rate_supplier_type,omitempty" gorm:"foreignKey:MarkupRateID;references:ID" swaggerignore:"true"`
	MarkupRateHotelSupplier    []MarkupRateHotelSupplier    `json:"markup_rate_hotel_supplier,omitempty" gorm:"foreignKey:MarkupRateID;references:ID" swaggerignore:"true"`
}

// MarkupRateAPI API
type MarkupRateAPI struct {
	MarkupCategoryID           *uuid.UUID `json:"markup_category_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid"`
	FeeTaxTypeID               *uuid.UUID `json:"fee_tax_type_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid"`
	IsSummary                  *bool      `json:"is_summary,omitempty"`
	Percent                    *float64   `json:"percent,omitempty" gorm:"type:decimal(19,4)" example:"1000.00"`
	CurrencyID                 *uuid.UUID `json:"currency_id,omitempty" gorm:"type:varchar(36);" format:"uuid"`
	Amount                     *float64   `json:"amount,omitempty" gorm:"type:decimal(19,4)" example:"1000.00"`
	ChargeTypeID               *uuid.UUID `json:"charge_type_id,omitempty" gorm:"type:varchar(36);" format:"uuid"`
	IsTaxInclusive             *bool      `json:"is_tax_inclusive,omitempty"`
	ParentMarkupRateID         *uuid.UUID `json:"parent_markup_rate_id,omitempty" gorm:"type:varchar(36);" format:"uuid"`
	Cascade                    *bool      `json:"cascade,omitempty"`
	ApplicationOrder           *int64     `json:"application_order" example:"10"`
	AssignAllProductTypes      *bool      `json:"assign_all_product_types,omitempty"`
	AssignAllProductCategories *bool      `json:"assign_all_product_categories,omitempty"`
	AssignAllDestinationGroups *bool      `json:"assign_all_destination_groups,omitempty"`
	AssignAllAirlineCategories *bool      `json:"assign_all_airline_categories,omitempty"`
	AssignAllAirlines          *bool      `json:"assign_all_airlines,omitempty"`
	AssignAllSupplierTypes     *bool      `json:"assign_all_supplier_types,omitempty"`
	AssignAllHotelSuppliers    *bool      `json:"assign_all_hotel_suppliers,omitempty"`
	AssignAllHotels            *bool      `json:"assign_all_hotels,omitempty"`
	Description                *string    `json:"description,omitempty" gorm:"type:text"`
}
