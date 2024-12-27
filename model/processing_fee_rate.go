package model

import (
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// ProcessingFeeRate schema
type ProcessingFeeRate struct {
	basic.Base
	basic.DataOwner
	ProcessingFeeRateAPI
	ProcessingFeeCategory             *ProcessingFeeCategory              `json:"processing_fee_category,omitempty" swaggerignore:"true"`
	FeeTaxType                        *FeeTaxType                         `json:"fee_tax_type,omitempty" swaggerignore:"true"`
	Currency                          *Currency                           `json:"currency,omitempty" swaggerignore:"true"`
	ChargeType                        *ChargeType                         `json:"charge_type,omitempty" swaggerignore:"true"`
	ProcessingFeeRateProductType      *ProcessingFeeRateProductType       `json:"processing_fee_rate_product_type,omitempty" gorm:"foreignKey:ProcessingFeeRateID;references:ID" swaggerignore:"true"`
	ProcessingFeeRateAirline          []ProcessingFeeRateAirline          `json:"processing_fee_rate_airline,omitempty" gorm:"foreignKey:ProcessingFeeRateID;references:ID" swaggerignore:"true"`
	ProcessingFeeRateDestinationGroup []ProcessingFeeRateDestinationGroup `json:"processing_fee_rate_destination_group,omitempty" gorm:"foreignKey:ProcessingFeeRateID;references:ID" swaggerignore:"true"`
}

// ProcessingFeeRateAPI schema
type ProcessingFeeRateAPI struct {
	ProcessingFeeCategoryID    *uuid.UUID `json:"processing_fee_category_id,omitempty" gorm:"type:varchar(36);not null"` // Processing Fee Category Id
	FeeTaxTypeID               *uuid.UUID `json:"fee_tax_type_id,omitempty" gorm:"type:varchar(36);not null"`            // Fee Tax Type Id
	CurrencyID                 *uuid.UUID `json:"currency_id,omitempty" gorm:"type:varchar(36)"`                         // Currency Id
	ChargeTypeID               *uuid.UUID `json:"charge_type_id,omitempty" gorm:"type:varchar(36)"`                      // Charge Type Id
	IsSummary                  *bool      `json:"is_summary,omitempty"`                                                  // Is Summary
	ParentProcessingFeeRateID  *uuid.UUID `json:"parent_processing_fee_rate_id,omitempty" gorm:"type:varchar(36)"`       // Parent Processing Fee Rate Id
	Percent                    *float64   `json:"percent,omitempty" gorm:"type:decimal(19,4)"`                           // Percent
	Amount                     *float64   `json:"amount,omitempty" gorm:"type:decimal(19,4)"`                            // Amount
	IsTaxInclusive             *bool      `json:"is_tax_inclusive,omitempty"`                                            // Is Tax Inclusive
	Cascade                    *bool      `json:"cascade,omitempty"`                                                     // Cascade
	ApplicationOrder           *int       `json:"application_order,omitempty"`                                           // Application Order
	AssignAllProductTypes      *bool      `json:"assign_all_product_types,omitempty"`                                    // Assign All Product Types
	AssignAllProductCategories *bool      `json:"assign_all_product_categories,omitempty"`                               // Assign All Product Categories
	AssignAllDestinationGroups *bool      `json:"assign_all_destination_groups,omitempty"`                               // Assign All Destination Groups
	AssignAllAirlineCategories *bool      `json:"assign_all_airline_categories,omitempty"`                               // Assign All Airline Categories
	AssignAllAirlines          *bool      `json:"assign_all_airlines,omitempty"`                                         // Assign All Airlines
	AssignAllSupplierTypes     *bool      `json:"assign_all_supplier_types,omitempty"`                                   // Assign All Supplier Types
	AssignAllHotelSuppliers    *bool      `json:"assign_all_hotel_suppliers,omitempty"`                                  // Assign All Hotel Suppliers
	IsIncluded                 *bool      `json:"is_included,omitempty"`                                                 // Is Included
	IsHidden                   *bool      `json:"is_hidden,omitempty"`                                                   // Is Hidden
	Description                *string    `json:"description,omitempty" gorm:"type:text"`                                // Description
}

// ProcessingFeeRateDataGet schema
type ProcessingFeeRateDataGet struct {
	ID                         *uuid.UUID       `json:"id,omitempty"`
	ProcessingFeeCategoryID    *uuid.UUID       `json:"processing_fee_category_id,omitempty"`                         // Processing Fee Category Id
	FeeTaxTypeID               *uuid.UUID       `json:"-"`                                                            // Fee Tax Type Id
	CurrencyID                 *uuid.UUID       `json:"-"`                                                            // Currency Id
	ChargeTypeID               *uuid.UUID       `json:"-"`                                                            // Charge Type Id
	CreatedAt                  *strfmt.DateTime `json:"created_at,omitempty" swaggertype:"string" format:"date-time"` // created at
	UpdatedAt                  *strfmt.DateTime `json:"updated_at,omitempty" swaggertype:"string" format:"date-time"` // updated at
	Sort                       *int64           `json:"sort,omitempty"`                                               // sort
	IsSummary                  *bool            `json:"is_summary,omitempty"`                                         // Is Summary
	ParentProcessingFeeRateID  *uuid.UUID       `json:"parent_processing_fee_rate_id,omitempty"`                      // Parent Processing Fee Rate Id
	Percent                    *float64         `json:"percent,omitempty"`                                            // Percent
	Amount                     *float64         `json:"amount,omitempty"`                                             // Amount
	IsTaxInclusive             *bool            `json:"is_tax_inclusive,omitempty"`                                   // Is Tax Inclusive
	Cascade                    *bool            `json:"cascade,omitempty"`                                            // Cascade
	ApplicationOrder           *int             `json:"application_order,omitempty"`                                  // Application Order
	AssignAllProductTypes      *bool            `json:"assign_all_product_types,omitempty"`                           // Assign All Product Types
	AssignAllProductCategories *bool            `json:"assign_all_product_categories,omitempty"`                      // Assign All Product Categories
	AssignAllDestinationGroups *bool            `json:"assign_all_destination_groups,omitempty"`                      // Assign All Destination Groups
	AssignAllAirlineCategories *bool            `json:"assign_all_airline_categories,omitempty"`                      // Assign All Airline Categories
	AssignAllAirlines          *bool            `json:"assign_all_airlines,omitempty"`                                // Assign All Airlines
	AssignAllSupplierTypes     *bool            `json:"assign_all_supplier_types,omitempty"`                          // Assign All Supplier Types
	AssignAllHotelSuppliers    *bool            `json:"assign_all_hotel_suppliers,omitempty"`                         // Assign All Hotel Suppliers
	IsIncluded                 *bool            `json:"is_included,omitempty"`                                        // Is Included
	IsHidden                   *bool            `json:"is_hidden,omitempty"`                                          // Is Hidden
	Description                *string          `json:"description,omitempty" gorm:"type:text"`
	ProcessingFee              *string          `json:"processing_fee,omitempty" gorm:"type:text"`
	ProcessingFeeRateDataGetAPI
}

// ProcessingFeeRateDataGetAPI schema
type ProcessingFeeRateDataGetAPI struct {
	ProcessingFeeCategory *ProcessingFeeCategory `json:"processing_fee_category,omitempty" gorm:"foreignKey:ID"`
	FeeTaxType            *FeeTaxType            `json:"fee_tax_type,omitempty" gorm:"foreignKey:FeeTaxTypeID;references:ID"`
	Currency              *Currency              `json:"currency,omitempty" gorm:"foreignKey:CurrencyID;references:ID"`
	ChargeType            *ChargeType            `json:"charge_type,omitempty" gorm:"foreignKey:ChargeTypeID;references:ID"`
}
